/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsHostType


// NsHostType :
type NsHostType struct {
   // InitiatorName
   InitiatorName string `json:"initiator_name,omitempty"`
   // SourceHostType
   SourceHostType string `json:"source_host_type,omitempty"`
   // DestinationHostType
   DestinationHostType string `json:"destination_host_type,omitempty"`
}
