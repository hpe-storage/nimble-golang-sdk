// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsObjectIDKV - A key value pair containing an object ID as a value.
// Export NsObjectIDKVFields for advance operations like search filter etc.
var NsObjectIDKVFields *NsObjectIDKV

func init(){
	Keyfield:= "key"
	IDfield:= "id"
		
	NsObjectIDKVFields= &NsObjectIDKV{
		Key:&Keyfield,
		ID: &IDfield,
	}
}

type NsObjectIDKV struct {
	// Key - Identifier of key-value pair.
 	Key *string `json:"key,omitempty"`
	// ID - An object ID.
 	ID *string `json:"id,omitempty"`
}
