module github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus

go 1.22.7

toolchain go1.23.1

retract v1.1.2 // Breaks customers in situations where close is slow/infinite.

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.17.0
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.8.0
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.10.0
	github.com/Azure/go-amqp v1.3.0
)

require (
	// used in tests only
	github.com/joho/godotenv v1.5.1
	github.com/stretchr/testify v1.10.0
)

require (
	github.com/coder/websocket v1.8.12
	github.com/golang/mock v1.6.0
)

require (
	github.com/Azure/azure-sdk-for-go/sdk/tracing/azotel v0.4.0 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.3.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	go.opentelemetry.io/otel v1.34.0 // indirect
	go.opentelemetry.io/otel/trace v1.34.0 // indirect
	golang.org/x/crypto v0.32.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/Azure/azure-sdk-for-go/sdk/azcore => github.com/Azure/azure-sdk-for-go/sdk/azcore v1.17.1-0.20250213005623-521dc9bd0933

replace github.com/Azure/azure-sdk-for-go/sdk/internal => /Users/karenchen/go/src/azure-sdk-for-go/sdk/internal
