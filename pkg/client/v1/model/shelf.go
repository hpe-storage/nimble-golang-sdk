/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/Shelf


// Shelf :
type Shelf struct {
   // ID
   ID string `json:"id,omitempty"`
   // ArrayName
   ArrayName string `json:"array_name,omitempty"`
   // ArrayID
   ArrayID string `json:"array_id,omitempty"`
   // PartialResponseOk
   PartialResponseOk bool `json:"partial_response_ok,omitempty"`
   // Serial
   Serial string `json:"serial,omitempty"`
   // Model
   Model string `json:"model,omitempty"`
   // ModelExt
   ModelExt string `json:"model_ext,omitempty"`
   // Activated
   Activated bool `json:"activated,omitempty"`
   // Driveset
   Driveset float64 `json:"driveset,omitempty"`
   // Force
   Force bool `json:"force,omitempty"`
   // AcceptForeign
   AcceptForeign bool `json:"accept_foreign,omitempty"`
   // AcceptDedupeImpact
   AcceptDedupeImpact bool `json:"accept_dedupe_impact,omitempty"`
}
