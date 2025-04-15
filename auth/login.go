package auth

import (
	"context"
	"fmt"

	"github.com/imroc/req/v3"
)

// LoginAsIntermediary implements Client.
func (i *AuthImpl) LoginAsIntermediary(ctx context.Context, onBehalfOf string) (*Auth, error) {
	if i.clientConfig == nil {
		return nil, fmt.Errorf("client config is not set")
	}

	auth := &Auth{}
	errRes := &ErrResponse{}
	req := req.C().R()

	res, err := req.
		SetContext(ctx).
		SetFormData(map[string]string{
			"client_id":     i.clientConfig.ID,
			"client_secret": i.clientConfig.Secret,
			"grant_type":    "client_credentials",
			"scope":         "InvoicingAPI",
		}).
		SetHeader("onbehalfof", onBehalfOf).
		SetSuccessResult(auth).
		SetErrorResult(errRes).
		Post(fmt.Sprintf("%s/connect/token", i.serviceConfig.GetBaseUrl()))

	if err != nil {
		return nil, fmt.Errorf("failed to login as intermediary: %w", err)
	}

	if res.IsErrorState() {
		return nil, fmt.Errorf("failed to login as intermediary: %s", errRes.Error)
	}

	return auth, nil
}

// LoginAsTaxPayer implements Client.
func (i *AuthImpl) LoginAsTaxPayer(ctx context.Context) (*Auth, error) {
	if i.clientConfig == nil {
		return nil, fmt.Errorf("client config is not set")
	}

	auth := &Auth{}
	errRes := &ErrResponse{}
	req := req.C().R()

	res, err := req.
		SetContext(ctx).
		SetFormData(map[string]string{
			"client_id":     i.clientConfig.ID,
			"client_secret": i.clientConfig.Secret,
			"grant_type":    "client_credentials",
			"scope":         "InvoicingAPI",
		}).
		SetSuccessResult(auth).
		SetErrorResult(errRes).
		Post(fmt.Sprintf("%s/connect/token", i.serviceConfig.GetBaseUrl()))

	if err != nil {
		return nil, fmt.Errorf("failed to login: %w", err)
	}

	if res.IsErrorState() {
		return nil, fmt.Errorf("failed to login: %s", errRes.Error)
	}

	return auth, nil
}
