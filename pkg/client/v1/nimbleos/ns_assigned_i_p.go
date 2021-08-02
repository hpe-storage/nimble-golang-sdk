// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAssignedIPFields provides field names to use in filter parameters, for example.
var NsAssignedIPFields *NsAssignedIPFieldHandles

func init() {
	NsAssignedIPFields = &NsAssignedIPFieldHandles{
		VlanId: "vlan_id",
		Ip:     "ip",
	}
}

// NsAssignedIP - IP address assignment details to network interface.
type NsAssignedIP struct {
	// VlanId - VLAN id of network interface.
	VlanId *int64 `json:"vlan_id,omitempty"`
	// Ip - Assigned IP address to network interface.
	Ip *string `json:"ip,omitempty"`
}

// NsAssignedIPFieldHandles provides a string representation for each AccessControlRecord field.
type NsAssignedIPFieldHandles struct {
	VlanId string
	Ip     string
}
