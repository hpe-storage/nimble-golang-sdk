// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// SubnetFields provides field names to use in filter parameters, for example.
var SubnetFields *SubnetFieldHandles

func init() {
	SubnetFields = &SubnetFieldHandles{
		ID:                 "id",
		Name:               "name",
		Network:            "network",
		Netmask:            "netmask",
		Type:               "type",
		AllowIscsi:         "allow_iscsi",
		AllowGroup:         "allow_group",
		DiscoveryIp:        "discovery_ip",
		Mtu:                "mtu",
		NetzoneType:        "netzone_type",
		VlanId:             "vlan_id",
		CreationTime:       "creation_time",
		LastModified:       "last_modified",
		Failover:           "failover",
		FailoverEnableTime: "failover_enable_time",
	}
}

// Subnet - Search subnets information. Many networking tasks require that objects such as replication partners are either on the same network or have a route to a secondary network. Subnets let you create logical addressing for selective routing.
type Subnet struct {
	// ID - Identifier for the initiator group.
	ID *string `json:"id,omitempty"`
	// Name - Name of subnet configuration.
	Name *string `json:"name,omitempty"`
	// Network - Subnet network address.
	Network *string `json:"network,omitempty"`
	// Netmask - Subnet netmask address.
	Netmask *string `json:"netmask,omitempty"`
	// Type - Subnet type. Options include 'mgmt', 'data', and 'mgmt,data'.
	Type *NsSubnetType `json:"type,omitempty"`
	// AllowIscsi - Subnet type.
	AllowIscsi *bool `json:"allow_iscsi,omitempty"`
	// AllowGroup - Subnet type.
	AllowGroup *bool `json:"allow_group,omitempty"`
	// DiscoveryIp - Subnet network address.
	DiscoveryIp *string `json:"discovery_ip,omitempty"`
	// Mtu - MTU for specified subnet. Valid MTU's are in the 512-16000 range.
	Mtu *int64 `json:"mtu,omitempty"`
	// NetzoneType - Specify Network Affinity Zone type for iSCSI enabled subnets. Valid types are Single, Bisect, and EvenOdd for iSCSI subnets.
	NetzoneType *NsNetZoneType `json:"netzone_type,omitempty"`
	// VlanId - VLAN ID for specified subnet. Valid ID's are in the 1-4094 range.
	VlanId *int64 `json:"vlan_id,omitempty"`
	// CreationTime - Time when this subnet configuration was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this subnet configuration was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
	// Failover - Failover setting of the subnet.
	Failover *bool `json:"failover,omitempty"`
	// FailoverEnableTime - Failover for this subnet will be enabled again at the time specified by failover_enable_time.
	FailoverEnableTime *int64 `json:"failover_enable_time,omitempty"`
}

// SubnetFieldHandles provides a string representation for each Subnet field.
type SubnetFieldHandles struct {
	ID                 string
	Name               string
	Network            string
	Netmask            string
	Type               string
	AllowIscsi         string
	AllowGroup         string
	DiscoveryIp        string
	Mtu                string
	NetzoneType        string
	VlanId             string
	CreationTime       string
	LastModified       string
	Failover           string
	FailoverEnableTime string
}
