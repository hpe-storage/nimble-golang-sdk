// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos


// NsArrayUpgrade - Array upgrade attributes.
// Export NsArrayUpgradeFields for advance operations like search filter etc.
var NsArrayUpgradeFields *NsArrayUpgrade

func init(){
 CtrlrAPortListfield:= "ctrlr_a_port_list"
 CtrlrBPortListfield:= "ctrlr_b_port_list"

 NsArrayUpgradeFields= &NsArrayUpgrade{
  CtrlrAPortList:   &CtrlrAPortListfield,
  CtrlrBPortList:   &CtrlrBPortListfield,
 }
}


type NsArrayUpgrade struct {
 // Type - Array upgrade type.
    Type *NsArrayUpgradeType `json:"type,omitempty"`
 // State - Array upgrade state.
    State *NsArrayUpgradeState `json:"state,omitempty"`
 // Stage - Array upgrade stage.
    Stage *NsArrayUpgradeStage `json:"stage,omitempty"`
 // CtrlrAPortList - Comma-separated list of array upgrade port names for controller A.
  CtrlrAPortList *string `json:"ctrlr_a_port_list,omitempty"`
 // CtrlrBPortList - Comma-separated list of array upgrade port names for controller B.
  CtrlrBPortList *string `json:"ctrlr_b_port_list,omitempty"`
 // Messages - List of error or info messages.
    Messages []*NsErrorWithArguments `json:"messages,omitempty"`
 // Metadata - Key-value pairs that augment an array upgrade attributes.
    Metadata []*NsKeyValue `json:"metadata,omitempty"`
}
