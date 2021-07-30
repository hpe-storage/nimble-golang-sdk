// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAssignedIP - IP address assignment details to network interface.

// Export NsAssignedIPFields provides field names to use in filter parameters, for example.
var NsAssignedIPFields *NsAssignedIPStringFields

func init() {
	fieldVlanId := "vlan_id"
	fieldIp := "ip"

	NsAssignedIPFields = &NsAssignedIPStringFields{
		VlanId: &fieldVlanId,
		Ip:     &fieldIp,
	}
}

type NsAssignedIP struct {
	// VlanId - VLAN id of network interface.
	VlanId *int64 `json:"vlan_id,omitempty"`
	// Ip - Assigned IP address to network interface.
	Ip *string `json:"ip,omitempty"`
}

// Struct for NsAssignedIPFields
type NsAssignedIPStringFields struct {
	VlanId *string
	Ip     *string
}
