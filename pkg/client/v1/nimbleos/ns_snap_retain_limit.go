// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

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
	RetainLimit int64 `json:"retain_limit,omitempty"`
	// RetainNum - Number of objects after group merge.
	RetainNum int64 `json:"retain_num,omitempty"`
}

// sdk internal struct
type nsSnapRetainLimit struct {
	// ObjType - Type of the object.
	ObjType *NsObjectType `json:"obj_type,omitempty"`
	// RetainLimit - Limit of the objects.
	RetainLimit *int64 `json:"retain_limit,omitempty"`
	// RetainNum - Number of objects after group merge.
	RetainNum *int64 `json:"retain_num,omitempty"`
}

// EncodeNsSnapRetainLimit - Transform NsSnapRetainLimit to nsSnapRetainLimit type
func EncodeNsSnapRetainLimit(request interface{}) (*nsSnapRetainLimit, error) {
	reqNsSnapRetainLimit := request.(*NsSnapRetainLimit)
	byte, err := json.Marshal(reqNsSnapRetainLimit)
	objPtr := &nsSnapRetainLimit{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsSnapRetainLimit Transform nsSnapRetainLimit to NsSnapRetainLimit type
func DecodeNsSnapRetainLimit(request interface{}) (*NsSnapRetainLimit, error) {
	reqNsSnapRetainLimit := request.(*nsSnapRetainLimit)
	byte, err := json.Marshal(reqNsSnapRetainLimit)
	obj := &NsSnapRetainLimit{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
