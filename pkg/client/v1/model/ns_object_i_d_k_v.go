/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsObjectIDKV - A key value pair containing an object ID as a value.
// Export NsObjectIDKVFields for advance operations like search filter etc.
var NsObjectIDKVFields *NsObjectIDKV

func init(){
	Keyfield:= "key"
	IDfield:= "id"
		
	NsObjectIDKVFields= &NsObjectIDKV{
		Key: &Keyfield,
		ID: &IDfield,
		
	}
}

type NsObjectIDKV struct {
   
    // Identifier of key-value pair.
    
 	Key *string `json:"key,omitempty"`
   
    // An object ID.
    
 	ID *string `json:"id,omitempty"`
}
