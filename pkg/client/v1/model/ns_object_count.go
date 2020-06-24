// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsObjectCount - Number of objects of a type in a given scope.
// Export NsObjectCountFields for advance operations like search filter etc.
var NsObjectCountFields *NsObjectCount

func init(){
	ScopeNamefield:= "scope_name"
		
	NsObjectCountFields= &NsObjectCount{
	ScopeName: &ScopeNamefield,
		
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
