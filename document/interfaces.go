package document

import "context"

type Provider interface {
	// AllDocumentTypes retrieves all document types using the provided context.
	//
	// ctx: The context.Context to use for the request.
	//
	// Returns:
	// - []DocumentType: A slice of DocumentType representing the retrieved document types.
	// - error: An error if the retrieval failed.
	AllDocumentTypes(ctx context.Context) ([]DocumentType, error)
}
