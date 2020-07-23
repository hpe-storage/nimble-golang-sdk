// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// InitiatorGroup - Manage initiator groups for initiator authentication. An initiator group is a set of initiators that can be configured as part of your ACL to access a specific volume through group membership.
// Export InitiatorGroupFields for advance operations like search filter etc.
var InitiatorGroupFields *InitiatorGroup

func init() {
	IDfield := "id"
	Namefield := "name"
	FullNamefield := "full_name"
	SearchNamefield := "search_name"
	Descriptionfield := "description"
	HostTypefield := "host_type"
	AppUuidfield := "app_uuid"

	InitiatorGroupFields = &InitiatorGroup{
		ID:          &IDfield,
		Name:        &Namefield,
		FullName:    &FullNamefield,
		SearchName:  &SearchNamefield,
		Description: &Descriptionfield,
		HostType:    &HostTypefield,
		AppUuid:     &AppUuidfield,
	}
}

type InitiatorGroup struct {
	// ID - Identifier for initiator group.
	ID *string `json:"id,omitempty"`
	// Name - Name of initiator group.
	Name *string `json:"name,omitempty"`
	// FullName - Initiator group's full name.
	FullName *string `json:"full_name,omitempty"`
	// SearchName - Initiator group name used for search.
	SearchName *string `json:"search_name,omitempty"`
	// Description - Text description of initiator group.
	Description *string `json:"description,omitempty"`
	// AccessProtocol - Initiator group access protocol.
	AccessProtocol *NsAccessProtocol `json:"access_protocol,omitempty"`
	// HostType - Initiator group host type. Available options are auto and hpux. The default option is auto. This attribute will be applied to all the initiators in the initiator group. Initiators with different host OSes should not be kept in the same initiator group having a non-default host type attribute.
	HostType *string `json:"host_type,omitempty"`
	// FcTdzPorts - List of target Fibre Channel ports with Target Driven Zoning configured on this initiator group.
	FcTdzPorts []*NsFcTdzPort `json:"fc_tdz_ports,omitempty"`
	// TargetSubnets - List of target subnet labels. If specified, discovery and access to volumes will be restricted to the specified subnets.
	TargetSubnets []*NsTargetSubnet `json:"target_subnets,omitempty"`
	// IscsiInitiators - List of iSCSI initiators. When create/update iscsi_initiators, either iqn or ip_address is always required with label.
	IscsiInitiators []*NsISCSIInitiator `json:"iscsi_initiators,omitempty"`
	// FcInitiators - List of FC initiators. When create/update fc_initiators, wwpn is required.
	FcInitiators []*NsFCInitiator `json:"fc_initiators,omitempty"`
	// CreationTime - Time when this initiator group was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this initiator group was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
	// AppUuid - Application identifier of initiator group.
	AppUuid *string `json:"app_uuid,omitempty"`
	// VolumeCount - Number of volumes that are accessible by the initiator group.
	VolumeCount *int64 `json:"volume_count,omitempty"`
	// VolumeList - List of volumes that are accessible by the initiator group.
	VolumeList []*NsVolumeSummaryWithAppCategory `json:"volume_list,omitempty"`
	// NumConnections - Total number of connections from initiators in the initiator group.
	NumConnections *int64 `json:"num_connections,omitempty"`
	// Metadata - Key-value pairs that augment an initiator group's attributes.
	Metadata []*NsKeyValue `json:"metadata,omitempty"`
}
