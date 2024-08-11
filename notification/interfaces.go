package notification

import "context"

type Provider interface {
	Notifications(ctx context.Context, filter *Filter) ([]Result, error)
}