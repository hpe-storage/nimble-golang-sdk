// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsCompleteSetupNetconfig - Information includes group, array and network configuration in complete setup.
// Export NsCompleteSetupNetconfigFields for advance operations like search filter etc.
var NsCompleteSetupNetconfigFields *NsCompleteSetupNetconfig

func init() {

	NsCompleteSetupNetconfigFields = &NsCompleteSetupNetconfig{
		DiscoveryIp:      "discovery_ip",
		DiscoveryNetmask: "discovery_netmask",
		MgmtIp:           "mgmt_ip",
		SecondaryMgmtIp:  "secondary_mgmt_ip",
		MgmtNetmask:      "mgmt_netmask",
	}
}

type NsCompleteSetupNetconfig struct {
	// Name - Network configuration name, possible values: 'active', 'backup', 'draft'.
	Name *NsNetConfigName `json:"name,omitempty"`
	// DiscoveryIp - The discovery IP address of the group.
	DiscoveryIp string `json:"discovery_ip,omitempty"`
	// DiscoveryNetmask - The netmask of the discovery IP address of the group.
	DiscoveryNetmask string `json:"discovery_netmask,omitempty"`
	// MgmtIp - The management IP address of the group.
	MgmtIp string `json:"mgmt_ip,omitempty"`
	// SecondaryMgmtIp - The secondary management IP address of the group.
	SecondaryMgmtIp string `json:"secondary_mgmt_ip,omitempty"`
	// MgmtNetmask - The netmask of the management IP address of the group.
	MgmtNetmask string `json:"mgmt_netmask,omitempty"`
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

// sdk internal struct
type nsCompleteSetupNetconfig struct {
	Name                           *NsNetConfigName `json:"name,omitempty"`
	DiscoveryIp                    *string          `json:"discovery_ip,omitempty"`
	DiscoveryNetmask               *string          `json:"discovery_netmask,omitempty"`
	MgmtIp                         *string          `json:"mgmt_ip,omitempty"`
	SecondaryMgmtIp                *string          `json:"secondary_mgmt_ip,omitempty"`
	MgmtNetmask                    *string          `json:"mgmt_netmask,omitempty"`
	IscsiAutomaticConnectionMethod *bool            `json:"iscsi_automatic_connection_method,omitempty"`
	IscsiConnectionRebalancing     *bool            `json:"iscsi_connection_rebalancing,omitempty"`
	RouteList                      []*NsRoute       `json:"route_list,omitempty"`
	SubnetList                     []*NsSubnet      `json:"subnet_list,omitempty"`
	ArrayList                      []*NsArrayNet    `json:"array_list,omitempty"`
}

// EncodeNsCompleteSetupNetconfig - Transform NsCompleteSetupNetconfig to nsCompleteSetupNetconfig type
func EncodeNsCompleteSetupNetconfig(request interface{}) (*nsCompleteSetupNetconfig, error) {
	reqNsCompleteSetupNetconfig := request.(*NsCompleteSetupNetconfig)
	byte, err := json.Marshal(reqNsCompleteSetupNetconfig)
	if err != nil {
		return nil, err
	}
	respNsCompleteSetupNetconfigPtr := &nsCompleteSetupNetconfig{}
	err = json.Unmarshal(byte, respNsCompleteSetupNetconfigPtr)
	return respNsCompleteSetupNetconfigPtr, err
}

// DecodeNsCompleteSetupNetconfig Transform nsCompleteSetupNetconfig to NsCompleteSetupNetconfig type
func DecodeNsCompleteSetupNetconfig(request interface{}) (*NsCompleteSetupNetconfig, error) {
	reqNsCompleteSetupNetconfig := request.(*nsCompleteSetupNetconfig)
	byte, err := json.Marshal(reqNsCompleteSetupNetconfig)
	if err != nil {
		return nil, err
	}
	respNsCompleteSetupNetconfig := &NsCompleteSetupNetconfig{}
	err = json.Unmarshal(byte, respNsCompleteSetupNetconfig)
	return respNsCompleteSetupNetconfig, err
}
