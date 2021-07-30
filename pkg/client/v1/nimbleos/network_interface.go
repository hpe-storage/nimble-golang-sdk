// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NetworkInterfaceFields provides field names to use in filter parameters, for example.
var NetworkInterfaceFields *NetworkInterfaceFieldHandles

func init() {
	fieldID := "id"
	fieldArrayNameOrSerial := "array_name_or_serial"
	fieldPartialResponseOk := "partial_response_ok"
	fieldArrayId := "array_id"
	fieldControllerName := "controller_name"
	fieldControllerId := "controller_id"
	fieldName := "name"
	fieldMac := "mac"
	fieldIsPresent := "is_present"
	fieldLinkSpeed := "link_speed"
	fieldLinkStatus := "link_status"
	fieldMtu := "mtu"
	fieldPort := "port"
	fieldSlot := "slot"
	fieldMaxLinkSpeed := "max_link_speed"
	fieldNicType := "nic_type"
	fieldIpList := "ip_list"

	NetworkInterfaceFields = &NetworkInterfaceFieldHandles{
		ID:                &fieldID,
		ArrayNameOrSerial: &fieldArrayNameOrSerial,
		PartialResponseOk: &fieldPartialResponseOk,
		ArrayId:           &fieldArrayId,
		ControllerName:    &fieldControllerName,
		ControllerId:      &fieldControllerId,
		Name:              &fieldName,
		Mac:               &fieldMac,
		IsPresent:         &fieldIsPresent,
		LinkSpeed:         &fieldLinkSpeed,
		LinkStatus:        &fieldLinkStatus,
		Mtu:               &fieldMtu,
		Port:              &fieldPort,
		Slot:              &fieldSlot,
		MaxLinkSpeed:      &fieldMaxLinkSpeed,
		NicType:           &fieldNicType,
		IpList:            &fieldIpList,
	}
}

// NetworkInterface - Manage per array network interface configuration.
type NetworkInterface struct {
	// ID - Identifier for the interface.
	ID *string `json:"id,omitempty"`
	// ArrayNameOrSerial - Name or serial of the array where the interface is hosted.
	ArrayNameOrSerial *string `json:"array_name_or_serial,omitempty"`
	// PartialResponseOk - Indicate that it is ok to provide partially available response.
	PartialResponseOk *bool `json:"partial_response_ok,omitempty"`
	// ArrayId - Identifier for the array.
	ArrayId *string `json:"array_id,omitempty"`
	// ControllerName - Name (A or B) of the controller where the interface is hosted.
	ControllerName *string `json:"controller_name,omitempty"`
	// ControllerId - Identifier of the controller where the interface is hosted.
	ControllerId *string `json:"controller_id,omitempty"`
	// Name - Name of the interface.
	Name *string `json:"name,omitempty"`
	// Mac - MAC address of the interface.
	Mac *string `json:"mac,omitempty"`
	// IsPresent - Whether this interface is present on this controller.
	IsPresent *bool `json:"is_present,omitempty"`
	// LinkSpeed - Speed of the link.
	LinkSpeed *NsPlatLinkSpeed `json:"link_speed,omitempty"`
	// LinkStatus - Status of the link.
	LinkStatus *NsPlatLinkStatus `json:"link_status,omitempty"`
	// Mtu - MTU on the link.
	Mtu *int64 `json:"mtu,omitempty"`
	// Port - Port number for this interface.
	Port *int64 `json:"port,omitempty"`
	// Slot - Slot number for this interface.
	Slot *int64 `json:"slot,omitempty"`
	// MaxLinkSpeed - Maximum speed of the link.
	MaxLinkSpeed *NsPlatLinkSpeed `json:"max_link_speed,omitempty"`
	// NicType - Interface type.
	NicType *NsPlatNicType `json:"nic_type,omitempty"`
	// IpList - List of IP addresses assigned to this network interface.
	IpList []*NsAssignedIP `json:"ip_list,omitempty"`
}

// NetworkInterfaceFieldHandles provides a string representation for each AccessControlRecord field.
type NetworkInterfaceFieldHandles struct {
	ID                *string
	ArrayNameOrSerial *string
	PartialResponseOk *string
	ArrayId           *string
	ControllerName    *string
	ControllerId      *string
	Name              *string
	Mac               *string
	IsPresent         *string
	LinkSpeed         *string
	LinkStatus        *string
	Mtu               *string
	Port              *string
	Slot              *string
	MaxLinkSpeed      *string
	NicType           *string
	IpList            *string
}
