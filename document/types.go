package document

import "time"

type DocumentType struct {
	ID                   int64                 `json:"id"`
	InvoiceTypeCode      int64                 `json:"invoiceTypeCode"`
	Description          string                `json:"description"`
	ActiveFrom           time.Time             `json:"activeFrom"`
	ActiveTo             *time.Time            `json:"activeTo"`
	DocumentTypeVersions []DocumentTypeVersion `json:"documentTypeVersions"`
}

type DocumentTypeVersion struct {
	ID            int64      `json:"id"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	ActiveFrom    time.Time  `json:"activeFrom"`
	ActiveTo      *time.Time `json:"activeTo"`
	VersionNumber float64    `json:"versionNumber"`
	Status        string     `json:"status"`
}
