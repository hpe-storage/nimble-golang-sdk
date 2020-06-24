/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsFCInitiator - Fibre Channel initiator.
// Export NsFCInitiatorFields for advance operations like search filter etc.
var NsFCInitiatorFields *NsFCInitiator

func init(){
	IDfield:= "id"
	InitiatorIDfield:= "initiator_id"
	Wwpnfield:= "wwpn"
	Aliasfield:= "alias"
		
	NsFCInitiatorFields= &NsFCInitiator{
		ID: &IDfield,
		InitiatorID: &InitiatorIDfield,
		Wwpn: &Wwpnfield,
		Alias: &Aliasfield,
		
	}
}

type NsFCInitiator struct {
   
    // Unique identifier of the Fibre Channel initiator.
    
 	ID *string `json:"id,omitempty"`
   
    // Unique identifier of the Fibre Channel initiator.
    
 	InitiatorID *string `json:"initiator_id,omitempty"`
   
    // WWPN (World Wide Port Name) of the Fibre Channel initiator.
    
 	Wwpn *string `json:"wwpn,omitempty"`
   
    // Alias of the Fibre Channel initiator.
    
 	Alias *string `json:"alias,omitempty"`
}
