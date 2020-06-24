/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// FibreChannelSession - Fibre Channel session is created when Fibre Channel initiator connects to this group.
// Export FibreChannelSessionFields for advance operations like search filter etc.
var FibreChannelSessionFields *FibreChannelSession

func init(){
	IDfield:= "id"
		
	FibreChannelSessionFields= &FibreChannelSession{
		ID: &IDfield,
		
	}
}

type FibreChannelSession struct {
   
    // Unique identifier of the Fibre Channel session.
    
 	ID *string `json:"id,omitempty"`
   
    // Information about the Fibre Channel initiator.
    
	InitiatorInfo *NsFcSessionInitiator `json:"initiator_info,omitempty"`
   
    // Information about the Fibre Channel target.
    
	TargetInfo *NsFcSessionTarget `json:"target_info,omitempty"`
}
