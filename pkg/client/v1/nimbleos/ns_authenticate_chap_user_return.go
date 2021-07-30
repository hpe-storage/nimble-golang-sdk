// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAuthenticateChapUserReturn - Response of chap user authentication.

// Export NsAuthenticateChapUserReturnFields provides field names to use in filter parameters, for example.
var NsAuthenticateChapUserReturnFields *NsAuthenticateChapUserReturnStringFields

func init() {
	fieldIsValid := "is_valid"

	NsAuthenticateChapUserReturnFields = &NsAuthenticateChapUserReturnStringFields{
		IsValid: &fieldIsValid,
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
