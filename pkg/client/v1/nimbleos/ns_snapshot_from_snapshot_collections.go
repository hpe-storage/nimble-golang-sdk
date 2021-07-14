// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos


// NsSnapshotFromSnapshotCollections - Snapshot as presented in snapshot collections object set.
// Export NsSnapshotFromSnapshotCollectionsFields for advance operations like search filter etc.
var NsSnapshotFromSnapshotCollectionsFields *NsSnapshotFromSnapshotCollections

func init(){
 IDfield:= "id"
 SnapcollIdfield:= "snapcoll_id"
 Namefield:= "name"
 SnapcollNamefield:= "snapcoll_name"
 VolIdfield:= "vol_id"
 VolNamefield:= "vol_name"
 SnapIdfield:= "snap_id"
 SnapNamefield:= "snap_name"
 ScheduleIdfield:= "schedule_id"
 ScheduleNamefield:= "schedule_name"

 NsSnapshotFromSnapshotCollectionsFields= &NsSnapshotFromSnapshotCollections{
  ID:           &IDfield,
  SnapcollId:   &SnapcollIdfield,
  Name:         &Namefield,
  SnapcollName: &SnapcollNamefield,
  VolId:        &VolIdfield,
  VolName:      &VolNamefield,
  SnapId:       &SnapIdfield,
  SnapName:     &SnapNamefield,
  ScheduleId:   &ScheduleIdfield,
  ScheduleName: &ScheduleNamefield,
 }
}


type NsSnapshotFromSnapshotCollections struct {
 // ID - Snapshot ID.
  ID *string `json:"id,omitempty"`
 // SnapcollId - ID of the snapshot collection which owns this snapshot.
  SnapcollId *string `json:"snapcoll_id,omitempty"`
 // Name - Snapshot name.
  Name *string `json:"name,omitempty"`
 // SnapcollName - Name of the snapshot collection which owns this snapshot.
  SnapcollName *string `json:"snapcoll_name,omitempty"`
 // VolId - ID of the volume that is parent to this snapshot.
  VolId *string `json:"vol_id,omitempty"`
 // VolName - Name of the volume that is parent to this snapshot.
  VolName *string `json:"vol_name,omitempty"`
 // SnapId - Snapshot ID.
  SnapId *string `json:"snap_id,omitempty"`
 // SnapName - Snapshot name.
  SnapName *string `json:"snap_name,omitempty"`
 // ScheduleId - ID of protection schedule.
  ScheduleId *string `json:"schedule_id,omitempty"`
 // ScheduleName - Name of protection schedule.
  ScheduleName *string `json:"schedule_name,omitempty"`
 // ExpiryTime - Unix timestamp indicating that the snapshot is considered expired by Snapshot Time-to-live(TTL). A value of 0 indicates that snapshot never expires.
    ExpiryTime *int64 `json:"expiry_time,omitempty"`
}
