// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// AccessControlRecord - Manage access control records for volumes.
// Export AccessControlRecordFields for advance operations like search filter etc.
var AccessControlRecordFields *AccessControlRecord

func init() {
	IDfield := "id"
	ChapUserIdfield := "chap_user_id"
	ChapUserNamefield := "chap_user_name"
	InitiatorGroupIdfield := "initiator_group_id"
	InitiatorGroupNamefield := "initiator_group_name"
	VolIdfield := "vol_id"
	VolNamefield := "vol_name"
	PeIdfield := "pe_id"
	PeNamefield := "pe_name"
	SnapIdfield := "snap_id"
	SnapNamefield := "snap_name"

	AccessControlRecordFields = &AccessControlRecord{
		ID:                 &IDfield,
		ChapUserId:         &ChapUserIdfield,
		ChapUserName:       &ChapUserNamefield,
		InitiatorGroupId:   &InitiatorGroupIdfield,
		InitiatorGroupName: &InitiatorGroupNamefield,
		VolId:              &VolIdfield,
		VolName:            &VolNamefield,
		PeId:               &PeIdfield,
		PeName:             &PeNamefield,
		SnapId:             &SnapIdfield,
		SnapName:           &SnapNamefield,
	}
}

type AccessControlRecord struct {
	// ID - Identifier for the access control record.
	ID *string `json:"id,omitempty"`
	// ApplyTo - Type of object this access control record applies to.
	ApplyTo *NsAccessApplyTo `json:"apply_to,omitempty"`
	// ChapUserId - Identifier for the CHAP user.
	ChapUserId *string `json:"chap_user_id,omitempty"`
	// ChapUserName - Name of the CHAP user.
	ChapUserName *string `json:"chap_user_name,omitempty"`
	// InitiatorGroupId - Identifier for the initiator group.
	InitiatorGroupId *string `json:"initiator_group_id,omitempty"`
	// InitiatorGroupName - Name of the initiator group.
	InitiatorGroupName *string `json:"initiator_group_name,omitempty"`
	// Lun - If this access control record applies to a regular volume, this attribute is the volume's LUN (Logical Unit Number). If the access protocol is iSCSI, the LUN will be 0. However, if the access protocol is Fibre Channel, the LUN will be in the range from 0 to 2047. If this record applies to a Virtual Volume, this attribute is the volume's secondary LUN in the range from 0 to 399999, for both iSCSI and Fibre Channel. If the record applies to a OpenstackV2 volume, the LUN will be in the range from 0 to 2047, for both iSCSI and Fibre Channel. If this record applies to a protocol endpoint or only a snapshot, this attribute is not meaningful and is set to null.
	Lun *int64 `json:"lun,omitempty"`
	// VolId - Identifier for the volume this access control record applies to.
	VolId *string `json:"vol_id,omitempty"`
	// VolName - Name of the volume this access control record applies to.
	VolName *string `json:"vol_name,omitempty"`
	// VolAgentType - External management agent type.
	VolAgentType *NsAgentType `json:"vol_agent_type,omitempty"`
	// PeId - Identifier for the protocol endpoint this access control record applies to.
	PeId *string `json:"pe_id,omitempty"`
	// PeName - Name of the protocol endpoint this access control record applies to.
	PeName *string `json:"pe_name,omitempty"`
	// PeLun - LUN (Logical Unit Number) to associate with this protocol endpoint. Valid LUNs are in the 0-2047 range.
	PeLun *int64 `json:"pe_lun,omitempty"`
	// SnapId - Identifier for the snapshot this access control record applies to.
	SnapId *string `json:"snap_id,omitempty"`
	// SnapName - Name of the snapshot this access control record applies to.
	SnapName *string `json:"snap_name,omitempty"`
	// PeIds - List of candidate protocol endpoints that may be used to access the Virtual Volume. One of them will be selected for the access control record. This field is required only when creating an access control record for a Virtual Volume.
	PeIds []*string `json:"pe_ids,omitempty"`
	// Snapluns - Information about the snapshot LUNs associated with this access control record. This field is meaningful when the online snapshot can be accessed as a LUN in the group.
	Snapluns []*NsSnapLunInfo `json:"snapluns,omitempty"`
	// CreationTime - Time when this access control record was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this access control record was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
	// AccessProtocol - Access protocol of the volume.
	AccessProtocol *NsAccessProtocol `json:"access_protocol,omitempty"`
}
