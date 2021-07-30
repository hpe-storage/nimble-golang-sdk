// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSnapRetainLimit - Limit for scheduled snapshot retainment params.

// Export NsSnapRetainLimitFields provides field names to use in filter parameters, for example.
var NsSnapRetainLimitFields *NsSnapRetainLimitStringFields

func init() {
	fieldObjType := "obj_type"
	fieldRetainLimit := "retain_limit"
	fieldRetainNum := "retain_num"

	NsSnapRetainLimitFields = &NsSnapRetainLimitStringFields{
		ObjType:     &fieldObjType,
		RetainLimit: &fieldRetainLimit,
		RetainNum:   &fieldRetainNum,
	}
}

type NsSnapRetainLimit struct {
	// ObjType - Type of the object.
	ObjType *NsObjectType `json:"obj_type,omitempty"`
	// RetainLimit - Limit of the objects.
	RetainLimit *int64 `json:"retain_limit,omitempty"`
	// RetainNum - Number of objects after group merge.
	RetainNum *int64 `json:"retain_num,omitempty"`
}

// Struct for NsSnapRetainLimitFields
type NsSnapRetainLimitStringFields struct {
	ObjType     *string
	RetainLimit *string
	RetainNum   *string
}
