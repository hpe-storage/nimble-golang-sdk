/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsISCSIInitiator - ISCSI initiator.
// Export NsISCSIInitiatorFields for advance operations like search filter etc.
var NsISCSIInitiatorFields *NsISCSIInitiator

func init(){
	IDfield:= "id"
	InitiatorIDfield:= "initiator_id"
	Labelfield:= "label"
	Iqnfield:= "iqn"
	IpAddressfield:= "ip_address"
		
	NsISCSIInitiatorFields= &NsISCSIInitiator{
		ID: &IDfield,
		InitiatorID: &InitiatorIDfield,
		Label: &Labelfield,
		Iqn: &Iqnfield,
		IpAddress: &IpAddressfield,
		
	}
}

type NsISCSIInitiator struct {
   
    // Unique identifier of the iSCSI initiator.
    
 	ID *string `json:"id,omitempty"`
   
    // Unique identifier of the iSCSI initiator.
    
 	InitiatorID *string `json:"initiator_id,omitempty"`
   
    // Unique label of the iSCSI initiator.
    
 	Label *string `json:"label,omitempty"`
   
    // IQN name of the iSCSI initiator.
    
 	Iqn *string `json:"iqn,omitempty"`
   
    // IP address of the iSCSI initiator.
    
 	IpAddress *string `json:"ip_address,omitempty"`
}
