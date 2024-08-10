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
	form.Add("client_id", c.Request.ClientID)
	form.Add("client_secret", c.Request.ClientSecret)
	form.Add("grant_type", c.Request.GrantType)
	form.Add("scope", c.Request.Scope)

	response := &Response{}

	err := requests.
		URL(c.Url).
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
	form.Add("client_id", c.Request.ClientID)
	form.Add("client_secret", c.Request.ClientSecret)
	form.Add("grant_type", c.Request.GrantType)
	form.Add("scope", c.Request.Scope)

	response := &Response{}

	err := requests.
		URL(c.Url).
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
