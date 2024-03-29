// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSnapRetainLimitFields provides field names to use in filter parameters, for example.
var NsSnapRetainLimitFields *NsSnapRetainLimitFieldHandles

func init() {
	NsSnapRetainLimitFields = &NsSnapRetainLimitFieldHandles{
		ObjType:     "obj_type",
		RetainLimit: "retain_limit",
		RetainNum:   "retain_num",
	}
}

// NsSnapRetainLimit - Limit for scheduled snapshot retainment params.
type NsSnapRetainLimit struct {
	// ObjType - Type of the object.
	ObjType *NsObjectType `json:"obj_type,omitempty"`
	// RetainLimit - Limit of the objects.
	RetainLimit *int64 `json:"retain_limit,omitempty"`
	// RetainNum - Number of objects after group merge.
	RetainNum *int64 `json:"retain_num,omitempty"`
}

// NsSnapRetainLimitFieldHandles provides a string representation for each NsSnapRetainLimit field.
type NsSnapRetainLimitFieldHandles struct {
	ObjType     string
	RetainLimit string
	RetainNum   string
}
