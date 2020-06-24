/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

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
   
    // Name of the scope.
    
 	ScopeName *string `json:"scope_name,omitempty"`
   
    // Number of objects.
    
   	ObjectCount *int64 `json:"object_count,omitempty"`
   
    // Override max object limit for this scope.
    
   	MaxLimitOverrIDe *int64 `json:"max_limit_override,omitempty"`
   
    // Override warning threshold for this scope.
    
   	WarningThresholdOverrIDe *int64 `json:"warning_threshold_override,omitempty"`
}
