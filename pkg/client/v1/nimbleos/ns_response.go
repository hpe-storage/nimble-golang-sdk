// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsResponseFields provides field names to use in filter parameters, for example.
var NsResponseFields *NsResponseFieldHandles

func init() {
	NsResponseFields = &NsResponseFieldHandles{
		Data:     "data",
		Messages: "messages",
	}
}

// NsResponse - May contain anything that any REST API response can contain.
type NsResponse struct {
	// Data - Response data.
	Data *NsObject `json:"data,omitempty"`
	// Messages - A list of error messages.
	Messages []*NsErrorWithArguments `json:"messages,omitempty"`
}

// NsResponseFieldHandles provides a string representation for each NsResponse field.
type NsResponseFieldHandles struct {
	Data     string
	Messages string
}
