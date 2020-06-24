/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsSnapshotFromSnapshotCollections - Snapshot as presented in snapshot collections object set.
// Export NsSnapshotFromSnapshotCollectionsFields for advance operations like search filter etc.
var NsSnapshotFromSnapshotCollectionsFields *NsSnapshotFromSnapshotCollections

func init(){
	IDfield:= "id"
	SnapcollIDfield:= "snapcoll_id"
	Namefield:= "name"
	SnapcollNamefield:= "snapcoll_name"
	VolIDfield:= "vol_id"
	VolNamefield:= "vol_name"
	SnapIDfield:= "snap_id"
	SnapNamefield:= "snap_name"
	ScheduleIDfield:= "schedule_id"
	ScheduleNamefield:= "schedule_name"
		
	NsSnapshotFromSnapshotCollectionsFields= &NsSnapshotFromSnapshotCollections{
		ID: &IDfield,
		SnapcollID: &SnapcollIDfield,
		Name: &Namefield,
		SnapcollName: &SnapcollNamefield,
		VolID: &VolIDfield,
		VolName: &VolNamefield,
		SnapID: &SnapIDfield,
		SnapName: &SnapNamefield,
		ScheduleID: &ScheduleIDfield,
		ScheduleName: &ScheduleNamefield,
		
	}
}

type NsSnapshotFromSnapshotCollections struct {
   
    // Snapshot ID.
    
 	ID *string `json:"id,omitempty"`
   
    // ID of the snapshot collection which owns this snapshot.
    
 	SnapcollID *string `json:"snapcoll_id,omitempty"`
   
    // Snapshot name.
    
 	Name *string `json:"name,omitempty"`
   
    // Name of the snapshot collection which owns this snapshot.
    
 	SnapcollName *string `json:"snapcoll_name,omitempty"`
   
    // ID of the volume that is parent to this snapshot.
    
 	VolID *string `json:"vol_id,omitempty"`
   
    // Name of the volume that is parent to this snapshot.
    
 	VolName *string `json:"vol_name,omitempty"`
   
    // Snapshot ID.
    
 	SnapID *string `json:"snap_id,omitempty"`
   
    // Snapshot name.
    
 	SnapName *string `json:"snap_name,omitempty"`
   
    // ID of protection schedule.
    
 	ScheduleID *string `json:"schedule_id,omitempty"`
   
    // Name of protection schedule.
    
 	ScheduleName *string `json:"schedule_name,omitempty"`
   
    // Unix timestamp indicating that the snapshot is considered expired by Snapshot Time-to-live(TTL). A value of 0 indicates that snapshot never expires.
    
   	ExpiryTime *int64 `json:"expiry_time,omitempty"`
}
