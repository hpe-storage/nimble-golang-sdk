// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAuthenticateChapUserReturn - Response of chap user authentication.
// Export NsAuthenticateChapUserReturnFields for advance operations like search filter etc.
var NsAuthenticateChapUserReturnFields *NsAuthenticateChapUserReturnStringFields

func init() {
	IsValidfield := "is_valid"

	NsAuthenticateChapUserReturnFields = &NsAuthenticateChapUserReturnStringFields{
		IsValid: &IsValidfield,
	}
}

type NsAuthenticateChapUserReturn struct {
	// IsValid - Chap user secret is valid.
	IsValid *bool `json:"is_valid,omitempty"`
}

// Struct for NsAuthenticateChapUserReturnFields
type NsAuthenticateChapUserReturnStringFields struct {
	IsValid *string
}
