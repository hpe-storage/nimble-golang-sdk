/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsAccessControlRecord - An access control record associated with a volume or snapshot or protocol endpoint.
// Export NsAccessControlRecordFields for advance operations like search filter etc.
var NsAccessControlRecordFields *NsAccessControlRecord

func init(){
	IDfield:= "id"
	AclIDfield:= "acl_id"
	InitiatorGroupIDfield:= "initiator_group_id"
	InitiatorGroupNamefield:= "initiator_group_name"
	ChapUserIDfield:= "chap_user_id"
	ChapUserNamefield:= "chap_user_name"
	PeIDfield:= "pe_id"
	PeNamefield:= "pe_name"
	VolIDfield:= "vol_id"
	VolNamefield:= "vol_name"
	SnapIDfield:= "snap_id"
	SnapNamefield:= "snap_name"
		
	NsAccessControlRecordFields= &NsAccessControlRecord{
		ID: &IDfield,
		AclID: &AclIDfield,
		InitiatorGroupID: &InitiatorGroupIDfield,
		InitiatorGroupName: &InitiatorGroupNamefield,
		ChapUserID: &ChapUserIDfield,
		ChapUserName: &ChapUserNamefield,
		PeID: &PeIDfield,
		PeName: &PeNamefield,
		VolID: &VolIDfield,
		VolName: &VolNamefield,
		SnapID: &SnapIDfield,
		SnapName: &SnapNamefield,
		
	}
}

type NsAccessControlRecord struct {
   
    // ID of access control record.
    
 	ID *string `json:"id,omitempty"`
   
    // ID of access control record.
    
 	AclID *string `json:"acl_id,omitempty"`
   
    // ID of initiator group.
    
 	InitiatorGroupID *string `json:"initiator_group_id,omitempty"`
   
    // Name of initiator group.
    
 	InitiatorGroupName *string `json:"initiator_group_name,omitempty"`
   
    // ID of chap user.
    
 	ChapUserID *string `json:"chap_user_id,omitempty"`
   
    // Name of chap user.
    
 	ChapUserName *string `json:"chap_user_name,omitempty"`
   
    // State of apply to.
    
   	ApplyTo *NsAccessApplyTo `json:"apply_to,omitempty"`
   
    // LUN of snapshot or volume. Secondary LUN if this is Virtual Volume.
    
   	Lun *int64 `json:"lun,omitempty"`
   
    // Access protocol of snapshot or volume.
    
   	AccessProtocol *NsAccessProtocol `json:"access_protocol,omitempty"`
   
    // Information about the snapshot LUNs. This information is available only for Fibre Channel.
    
   	Snapluns []*NsSnapLunInfo `json:"snapluns,omitempty"`
   
    // ID of protocol endpoint.
    
 	PeID *string `json:"pe_id,omitempty"`
   
    // Name of protocol endpoint.
    
 	PeName *string `json:"pe_name,omitempty"`
   
    // LUN of protocol endpoint.
    
   	PeLun *int64 `json:"pe_lun,omitempty"`
   
    // ID of volume.
    
 	VolID *string `json:"vol_id,omitempty"`
   
    // Name of volume.
    
 	VolName *string `json:"vol_name,omitempty"`
   
    // ID of snapshot.
    
 	SnapID *string `json:"snap_id,omitempty"`
   
    // Name of snapshot.
    
 	SnapName *string `json:"snap_name,omitempty"`
}
