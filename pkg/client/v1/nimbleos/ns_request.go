// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsRequest - May contain anything that any REST API request can contain.
// Export NsRequestFields for advance operations like search filter etc.
var NsRequestFields *NsRequestStringFields

func init() {
	Datafield := "data"
	Methodfield := "method"
	Pathfield := "path"

	NsRequestFields = &NsRequestStringFields{
		Data:   &Datafield,
		Method: &Methodfield,
		Path:   &Pathfield,
	}
}

type NsRequest struct {
	// Data - Request data.
	Data *NsObject `json:"data,omitempty"`
	// Method - HTTP method.
	Method *int64 `json:"method,omitempty"`
	// Path - Path which identifies the target resource.
	Path *string `json:"path,omitempty"`
}

// Struct for NsRequestFields
type NsRequestStringFields struct {
	Data   *string
	Method *string
	Path   *string
}
