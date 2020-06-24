/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsHostType - Host type attribute.
// Export NsHostTypeFields for advance operations like search filter etc.
var NsHostTypeFields *NsHostType

func init(){
	InitiatorNamefield:= "initiator_name"
	SourceHostTypefield:= "source_host_type"
	DestinationHostTypefield:= "destination_host_type"
		
	NsHostTypeFields= &NsHostType{
		InitiatorName: &InitiatorNamefield,
		SourceHostType: &SourceHostTypefield,
		DestinationHostType: &DestinationHostTypefield,
		
	}
}

type NsHostType struct {
   
    // The initiator for which the conflict exists.
    
 	InitiatorName *string `json:"initiator_name,omitempty"`
   
    // Source initiator group.
    
	SourceInitiatorGroup []*string `json:"source_initiator_group,omitempty"`
   
    // Source host type.
    
 	SourceHostType *string `json:"source_host_type,omitempty"`
   
    // Destination initiator group.
    
	DestinationInitiatorGroup []*string `json:"destination_initiator_group,omitempty"`
   
    // Destination Host type.
    
 	DestinationHostType *string `json:"destination_host_type,omitempty"`
}
