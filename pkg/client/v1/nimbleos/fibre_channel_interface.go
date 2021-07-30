// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export FibreChannelInterfaceFields provides field names to use in filter parameters, for example.
var FibreChannelInterfaceFields *FibreChannelInterfaceFieldHandles

func init() {
	fieldID := "id"
	fieldArrayNameOrSerial := "array_name_or_serial"
	fieldPartialResponseOk := "partial_response_ok"
	fieldControllerName := "controller_name"
	fieldFcPortId := "fc_port_id"
	fieldName := "name"
	fieldWwnn := "wwnn"
	fieldWwpn := "wwpn"
	fieldPeerzone := "peerzone"
	fieldOnline := "online"
	fieldFirmwareVersion := "firmware_version"
	fieldLogicalPortNumber := "logical_port_number"
	fieldFcPortName := "fc_port_name"
	fieldBusLocation := "bus_location"
	fieldSlot := "slot"
	fieldOrientation := "orientation"
	fieldPort := "port"
	fieldLinkInfo := "link_info"
	fieldFabricInfo := "fabric_info"

	FibreChannelInterfaceFields = &FibreChannelInterfaceFieldHandles{
		ID:                &fieldID,
		ArrayNameOrSerial: &fieldArrayNameOrSerial,
		PartialResponseOk: &fieldPartialResponseOk,
		ControllerName:    &fieldControllerName,
		FcPortId:          &fieldFcPortId,
		Name:              &fieldName,
		Wwnn:              &fieldWwnn,
		Wwpn:              &fieldWwpn,
		Peerzone:          &fieldPeerzone,
		Online:            &fieldOnline,
		FirmwareVersion:   &fieldFirmwareVersion,
		LogicalPortNumber: &fieldLogicalPortNumber,
		FcPortName:        &fieldFcPortName,
		BusLocation:       &fieldBusLocation,
		Slot:              &fieldSlot,
		Orientation:       &fieldOrientation,
		Port:              &fieldPort,
		LinkInfo:          &fieldLinkInfo,
		FabricInfo:        &fieldFabricInfo,
	}
}

// FibreChannelInterface - Represent information of specified Fibre Channel interfaces. Fibre Channel interfaces are hosted on Fibre Channel ports to provide data access.
type FibreChannelInterface struct {
	// ID - Identifier for the Fibre Channel interface.
	ID *string `json:"id,omitempty"`
	// ArrayNameOrSerial - Name or serial number of array where the interface is hosted.
	ArrayNameOrSerial *string `json:"array_name_or_serial,omitempty"`
	// PartialResponseOk - Indicate that it is ok to provide partially available response.
	PartialResponseOk *bool `json:"partial_response_ok,omitempty"`
	// ControllerName - Name (A or B) of the controller where the interface is hosted.
	ControllerName *string `json:"controller_name,omitempty"`
	// FcPortId - ID of the port with which the interface is associated.
	FcPortId *string `json:"fc_port_id,omitempty"`
	// Name - Name of Fibre Channel interface.
	Name *string `json:"name,omitempty"`
	// Wwnn - WWNN (World Wide Node Name) for this Fibre Channel interface.
	Wwnn *string `json:"wwnn,omitempty"`
	// Wwpn - WWPN (World Wide Port Name) for this Fibre Channel interface.
	Wwpn *string `json:"wwpn,omitempty"`
	// Peerzone - Active peer zone for this Fibre Channel interface.
	Peerzone *string `json:"peerzone,omitempty"`
	// Online - Identify whether the Fibre Channel interface is online.
	Online *bool `json:"online,omitempty"`
	// FirmwareVersion - Version of the Fibre Channel firmware.
	FirmwareVersion *string `json:"firmware_version,omitempty"`
	// LogicalPortNumber - Logical port number for the Fibre Channel port.
	LogicalPortNumber *int64 `json:"logical_port_number,omitempty"`
	// FcPortName - Name of Fibre Channel port.
	FcPortName *string `json:"fc_port_name,omitempty"`
	// BusLocation - PCI bus location of the HBA for this Fibre Channel port.
	BusLocation *string `json:"bus_location,omitempty"`
	// Slot - HBA slot number for this Fibre Channel port.
	Slot *int64 `json:"slot,omitempty"`
	// Orientation - Orientation of FC ports on a HBA. An orientation of 'right_to_left' indicates that ports are ordered as 3,2,1,0 on the slot. Possible values: 'left_to_right', 'right_to_left'.
	Orientation *NsPlatOrientation `json:"orientation,omitempty"`
	// Port - HBA port number for this Fibre Channel port.
	Port *int64 `json:"port,omitempty"`
	// LinkInfo - Information about the Fibre Channel link at which interface is operating.
	LinkInfo *NsFcLinkInfo `json:"link_info,omitempty"`
	// FabricInfo - Fibre Channel fabric information.
	FabricInfo *NsFcFabricInfo `json:"fabric_info,omitempty"`
}

// FibreChannelInterfaceFieldHandles provides a string representation for each AccessControlRecord field.
type FibreChannelInterfaceFieldHandles struct {
	ID                *string
	ArrayNameOrSerial *string
	PartialResponseOk *string
	ControllerName    *string
	FcPortId          *string
	Name              *string
	Wwnn              *string
	Wwpn              *string
	Peerzone          *string
	Online            *string
	FirmwareVersion   *string
	LogicalPortNumber *string
	FcPortName        *string
	BusLocation       *string
	Slot              *string
	Orientation       *string
	Port              *string
	LinkInfo          *string
	FabricInfo        *string
}
