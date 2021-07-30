// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsObjectLimitFields provides field names to use in filter parameters, for example.
var NsObjectLimitFields *NsObjectLimitFieldHandles

func init() {
	fieldObjType := "obj_type"
	fieldObjLimit := "obj_limit"
	fieldObjNum := "obj_num"

	NsObjectLimitFields = &NsObjectLimitFieldHandles{
		ObjType:  &fieldObjType,
		ObjLimit: &fieldObjLimit,
		ObjNum:   &fieldObjNum,
	}
}

// NsObjectLimit - Limits (Max total number of objects) for object of a given type.
type NsObjectLimit struct {
	// ObjType - Type of the object.
	ObjType *NsObjectType `json:"obj_type,omitempty"`
	// ObjLimit - Limit of the objects.
	ObjLimit *int64 `json:"obj_limit,omitempty"`
	// ObjNum - Number of objects after group merge.
	ObjNum *int64 `json:"obj_num,omitempty"`
}

// NsObjectLimitFieldHandles provides a string representation for each AccessControlRecord field.
type NsObjectLimitFieldHandles struct {
	ObjType  *string
	ObjLimit *string
	ObjNum   *string
}
