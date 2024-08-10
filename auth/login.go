package auth

import (
	"context"
	"fmt"
	"net/url"

	"github.com/carlmjohnson/requests"
)

// LoginAsIntermediary implements Authenticator.
func (c *Client) LoginAsIntermediary(ctx context.Context, onBehalfOf string) (*Response, error) {
	form := url.Values{}
	form.Add("client_id", c.Conf.ClientID)
	form.Add("client_secret", c.Conf.ClientSecret)
	form.Add("grant_type", "client_credentials")
	form.Add("scope", "InvoicingAPI")

	response := &Response{}

	err := requests.
		URL(c.Conf.Url).
		Method("POST").
		Path("/connect/token").
		Header("onbehalfof", onBehalfOf).
		BodyForm(form).
		ToJSON(response).
		Fetch(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to login: %w", err)
	}

	c.Token = response.AccessToken

	return response, nil
}

// LoginAsTaxPayer implements Authenticator.
func (c *Client) LoginAsTaxPayer(ctx context.Context) (*Response, error) {
	form := url.Values{}
	form.Add("client_id", c.Conf.ClientID)
	form.Add("client_secret", c.Conf.ClientSecret)
	form.Add("grant_type", "client_credentials")
	form.Add("scope", "InvoicingAPI")

	response := &Response{}

	err := requests.
		URL(c.Conf.Url).
		Method("POST").
		Path("/connect/token").
		BodyForm(form).
		ToJSON(response).
		Fetch(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to login: %w", err)
	}

	c.Token = response.AccessToken

	return response, nil
}
