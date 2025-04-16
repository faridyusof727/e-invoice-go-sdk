package platform

import "context"

type DocumentTypeClient interface {
	// AllDocumentTypes retrieves a list of all available document types.
	// It returns a slice of DocumentType and an error if the retrieval fails.
	AllDocumentTypes(ctx context.Context) ([]DocumentType, error)
}
