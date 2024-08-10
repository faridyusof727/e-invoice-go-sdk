package auth

import (
	"github.com/faridyusof727/e-invoice-go-sdk/configs"
)

type Client struct {
	Conf  *configs.Config
	Token string
}

func New(c *configs.Config) Authenticator {
	if c == nil {
		c = &configs.Config{}
		c.Prod()
	}

	return &Client{
		Conf: c,
	}
}

// AccessToken implements Authenticator.
func (c *Client) AccessToken() string {
	return c.Token
}

// Config implements Authenticator.
func (c *Client) Config() *configs.Config {
	return c.Conf
}
