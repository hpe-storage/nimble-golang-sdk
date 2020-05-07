/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/Job


// Job :
type Job struct {
   // CompletionTime
   CompletionTime float64 `json:"completion_time,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // CurrentPhase
   CurrentPhase float64 `json:"current_phase,omitempty"`
   // CurrentPhaseDescription
   CurrentPhaseDescription string `json:"current_phase_description,omitempty"`
   // Description
   Description string `json:"description,omitempty"`
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
   // ObjectID
   ObjectID string `json:"object_id,omitempty"`
   // ParentJobID
   ParentJobID string `json:"parent_job_id,omitempty"`
   // PercentComplete
   PercentComplete float64 `json:"percent_complete,omitempty"`
   // TotalPhases
   TotalPhases float64 `json:"total_phases,omitempty"`
}
