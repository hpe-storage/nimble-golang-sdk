// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsSupportPasswordObject - Support password blob for a user.
// Export NsSupportPasswordObjectFields for advance operations like search filter etc.
var NsSupportPasswordObjectFields *NsSupportPasswordObject

func init(){
	Usernamefield:= "username"
	Blobfield:= "blob"
		
	NsSupportPasswordObjectFields= &NsSupportPasswordObject{
	Username: &Usernamefield,
	Blob: &Blobfield,
		
	}
}

type NsSupportPasswordObject struct {
	// Username - The username for the account the password blob relates to.
 	Username *string `json:"username,omitempty"`
	// Blob - The ciphertext blob holding the randomly produced password.
 	Blob *string `json:"blob,omitempty"`
}
