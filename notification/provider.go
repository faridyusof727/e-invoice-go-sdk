package notification

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/carlmjohnson/requests"
	"github.com/faridyusof727/e-invoice-go-sdk/constants"
)

// Notifications implements Provider.
func (c *Client) Notifications(ctx context.Context, filter *Filter) ([]Result, error) {

	dataResponse := &struct {
		Result []Result `json:"result"`
	}{}

	req := requests.
		URL(c.authClient.Config().Url).
		Method("GET").
		Path("/api/v1.0/notifications/taxpayer").
		Header("Accept", constants.HttpHeaderContentTypeJson).
		Header("Authorization", fmt.Sprintf("Bearer %s", c.authClient.AccessToken())).
		Header("Accept-Language", "en").
		Header("Content-Type", constants.HttpHeaderContentTypeJson)

	if filter != nil {
		if filter.DateFrom != nil {
			req.Param("dateFrom", filter.DateFrom.UTC().Format(time.RFC3339))
		}
		if filter.DateTo != nil {
			req.Param("dateTo", filter.DateTo.UTC().Format(time.RFC3339))
		}
		if filter.Type != nil {
			req.Param("type", *filter.Type)
		}
		if filter.Language != nil {
			req.Param("language", *filter.Language)
		}
		if filter.Status != nil {
			req.Param("status", *filter.Status)
		}
		if filter.Channel != nil {
			req.Param("channel", *filter.Channel)
		}
		if filter.PageNo != nil {
			req.Param("pageNo", strconv.Itoa(*filter.PageNo))
		}
		if filter.PageSize != nil {
			req.Param("pageSize", strconv.Itoa(*filter.PageSize))
		}
	}

	err := req.ToJSON(&dataResponse).Fetch(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get notification: %w", err)
	}

	return dataResponse.Result, nil
}
