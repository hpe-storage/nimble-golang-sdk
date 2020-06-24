/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsISCSIIQN - ISCSI IQN.
// Export NsISCSIIQNFields for advance operations like search filter etc.
var NsISCSIIQNFields *NsISCSIIQN

func init(){
	Namefield:= "name"
		
	NsISCSIIQNFields= &NsISCSIIQN{
		Name: &Namefield,
		
	}
}

type NsISCSIIQN struct {
   
    // IQN name of the iSCSI initiator.
    
 	Name *string `json:"name,omitempty"`
}
