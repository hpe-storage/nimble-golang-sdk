/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// AccessControlRecord - Manage access control records for volumes.
// Export AccessControlRecordFields for advance operations like search filter etc.
var AccessControlRecordFields *AccessControlRecord

func init(){
	IDfield:= "id"
	ChapUserIDfield:= "chap_user_id"
	ChapUserNamefield:= "chap_user_name"
	InitiatorGroupIDfield:= "initiator_group_id"
	InitiatorGroupNamefield:= "initiator_group_name"
	VolIDfield:= "vol_id"
	VolNamefield:= "vol_name"
	PeIDfield:= "pe_id"
	PeNamefield:= "pe_name"
	SnapIDfield:= "snap_id"
	SnapNamefield:= "snap_name"
		
	AccessControlRecordFields= &AccessControlRecord{
		ID: &IDfield,
		ChapUserID: &ChapUserIDfield,
		ChapUserName: &ChapUserNamefield,
		InitiatorGroupID: &InitiatorGroupIDfield,
		InitiatorGroupName: &InitiatorGroupNamefield,
		VolID: &VolIDfield,
		VolName: &VolNamefield,
		PeID: &PeIDfield,
		PeName: &PeNamefield,
		SnapID: &SnapIDfield,
		SnapName: &SnapNamefield,
		
	}
}

type AccessControlRecord struct {
   
    // Identifier for the access control record.
    
 	ID *string `json:"id,omitempty"`
   
    // Type of object this access control record applies to.
    
   	ApplyTo *NsAccessApplyTo `json:"apply_to,omitempty"`
   
    // Identifier for the CHAP user.
    
 	ChapUserID *string `json:"chap_user_id,omitempty"`
   
    // Name of the CHAP user.
    
 	ChapUserName *string `json:"chap_user_name,omitempty"`
   
    // Identifier for the initiator group.
    
 	InitiatorGroupID *string `json:"initiator_group_id,omitempty"`
   
    // Name of the initiator group.
    
 	InitiatorGroupName *string `json:"initiator_group_name,omitempty"`
   
    // If this access control record applies to a regular volume, this attribute is the volume's LUN (Logical Unit Number). If the access protocol is iSCSI, the LUN will be 0. However, if the access protocol is Fibre Channel, the LUN will be in the range from 0 to 2047. If this record applies to a Virtual Volume, this attribute is the volume's secondary LUN in the range from 0 to 399999, for both iSCSI and Fibre Channel. If the record applies to a OpenstackV2 volume, the LUN will be in the range from 0 to 2047, for both iSCSI and Fibre Channel. If this record applies to a protocol endpoint or only a snapshot, this attribute is not meaningful and is set to null.
    
   	Lun *int64 `json:"lun,omitempty"`
   
    // Identifier for the volume this access control record applies to.
    
 	VolID *string `json:"vol_id,omitempty"`
   
    // Name of the volume this access control record applies to.
    
 	VolName *string `json:"vol_name,omitempty"`
   
    // External management agent type.
    
   	VolAgentType *NsAgentType `json:"vol_agent_type,omitempty"`
   
    // Identifier for the protocol endpoint this access control record applies to.
    
 	PeID *string `json:"pe_id,omitempty"`
   
    // Name of the protocol endpoint this access control record applies to.
    
 	PeName *string `json:"pe_name,omitempty"`
   
    // LUN (Logical Unit Number) to associate with this protocol endpoint. Valid LUNs are in the 0-2047 range.
    
   	PeLun *int64 `json:"pe_lun,omitempty"`
   
    // Identifier for the snapshot this access control record applies to.
    
 	SnapID *string `json:"snap_id,omitempty"`
   
    // Name of the snapshot this access control record applies to.
    
 	SnapName *string `json:"snap_name,omitempty"`
   
    // List of candidate protocol endpoints that may be used to access the Virtual Volume. One of them will be selected for the access control record. This field is required only when creating an access control record for a Virtual Volume.
    
	PeIDs []*string `json:"pe_ids,omitempty"`
   
    // Information about the snapshot LUNs associated with this access control record. This field is meaningful when the online snapshot can be accessed as a LUN in the group.
    
   	Snapluns []*NsSnapLunInfo `json:"snapluns,omitempty"`
   
    // Time when this access control record was created.
    
   	CreationTime *int64 `json:"creation_time,omitempty"`
   
    // Time when this access control record was last modified.
    
   	LastModified *int64 `json:"last_modified,omitempty"`
   
    // Access protocol of the volume.
    
   	AccessProtocol *NsAccessProtocol `json:"access_protocol,omitempty"`
}
