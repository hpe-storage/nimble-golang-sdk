// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsObjectNameListWithTypeFields provides field names to use in filter parameters, for example.
var NsObjectNameListWithTypeFields *NsObjectNameListWithTypeFieldHandles

func init() {
	fieldObjType := "obj_type"
	fieldObjNameList := "obj_name_list"

	NsObjectNameListWithTypeFields = &NsObjectNameListWithTypeFieldHandles{
		ObjType:     &fieldObjType,
		ObjNameList: &fieldObjNameList,
	}
}

// NsObjectNameListWithType - List of objects of a given type.
type NsObjectNameListWithType struct {
	// ObjType - Type of the object.
	ObjType *NsObjectType `json:"obj_type,omitempty"`
	// ObjNameList - List of object names.
	ObjNameList []*string `json:"obj_name_list,omitempty"`
}

// NsObjectNameListWithTypeFieldHandles provides a string representation for each AccessControlRecord field.
type NsObjectNameListWithTypeFieldHandles struct {
	ObjType     *string
	ObjNameList *string
}
