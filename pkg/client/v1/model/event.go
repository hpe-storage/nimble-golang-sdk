/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/Event


// Event :
type Event struct {
   // ID
   ID string `json:"id,omitempty"`
   // Type
   Type float64 `json:"type,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // Scope
   Scope string `json:"scope,omitempty"`
   // Target
   Target string `json:"target,omitempty"`
   // Timestamp
   Timestamp float64 `json:"timestamp,omitempty"`
   // Summary
   Summary string `json:"summary,omitempty"`
   // Activity
   Activity string `json:"activity,omitempty"`
   // AlarmID
   AlarmID string `json:"alarm_id,omitempty"`
   // TenantID
   TenantID string `json:"tenant_id,omitempty"`
}
