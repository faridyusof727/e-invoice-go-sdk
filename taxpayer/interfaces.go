package taxpayer

import "context"

type Client interface {
	// Client represents an interface for validating taxpayer information.
	// It provides a method to validate a taxpayer's identification details.
	//
	// Validate validates the taxpayer's identification details based on the provided
	// taxpayer identification number (TIN), identification type, and identification number.
	//
	// Parameters:
	//   - ctx: The context for the validation request, which can be used to control
	//          cancellation and timeouts.
	//   - tin: The taxpayer identification number to be validated.
	//   - idType: The type of identification being used (e.g., passport, national ID).
	//   - idNumber: The identification number corresponding to the specified ID type.
	//
	// Returns:
	//   - error: An error if the validation fails, or nil if the validation is successful.
	Validate(ctx context.Context, tin string, idType IDType, idNumber string) error
}
