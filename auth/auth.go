package auth

import "github.com/faridyusof727/e-invoice-go-sdk/configs"

type Client struct {
	Request *configs.Request
	Url     string
}

// New creates a new Authenticator instance.
//
// It takes a variable number of functions that modify the Options.
// Returns an Authenticator instance.
func New(o ...func(o *configs.Options)) Authenticator {
	opts := &configs.Options{}
	opts.Url = "https://api.myinvois.hasil.gov.my"

	for _, f := range o {
		f(opts)
	}

	return &Client{
		Request: opts.Request,
		Url:     opts.Url,
	}
}
