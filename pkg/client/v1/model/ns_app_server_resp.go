/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsAppServerResp


// NsAppServerResp :
type NsAppServerResp struct {
   // GeneralError
   GeneralError string `json:"general_error,omitempty"`
   // HasAssocVols
   HasAssocVols bool `json:"has_assoc_vols,omitempty"`
}
