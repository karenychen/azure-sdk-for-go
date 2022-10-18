//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_CreateOrUpdate.json
func ExampleActivityLogAlertsClient_CreateOrUpdate_createOrUpdateAnActivityLogAlertRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewActivityLogAlertsClient("187f412d-1758-44d9-b052-169e2564721d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "MyResourceGroup", "SampleActivityLogAlertRule", armmonitor.ActivityLogAlertResource{
		Location: to.Ptr("Global"),
		Tags:     map[string]*string{},
		Properties: &armmonitor.AlertRuleProperties{
			Description: to.Ptr("Description of sample Activity Log Alert rule."),
			Actions: &armmonitor.ActionList{
				ActionGroups: []*armmonitor.ActionGroupAutoGenerated{
					{
						ActionGroupID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/actionGroups/SampleActionGroup"),
						WebhookProperties: map[string]*string{
							"sampleWebhookProperty": to.Ptr("SamplePropertyValue"),
						},
					}},
			},
			Condition: &armmonitor.AlertRuleAllOfCondition{
				AllOf: []*armmonitor.AlertRuleAnyOfOrLeafCondition{
					{
						Equals: to.Ptr("Administrative"),
						Field:  to.Ptr("category"),
					},
					{
						Equals: to.Ptr("Error"),
						Field:  to.Ptr("level"),
					}},
			},
			Enabled: to.Ptr(true),
			Scopes: []*string{
				to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_CreateOrUpdateRuleWithAnyOfCondition.json
func ExampleActivityLogAlertsClient_CreateOrUpdate_createOrUpdateAnActivityLogAlertRuleWithAnyOfCondition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewActivityLogAlertsClient("187f412d-1758-44d9-b052-169e2564721d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "MyResourceGroup", "SampleActivityLogAlertRuleWithAnyOfCondition", armmonitor.ActivityLogAlertResource{
		Location: to.Ptr("Global"),
		Tags:     map[string]*string{},
		Properties: &armmonitor.AlertRuleProperties{
			Description: to.Ptr("Description of sample Activity Log Alert rule with 'anyOf' condition."),
			Actions: &armmonitor.ActionList{
				ActionGroups: []*armmonitor.ActionGroupAutoGenerated{
					{
						ActionGroupID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/actionGroups/SampleActionGroup"),
						WebhookProperties: map[string]*string{
							"sampleWebhookProperty": to.Ptr("SamplePropertyValue"),
						},
					}},
			},
			Condition: &armmonitor.AlertRuleAllOfCondition{
				AllOf: []*armmonitor.AlertRuleAnyOfOrLeafCondition{
					{
						Equals: to.Ptr("ServiceHealth"),
						Field:  to.Ptr("category"),
					},
					{
						AnyOf: []*armmonitor.AlertRuleLeafCondition{
							{
								Equals: to.Ptr("Incident"),
								Field:  to.Ptr("properties.incidentType"),
							},
							{
								Equals: to.Ptr("Maintenance"),
								Field:  to.Ptr("properties.incidentType"),
							}},
					}},
			},
			Enabled: to.Ptr(true),
			Scopes: []*string{
				to.Ptr("subscriptions/187f412d-1758-44d9-b052-169e2564721d")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_CreateOrUpdateRuleWithContainsAny.json
func ExampleActivityLogAlertsClient_CreateOrUpdate_createOrUpdateAnActivityLogAlertRuleWithContainsAny() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewActivityLogAlertsClient("187f412d-1758-44d9-b052-169e2564721d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "MyResourceGroup", "SampleActivityLogAlertRuleWithContainsAny", armmonitor.ActivityLogAlertResource{
		Location: to.Ptr("Global"),
		Tags:     map[string]*string{},
		Properties: &armmonitor.AlertRuleProperties{
			Description: to.Ptr("Description of sample Activity Log Alert rule with 'containsAny'."),
			Actions: &armmonitor.ActionList{
				ActionGroups: []*armmonitor.ActionGroupAutoGenerated{
					{
						ActionGroupID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/actionGroups/SampleActionGroup"),
						WebhookProperties: map[string]*string{
							"sampleWebhookProperty": to.Ptr("SamplePropertyValue"),
						},
					}},
			},
			Condition: &armmonitor.AlertRuleAllOfCondition{
				AllOf: []*armmonitor.AlertRuleAnyOfOrLeafCondition{
					{
						Equals: to.Ptr("ServiceHealth"),
						Field:  to.Ptr("category"),
					},
					{
						ContainsAny: []*string{
							to.Ptr("North Europe"),
							to.Ptr("West Europe")},
						Field: to.Ptr("properties.impactedServices[*].ImpactedRegions[*].RegionName"),
					}},
			},
			Enabled: to.Ptr(true),
			Scopes: []*string{
				to.Ptr("subscriptions/187f412d-1758-44d9-b052-169e2564721d")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_Get.json
func ExampleActivityLogAlertsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewActivityLogAlertsClient("187f412d-1758-44d9-b052-169e2564721d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "MyResourceGroup", "SampleActivityLogAlertRule", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_Delete.json
func ExampleActivityLogAlertsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewActivityLogAlertsClient("187f412d-1758-44d9-b052-169e2564721d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx, "MyResourceGroup", "SampleActivityLogAlertRule", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_Update.json
func ExampleActivityLogAlertsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewActivityLogAlertsClient("187f412d-1758-44d9-b052-169e2564721d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "MyResourceGroup", "SampleActivityLogAlertRule", armmonitor.AlertRulePatchObject{
		Properties: &armmonitor.AlertRulePatchProperties{
			Enabled: to.Ptr(false),
		},
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
			"key2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_ListBySubscriptionId.json
func ExampleActivityLogAlertsClient_NewListBySubscriptionIDPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewActivityLogAlertsClient("187f412d-1758-44d9-b052-169e2564721d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListBySubscriptionIDPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_ListByResourceGroupName.json
func ExampleActivityLogAlertsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewActivityLogAlertsClient("187f412d-1758-44d9-b052-169e2564721d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("MyResourceGroup", nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}