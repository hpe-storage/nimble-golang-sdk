// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsRequestFields provides field names to use in filter parameters, for example.
var NsRequestFields *NsRequestFieldHandles

func init() {
	NsRequestFields = &NsRequestFieldHandles{
		Data:   "data",
		Method: "method",
		Path:   "path",
	}
}

// NsRequest - May contain anything that any REST API request can contain.
type NsRequest struct {
	// Data - Request data.
	Data *NsObject `json:"data,omitempty"`
	// Method - HTTP method.
	Method *int64 `json:"method,omitempty"`
	// Path - Path which identifies the target resource.
	Path *string `json:"path,omitempty"`
}

// NsRequestFieldHandles provides a string representation for each NsRequest field.
type NsRequestFieldHandles struct {
	Data   string
	Method string
	Path   string
}
