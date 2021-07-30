// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsReplPairListWithObjectTypeFields provides field names to use in filter parameters, for example.
var NsReplPairListWithObjectTypeFields *NsReplPairListWithObjectTypeFieldHandles

func init() {
	fieldObjType := "obj_type"
	fieldReplList := "repl_list"

	NsReplPairListWithObjectTypeFields = &NsReplPairListWithObjectTypeFieldHandles{
		ObjType:  &fieldObjType,
		ReplList: &fieldReplList,
	}
}

// NsReplPairListWithObjectType - Replicated objects of the specified type.
type NsReplPairListWithObjectType struct {
	// ObjType - Type of the replicated object.
	ObjType *NsObjectType `json:"obj_type,omitempty"`
	// ReplList - List of replicated objects of this type.
	ReplList []*NsReplPair `json:"repl_list,omitempty"`
}

// NsReplPairListWithObjectTypeFieldHandles provides a string representation for each AccessControlRecord field.
type NsReplPairListWithObjectTypeFieldHandles struct {
	ObjType  *string
	ReplList *string
}
