// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsRequest - May contain anything that any REST API request can contain.

// Export NsRequestFields provides field names to use in filter parameters, for example.
var NsRequestFields *NsRequestStringFields

func init() {
	fieldData := "data"
	fieldMethod := "method"
	fieldPath := "path"

	NsRequestFields = &NsRequestStringFields{
		Data:   &fieldData,
		Method: &fieldMethod,
		Path:   &fieldPath,
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
