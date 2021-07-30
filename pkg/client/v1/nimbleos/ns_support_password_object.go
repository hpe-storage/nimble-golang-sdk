// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSupportPasswordObject - Support password blob for a user.

// Export NsSupportPasswordObjectFields provides field names to use in filter parameters, for example.
var NsSupportPasswordObjectFields *NsSupportPasswordObjectStringFields

func init() {
	fieldUsername := "username"
	fieldBlob := "blob"

	NsSupportPasswordObjectFields = &NsSupportPasswordObjectStringFields{
		Username: &fieldUsername,
		Blob:     &fieldBlob,
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
