# E-Invoice Go SDK

A Go SDK for interacting with the E-Invoice system. This SDK provides a simple and efficient way to authenticate and interact with the E-Invoice API.

## Requirements

- Go 1.22.5 or higher

## Installation

```bash
go get github.com/faridyusof727/e-invoice-go-sdk
```

## Features

- Authentication client for E-Invoice system
- Support for both sandbox and production environments
- Easy configuration management
- Intermediary login support
- Document type client for retrieving available document types

## Usage

### Basic Setup

```go
import (
    "github.com/faridyusof727/e-invoice-go-sdk/auth"
    "github.com/faridyusof727/e-invoice-go-sdk/config"
)

// Create client configuration
clientConfig := &config.Client{
    ID:     "your-client-id",
    Secret: "your-client-secret",
}

// Create service configuration
serviceConfig := &config.Service{
    Environment: config.EnvironmentSandbox, // or config.EnvironmentProduction for production
}

// Initialize auth client
authClient := auth.New(
    auth.WithClientConfig(clientConfig),
    auth.WithServiceConfig(serviceConfig),
)
```

### Login as Intermediary

```go
auth, err := authClient.LoginAsIntermediary(context.Background(), "your-tax-id")
if err != nil {
    // Handle error
}

// Use the access token
token := auth.AccessToken
```

### Login as Tax Payer

```go
auth, err := authClient.LoginAsTaxPayer(context.Background())
if err != nil {
    // Handle error
}

// Use the access token
token := auth.AccessToken
```

### Retrieve All Document Types

```go
import (
    "context"
    "github.com/faridyusof727/e-invoice-go-sdk/platform"
    "github.com/faridyusof727/e-invoice-go-sdk/auth"
    "github.com/faridyusof727/e-invoice-go-sdk/config"
)

// Assume you have already initialized serviceConfig and authClient as shown above

// Perform login to get an auth token
auth, err := authClient.LoginAsTaxPayer(context.Background())
if err != nil {
    // Handle error
}

// Create a DocumentType client
docTypeClient, err := platform.NewDocumentType(serviceConfig, auth)
if err != nil {
    // Handle error
}

// Retrieve all document types
docTypes, err := docTypeClient.AllDocumentTypes(context.Background())
if err != nil {
    // Handle error
}

// Use the list of document types
for _, docType := range docTypes {
    // Process each docType
}
```

## Configuration

The SDK supports two environments:

- Sandbox (for testing)
- Production (for live usage)

Configure the environment using the `config.Service` struct:

```go
serviceConfig := &config.Service{
    Environment: config.EnvironmentSandbox, // or config.EnvironmentProduction
}
```

## Development

### Running Tests

```bash
go test ./...
```

## License

[License information to be added]

## Contributing

[Contributing guidelines to be added]
