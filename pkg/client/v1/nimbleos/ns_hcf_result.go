// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsHcfResult - Results from health check of a single element.
// Export NsHcfResultFields for advance operations like search filter etc.
var NsHcfResultFields *NsHcfResult

func init() {
	ElementNamefield := "element_name"

	NsHcfResultFields = &NsHcfResult{
		ElementName: &ElementNamefield,
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
