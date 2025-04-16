package platform

type ErrResponse struct {
	Error string `json:"error"`
}
type Result struct {
	Result []DocumentType `json:"result"`
}

type DocumentType struct {
	ID                   int64                 `json:"id"`
	InvoiceTypeCode      string                `json:"invoiceTypeCode"`
	Name                 string                `json:"name"`
	Description          string                `json:"description"`
	ActiveFrom           string                `json:"activeFrom"`
	ActiveTo             interface{}           `json:"activeTo"`
	DocumentTypeVersions []DocumentTypeVersion `json:"documentTypeVersions"`
}

type DocumentTypeVersion struct {
	ID            int64       `json:"id"`
	Name          Description `json:"name"`
	Description   Description `json:"description"`
	ActiveFrom    string      `json:"activeFrom"`
	ActiveTo      interface{} `json:"activeTo"`
	VersionNumber string      `json:"versionNumber"`
	Status        Status      `json:"status"`
}

type Description string

const (
	Version1 Description = "Version 1"
	Version2 Description = "Version 2"
)

type Status string

const (
	Published Status = "Published"
)
