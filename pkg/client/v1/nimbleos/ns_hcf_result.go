// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsHcfResult - Results from health check of a single element.

// Export NsHcfResultFields provides field names to use in filter parameters, for example.
var NsHcfResultFields *NsHcfResultStringFields

func init() {
	fieldElementName := "element_name"
	fieldErrorList := "error_list"
	fieldMessages := "messages"

	NsHcfResultFields = &NsHcfResultStringFields{
		ElementName: &fieldElementName,
		ErrorList:   &fieldErrorList,
		Messages:    &fieldMessages,
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
