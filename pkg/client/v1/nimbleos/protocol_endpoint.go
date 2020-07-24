// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// ProtocolEndpoint - Protocol endpoints are administrative logical units (LUs) in an LU conglomerate to be used with VMware Virtual Volumes.
// Export ProtocolEndpointFields for advance operations like search filter etc.
var ProtocolEndpointFields *ProtocolEndpoint

func init() {
	IDfield := "id"
	Namefield := "name"
	Descriptionfield := "description"
	PoolNamefield := "pool_name"
	PoolIdfield := "pool_id"
	SerialNumberfield := "serial_number"
	TargetNamefield := "target_name"

	ProtocolEndpointFields = &ProtocolEndpoint{
		ID:           &IDfield,
		Name:         &Namefield,
		Description:  &Descriptionfield,
		PoolName:     &PoolNamefield,
		PoolId:       &PoolIdfield,
		SerialNumber: &SerialNumberfield,
		TargetName:   &TargetNamefield,
	}
}

type ProtocolEndpoint struct {
	// ID - Identifier for the protocol endpoint.
	ID *string `json:"id,omitempty"`
	// Name - Name of the protocol endpoint.
	Name *string `json:"name,omitempty"`
	// Description - Text description of the protocol endpoint.
	Description *string `json:"description,omitempty"`
	// PoolName - Name of the pool where the protocol endpoint resides. If pool option is not specified, protocol endpoint is assigned to the default pool.
	PoolName *string `json:"pool_name,omitempty"`
	// PoolId - Identifier associated with the pool in the storage pool table.
	PoolId *string `json:"pool_id,omitempty"`
	// State - Operational state of protocol endpoint.
	State *NsPEOpStateType `json:"state,omitempty"`
	// SerialNumber - Identifier associated with the protocol endpoint for the SCSI protocol.
	SerialNumber *string `json:"serial_number,omitempty"`
	// TargetName - The iSCSI Qualified Name (IQN) or the Fibre Channel World Wide Node Name (WWNN) of the target protocol endpoint.
	TargetName *string `json:"target_name,omitempty"`
	// GroupSpecificIds - External UID is used to compute the serial number and IQN which never change even if the running group changes (e.g. after group merge). Group-specific IDs determine whether external UID is used for computing serial number and IQN.
	GroupSpecificIds *bool `json:"group_specific_ids,omitempty"`
	// CreationTime - Time when this protocol endpoint was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this protocol endpoint was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
	// NumConnections - Number of connections via this protocol endpoint.
	NumConnections *int64 `json:"num_connections,omitempty"`
	// NumIscsiConnections - Number of iSCSI connections via this protocol endpoint.
	NumIscsiConnections *int64 `json:"num_iscsi_connections,omitempty"`
	// NumFcConnections - Number of FC connections via this protocol endpoint.
	NumFcConnections *int64 `json:"num_fc_connections,omitempty"`
	// AccessControlRecords - List of access control records that apply to this protocol endpoint.
	AccessControlRecords []*NsAccessControlRecord `json:"access_control_records,omitempty"`
	// IscsiSessions - List of iSCSI sessions connected to this protocol endpoint.
	IscsiSessions []*NsISCSISession `json:"iscsi_sessions,omitempty"`
	// FcSessions - List of FC sessions connected to this protocol endpoint.
	FcSessions []*NsFCSession `json:"fc_sessions,omitempty"`
	// AccessProtocol - Access protocol of the protocol endpoint. Only initiator groups with the same access protocol can access the protocol endpoint. If not specified in the creation request, it will be the access protocol supported by the group. If the group supports multiple protocols, the default will be Fibre Channel.
	AccessProtocol *NsAccessProtocol `json:"access_protocol,omitempty"`
}
