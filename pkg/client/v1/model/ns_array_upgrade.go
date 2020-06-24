// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsArrayUpgrade - Array upgrade attributes.
// Export NsArrayUpgradeFields for advance operations like search filter etc.
var NsArrayUpgradeFields *NsArrayUpgrade

func init(){
		
	NsArrayUpgradeFields= &NsArrayUpgrade{
		
	}
}

type NsArrayUpgrade struct {
	// Type - Array upgrade type.
   	Type *NsArrayUpgradeType `json:"type,omitempty"`
	// State - Array upgrade state.
   	State *NsArrayUpgradeState `json:"state,omitempty"`
	// Stage - Array upgrade stage.
   	Stage *NsArrayUpgradeStage `json:"stage,omitempty"`
	// Messages - List of error or info messages.
   	Messages []*NsErrorWithArguments `json:"messages,omitempty"`
}
