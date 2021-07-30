// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsResponse - May contain anything that any REST API response can contain.

// Export NsResponseFields provides field names to use in filter parameters, for example.
var NsResponseFields *NsResponseStringFields

func init() {
	fieldData := "data"
	fieldMessages := "messages"

	NsResponseFields = &NsResponseStringFields{
		Data:     &fieldData,
		Messages: &fieldMessages,
	}
}

type NsResponse struct {
	// Data - Response data.
	Data *NsObject `json:"data,omitempty"`
	// Messages - A list of error messages.
	Messages []*NsErrorWithArguments `json:"messages,omitempty"`
}

// Struct for NsResponseFields
type NsResponseStringFields struct {
	Data     *string
	Messages *string
}
