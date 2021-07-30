// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsObjectCountFields provides field names to use in filter parameters, for example.
var NsObjectCountFields *NsObjectCountFieldHandles

func init() {
	NsObjectCountFields = &NsObjectCountFieldHandles{
		ScopeName:                "scope_name",
		ObjectCount:              "object_count",
		MaxLimitOverride:         "max_limit_override",
		WarningThresholdOverride: "warning_threshold_override",
	}
}

// NsObjectCount - Number of objects of a type in a given scope.
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

// NsObjectCountFieldHandles provides a string representation for each AccessControlRecord field.
type NsObjectCountFieldHandles struct {
	ScopeName                string
	ObjectCount              string
	MaxLimitOverride         string
	WarningThresholdOverride string
}
