/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NetworkInterface - Manage per array network interface configuration.
// Export NetworkInterfaceFields for advance operations like search filter etc.
var NetworkInterfaceFields *NetworkInterface

func init(){
	IDfield:= "id"
	ArrayNameOrSerialfield:= "array_name_or_serial"
	ArrayIDfield:= "array_id"
	ControllerNamefield:= "controller_name"
	ControllerIDfield:= "controller_id"
	Namefield:= "name"
	Macfield:= "mac"
		
	NetworkInterfaceFields= &NetworkInterface{
		ID: &IDfield,
		ArrayNameOrSerial: &ArrayNameOrSerialfield,
		ArrayID: &ArrayIDfield,
		ControllerName: &ControllerNamefield,
		ControllerID: &ControllerIDfield,
		Name: &Namefield,
		Mac: &Macfield,
		
	}
}

type NetworkInterface struct {
   
    // Identifier for the interface.
    
 	ID *string `json:"id,omitempty"`
   
    // Name or serial of the array where the interface is hosted.
    
 	ArrayNameOrSerial *string `json:"array_name_or_serial,omitempty"`
   
    // Indicate that it is ok to provide partially available response.
    
 	PartialResponseOk *bool `json:"partial_response_ok,omitempty"`
   
    // Identifier for the array.
    
 	ArrayID *string `json:"array_id,omitempty"`
   
    // Name (A or B) of the controller where the interface is hosted.
    
 	ControllerName *string `json:"controller_name,omitempty"`
   
    // Identifier of the controller where the interface is hosted.
    
 	ControllerID *string `json:"controller_id,omitempty"`
   
    // Name of the interface.
    
 	Name *string `json:"name,omitempty"`
   
    // MAC address of the interface.
    
 	Mac *string `json:"mac,omitempty"`
   
    // Whether this interface is present on this controller.
    
 	IsPresent *bool `json:"is_present,omitempty"`
   
    // Speed of the link.
    
   	LinkSpeed *NsPlatLinkSpeed `json:"link_speed,omitempty"`
   
    // Status of the link.
    
   	LinkStatus *NsPlatLinkStatus `json:"link_status,omitempty"`
   
    // MTU on the link.
    
   	Mtu *int64 `json:"mtu,omitempty"`
   
    // Port number for this interface.
    
   	Port *int64 `json:"port,omitempty"`
   
    // Slot number for this interface.
    
   	Slot *int64 `json:"slot,omitempty"`
   
    // Maximum speed of the link.
    
   	MaxLinkSpeed *NsPlatLinkSpeed `json:"max_link_speed,omitempty"`
   
    // Interface type.
    
   	NicType *NsPlatNicType `json:"nic_type,omitempty"`
   
    // List of IP addresses assigned to this network interface.
    
   	IpList []*NsAssignedIP `json:"ip_list,omitempty"`
}
