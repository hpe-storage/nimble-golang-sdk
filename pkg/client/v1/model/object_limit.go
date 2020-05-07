/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/ObjectLimit


// ObjectLimit :
type ObjectLimit struct {
   // ID
   ID string `json:"id,omitempty"`
   // ObjectType
   ObjectType float64 `json:"object_type,omitempty"`
   // ObjectTypeName
   ObjectTypeName string `json:"object_type_name,omitempty"`
   // ScopeType
   ScopeType float64 `json:"scope_type,omitempty"`
   // ScopeTypeName
   ScopeTypeName string `json:"scope_type_name,omitempty"`
   // WarningThreshold
   WarningThreshold float64 `json:"warning_threshold,omitempty"`
   // MaxLimit
   MaxLimit float64 `json:"max_limit,omitempty"`
}
