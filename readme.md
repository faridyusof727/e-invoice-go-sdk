# E-Invoice LHDN Go Wrapper

## Overview

The E-Invoice LHDN Go Wrapper is an open-source project designed to simplify the integration with the LHDN e-invoicing system. This project aims to provide a straightforward and efficient way to interact with the [LHDN e-invoice API](https://sdk.myinvois.hasil.gov.my/api/).

## Usage

### Authentication

In E-Invoice, there are two types of auth

* Login as Taxpayer System
* Login as Intermediary System

#### Login as Taxpayer System

```go
import (
    "context"
    
    "github.com/faridyusof727/e-invoice-go-sdk/auth"
    "github.com/faridyusof727/e-invoice-go-sdk/configs"
)

// Configuration 
config := &configs.Config{
    ClientID:     "<your-client-id-here>",
    ClientSecret: "<your-client-secret-here>",
}

// If the integration is to Sandbox environment
config.SandBox()

// Otherwise
config.Prod()

authClient := auth.New(config)
data, err := authClient.LoginAsTaxPayer(context.TODO())

if err != nil {
    panic(err)
}
```

#### Login as Intermediary System

> In Malaysia, the Tax Identification Number (TIN), also known as the Income Tax Number, is a unique identifier assigned to individuals and entities who are registered taxpayers with the Inland Revenue Board of Malaysia (LHDN).

This function is used to authenticate the ERP system associated with an intermediary that is representing a taxpayer (acting on behalf of a specific taxpayer) calling and issue access token.

```go
tin := "<your-tin-num>"
data, err := authClient.LoginAsIntermediary(context.TODO(), tin)

if err != nil {
    panic(err)
}
```

#### Getting the Access Token

Although, there's no need to use this function, but for whatever reasons, if you need to get the access token on demand, you can always use:

```go
token := authClient.AccessToken()
```

### Platform: Document Lookups

#### Get All Document Types

This allows taxpayer's systems to retrieve list of document types published by the MyInvois System. There are multiple types of documents supported by the solution and this allows taxpayer ERP systems to retrieve their definitions through this function call.

```go
import (
    "context"
    
    "github.com/faridyusof727/e-invoice-go-sdk/document_type"
)

// `authClient` is from previous initialization
doc, err := document_type.New(authClient)

if err != nil {
    panic(err)
}

d, err := doc.AllDocumentTypes(context.Background())
if err != nil {
    panic(err)
}
```

#### Get Document Type

This function allows taxpayer ERP system to retrieve the details of single document type that contains structure definitions of the document.

```go
documentId := 1
d, err := doc.DocumentType(context.Background(), documentId)
if err != nil {
    panic(err)
}
```

Note: This repository is still in progress. Contributors are welcomed.
