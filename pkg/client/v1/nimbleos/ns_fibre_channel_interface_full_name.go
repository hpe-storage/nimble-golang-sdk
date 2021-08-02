// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFibreChannelInterfaceFullNameFields provides field names to use in filter parameters, for example.
var NsFibreChannelInterfaceFullNameFields *NsFibreChannelInterfaceFullNameFieldHandles

func init() {
	NsFibreChannelInterfaceFullNameFields = &NsFibreChannelInterfaceFullNameFieldHandles{
		ArrayName: "array_name",
		CtrlrName: "ctrlr_name",
		IntfName:  "intf_name",
	}
}

// NsFibreChannelInterfaceFullName - Fully qualified name information for a Fibre Channel interface (array/controller/interface).
type NsFibreChannelInterfaceFullName struct {
	// ArrayName - Array name.
	ArrayName *string `json:"array_name,omitempty"`
	// CtrlrName - Controller name.
	CtrlrName *string `json:"ctrlr_name,omitempty"`
	// IntfName - Fibre Channel interface name.
	IntfName *string `json:"intf_name,omitempty"`
}

// NsFibreChannelInterfaceFullNameFieldHandles provides a string representation for each AccessControlRecord field.
type NsFibreChannelInterfaceFullNameFieldHandles struct {
	ArrayName string
	CtrlrName string
	IntfName  string
}
