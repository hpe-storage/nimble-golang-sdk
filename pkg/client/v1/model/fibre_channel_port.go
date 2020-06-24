// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// FibreChannelPort - Fibre Channel ports provide data access. This API provides the list of all Fibre Channel ports configured on the arrays.
// Export FibreChannelPortFields for advance operations like search filter etc.
var FibreChannelPortFields *FibreChannelPort

func init(){
	IDfield:= "id"
	ArrayNameOrSerialfield:= "array_name_or_serial"
	ControllerNamefield:= "controller_name"
	FcPortNamefield:= "fc_port_name"
	BusLocationfield:= "bus_location"
		
	FibreChannelPortFields= &FibreChannelPort{
	ID: &IDfield,
	ArrayNameOrSerial: &ArrayNameOrSerialfield,
	ControllerName: &ControllerNamefield,
	FcPortName: &FcPortNamefield,
	BusLocation: &BusLocationfield,
		
	}
}

type FibreChannelPort struct {
	// ID - Identifier for the Fibre Channel port.
 	ID *string `json:"id,omitempty"`
	// ArrayNameOrSerial - Name or serial number of the array.
 	ArrayNameOrSerial *string `json:"array_name_or_serial,omitempty"`
	// ControllerName - Name (A or B) of the controller to which the port belongs.
 	ControllerName *string `json:"controller_name,omitempty"`
	// FcPortName - Name of the Fibre Channel port.
 	FcPortName *string `json:"fc_port_name,omitempty"`
	// BusLocation - PCI bus location of the HBA (Host Bus Adapter) for this Fibre Channel port.
 	BusLocation *string `json:"bus_location,omitempty"`
	// Port - HBA (Host Bus Adapter) port number for this Fibre Channel port.
   	Port *int64 `json:"port,omitempty"`
	// Slot - HBA (Host Bus Adapter) slot number for this Fibre Channel port.
   	Slot *int64 `json:"slot,omitempty"`
	// Orientation - Orientation of FC ports on a HBA. An orientation of 'right_to_left' indicates that ports are ordered as 3,2,1,0 on the slot. Possible values: 'left_to_right', 'right_to_left'.
   	Orientation *NsPlatOrientation `json:"orientation,omitempty"`
	// LinkInfo - Information about the Fibre Channel link associated with this port.
	LinkInfo *NsFcLinkInfo `json:"link_info,omitempty"`
	// RxPower - SFP RX power in uW.
   	RxPower *int64 `json:"rx_power,omitempty"`
	// TxPower - SFP TX power in uW.
   	TxPower *int64 `json:"tx_power,omitempty"`
}
