// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsAccessControlRecordFields provides field names to use in filter parameters, for example.
var NsAccessControlRecordFields *NsAccessControlRecordFieldHandles

func init() {
	NsAccessControlRecordFields = &NsAccessControlRecordFieldHandles{
		ID:                 "id",
		AclId:              "acl_id",
		InitiatorGroupId:   "initiator_group_id",
		InitiatorGroupName: "initiator_group_name",
		ChapUserId:         "chap_user_id",
		ChapUserName:       "chap_user_name",
		ApplyTo:            "apply_to",
		Lun:                "lun",
		AccessProtocol:     "access_protocol",
		Snapluns:           "snapluns",
		PeId:               "pe_id",
		PeName:             "pe_name",
		PeLun:              "pe_lun",
		VolId:              "vol_id",
		VolName:            "vol_name",
		SnapId:             "snap_id",
		SnapName:           "snap_name",
	}
}

// NsAccessControlRecord - An access control record associated with a volume or snapshot or protocol endpoint.
type NsAccessControlRecord struct {
	// ID - ID of access control record.
	ID *string `json:"id,omitempty"`
	// AclId - ID of access control record.
	AclId *string `json:"acl_id,omitempty"`
	// InitiatorGroupId - ID of initiator group.
	InitiatorGroupId *string `json:"initiator_group_id,omitempty"`
	// InitiatorGroupName - Name of initiator group.
	InitiatorGroupName *string `json:"initiator_group_name,omitempty"`
	// ChapUserId - ID of chap user.
	ChapUserId *string `json:"chap_user_id,omitempty"`
	// ChapUserName - Name of chap user.
	ChapUserName *string `json:"chap_user_name,omitempty"`
	// ApplyTo - State of apply to.
	ApplyTo *NsAccessApplyTo `json:"apply_to,omitempty"`
	// Lun - LUN of snapshot or volume. Secondary LUN if this is Virtual Volume.
	Lun *int64 `json:"lun,omitempty"`
	// AccessProtocol - Access protocol of snapshot or volume.
	AccessProtocol *NsAccessProtocol `json:"access_protocol,omitempty"`
	// Snapluns - Information about the snapshot LUNs. This information is available only for Fibre Channel.
	Snapluns []*NsSnapLunInfo `json:"snapluns,omitempty"`
	// PeId - ID of protocol endpoint.
	PeId *string `json:"pe_id,omitempty"`
	// PeName - Name of protocol endpoint.
	PeName *string `json:"pe_name,omitempty"`
	// PeLun - LUN of protocol endpoint.
	PeLun *int64 `json:"pe_lun,omitempty"`
	// VolId - ID of volume.
	VolId *string `json:"vol_id,omitempty"`
	// VolName - Name of volume.
	VolName *string `json:"vol_name,omitempty"`
	// SnapId - ID of snapshot.
	SnapId *string `json:"snap_id,omitempty"`
	// SnapName - Name of snapshot.
	SnapName *string `json:"snap_name,omitempty"`
}

// NsAccessControlRecordFieldHandles provides a string representation for each AccessControlRecord field.
type NsAccessControlRecordFieldHandles struct {
	ID                 string
	AclId              string
	InitiatorGroupId   string
	InitiatorGroupName string
	ChapUserId         string
	ChapUserName       string
	ApplyTo            string
	Lun                string
	AccessProtocol     string
	Snapluns           string
	PeId               string
	PeName             string
	PeLun              string
	VolId              string
	VolName            string
	SnapId             string
	SnapName           string
}
