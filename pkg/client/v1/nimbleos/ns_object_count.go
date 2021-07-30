// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsObjectCount - Number of objects of a type in a given scope.

// Export NsObjectCountFields provides field names to use in filter parameters, for example.
var NsObjectCountFields *NsObjectCountStringFields

func init() {
	fieldScopeName := "scope_name"
	fieldObjectCount := "object_count"
	fieldMaxLimitOverride := "max_limit_override"
	fieldWarningThresholdOverride := "warning_threshold_override"

	NsObjectCountFields = &NsObjectCountStringFields{
		ScopeName:                &fieldScopeName,
		ObjectCount:              &fieldObjectCount,
		MaxLimitOverride:         &fieldMaxLimitOverride,
		WarningThresholdOverride: &fieldWarningThresholdOverride,
	}
}

type NsObjectCount struct {
	// ScopeName - Name of the scope.
	ScopeName *string `json:"scope_name,omitempty"`
	// ObjectCount - Number of objects.
	ObjectCount *int64 `json:"object_count,omitempty"`
	// MaxLimitOverride - Override max object limit for this scope.
	MaxLimitOverride *int64 `json:"max_limit_override,omitempty"`
	// WarningThresholdOverride - Override warning threshold for this scope.
	WarningThresholdOverride *int64 `json:"warning_threshold_override,omitempty"`
}

// Struct for NsObjectCountFields
type NsObjectCountStringFields struct {
	ScopeName                *string
	ObjectCount              *string
	MaxLimitOverride         *string
	WarningThresholdOverride *string
}
