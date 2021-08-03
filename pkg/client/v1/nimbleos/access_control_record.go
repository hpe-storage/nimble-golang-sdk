// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// AccessControlRecordFields provides field names to use in filter parameters, for example.
var AccessControlRecordFields *AccessControlRecordFieldHandles

func init() {
	AccessControlRecordFields = &AccessControlRecordFieldHandles{
		ID:                 "id",
		ApplyTo:            "apply_to",
		ChapUserId:         "chap_user_id",
		ChapUserName:       "chap_user_name",
		InitiatorGroupId:   "initiator_group_id",
		InitiatorGroupName: "initiator_group_name",
		Lun:                "lun",
		VolId:              "vol_id",
		VolName:            "vol_name",
		VolAgentType:       "vol_agent_type",
		PeId:               "pe_id",
		PeName:             "pe_name",
		PeLun:              "pe_lun",
		SnapId:             "snap_id",
		SnapName:           "snap_name",
		PeIds:              "pe_ids",
		Snapluns:           "snapluns",
		CreationTime:       "creation_time",
		LastModified:       "last_modified",
		AccessProtocol:     "access_protocol",
	}
}

// AccessControlRecord - Manage access control records for volumes.
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

// AccessControlRecordFieldHandles provides a string representation for each AccessControlRecord field.
type AccessControlRecordFieldHandles struct {
	ID                 string
	ApplyTo            string
	ChapUserId         string
	ChapUserName       string
	InitiatorGroupId   string
	InitiatorGroupName string
	Lun                string
	VolId              string
	VolName            string
	VolAgentType       string
	PeId               string
	PeName             string
	PeLun              string
	SnapId             string
	SnapName           string
	PeIds              string
	Snapluns           string
	CreationTime       string
	LastModified       string
	AccessProtocol     string
}
