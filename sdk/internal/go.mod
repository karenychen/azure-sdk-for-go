module github.com/Azure/azure-sdk-for-go/sdk/internal

go 1.22.7

toolchain go1.23.1

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.17.0
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.8.2
	github.com/Azure/azure-sdk-for-go/sdk/tracing/azotel v0.4.0
	github.com/stretchr/testify v1.10.0
	golang.org/x/net v0.35.0
	golang.org/x/text v0.22.0
)

require (
	github.com/AzureAD/microsoft-authentication-library-for-go v1.3.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	go.opentelemetry.io/otel v1.33.0 // indirect
	go.opentelemetry.io/otel/trace v1.33.0 // indirect
	golang.org/x/crypto v0.33.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/Azure/azure-sdk-for-go/sdk/azcore => github.com/Azure/azure-sdk-for-go/sdk/azcore v1.17.1-0.20250219025933-3121a844ea6f

replace github.com/Azure/azure-sdk-for-go/sdk/tracing/azotel => github.com/Azure/azure-sdk-for-go/sdk/tracing/azotel v0.4.1-0.20250214223549-8a6d539f1772
