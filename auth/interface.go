package auth

import "context"

// Client interface defines methods for authentication operations.
//
// It provides functionality to authenticate as different entities 
// in the e-invoice system and retrieve authentication tokens.
type Client interface {
	// LoginAsTaxPayer authenticates as a tax payer and returns authentication details.
	//
	// The method uses the context for cancellation and timeout control.
	// Returns Auth object containing tokens and authentication details or an error
	// if authentication fails.
	LoginAsTaxPayer(ctx context.Context) (*Auth, error)

	// LoginAsIntermediary authenticates as an intermediary entity on behalf of a tax payer.
	//
	// Parameters:
	//   - ctx: Context for cancellation and timeout control
	//   - onBehalfOf: Identifier of the tax payer on whose behalf the authentication is performed
	//
	// Returns Auth object containing tokens and authentication details or an error
	// if authentication fails.
	LoginAsIntermediary(ctx context.Context, onBehalfOf string) (*Auth, error)
}
