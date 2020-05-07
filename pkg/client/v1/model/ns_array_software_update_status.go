/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsArraySoftwareUpdateStatus


// NsArraySoftwareUpdateStatus :
type NsArraySoftwareUpdateStatus struct {
   // ArrayName
   ArrayName string `json:"array_name,omitempty"`
   // Error
   Error string `json:"error,omitempty"`
   // CtrlrErrorMask
   CtrlrErrorMask float64 `json:"ctrlr_error_mask,omitempty"`
}
