// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsCompleteSetupNetconfigFields provides field names to use in filter parameters, for example.
var NsCompleteSetupNetconfigFields *NsCompleteSetupNetconfigFieldHandles

func init() {
	NsCompleteSetupNetconfigFields = &NsCompleteSetupNetconfigFieldHandles{
		Name:                           "name",
		DiscoveryIp:                    "discovery_ip",
		DiscoveryNetmask:               "discovery_netmask",
		MgmtIp:                         "mgmt_ip",
		SecondaryMgmtIp:                "secondary_mgmt_ip",
		MgmtNetmask:                    "mgmt_netmask",
		IscsiAutomaticConnectionMethod: "iscsi_automatic_connection_method",
		IscsiConnectionRebalancing:     "iscsi_connection_rebalancing",
		RouteList:                      "route_list",
		SubnetList:                     "subnet_list",
		ArrayList:                      "array_list",
	}
}

// NsCompleteSetupNetconfig - Information includes group, array and network configuration in complete setup.
type NsCompleteSetupNetconfig struct {
	// Name - Network configuration name, possible values: 'active', 'backup', 'draft'.
	Name *NsNetConfigName `json:"name,omitempty"`
	// DiscoveryIp - The discovery IP address of the group.
	DiscoveryIp *string `json:"discovery_ip,omitempty"`
	// DiscoveryNetmask - The netmask of the discovery IP address of the group.
	DiscoveryNetmask *string `json:"discovery_netmask,omitempty"`
	// MgmtIp - The management IP address of the group.
	MgmtIp *string `json:"mgmt_ip,omitempty"`
	// SecondaryMgmtIp - The secondary management IP address of the group.
	SecondaryMgmtIp *string `json:"secondary_mgmt_ip,omitempty"`
	// MgmtNetmask - The netmask of the management IP address of the group.
	MgmtNetmask *string `json:"mgmt_netmask,omitempty"`
	// IscsiAutomaticConnectionMethod - Iscsi has atomatic connection method.
	IscsiAutomaticConnectionMethod *bool `json:"iscsi_automatic_connection_method,omitempty"`
	// IscsiConnectionRebalancing - Iscsi connection rebalanced.
	IscsiConnectionRebalancing *bool `json:"iscsi_connection_rebalancing,omitempty"`
	// RouteList - Network route list.
	RouteList []*NsRoute `json:"route_list,omitempty"`
	// SubnetList - Subnet list.
	SubnetList []*NsSubnet `json:"subnet_list,omitempty"`
	// ArrayList - Array list.
	ArrayList []*NsArrayNet `json:"array_list,omitempty"`
}

// NsCompleteSetupNetconfigFieldHandles provides a string representation for each AccessControlRecord field.
type NsCompleteSetupNetconfigFieldHandles struct {
	Name                           string
	DiscoveryIp                    string
	DiscoveryNetmask               string
	MgmtIp                         string
	SecondaryMgmtIp                string
	MgmtNetmask                    string
	IscsiAutomaticConnectionMethod string
	IscsiConnectionRebalancing     string
	RouteList                      string
	SubnetList                     string
	ArrayList                      string
}
