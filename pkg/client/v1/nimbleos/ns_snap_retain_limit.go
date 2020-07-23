// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSnapRetainLimit - Limit for scheduled snapshot retainment params.
// Export NsSnapRetainLimitFields for advance operations like search filter etc.
var NsSnapRetainLimitFields *NsSnapRetainLimit

func init() {

	NsSnapRetainLimitFields = &NsSnapRetainLimit{}
}

type NsSnapRetainLimit struct {
	// ObjType - Type of the object.
	ObjType *NsObjectType `json:"obj_type,omitempty"`
	// RetainLimit - Limit of the objects.
	RetainLimit *int64 `json:"retain_limit,omitempty"`
	// RetainNum - Number of objects after group merge.
	RetainNum *int64 `json:"retain_num,omitempty"`
}
