// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NetworkConfigFields provides field names to use in filter parameters, for example.
var NetworkConfigFields *NetworkConfigFieldHandles

func init() {
	NetworkConfigFields = &NetworkConfigFieldHandles{
		ID:                             "id",
		Name:                           "name",
		MgmtIp:                         "mgmt_ip",
		SecondaryMgmtIp:                "secondary_mgmt_ip",
		Role:                           "role",
		IscsiAutomaticConnectionMethod: "iscsi_automatic_connection_method",
		IscsiConnectionRebalancing:     "iscsi_connection_rebalancing",
		RouteList:                      "route_list",
		SubnetList:                     "subnet_list",
		ArrayList:                      "array_list",
		GroupLeaderArray:               "group_leader_array",
		CreationTime:                   "creation_time",
		LastModified:                   "last_modified",
		ActiveSince:                    "active_since",
		LastActive:                     "last_active",
		IgnoreValidationMask:           "ignore_validation_mask",
	}
}

// NetworkConfig - Manage group wide network configuration. The three possible network configurations include active, backup and an optional draft configuration.
type NetworkConfig struct {
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

// NetworkConfigFieldHandles provides a string representation for each AccessControlRecord field.
type NetworkConfigFieldHandles struct {
	ID                             string
	Name                           string
	MgmtIp                         string
	SecondaryMgmtIp                string
	Role                           string
	IscsiAutomaticConnectionMethod string
	IscsiConnectionRebalancing     string
	RouteList                      string
	SubnetList                     string
	ArrayList                      string
	GroupLeaderArray               string
	CreationTime                   string
	LastModified                   string
	ActiveSince                    string
	LastActive                     string
	IgnoreValidationMask           string
}
