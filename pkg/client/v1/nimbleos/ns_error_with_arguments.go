// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsErrorWithArgumentsFields provides field names to use in filter parameters, for example.
var NsErrorWithArgumentsFields *NsErrorWithArgumentsFieldHandles

func init() {
	NsErrorWithArgumentsFields = &NsErrorWithArgumentsFieldHandles{
		Code:     "code",
		Severity: "severity",
		Text:     "text",
	}
}

// NsErrorWithArguments - Non-negative integer in range [0,9000].
type NsErrorWithArguments struct {
	// Code - Error code.
	Code *string `json:"code,omitempty"`
	// Severity - Severity of the error.
	Severity *NsApiSeverityLevel `json:"severity,omitempty"`
	// Text - Full error message with argument populated.
	Text *string `json:"text,omitempty"`
}

// NsErrorWithArgumentsFieldHandles provides a string representation for each AccessControlRecord field.
type NsErrorWithArgumentsFieldHandles struct {
	Code     string
	Severity string
	Text     string
}
