/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsGenericResp


// NsGenericResp :
type NsGenericResp struct {
   // GenericError
   GenericError string `json:"generic_error,omitempty"`
   // GenericErrorMessage
   GenericErrorMessage string `json:"generic_error_message,omitempty"`
   // ConnStatusOk
   ConnStatusOk bool `json:"conn_status_ok,omitempty"`
   // ConnMessage
   ConnMessage string `json:"conn_message,omitempty"`
}
