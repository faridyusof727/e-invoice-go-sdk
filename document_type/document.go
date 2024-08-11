package document_type

import (
	"errors"

	"github.com/faridyusof727/e-invoice-go-sdk/auth"
)

type Client struct {
	authClient auth.Authenticator
}

func New(authClient auth.Authenticator) (Provider, error) {
	if authClient.Config() == nil {
		return nil, errors.New("auth client config is nil")
	}

	return &Client{
		authClient: authClient,
	}, nil
}
