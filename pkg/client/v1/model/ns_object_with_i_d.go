// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsObjectWithID - An object with an ID.
// Export NsObjectWithIDFields for advance operations like search filter etc.
var NsObjectWithIDFields *NsObjectWithID

func init(){
	IDfield:= "id"
		
	NsObjectWithIDFields= &NsObjectWithID{
		ID:&IDfield,
	}
}

type NsObjectWithID struct {
	// ID - ID of object.
 	ID *string `json:"id,omitempty"`
}
