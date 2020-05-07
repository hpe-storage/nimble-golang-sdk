/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/HealthCheck


// HealthCheck :
type HealthCheck struct {
   // ID
   ID string `json:"id,omitempty"`
   // OnDemand
   OnDemand bool `json:"on_demand,omitempty"`
   // ArrayID
   ArrayID string `json:"array_id,omitempty"`
   // CtrlrID
   CtrlrID string `json:"ctrlr_id,omitempty"`
}
