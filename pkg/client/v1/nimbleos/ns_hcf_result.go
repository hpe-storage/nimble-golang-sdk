// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsHcfResultFields provides field names to use in filter parameters, for example.
var NsHcfResultFields *NsHcfResultFieldHandles

func init() {
	fieldElementName := "element_name"
	fieldErrorList := "error_list"
	fieldMessages := "messages"

	NsHcfResultFields = &NsHcfResultFieldHandles{
		ElementName: &fieldElementName,
		ErrorList:   &fieldErrorList,
		Messages:    &fieldMessages,
	}
}

// NsHcfResult - Results from health check of a single element.
type NsHcfResult struct {
	// ElementName - Name of the element.
	ElementName *string `json:"element_name,omitempty"`
	// ErrorList - List of health check errors for this element.
	ErrorList []*string `json:"error_list,omitempty"`
	// Messages - A list of error messages.
	Messages []*NsErrorWithArguments `json:"messages,omitempty"`
}

// NsHcfResultFieldHandles provides a string representation for each AccessControlRecord field.
type NsHcfResultFieldHandles struct {
	ElementName *string
	ErrorList   *string
	Messages    *string
}
