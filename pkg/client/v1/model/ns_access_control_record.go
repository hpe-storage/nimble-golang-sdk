// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsAccessControlRecord - An access control record associated with a volume or snapshot or protocol endpoint.
// Export NsAccessControlRecordFields for advance operations like search filter etc.
var NsAccessControlRecordFields *NsAccessControlRecord

func init() {

	NsAccessControlRecordFields = &NsAccessControlRecord{
		ID:                 "id",
		AclId:              "acl_id",
		InitiatorGroupId:   "initiator_group_id",
		InitiatorGroupName: "initiator_group_name",
		ChapUserId:         "chap_user_id",
		ChapUserName:       "chap_user_name",
		PeId:               "pe_id",
		PeName:             "pe_name",
		VolId:              "vol_id",
		VolName:            "vol_name",
		SnapId:             "snap_id",
		SnapName:           "snap_name",
	}
}

type NsAccessControlRecord struct {
	// ID - ID of access control record.
	ID string `json:"id,omitempty"`
	// AclId - ID of access control record.
	AclId string `json:"acl_id,omitempty"`
	// InitiatorGroupId - ID of initiator group.
	InitiatorGroupId string `json:"initiator_group_id,omitempty"`
	// InitiatorGroupName - Name of initiator group.
	InitiatorGroupName string `json:"initiator_group_name,omitempty"`
	// ChapUserId - ID of chap user.
	ChapUserId string `json:"chap_user_id,omitempty"`
	// ChapUserName - Name of chap user.
	ChapUserName string `json:"chap_user_name,omitempty"`
	// ApplyTo - State of apply to.
	ApplyTo *NsAccessApplyTo `json:"apply_to,omitempty"`
	// Lun - LUN of snapshot or volume. Secondary LUN if this is Virtual Volume.
	Lun int64 `json:"lun,omitempty"`
	// AccessProtocol - Access protocol of snapshot or volume.
	AccessProtocol *NsAccessProtocol `json:"access_protocol,omitempty"`
	// Snapluns - Information about the snapshot LUNs. This information is available only for Fibre Channel.
	Snapluns []*NsSnapLunInfo `json:"snapluns,omitempty"`
	// PeId - ID of protocol endpoint.
	PeId string `json:"pe_id,omitempty"`
	// PeName - Name of protocol endpoint.
	PeName string `json:"pe_name,omitempty"`
	// PeLun - LUN of protocol endpoint.
	PeLun int64 `json:"pe_lun,omitempty"`
	// VolId - ID of volume.
	VolId string `json:"vol_id,omitempty"`
	// VolName - Name of volume.
	VolName string `json:"vol_name,omitempty"`
	// SnapId - ID of snapshot.
	SnapId string `json:"snap_id,omitempty"`
	// SnapName - Name of snapshot.
	SnapName string `json:"snap_name,omitempty"`
}

// sdk internal struct
type nsAccessControlRecord struct {
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

// EncodeNsAccessControlRecord - Transform NsAccessControlRecord to nsAccessControlRecord type
func EncodeNsAccessControlRecord(request interface{}) (*nsAccessControlRecord, error) {
	reqNsAccessControlRecord := request.(*NsAccessControlRecord)
	byte, err := json.Marshal(reqNsAccessControlRecord)
	objPtr := &nsAccessControlRecord{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsAccessControlRecord Transform nsAccessControlRecord to NsAccessControlRecord type
func DecodeNsAccessControlRecord(request interface{}) (*NsAccessControlRecord, error) {
	reqNsAccessControlRecord := request.(*nsAccessControlRecord)
	byte, err := json.Marshal(reqNsAccessControlRecord)
	obj := &NsAccessControlRecord{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
