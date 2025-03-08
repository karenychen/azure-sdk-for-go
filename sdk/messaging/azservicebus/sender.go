// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azservicebus

import (
	"context"
	"errors"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/amqpwrap"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/tracing"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/utils"
	"github.com/Azure/go-amqp"
)

type (
	// Sender is used to send messages as well as schedule them to be delivered at a later date.
	Sender struct {
		tracer         tracing.Tracer
		queueOrTopic   string
		cleanupOnClose func()
		links          internal.AMQPLinks
		retryOptions   RetryOptions
	}
)

// MessageBatchOptions contains options for the `Sender.NewMessageBatch` function.
type MessageBatchOptions struct {
	// MaxBytes overrides the max size (in bytes) for a batch.
	// By default NewMessageBatch will use the max message size provided by the service.
	MaxBytes uint64
}

// NewMessageBatch can be used to create a batch that contain multiple
// messages. Sending a batch of messages is more efficient than sending the
// messages one at a time.
// If the operation fails it can return an [*azservicebus.Error] type if the failure is actionable.
func (s *Sender) NewMessageBatch(ctx context.Context, options *MessageBatchOptions) (*MessageBatch, error) {
	var batch *MessageBatch

	err := s.links.Retry(ctx, EventSender, "send", func(ctx context.Context, lwid *internal.LinksWithID, args *utils.RetryFnArgs) error {
		maxBytes := lwid.Sender.MaxMessageSize()

		if options != nil && options.MaxBytes != 0 {
			maxBytes = options.MaxBytes
		}

		batch = newMessageBatch(maxBytes)
		return nil
	}, s.retryOptions, nil)

	if err != nil {
		return nil, internal.TransformError(err)
	}

	return batch, nil
}

// SendMessageOptions contains optional parameters for the SendMessage function.
type SendMessageOptions struct {
	// For future expansion
}

// SendMessage sends a Message to a queue or topic.
// If the operation fails it can return:
//   - [ErrMessageTooLarge] if the message is larger than the maximum allowed link size.
//   - An [*azservicebus.Error] type if the failure is actionable.
func (s *Sender) SendMessage(ctx context.Context, message *Message, options *SendMessageOptions) error {
	return s.sendMessage(ctx, message)
}

// SendAMQPAnnotatedMessageOptions contains optional parameters for the SendAMQPAnnotatedMessage function.
type SendAMQPAnnotatedMessageOptions struct {
	// For future expansion
}

// SendAMQPAnnotatedMessage sends an AMQPMessage to a queue or topic.
// Using an AMQPMessage allows for advanced use cases, like payload encoding, as well as better
// interoperability with pure AMQP clients.
// If the operation fails it can return:
//   - [ErrMessageTooLarge] if the message is larger than the maximum allowed link size.
//   - An [*azservicebus.Error] type if the failure is actionable.
func (s *Sender) SendAMQPAnnotatedMessage(ctx context.Context, message *AMQPAnnotatedMessage, options *SendAMQPAnnotatedMessageOptions) error {
	return s.sendMessage(ctx, message)
}

// SendMessageBatchOptions contains optional parameters for the SendMessageBatch function.
type SendMessageBatchOptions struct {
	// For future expansion
}

// SendMessageBatch sends a MessageBatch to a queue or topic.
// Message batches can be created using [Sender.NewMessageBatch].
// If the operation fails it can return an [*azservicebus.Error] type if the failure is actionable.
func (s *Sender) SendMessageBatch(ctx context.Context, batch *MessageBatch, options *SendMessageBatchOptions) (err error) {
	// We start a trace span for the send operation, but do not start a creation span for each message.
	// This needs to be done through a future fix to enable injecting context when we add to the MessageBatch.
	err = s.links.Retry(ctx, EventSender, "SendMessageBatch", func(ctx context.Context, lwid *internal.LinksWithID, args *utils.RetryFnArgs) error {
		return lwid.Sender.Send(ctx, batch.toAMQPMessage(), nil)
	}, s.retryOptions, &tracing.StartSpanOptions{
		Tracer:        s.tracer,
		OperationName: tracing.SendOperationName,
		Attributes:    getMessageBatchSpanAttributes(int(batch.NumMessages())),
	})

	return internal.TransformError(err)
}

// ScheduleMessagesOptions contains optional parameters for the ScheduleMessages function.
type ScheduleMessagesOptions struct {
	// For future expansion
}

// ScheduleMessages schedules a slice of Messages to appear on Service Bus Queue/Subscription at a later time.
// Returns the sequence numbers of the messages that were scheduled.  Messages that haven't been
// delivered can be cancelled using `Receiver.CancelScheduleMessage(s)`
// If the operation fails it can return an [*azservicebus.Error] type if the failure is actionable.
func (s *Sender) ScheduleMessages(ctx context.Context, messages []*Message, scheduledEnqueueTime time.Time, options *ScheduleMessagesOptions) ([]int64, error) {
	sequenceNumbers, err := scheduleMessages(ctx, s.tracer, s.links, s.retryOptions, messages, scheduledEnqueueTime)
	return sequenceNumbers, err
}

// ScheduleAMQPAnnotatedMessagesOptions contains optional parameters for the ScheduleAMQPAnnotatedMessages function.
type ScheduleAMQPAnnotatedMessagesOptions struct {
	// For future expansion
}

// ScheduleAMQPAnnotatedMessages schedules a slice of Messages to appear on Service Bus Queue/Subscription at a later time.
// Returns the sequence numbers of the messages that were scheduled.  Messages that haven't been
// delivered can be cancelled using `Receiver.CancelScheduleMessage(s)`
// If the operation fails it can return an [*azservicebus.Error] type if the failure is actionable.
func (s *Sender) ScheduleAMQPAnnotatedMessages(ctx context.Context, messages []*AMQPAnnotatedMessage, scheduledEnqueueTime time.Time, options *ScheduleAMQPAnnotatedMessagesOptions) ([]int64, error) {
	sequenceNumbers, err := scheduleMessages(ctx, s.tracer, s.links, s.retryOptions, messages, scheduledEnqueueTime)
	return sequenceNumbers, err
}

func scheduleMessages[T amqpCompatibleMessage](ctx context.Context, tracer tracing.Tracer, links internal.AMQPLinks, retryOptions RetryOptions, messages []T, scheduledEnqueueTime time.Time) ([]int64, error) {
	var err error
	scheduleCtx, endSpan := tracing.StartSpan(ctx, &tracing.StartSpanOptions{
		Tracer:        tracer,
		OperationName: tracing.ScheduleOperationName,
		Attributes:    getMessageBatchSpanAttributes(len(messages)),
	})
	defer func() { endSpan(err) }()
	scheduleSpan := tracer.SpanFromContext(scheduleCtx)

	var amqpMessages []*amqp.Message

	for _, m := range messages {
		amqpMessage := m.toAMQPMessage()
		createMessageSpan(ctx, tracer, scheduleSpan, amqpMessage)
		amqpMessages = append(amqpMessages, amqpMessage)
	}

	var sequenceNumbers []int64

	err = links.Retry(ctx, EventSender, "ScheduleMessages", func(ctx context.Context, lwv *internal.LinksWithID, args *utils.RetryFnArgs) error {
		sn, err := internal.ScheduleMessages(ctx, lwv.RPC, lwv.Sender.LinkName(), scheduledEnqueueTime, amqpMessages)

		if err != nil {
			return err
		}
		sequenceNumbers = sn
		return nil
	}, retryOptions, nil)

	return sequenceNumbers, internal.TransformError(err)
}

// MessageBatch changes

// CancelScheduledMessagesOptions contains optional parameters for the CancelScheduledMessages function.
type CancelScheduledMessagesOptions struct {
	// For future expansion
}

// CancelScheduledMessages cancels multiple messages that were scheduled.
// If the operation fails it can return an [*azservicebus.Error] type if the failure is actionable.
func (s *Sender) CancelScheduledMessages(ctx context.Context, sequenceNumbers []int64, options *CancelScheduledMessagesOptions) error {
	err := s.links.Retry(ctx, EventSender, "CancelScheduledMessages", func(ctx context.Context, lwv *internal.LinksWithID, args *utils.RetryFnArgs) error {
		return internal.CancelScheduledMessages(ctx, lwv.RPC, lwv.Sender.LinkName(), sequenceNumbers)
	}, s.retryOptions, &tracing.StartSpanOptions{
		Tracer:        s.tracer,
		OperationName: tracing.CancelScheduledOperationName,
		Attributes:    getMessageBatchSpanAttributes(len(sequenceNumbers)),
	})

	return internal.TransformError(err)
}

// Close permanently closes the Sender.
func (s *Sender) Close(ctx context.Context) error {
	s.cleanupOnClose()
	return s.links.Close(ctx, true)
}

func (s *Sender) sendMessage(ctx context.Context, message amqpCompatibleMessage) error {
	err := s.links.Retry(ctx, EventSender, "SendMessage", func(ctx context.Context, lwid *internal.LinksWithID, args *utils.RetryFnArgs) error {
		msg := message.toAMQPMessage()
		s.tracer.Inject(ctx, msg)
		return lwid.Sender.Send(ctx, msg, nil)
	}, RetryOptions(s.retryOptions), &tracing.StartSpanOptions{
		Tracer:        s.tracer,
		OperationName: tracing.SendOperationName,
		Attributes:    getMessageSpanAttributes(message.toAMQPMessage()),
	})

	if amqpErr := (*amqp.Error)(nil); errors.As(err, &amqpErr) && amqpErr.Condition == amqp.ErrCondMessageSizeExceeded {
		return ErrMessageTooLarge
	}

	return internal.TransformError(err)
}

func createMessageSpan(ctx context.Context, tracer tracing.Tracer, sendOrScheduleSpan tracing.Span, message *amqp.Message) {
	// derive a new context to be used for creation spans
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ctx, endSpan := tracing.StartSpan(ctx, &tracing.StartSpanOptions{
		Tracer:        tracer,
		OperationName: tracing.CreateOperationName,
		Attributes:    getMessageSpanAttributes(message),
	})
	defer func() { endSpan(nil) }()
	sendOrScheduleSpan.AddLink(tracer.LinkFromContext(ctx, getMessageIDAttribute(message)...))
	tracer.Inject(ctx, message)
}

func (sender *Sender) createSenderLink(ctx context.Context, session amqpwrap.AMQPSession) (amqpwrap.AMQPSenderCloser, amqpwrap.AMQPReceiverCloser, error) {
	amqpSender, err := session.NewSender(
		ctx,
		sender.queueOrTopic,
		&amqp.SenderOptions{
			SettlementMode:              amqp.SenderSettleModeMixed.Ptr(),
			RequestedReceiverSettleMode: amqp.ReceiverSettleModeFirst.Ptr(),
		})

	if err != nil {
		return nil, nil, err
	}

	return amqpSender, nil, nil
}

type newSenderArgs struct {
	tracer         tracing.Tracer
	ns             internal.NamespaceForAMQPLinks
	queueOrTopic   string
	cleanupOnClose func()
	retryOptions   RetryOptions
}

func newSender(args newSenderArgs) (*Sender, error) {
	if err := args.ns.Check(); err != nil {
		return nil, err
	}

	sender := &Sender{
		tracer:         args.tracer,
		queueOrTopic:   args.queueOrTopic,
		cleanupOnClose: args.cleanupOnClose,
		retryOptions:   args.retryOptions,
	}

	sender.links = internal.NewAMQPLinks(internal.NewAMQPLinksArgs{
		NS:                  args.ns,
		EntityPath:          args.queueOrTopic,
		CreateLinkFunc:      sender.createSenderLink,
		GetRecoveryKindFunc: internal.GetRecoveryKind,
	})

	return sender, nil
}

// amqpCompatibleMessage is implemented by all the messages that can be
// converted to amqp.Message
// Implemented by AMQPMessage, MessageBatch and Message.
type amqpCompatibleMessage interface {
	toAMQPMessage() *amqp.Message
}
