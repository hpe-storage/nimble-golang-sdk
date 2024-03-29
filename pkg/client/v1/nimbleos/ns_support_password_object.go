// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSupportPasswordObjectFields provides field names to use in filter parameters, for example.
var NsSupportPasswordObjectFields *NsSupportPasswordObjectFieldHandles

func init() {
	NsSupportPasswordObjectFields = &NsSupportPasswordObjectFieldHandles{
		Username: "username",
		Blob:     "blob",
	}
}

// NsSupportPasswordObject - Support password blob for a user.
type NsSupportPasswordObject struct {
	// Username - The username for the account the password blob relates to.
	Username *string `json:"username,omitempty"`
	// Blob - The ciphertext blob holding the randomly produced password.
	Blob *string `json:"blob,omitempty"`
}

// NsSupportPasswordObjectFieldHandles provides a string representation for each NsSupportPasswordObject field.
type NsSupportPasswordObjectFieldHandles struct {
	Username string
	Blob     string
}
