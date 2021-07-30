// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsCompleteSetupNetconfig - Information includes group, array and network configuration in complete setup.

// Export NsCompleteSetupNetconfigFields provides field names to use in filter parameters, for example.
var NsCompleteSetupNetconfigFields *NsCompleteSetupNetconfigStringFields

func init() {
	fieldName := "name"
	fieldDiscoveryIp := "discovery_ip"
	fieldDiscoveryNetmask := "discovery_netmask"
	fieldMgmtIp := "mgmt_ip"
	fieldSecondaryMgmtIp := "secondary_mgmt_ip"
	fieldMgmtNetmask := "mgmt_netmask"
	fieldIscsiAutomaticConnectionMethod := "iscsi_automatic_connection_method"
	fieldIscsiConnectionRebalancing := "iscsi_connection_rebalancing"
	fieldRouteList := "route_list"
	fieldSubnetList := "subnet_list"
	fieldArrayList := "array_list"

	NsCompleteSetupNetconfigFields = &NsCompleteSetupNetconfigStringFields{
		Name:                           &fieldName,
		DiscoveryIp:                    &fieldDiscoveryIp,
		DiscoveryNetmask:               &fieldDiscoveryNetmask,
		MgmtIp:                         &fieldMgmtIp,
		SecondaryMgmtIp:                &fieldSecondaryMgmtIp,
		MgmtNetmask:                    &fieldMgmtNetmask,
		IscsiAutomaticConnectionMethod: &fieldIscsiAutomaticConnectionMethod,
		IscsiConnectionRebalancing:     &fieldIscsiConnectionRebalancing,
		RouteList:                      &fieldRouteList,
		SubnetList:                     &fieldSubnetList,
		ArrayList:                      &fieldArrayList,
	}
}

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

// Struct for NsCompleteSetupNetconfigFields
type NsCompleteSetupNetconfigStringFields struct {
	Name                           *string
	DiscoveryIp                    *string
	DiscoveryNetmask               *string
	MgmtIp                         *string
	SecondaryMgmtIp                *string
	MgmtNetmask                    *string
	IscsiAutomaticConnectionMethod *string
	IscsiConnectionRebalancing     *string
	RouteList                      *string
	SubnetList                     *string
	ArrayList                      *string
}
