/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsKeyValue - Key-value pair.
// Export NsKeyValueFields for advance operations like search filter etc.
var NsKeyValueFields *NsKeyValue

func init(){
	Keyfield:= "key"
	Valuefield:= "value"
		
	NsKeyValueFields= &NsKeyValue{
		Key: &Keyfield,
		Value: &Valuefield,
		
	}
}

type NsKeyValue struct {
   
    // Identifier of key-value pair.
    
 	Key *string `json:"key,omitempty"`
   
    // Value of key-value pair.
    
 	Value *string `json:"value,omitempty"`
}
