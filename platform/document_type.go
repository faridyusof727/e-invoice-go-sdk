package platform

import (
	"context"
	"fmt"

	"github.com/faridyusof727/e-invoice-go-sdk/auth"
	"github.com/faridyusof727/e-invoice-go-sdk/config"
	"github.com/imroc/req/v3"
)

type DocumentTypeImpl struct {
	serviceConfig *config.Service
	auth          *auth.Auth
}

// NewDocumentType creates a new DocumentType client using the provided auth client and service config.
func NewDocumentType(serviceConfig *config.Service, auth *auth.Auth) (DocumentTypeClient, error) {
	return &DocumentTypeImpl{
		serviceConfig: serviceConfig,
		auth:          auth,
	}, nil
}

// AllDocumentTypes fetches all document types from the platform API.
func (c *DocumentTypeImpl) AllDocumentTypes(ctx context.Context) ([]DocumentType, error) {
	baseURL := c.serviceConfig.GetBaseUrl()

	dataResponse := &Result{}
	errRes := &ErrResponse{}

	res, err := req.C().R().
		SetContext(ctx).
		SetHeaders(map[string]string{
			"Authorization":   fmt.Sprintf("Bearer %s", c.auth.AccessToken),
			"Acccept":         "application/json",
			"Content-Type":    "application/json",
			"Accept-Language": "en",
		}).
		SetSuccessResult(dataResponse).
		SetErrorResult(errRes).
		Get(fmt.Sprintf("%s/api/v1.0/documenttypes", baseURL))

	if err != nil {
		return nil, fmt.Errorf("failed to fetch document types: %w", err)
	}

	if res.IsErrorState() {
		return nil, fmt.Errorf("failed to fetch document types: %s", errRes.Error)
	}

	return dataResponse.Result, nil
}
