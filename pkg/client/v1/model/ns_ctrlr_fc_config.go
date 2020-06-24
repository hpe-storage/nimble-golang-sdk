/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsCtrlrFcConfig - Controller Fibre Channel configuration.
// Export NsCtrlrFcConfigFields for advance operations like search filter etc.
var NsCtrlrFcConfigFields *NsCtrlrFcConfig

func init(){
		
	NsCtrlrFcConfigFields= &NsCtrlrFcConfig{
		
	}
}

type NsCtrlrFcConfig struct {
   
    // List of Fibre Channel ports.
    
   	FcPortList []*NsFcPortInfo `json:"fc_port_list,omitempty"`
   
    // List of Fibre Channel interfaces.
    
   	FcInterfaceList []*NsFcInterfaceInfo `json:"fc_interface_list,omitempty"`
}
