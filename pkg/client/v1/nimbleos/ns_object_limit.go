// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsObjectLimit - Limits (Max total number of objects) for object of a given type.
// Export NsObjectLimitFields for advance operations like search filter etc.
var NsObjectLimitFields *NsObjectLimitStringFields

func init() {
	ObjTypefield := "obj_type"
	ObjLimitfield := "obj_limit"
	ObjNumfield := "obj_num"

	NsObjectLimitFields = &NsObjectLimitStringFields{
		ObjType:  &ObjTypefield,
		ObjLimit: &ObjLimitfield,
		ObjNum:   &ObjNumfield,
	}
}

type NsObjectLimit struct {
	// ObjType - Type of the object.
	ObjType *NsObjectType `json:"obj_type,omitempty"`
	// ObjLimit - Limit of the objects.
	ObjLimit *int64 `json:"obj_limit,omitempty"`
	// ObjNum - Number of objects after group merge.
	ObjNum *int64 `json:"obj_num,omitempty"`
}

// Struct for NsObjectLimitFields
type NsObjectLimitStringFields struct {
	ObjType  *string
	ObjLimit *string
	ObjNum   *string
}
