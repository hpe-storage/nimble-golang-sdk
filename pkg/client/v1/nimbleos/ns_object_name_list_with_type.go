// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsObjectNameListWithType - List of objects of a given type.

// Export NsObjectNameListWithTypeFields provides field names to use in filter parameters, for example.
var NsObjectNameListWithTypeFields *NsObjectNameListWithTypeStringFields

func init() {
	fieldObjType := "obj_type"
	fieldObjNameList := "obj_name_list"

	NsObjectNameListWithTypeFields = &NsObjectNameListWithTypeStringFields{
		ObjType:     &fieldObjType,
		ObjNameList: &fieldObjNameList,
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
