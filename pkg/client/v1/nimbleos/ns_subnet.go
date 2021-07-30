// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsSubnetFields provides field names to use in filter parameters, for example.
var NsSubnetFields *NsSubnetFieldHandles

func init() {
	fieldLabel := "label"
	fieldNetwork := "network"
	fieldNetmask := "netmask"
	fieldNetzoneType := "netzone_type"
	fieldType := "type"
	fieldAllowIscsi := "allow_iscsi"
	fieldAllowGroup := "allow_group"
	fieldDiscoveryIp := "discovery_ip"
	fieldMtu := "mtu"
	fieldVlanId := "vlan_id"
	fieldFailover := "failover"
	fieldFailoverEnableTime := "failover_enable_time"

	NsSubnetFields = &NsSubnetFieldHandles{
		Label:              &fieldLabel,
		Network:            &fieldNetwork,
		Netmask:            &fieldNetmask,
		NetzoneType:        &fieldNetzoneType,
		Type:               &fieldType,
		AllowIscsi:         &fieldAllowIscsi,
		AllowGroup:         &fieldAllowGroup,
		DiscoveryIp:        &fieldDiscoveryIp,
		Mtu:                &fieldMtu,
		VlanId:             &fieldVlanId,
		Failover:           &fieldFailover,
		FailoverEnableTime: &fieldFailoverEnableTime,
	}
}

// NsSubnet - A subnet configuration.
type NsSubnet struct {
	// Label - Subnet label.
	Label *string `json:"label,omitempty"`
	// Network - Network IP address.
	Network *string `json:"network,omitempty"`
	// Netmask - Subnet netmask address.
	Netmask *string `json:"netmask,omitempty"`
	// NetzoneType - Netzone type.
	NetzoneType *NsNetZoneType `json:"netzone_type,omitempty"`
	// Type - Subnet type.
	Type *NsSubnetType `json:"type,omitempty"`
	// AllowIscsi - Allow iSCSI.
	AllowIscsi *bool `json:"allow_iscsi,omitempty"`
	// AllowGroup - Allow group.
	AllowGroup *bool `json:"allow_group,omitempty"`
	// DiscoveryIp - Discovery IP address.
	DiscoveryIp *string `json:"discovery_ip,omitempty"`
	// Mtu - MTU for specified subnet.
	Mtu *int64 `json:"mtu,omitempty"`
	// VlanId - VLAN ID for specified subnet.
	VlanId *int64 `json:"vlan_id,omitempty"`
	// Failover - Failover setting of the subnet.
	Failover *bool `json:"failover,omitempty"`
	// FailoverEnableTime - Failover for this subnet will be enabled again at the time specified by failover_enable_time.
	FailoverEnableTime *int64 `json:"failover_enable_time,omitempty"`
}

// NsSubnetFieldHandles provides a string representation for each AccessControlRecord field.
type NsSubnetFieldHandles struct {
	Label              *string
	Network            *string
	Netmask            *string
	NetzoneType        *string
	Type               *string
	AllowIscsi         *string
	AllowGroup         *string
	DiscoveryIp        *string
	Mtu                *string
	VlanId             *string
	Failover           *string
	FailoverEnableTime *string
}
