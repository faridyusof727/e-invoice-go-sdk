package document

import (
	"github.com/faridyusof727/e-invoice-go-sdk/auth"
	"github.com/faridyusof727/e-invoice-go-sdk/configs"
)

type Client struct {
	authClient auth.Authenticator
	baseUrl    string
}

func New(authClient auth.Authenticator, o ...func(o *configs.Options)) Provider {
	opts := &configs.Options{}
	opts.Url = "https://api.myinvois.hasil.gov.my"

	for _, f := range o {
		f(opts)
	}

	return &Client{
		authClient: authClient,
		baseUrl:    opts.Url,
	}
}
