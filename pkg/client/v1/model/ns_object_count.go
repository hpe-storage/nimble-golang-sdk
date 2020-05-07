/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsObjectCount


// NsObjectCount :
type NsObjectCount struct {
   // ScopeName
   ScopeName string `json:"scope_name,omitempty"`
   // ObjectCount
   ObjectCount float64 `json:"object_count,omitempty"`
   // MaxLimitOverrIDe
   MaxLimitOverrIDe float64 `json:"max_limit_override,omitempty"`
   // WarningThresholdOverrIDe
   WarningThresholdOverrIDe float64 `json:"warning_threshold_override,omitempty"`
}
