// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// ObjectLimit - List the maximum limits and warning thresholds for number of objects in the storage group.
// Export ObjectLimitFields for advance operations like search filter etc.
var ObjectLimitFields *ObjectLimit

func init() {

	ObjectLimitFields = &ObjectLimit{
		ID:             "id",
		ObjectTypeName: "object_type_name",
		ScopeTypeName:  "scope_type_name",
	}
}

type ObjectLimit struct {
	// ID - Identifier for the object limit.
	ID string `json:"id,omitempty"`
	// ObjectType - Unique ID of the object type.
	ObjectType int64 `json:"object_type,omitempty"`
	// ObjectTypeName - Name of object type.
	ObjectTypeName string `json:"object_type_name,omitempty"`
	// ScopeType - Unique ID of the scope type for the object limit.
	ScopeType int64 `json:"scope_type,omitempty"`
	// ScopeTypeName - Scope type name for the object limit.
	ScopeTypeName string `json:"scope_type_name,omitempty"`
	// WarningThreshold - Warning threshold for the object limit. A warning alert will be generated if the total number of objects in given scope crosses this threshold.
	WarningThreshold int64 `json:"warning_threshold,omitempty"`
	// MaxLimit - Total number of objects for a given scope is not allowed to cross this limit. Any user action attempting to create objects beyond this maximum limit will fail.
	MaxLimit int64 `json:"max_limit,omitempty"`
	// ObjectCounts - Current object counts for objects in given scope.
	ObjectCounts []*NsObjectCount `json:"object_counts,omitempty"`
}

// sdk internal struct
type objectLimit struct {
	// ID - Identifier for the object limit.
	ID *string `json:"id,omitempty"`
	// ObjectType - Unique ID of the object type.
	ObjectType *int64 `json:"object_type,omitempty"`
	// ObjectTypeName - Name of object type.
	ObjectTypeName *string `json:"object_type_name,omitempty"`
	// ScopeType - Unique ID of the scope type for the object limit.
	ScopeType *int64 `json:"scope_type,omitempty"`
	// ScopeTypeName - Scope type name for the object limit.
	ScopeTypeName *string `json:"scope_type_name,omitempty"`
	// WarningThreshold - Warning threshold for the object limit. A warning alert will be generated if the total number of objects in given scope crosses this threshold.
	WarningThreshold *int64 `json:"warning_threshold,omitempty"`
	// MaxLimit - Total number of objects for a given scope is not allowed to cross this limit. Any user action attempting to create objects beyond this maximum limit will fail.
	MaxLimit *int64 `json:"max_limit,omitempty"`
	// ObjectCounts - Current object counts for objects in given scope.
	ObjectCounts []*NsObjectCount `json:"object_counts,omitempty"`
}

// EncodeObjectLimit - Transform ObjectLimit to objectLimit type
func EncodeObjectLimit(request interface{}) (*objectLimit, error) {
	reqObjectLimit := request.(*ObjectLimit)
	byte, err := json.Marshal(reqObjectLimit)
	objPtr := &objectLimit{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeObjectLimit Transform objectLimit to ObjectLimit type
func DecodeObjectLimit(request interface{}) (*ObjectLimit, error) {
	reqObjectLimit := request.(*objectLimit)
	byte, err := json.Marshal(reqObjectLimit)
	obj := &ObjectLimit{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
