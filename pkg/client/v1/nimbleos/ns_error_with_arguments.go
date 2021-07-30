// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsErrorWithArguments - Non-negative integer in range [0,9000].

// Export NsErrorWithArgumentsFields provides field names to use in filter parameters, for example.
var NsErrorWithArgumentsFields *NsErrorWithArgumentsStringFields

func init() {
	fieldCode := "code"
	fieldSeverity := "severity"
	fieldText := "text"

	NsErrorWithArgumentsFields = &NsErrorWithArgumentsStringFields{
		Code:     &fieldCode,
		Severity: &fieldSeverity,
		Text:     &fieldText,
	}
}

type NsErrorWithArguments struct {
	// Code - Error code.
	Code *string `json:"code,omitempty"`
	// Severity - Severity of the error.
	Severity *NsApiSeverityLevel `json:"severity,omitempty"`
	// Text - Full error message with argument populated.
	Text *string `json:"text,omitempty"`
}

// Struct for NsErrorWithArgumentsFields
type NsErrorWithArgumentsStringFields struct {
	Code     *string
	Severity *string
	Text     *string
}
