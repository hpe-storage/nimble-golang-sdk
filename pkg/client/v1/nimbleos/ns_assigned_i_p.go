// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsAssignedIPFields provides field names to use in filter parameters, for example.
var NsAssignedIPFields *NsAssignedIPFieldHandles

func init() {
	fieldVlanId := "vlan_id"
	fieldIp := "ip"

	NsAssignedIPFields = &NsAssignedIPFieldHandles{
		VlanId: &fieldVlanId,
		Ip:     &fieldIp,
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
	VlanId *string
	Ip     *string
}
