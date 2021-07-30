// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsReplPairListWithObjectType - Replicated objects of the specified type.
// Export NsReplPairListWithObjectTypeFields for advance operations like search filter etc.
var NsReplPairListWithObjectTypeFields *NsReplPairListWithObjectTypeStringFields

func init() {
	ObjTypefield := "obj_type"
	ReplListfield := "repl_list"

	NsReplPairListWithObjectTypeFields = &NsReplPairListWithObjectTypeStringFields{
		ObjType:  &ObjTypefield,
		ReplList: &ReplListfield,
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
