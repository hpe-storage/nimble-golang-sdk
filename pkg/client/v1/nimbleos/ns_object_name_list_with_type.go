// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsObjectNameListWithType - List of objects of a given type.
// Export NsObjectNameListWithTypeFields for advance operations like search filter etc.
var NsObjectNameListWithTypeFields *NsObjectNameListWithTypeStringFields

func init() {
	ObjTypefield := "obj_type"
	ObjNameListfield := "obj_name_list"

	NsObjectNameListWithTypeFields = &NsObjectNameListWithTypeStringFields{
		ObjType:     &ObjTypefield,
		ObjNameList: &ObjNameListfield,
	}
}

type NsObjectNameListWithType struct {
	// ObjType - Type of the object.
	ObjType *NsObjectType `json:"obj_type,omitempty"`
	// ObjNameList - List of object names.
	ObjNameList []*string `json:"obj_name_list,omitempty"`
}

// Struct for NsObjectNameListWithTypeFields
type NsObjectNameListWithTypeStringFields struct {
	ObjType     *string
	ObjNameList *string
}
