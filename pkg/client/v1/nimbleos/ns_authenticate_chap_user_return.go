// Copyright 2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAuthenticateChapUserReturn - Response of chap user authentication.
// Export NsAuthenticateChapUserReturnFields for advance operations like search filter etc.
var NsAuthenticateChapUserReturnFields *NsAuthenticateChapUserReturn

func init() {

	NsAuthenticateChapUserReturnFields = &NsAuthenticateChapUserReturn{}
}

type NsAuthenticateChapUserReturn struct {
	// IsValid - Chap user secret is valid.
	IsValid *bool `json:"is_valid,omitempty"`
}
