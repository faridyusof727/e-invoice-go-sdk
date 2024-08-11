package notification

import "time"

type Filter struct {
	DateFrom *time.Time
	DateTo   *time.Time
	Type     *string
	Language *string
	Status   *string
	Channel  *string
	PageNo   *int
	PageSize *int
}
