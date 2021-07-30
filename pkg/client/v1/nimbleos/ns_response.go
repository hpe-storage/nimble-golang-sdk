// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsResponse - May contain anything that any REST API response can contain.
// Export NsResponseFields for advance operations like search filter etc.
var NsResponseFields *NsResponseStringFields

func init() {
	Datafield := "data"
	Messagesfield := "messages"

	NsResponseFields = &NsResponseStringFields{
		Data:     &Datafield,
		Messages: &Messagesfield,
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
