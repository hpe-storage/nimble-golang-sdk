// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsReplPairListWithObjectType - Replicated objects of the specified type.

// Export NsReplPairListWithObjectTypeFields provides field names to use in filter parameters, for example.
var NsReplPairListWithObjectTypeFields *NsReplPairListWithObjectTypeStringFields

func init() {
	fieldObjType := "obj_type"
	fieldReplList := "repl_list"

	NsReplPairListWithObjectTypeFields = &NsReplPairListWithObjectTypeStringFields{
		ObjType:  &fieldObjType,
		ReplList: &fieldReplList,
	}
}

type NsReplPairListWithObjectType struct {
	// ObjType - Type of the replicated object.
	ObjType *NsObjectType `json:"obj_type,omitempty"`
	// ReplList - List of replicated objects of this type.
	ReplList []*NsReplPair `json:"repl_list,omitempty"`
}

// Struct for NsReplPairListWithObjectTypeFields
type NsReplPairListWithObjectTypeStringFields struct {
	ObjType  *string
	ReplList *string
}
