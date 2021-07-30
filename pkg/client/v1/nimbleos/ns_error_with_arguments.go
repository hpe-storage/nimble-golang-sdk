// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsErrorWithArguments - Non-negative integer in range [0,9000].
// Export NsErrorWithArgumentsFields for advance operations like search filter etc.
var NsErrorWithArgumentsFields *NsErrorWithArgumentsStringFields

func init() {
	Codefield := "code"
	Severityfield := "severity"
	Textfield := "text"

	NsErrorWithArgumentsFields = &NsErrorWithArgumentsStringFields{
		Code:     &Codefield,
		Severity: &Severityfield,
		Text:     &Textfield,
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
