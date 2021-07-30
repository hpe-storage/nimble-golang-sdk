// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsHcfResult - Results from health check of a single element.
// Export NsHcfResultFields for advance operations like search filter etc.
var NsHcfResultFields *NsHcfResultStringFields

func init() {
	ElementNamefield := "element_name"
	ErrorListfield := "error_list"
	Messagesfield := "messages"

	NsHcfResultFields = &NsHcfResultStringFields{
		ElementName: &ElementNamefield,
		ErrorList:   &ErrorListfield,
		Messages:    &Messagesfield,
	}
}

type NsHcfResult struct {
	// ElementName - Name of the element.
	ElementName *string `json:"element_name,omitempty"`
	// ErrorList - List of health check errors for this element.
	ErrorList []*string `json:"error_list,omitempty"`
	// Messages - A list of error messages.
	Messages []*NsErrorWithArguments `json:"messages,omitempty"`
}

// Struct for NsHcfResultFields
type NsHcfResultStringFields struct {
	ElementName *string
	ErrorList   *string
	Messages    *string
}
