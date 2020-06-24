// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsAccessControlRecord - An access control record associated with a volume or snapshot or protocol endpoint.
// Export NsAccessControlRecordFields for advance operations like search filter etc.
var NsAccessControlRecordFields *NsAccessControlRecord

func init(){
	IDfield:= "id"
	AclIdfield:= "acl_id"
	InitiatorGroupIdfield:= "initiator_group_id"
	InitiatorGroupNamefield:= "initiator_group_name"
	ChapUserIdfield:= "chap_user_id"
	ChapUserNamefield:= "chap_user_name"
	PeIdfield:= "pe_id"
	PeNamefield:= "pe_name"
	VolIdfield:= "vol_id"
	VolNamefield:= "vol_name"
	SnapIdfield:= "snap_id"
	SnapNamefield:= "snap_name"
		
	NsAccessControlRecordFields= &NsAccessControlRecord{
	ID: &IDfield,
	AclId: &AclIdfield,
	InitiatorGroupId: &InitiatorGroupIdfield,
	InitiatorGroupName: &InitiatorGroupNamefield,
	ChapUserId: &ChapUserIdfield,
	ChapUserName: &ChapUserNamefield,
	PeId: &PeIdfield,
	PeName: &PeNamefield,
	VolId: &VolIdfield,
	VolName: &VolNamefield,
	SnapId: &SnapIdfield,
	SnapName: &SnapNamefield,
		
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
