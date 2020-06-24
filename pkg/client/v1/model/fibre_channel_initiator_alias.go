/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// FibreChannelInitiatorAlias - This API provides the alias information for Fibre Channel initiators.
// Export FibreChannelInitiatorAliasFields for advance operations like search filter etc.
var FibreChannelInitiatorAliasFields *FibreChannelInitiatorAlias

func init(){
	IDfield:= "id"
	Aliasfield:= "alias"
	Wwpnfield:= "wwpn"
		
	FibreChannelInitiatorAliasFields= &FibreChannelInitiatorAlias{
		ID: &IDfield,
		Alias: &Aliasfield,
		Wwpn: &Wwpnfield,
		
	}
}

type FibreChannelInitiatorAlias struct {
   
    // Unique identifier for the Fibre Channel initiator alias.
    
 	ID *string `json:"id,omitempty"`
   
    // Alias of the Fibre Channel initiator.
    
 	Alias *string `json:"alias,omitempty"`
   
    // WWPN (World Wide Port Name) of the Fibre Channel initiator.
    
 	Wwpn *string `json:"wwpn,omitempty"`
   
    // Source of the Fibre Channel initiator alias.
    
   	Source *SmFcInitiatorAliasSource `json:"source,omitempty"`
}
