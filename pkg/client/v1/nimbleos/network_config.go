// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NetworkConfig - Manage group wide network configuration. The three possible network configurations include active, backup and an optional draft configuration.
// Export NetworkConfigFields for advance operations like search filter etc.
var NetworkConfigFields *NetworkConfig

func init() {

	NetworkConfigFields = &NetworkConfig{
		ID:               "id",
		MgmtIp:           "mgmt_ip",
		SecondaryMgmtIp:  "secondary_mgmt_ip",
		GroupLeaderArray: "group_leader_array",
	}
}

type NetworkConfig struct {
	// ID - Identifier for network configuration.
	ID string `json:"id,omitempty"`
	// Name - Name of the network configuration. Use the name 'draft' when creating a draft configuration.
	Name *NsNetConfigName `json:"name,omitempty"`
	// MgmtIp - Management IP address for the Group.
	MgmtIp string `json:"mgmt_ip,omitempty"`
	// SecondaryMgmtIp - Secondary management IP address for the Group.
	SecondaryMgmtIp string `json:"secondary_mgmt_ip,omitempty"`
	// Role - Role of network configuration.
	Role *NsNetConfigRole `json:"role,omitempty"`
	// IscsiAutomaticConnectionMethod - Whether automatic connection method is enabled. Enabling this means means redirecting connections from the specified iSCSI discovery IP address to the best data IP address based on connection counts.
	IscsiAutomaticConnectionMethod *bool `json:"iscsi_automatic_connection_method,omitempty"`
	// IscsiConnectionRebalancing - Whether rebalancing is enabled. Enabling this means rebalancing iSCSI connections by periodically breaking existing connections that are out-of-balance, allowing the host to reconnect to a more appropriate data IP address.
	IscsiConnectionRebalancing *bool `json:"iscsi_connection_rebalancing,omitempty"`
	// RouteList - List of static routes.
	RouteList []*NsRoute `json:"route_list,omitempty"`
	// SubnetList - List of subnet configs.
	SubnetList []*NsSubnet `json:"subnet_list,omitempty"`
	// ArrayList - List of array network configs.
	ArrayList []*NsArrayNet `json:"array_list,omitempty"`
	// GroupLeaderArray - Name of the group leader array.
	GroupLeaderArray string `json:"group_leader_array,omitempty"`
	// CreationTime - Time when this net configuration was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this network configuration was last modified.
	LastModified int64 `json:"last_modified,omitempty"`
	// ActiveSince - Start time of activity.
	ActiveSince int64 `json:"active_since,omitempty"`
	// LastActive - Time of last activity.
	LastActive int64 `json:"last_active,omitempty"`
	// IgnoreValidationMask - Indicates whether to ignore the validation.
	IgnoreValidationMask int64 `json:"ignore_validation_mask,omitempty"`
}

// sdk internal struct
type networkConfig struct {
	// ID - Identifier for network configuration.
	ID *string `json:"id,omitempty"`
	// Name - Name of the network configuration. Use the name 'draft' when creating a draft configuration.
	Name *NsNetConfigName `json:"name,omitempty"`
	// MgmtIp - Management IP address for the Group.
	MgmtIp *string `json:"mgmt_ip,omitempty"`
	// SecondaryMgmtIp - Secondary management IP address for the Group.
	SecondaryMgmtIp *string `json:"secondary_mgmt_ip,omitempty"`
	// Role - Role of network configuration.
	Role *NsNetConfigRole `json:"role,omitempty"`
	// IscsiAutomaticConnectionMethod - Whether automatic connection method is enabled. Enabling this means means redirecting connections from the specified iSCSI discovery IP address to the best data IP address based on connection counts.
	IscsiAutomaticConnectionMethod *bool `json:"iscsi_automatic_connection_method,omitempty"`
	// IscsiConnectionRebalancing - Whether rebalancing is enabled. Enabling this means rebalancing iSCSI connections by periodically breaking existing connections that are out-of-balance, allowing the host to reconnect to a more appropriate data IP address.
	IscsiConnectionRebalancing *bool `json:"iscsi_connection_rebalancing,omitempty"`
	// RouteList - List of static routes.
	RouteList []*NsRoute `json:"route_list,omitempty"`
	// SubnetList - List of subnet configs.
	SubnetList []*NsSubnet `json:"subnet_list,omitempty"`
	// ArrayList - List of array network configs.
	ArrayList []*NsArrayNet `json:"array_list,omitempty"`
	// GroupLeaderArray - Name of the group leader array.
	GroupLeaderArray *string `json:"group_leader_array,omitempty"`
	// CreationTime - Time when this net configuration was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this network configuration was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
	// ActiveSince - Start time of activity.
	ActiveSince *int64 `json:"active_since,omitempty"`
	// LastActive - Time of last activity.
	LastActive *int64 `json:"last_active,omitempty"`
	// IgnoreValidationMask - Indicates whether to ignore the validation.
	IgnoreValidationMask *int64 `json:"ignore_validation_mask,omitempty"`
}

// EncodeNetworkConfig - Transform NetworkConfig to networkConfig type
func EncodeNetworkConfig(request interface{}) (*networkConfig, error) {
	reqNetworkConfig := request.(*NetworkConfig)
	byte, err := json.Marshal(reqNetworkConfig)
	objPtr := &networkConfig{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNetworkConfig Transform networkConfig to NetworkConfig type
func DecodeNetworkConfig(request interface{}) (*NetworkConfig, error) {
	reqNetworkConfig := request.(*networkConfig)
	byte, err := json.Marshal(reqNetworkConfig)
	obj := &NetworkConfig{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
