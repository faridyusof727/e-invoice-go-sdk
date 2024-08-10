package document

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
)

// AllDocumentTypes implements Provider.
func (c *Client) AllDocumentTypes(ctx context.Context) ([]DocumentType, error) {
	dataResponse := &struct {
		Result []DocumentType `json:"result"`
	}{}

	authResponse, err := c.authClient.LoginAsTaxPayer(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to login: %w", err)
	}

	err = requests.
		URL(c.baseUrl).
		Method("GET").
		Path("/api/v1.0/documenttypes").
		Header("Accept", "application/json").
		Header("Authorization", fmt.Sprintf("Bearer %s", authResponse.AccessToken)).
		Header("Accept-Language", "en").
		Header("Content-Type", "application/json").
		ToJSON(dataResponse).
		Fetch(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all document types: %w", err)
	}

	return dataResponse.Result, nil
}
