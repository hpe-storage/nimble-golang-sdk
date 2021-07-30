// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsObjectCount - Number of objects of a type in a given scope.
// Export NsObjectCountFields for advance operations like search filter etc.
var NsObjectCountFields *NsObjectCountStringFields

func init() {
	ScopeNamefield := "scope_name"
	ObjectCountfield := "object_count"
	MaxLimitOverridefield := "max_limit_override"
	WarningThresholdOverridefield := "warning_threshold_override"

	NsObjectCountFields = &NsObjectCountStringFields{
		ScopeName:                &ScopeNamefield,
		ObjectCount:              &ObjectCountfield,
		MaxLimitOverride:         &MaxLimitOverridefield,
		WarningThresholdOverride: &WarningThresholdOverridefield,
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
