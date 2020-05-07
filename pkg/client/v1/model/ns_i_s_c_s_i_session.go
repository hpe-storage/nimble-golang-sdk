/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsISCSISession


// NsISCSISession :
type NsISCSISession struct {
   // ID
   ID string `json:"id,omitempty"`
   // SessionID
   SessionID string `json:"session_id,omitempty"`
   // InitiatorName
   InitiatorName string `json:"initiator_name,omitempty"`
   // NumConnections
   NumConnections float64 `json:"num_connections,omitempty"`
   // PrKey
   PrKey  int32 `json:"pr_key,omitempty"`
   // InitiatorIpAddr
   InitiatorIpAddr string `json:"initiator_ip_addr,omitempty"`
   // TargetIpAddr
   TargetIpAddr string `json:"target_ip_addr,omitempty"`
   // HeaderDigestEnabled
   HeaderDigestEnabled bool `json:"header_digest_enabled,omitempty"`
   // DataDigestEnabled
   DataDigestEnabled bool `json:"data_digest_enabled,omitempty"`
}
