/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

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
   
    // The username for the account the password blob relates to.
    
 	Username *string `json:"username,omitempty"`
   
    // The ciphertext blob holding the randomly produced password.
    
 	Blob *string `json:"blob,omitempty"`
}
