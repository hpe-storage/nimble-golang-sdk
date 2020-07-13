// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsObjectCount - Number of objects of a type in a given scope.
// Export NsObjectCountFields for advance operations like search filter etc.
var NsObjectCountFields *NsObjectCount

func init() {

	NsObjectCountFields = &NsObjectCount{
		ScopeName: "scope_name",
	}
}

type NsObjectCount struct {
	// ScopeName - Name of the scope.
	ScopeName string `json:"scope_name,omitempty"`
	// ObjectCount - Number of objects.
	ObjectCount int64 `json:"object_count,omitempty"`
	// MaxLimitOverride - Override max object limit for this scope.
	MaxLimitOverride int64 `json:"max_limit_override,omitempty"`
	// WarningThresholdOverride - Override warning threshold for this scope.
	WarningThresholdOverride int64 `json:"warning_threshold_override,omitempty"`
}

// sdk internal struct
type nsObjectCount struct {
	// ScopeName - Name of the scope.
	ScopeName *string `json:"scope_name,omitempty"`
	// ObjectCount - Number of objects.
	ObjectCount *int64 `json:"object_count,omitempty"`
	// MaxLimitOverride - Override max object limit for this scope.
	MaxLimitOverride *int64 `json:"max_limit_override,omitempty"`
	// WarningThresholdOverride - Override warning threshold for this scope.
	WarningThresholdOverride *int64 `json:"warning_threshold_override,omitempty"`
}

// EncodeNsObjectCount - Transform NsObjectCount to nsObjectCount type
func EncodeNsObjectCount(request interface{}) (*nsObjectCount, error) {
	reqNsObjectCount := request.(*NsObjectCount)
	byte, err := json.Marshal(reqNsObjectCount)
	objPtr := &nsObjectCount{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsObjectCount Transform nsObjectCount to NsObjectCount type
func DecodeNsObjectCount(request interface{}) (*NsObjectCount, error) {
	reqNsObjectCount := request.(*nsObjectCount)
	byte, err := json.Marshal(reqNsObjectCount)
	obj := &NsObjectCount{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
