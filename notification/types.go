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

type DeliveryAttempt struct {
	AttemptDateTime time.Time `json:"attemptDateTime"`
	Status          string    `json:"status"`
	StatusDetails   string    `json:"statusDetails"`
}

type Metadata struct {
	TotalPages int `json:"totalPages"`
	TotalCount int `json:"totalCount"`
}

type Notification struct {
	NotificationId    string            `json:"notificationId"`
	ReceivedDateTime  time.Time         `json:"receivedDateTime"`
	DeliveredDateTime time.Time         `json:"deliveredDateTime"`
	TypeId            string            `json:"typeId"`
	TypeName          string            `json:"typeName"`
	FinalMessage      string            `json:"finalMessage"`
	Channel           string            `json:"channel"`
	Address           string            `json:"address"`
	Language          string            `json:"language"`
	Status            string            `json:"status"`
	DeliveryAttempts  []DeliveryAttempt `json:"deliveryAttempts"`
}

type Result struct {
	Notifications []Notification `json:"notifications"`
	Metadata      Metadata       `json:"metadata"`
}
