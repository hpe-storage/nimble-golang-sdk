// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsAuthenticateChapUserReturnFields provides field names to use in filter parameters, for example.
var NsAuthenticateChapUserReturnFields *NsAuthenticateChapUserReturnFieldHandles

func init() {
	NsAuthenticateChapUserReturnFields = &NsAuthenticateChapUserReturnFieldHandles{
		IsValid: "is_valid",
	}
}

// NsAuthenticateChapUserReturn - Response of chap user authentication.
type NsAuthenticateChapUserReturn struct {
	// IsValid - Chap user secret is valid.
	IsValid *bool `json:"is_valid,omitempty"`
}

// NsAuthenticateChapUserReturnFieldHandles provides a string representation for each AccessControlRecord field.
type NsAuthenticateChapUserReturnFieldHandles struct {
	IsValid string
}
