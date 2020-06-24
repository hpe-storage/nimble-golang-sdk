// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsReplPairListWithObjectType - Replicated objects of the specified type.
// Export NsReplPairListWithObjectTypeFields for advance operations like search filter etc.
var NsReplPairListWithObjectTypeFields *NsReplPairListWithObjectType

func init(){
		
	NsReplPairListWithObjectTypeFields= &NsReplPairListWithObjectType{
		
	}
}

type NsReplPairListWithObjectType struct {
	// ObjType - Type of the replicated object.
   	ObjType *NsObjectType `json:"obj_type,omitempty"`
	// ReplList - List of replicated objects of this type.
   	ReplList []*NsReplPair `json:"repl_list,omitempty"`
}
