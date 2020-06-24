/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsVolumeSnapshotAttribute - Snapshot attributes that could be specified for individual snapshots during snapshot collection creation.
// Export NsVolumeSnapshotAttributeFields for advance operations like search filter etc.
var NsVolumeSnapshotAttributeFields *NsVolumeSnapshotAttribute

func init(){
	VolIDfield:= "vol_id"
	AppUuIDfield:= "app_uuid"
		
	NsVolumeSnapshotAttributeFields= &NsVolumeSnapshotAttribute{
		VolID: &VolIDfield,
		AppUuID: &AppUuIDfield,
		
	}
}

type NsVolumeSnapshotAttribute struct {
   
    // ID of the volume on which snapshot will be created.
    
 	VolID *string `json:"vol_id,omitempty"`
   
    // Key-value pairs that augment a snapshot's attributes.
    
   	Metadata []*NsKeyValue `json:"metadata,omitempty"`
   
    // Application identifier of snapshot.
    
 	AppUuID *string `json:"app_uuid,omitempty"`
}
