// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAccessControlRecord - An access control record associated with a volume or snapshot or protocol endpoint.
// Export NsAccessControlRecordFields for advance operations like search filter etc.
var NsAccessControlRecordFields *NsAccessControlRecordStringFields

func init() {
	IDfield := "id"
	AclIdfield := "acl_id"
	InitiatorGroupIdfield := "initiator_group_id"
	InitiatorGroupNamefield := "initiator_group_name"
	ChapUserIdfield := "chap_user_id"
	ChapUserNamefield := "chap_user_name"
	ApplyTofield := "apply_to"
	Lunfield := "lun"
	AccessProtocolfield := "access_protocol"
	Snaplunsfield := "snapluns"
	PeIdfield := "pe_id"
	PeNamefield := "pe_name"
	PeLunfield := "pe_lun"
	VolIdfield := "vol_id"
	VolNamefield := "vol_name"
	SnapIdfield := "snap_id"
	SnapNamefield := "snap_name"

	NsAccessControlRecordFields = &NsAccessControlRecordStringFields{
		ID:                 &IDfield,
		AclId:              &AclIdfield,
		InitiatorGroupId:   &InitiatorGroupIdfield,
		InitiatorGroupName: &InitiatorGroupNamefield,
		ChapUserId:         &ChapUserIdfield,
		ChapUserName:       &ChapUserNamefield,
		ApplyTo:            &ApplyTofield,
		Lun:                &Lunfield,
		AccessProtocol:     &AccessProtocolfield,
		Snapluns:           &Snaplunsfield,
		PeId:               &PeIdfield,
		PeName:             &PeNamefield,
		PeLun:              &PeLunfield,
		VolId:              &VolIdfield,
		VolName:            &VolNamefield,
		SnapId:             &SnapIdfield,
		SnapName:           &SnapNamefield,
	}
}

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

// Struct for NsAccessControlRecordFields
type NsAccessControlRecordStringFields struct {
	ID                 *string
	AclId              *string
	InitiatorGroupId   *string
	InitiatorGroupName *string
	ChapUserId         *string
	ChapUserName       *string
	ApplyTo            *string
	Lun                *string
	AccessProtocol     *string
	Snapluns           *string
	PeId               *string
	PeName             *string
	PeLun              *string
	VolId              *string
	VolName            *string
	SnapId             *string
	SnapName           *string
}
