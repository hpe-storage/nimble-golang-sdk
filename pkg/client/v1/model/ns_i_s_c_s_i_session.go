/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsISCSISession - ISCSI initiator session information.
// Export NsISCSISessionFields for advance operations like search filter etc.
var NsISCSISessionFields *NsISCSISession

func init(){
	IDfield:= "id"
	SessionIDfield:= "session_id"
	InitiatorNamefield:= "initiator_name"
	InitiatorIpAddrfield:= "initiator_ip_addr"
	TargetIpAddrfield:= "target_ip_addr"
		
	NsISCSISessionFields= &NsISCSISession{
		ID: &IDfield,
		SessionID: &SessionIDfield,
		InitiatorName: &InitiatorNamefield,
		InitiatorIpAddr: &InitiatorIpAddrfield,
		TargetIpAddr: &TargetIpAddrfield,
		
	}
}

type NsISCSISession struct {
   
    // Unique identifier of the iSCSI session.
    
 	ID *string `json:"id,omitempty"`
   
    // Unique identifier of the iSCSI session.
    
 	SessionID *string `json:"session_id,omitempty"`
   
    // Name of the iSCSI initiator (IQN).
    
 	InitiatorName *string `json:"initiator_name,omitempty"`
   
    // Number of connections in iSCSI session.
    
   	NumConnections *int64 `json:"num_connections,omitempty"`
   
    // Registered persistent reservation key for the I-T nexus on this LU (if there is no registration, its value will be zero).
    
  	PrKey  *int64 `json:"pr_key,omitempty"`
   
    // IP address of the iSCSI initiator.
    
 	InitiatorIpAddr *string `json:"initiator_ip_addr,omitempty"`
   
    // Target IP address of the iSCSI initiator.
    
 	TargetIpAddr *string `json:"target_ip_addr,omitempty"`
   
    // Indicate whether header digest is enabled.
    
 	HeaderDigestEnabled *bool `json:"header_digest_enabled,omitempty"`
   
    // Indicate whether data digest is enabled.
    
 	DataDigestEnabled *bool `json:"data_digest_enabled,omitempty"`
}
