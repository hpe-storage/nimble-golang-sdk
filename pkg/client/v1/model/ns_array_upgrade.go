/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsArrayUpgrade - Array upgrade attributes.
// Export NsArrayUpgradeFields for advance operations like search filter etc.
var NsArrayUpgradeFields *NsArrayUpgrade

func init(){
		
	NsArrayUpgradeFields= &NsArrayUpgrade{
		
	}
}

type NsArrayUpgrade struct {
   
    // Array upgrade type.
    
   	Type *NsArrayUpgradeType `json:"type,omitempty"`
   
    // Array upgrade state.
    
   	State *NsArrayUpgradeState `json:"state,omitempty"`
   
    // Array upgrade stage.
    
   	Stage *NsArrayUpgradeStage `json:"stage,omitempty"`
   
    // List of error or info messages.
    
   	Messages []*NsErrorWithArguments `json:"messages,omitempty"`
}
