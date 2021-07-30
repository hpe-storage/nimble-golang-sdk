// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSupportPasswordObject - Support password blob for a user.
// Export NsSupportPasswordObjectFields for advance operations like search filter etc.
var NsSupportPasswordObjectFields *NsSupportPasswordObjectStringFields

func init() {
	Usernamefield := "username"
	Blobfield := "blob"

	NsSupportPasswordObjectFields = &NsSupportPasswordObjectStringFields{
		Username: &Usernamefield,
		Blob:     &Blobfield,
	}
}

type NsSupportPasswordObject struct {
	// Username - The username for the account the password blob relates to.
	Username *string `json:"username,omitempty"`
	// Blob - The ciphertext blob holding the randomly produced password.
	Blob *string `json:"blob,omitempty"`
}

// Struct for NsSupportPasswordObjectFields
type NsSupportPasswordObjectStringFields struct {
	Username *string
	Blob     *string
}
