/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// FibreChannelInterface - Represent information of specified Fibre Channel interfaces. Fibre Channel interfaces are hosted on Fibre Channel ports to provide data access.
// Export FibreChannelInterfaceFields for advance operations like search filter etc.
var FibreChannelInterfaceFields *FibreChannelInterface

func init(){
	IDfield:= "id"
	ArrayNameOrSerialfield:= "array_name_or_serial"
	ControllerNamefield:= "controller_name"
	FcPortIDfield:= "fc_port_id"
	Namefield:= "name"
	Wwnnfield:= "wwnn"
	Wwpnfield:= "wwpn"
	Peerzonefield:= "peerzone"
	FirmwareVersionfield:= "firmware_version"
	FcPortNamefield:= "fc_port_name"
	BusLocationfield:= "bus_location"
		
	FibreChannelInterfaceFields= &FibreChannelInterface{
		ID: &IDfield,
		ArrayNameOrSerial: &ArrayNameOrSerialfield,
		ControllerName: &ControllerNamefield,
		FcPortID: &FcPortIDfield,
		Name: &Namefield,
		Wwnn: &Wwnnfield,
		Wwpn: &Wwpnfield,
		Peerzone: &Peerzonefield,
		FirmwareVersion: &FirmwareVersionfield,
		FcPortName: &FcPortNamefield,
		BusLocation: &BusLocationfield,
		
	}
}

type FibreChannelInterface struct {
   
    // Identifier for the Fibre Channel interface.
    
 	ID *string `json:"id,omitempty"`
   
    // Name or serial number of array where the interface is hosted.
    
 	ArrayNameOrSerial *string `json:"array_name_or_serial,omitempty"`
   
    // Indicate that it is ok to provide partially available response.
    
 	PartialResponseOk *bool `json:"partial_response_ok,omitempty"`
   
    // Name (A or B) of the controller where the interface is hosted.
    
 	ControllerName *string `json:"controller_name,omitempty"`
   
    // ID of the port with which the interface is associated.
    
 	FcPortID *string `json:"fc_port_id,omitempty"`
   
    // Name of Fibre Channel interface.
    
 	Name *string `json:"name,omitempty"`
   
    // WWNN (World Wide Node Name) for this Fibre Channel interface.
    
 	Wwnn *string `json:"wwnn,omitempty"`
   
    // WWPN (World Wide Port Name) for this Fibre Channel interface.
    
 	Wwpn *string `json:"wwpn,omitempty"`
   
    // Active peer zone for this Fibre Channel interface.
    
 	Peerzone *string `json:"peerzone,omitempty"`
   
    // Identify whether the Fibre Channel interface is online.
    
 	Online *bool `json:"online,omitempty"`
   
    // Version of the Fibre Channel firmware.
    
 	FirmwareVersion *string `json:"firmware_version,omitempty"`
   
    // Logical port number for the Fibre Channel port.
    
   	LogicalPortNumber *int64 `json:"logical_port_number,omitempty"`
   
    // Name of Fibre Channel port.
    
 	FcPortName *string `json:"fc_port_name,omitempty"`
   
    // PCI bus location of the HBA for this Fibre Channel port.
    
 	BusLocation *string `json:"bus_location,omitempty"`
   
    // HBA slot number for this Fibre Channel port.
    
   	Slot *int64 `json:"slot,omitempty"`
   
    // Orientation of FC ports on a HBA. An orientation of 'right_to_left' indicates that ports are ordered as 3,2,1,0 on the slot. Possible values: 'left_to_right', 'right_to_left'.
    
   	Orientation *NsPlatOrientation `json:"orientation,omitempty"`
   
    // HBA port number for this Fibre Channel port.
    
   	Port *int64 `json:"port,omitempty"`
   
    // Information about the Fibre Channel link at which interface is operating.
    
	LinkInfo *NsFcLinkInfo `json:"link_info,omitempty"`
   
    // Fibre Channel fabric information.
    
	FabricInfo *NsFcFabricInfo `json:"fabric_info,omitempty"`
}
