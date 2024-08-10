# E-Invoice LHDN Go Wrapper

## Overview

The E-Invoice LHDN Go Wrapper is an open-source project designed to simplify the integration with the LHDN e-invoicing system. This project aims to provide a straightforward and efficient way to interact with the [LHDN e-invoice API](https://sdk.myinvois.hasil.gov.my/api/).

## Usage

```go
import (
  "context"

  "github.com/faridyusof727/e-invoice-go-sdk/auth"
)

func AuthExample() {
  request := &auth.Request{
    ClientID:     "xxx",
    ClientSecret: "xxx",
    GrantType:    "client_credentials",
    Scope:        "InvoicingAPI",
  }

  authClient := auth.New(
    auth.WithRequest(request),
    auth.IsSandBox(true), // set to false or just remove this line for prod
  )

  data, err := authClient.LoginAsIntermediary(context.Background(), "IGXXXXXXXXXXXX")
  // or
  data, err := authClient.LoginAsTaxPayer(context.Background())

  if err != nil {
    panic(err)
  }

  // data.AccessToken <-- this would be your access token
}
```

Note: This repository is still in progress. Contributors are welcomed.
