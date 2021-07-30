// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// ObjectLimit - List the maximum limits and warning thresholds for number of objects in the storage group.
// Export ObjectLimitFields for advance operations like search filter etc.
var ObjectLimitFields *ObjectLimitStringFields

func init() {
	IDfield := "id"
	ObjectTypefield := "object_type"
	ObjectTypeNamefield := "object_type_name"
	ScopeTypefield := "scope_type"
	ScopeTypeNamefield := "scope_type_name"
	WarningThresholdfield := "warning_threshold"
	MaxLimitfield := "max_limit"
	ObjectCountsfield := "object_counts"

	ObjectLimitFields = &ObjectLimitStringFields{
		ID:               &IDfield,
		ObjectType:       &ObjectTypefield,
		ObjectTypeName:   &ObjectTypeNamefield,
		ScopeType:        &ScopeTypefield,
		ScopeTypeName:    &ScopeTypeNamefield,
		WarningThreshold: &WarningThresholdfield,
		MaxLimit:         &MaxLimitfield,
		ObjectCounts:     &ObjectCountsfield,
	}
}

type ObjectLimit struct {
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

// Struct for ObjectLimitFields
type ObjectLimitStringFields struct {
	ID               *string
	ObjectType       *string
	ObjectTypeName   *string
	ScopeType        *string
	ScopeTypeName    *string
	WarningThreshold *string
	MaxLimit         *string
	ObjectCounts     *string
}
