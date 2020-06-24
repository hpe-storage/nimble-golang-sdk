/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsFcSessionTarget - Information of the Fibre Channel session target.
// Export NsFcSessionTargetFields for advance operations like search filter etc.
var NsFcSessionTargetFields *NsFcSessionTarget

func init(){
	TargetPortArrayNamefield:= "target_port_array_name"
	TargetPortCtrlrNamefield:= "target_port_ctrlr_name"
	TargetPortInterfaceNamefield:= "target_port_interface_name"
	TargetWwnnfield:= "target_wwnn"
	TargetWwpnfield:= "target_wwpn"
	TargetFcIDfield:= "target_fcid"
		
	NsFcSessionTargetFields= &NsFcSessionTarget{
		TargetPortArrayName: &TargetPortArrayNamefield,
		TargetPortCtrlrName: &TargetPortCtrlrNamefield,
		TargetPortInterfaceName: &TargetPortInterfaceNamefield,
		TargetWwnn: &TargetWwnnfield,
		TargetWwpn: &TargetWwpnfield,
		TargetFcID: &TargetFcIDfield,
		
	}
}

type NsFcSessionTarget struct {
   
    // Name of the array hosting the Fibre Channel target port.
    
 	TargetPortArrayName *string `json:"target_port_array_name,omitempty"`
   
    // Name (A or B) of the controller to which the port belongs.
    
 	TargetPortCtrlrName *string `json:"target_port_ctrlr_name,omitempty"`
   
    // Name of the interface hosted on the Fibre Channel target port.
    
 	TargetPortInterfaceName *string `json:"target_port_interface_name,omitempty"`
   
    // WWNN (World Wide Node Name) of the Fibre Channel target port.
    
 	TargetWwnn *string `json:"target_wwnn,omitempty"`
   
    // WWPN (World Wide Port Name) of the Fibre Channel target port.
    
 	TargetWwpn *string `json:"target_wwpn,omitempty"`
   
    // FCID assigned to the Fibre Channel target port.
    
 	TargetFcID *string `json:"target_fcid,omitempty"`
}
