package notification

import "context"

// Notifications retrieves a list of notification.
//
// ctx: The context.Context to use for the request.
// filter: criteria to filter notifications.
// Returns:
// - *Result: The retrieved notification.
// - error: An error if the retrieval failed.
type Provider interface {
	Notifications(ctx context.Context, filter *Filter) ([]Result, error)
}
