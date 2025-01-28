package taxpayer

import (
	"context"
	"errors"
	"fmt"

	"github.com/carlmjohnson/requests"
	"github.com/faridyusof727/e-invoice-go-sdk/auth"
	"github.com/faridyusof727/e-invoice-go-sdk/constants"
)

type ClientImpl struct {
	authClient auth.Authenticator
}

// Validate implements Client.
func (c *ClientImpl) Validate(ctx context.Context, tin string, idType IDType, idNumber string) error {
	err := requests.
		URL(c.authClient.Config().Url).
		Method("GET").
		Path(fmt.Sprintf("/api/v1.0/taxpayer/validate/%s", tin)).
		Param("idType", string(idType)).
		Param("idValue", idNumber).
		Header("Accept", constants.HttpHeaderContentTypeJson).
		Header("Authorization", fmt.Sprintf("Bearer %s", c.authClient.AccessToken())).
		Header("Accept-Language", "en").
		Header("Content-Type", constants.HttpHeaderContentTypeJson).
		CheckStatus(200).
		Fetch(ctx)
	if err != nil {
		return fmt.Errorf("failed to validate taxpayer: %w", err)
	}

	return nil
}

func NewClient(authClient auth.Authenticator) (Client, error) {
	if authClient.Config() == nil {
		return nil, errors.New("auth client config is nil")
	}

	return &ClientImpl{
		authClient: authClient,
	}, nil
}
