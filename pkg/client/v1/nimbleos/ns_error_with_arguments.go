// Copyright 2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsErrorWithArguments - Non-negative integer in range [0,9000].
// Export NsErrorWithArgumentsFields for advance operations like search filter etc.
var NsErrorWithArgumentsFields *NsErrorWithArguments

func init() {
	Codefield := "code"
	Textfield := "text"

	NsErrorWithArgumentsFields = &NsErrorWithArguments{
		Code: &Codefield,
		Text: &Textfield,
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
