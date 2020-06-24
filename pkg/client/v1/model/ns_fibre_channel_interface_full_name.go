/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsFibreChannelInterfaceFullName - Fully qualified name information for a Fibre Channel interface (array/controller/interface).
// Export NsFibreChannelInterfaceFullNameFields for advance operations like search filter etc.
var NsFibreChannelInterfaceFullNameFields *NsFibreChannelInterfaceFullName

func init(){
	ArrayNamefield:= "array_name"
	CtrlrNamefield:= "ctrlr_name"
	IntfNamefield:= "intf_name"
		
	NsFibreChannelInterfaceFullNameFields= &NsFibreChannelInterfaceFullName{
		ArrayName: &ArrayNamefield,
		CtrlrName: &CtrlrNamefield,
		IntfName: &IntfNamefield,
		
	}
}

type NsFibreChannelInterfaceFullName struct {
   
    // Array name.
    
 	ArrayName *string `json:"array_name,omitempty"`
   
    // Controller name.
    
 	CtrlrName *string `json:"ctrlr_name,omitempty"`
   
    // Fibre Channel interface name.
    
 	IntfName *string `json:"intf_name,omitempty"`
}
