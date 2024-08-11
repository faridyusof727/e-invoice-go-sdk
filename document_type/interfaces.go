package document_type

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

	// DocumentType retrieves a document type by its ID.
	//
	// ctx: The context.Context to use for the request.
	// id: The ID of the document type to retrieve.
	// Returns:
	// - *DocumentType: The retrieved document type.
	// - error: An error if the retrieval failed.
	DocumentType(ctx context.Context, id int64) (*DocumentType, error)
}
