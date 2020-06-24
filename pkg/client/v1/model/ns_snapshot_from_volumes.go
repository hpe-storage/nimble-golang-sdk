/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsSnapshotFromVolumes - Snapshot as presented in volumes object set.
// Export NsSnapshotFromVolumesFields for advance operations like search filter etc.
var NsSnapshotFromVolumesFields *NsSnapshotFromVolumes

func init(){
	IDfield:= "id"
	SnapIDfield:= "snap_id"
	Namefield:= "name"
	SnapNamefield:= "snap_name"
		
	NsSnapshotFromVolumesFields= &NsSnapshotFromVolumes{
		ID: &IDfield,
		SnapID: &SnapIDfield,
		Name: &Namefield,
		SnapName: &SnapNamefield,
		
	}
}

type NsSnapshotFromVolumes struct {
   
    // Snapshot id.
    
 	ID *string `json:"id,omitempty"`
   
    // Snapshot id.
    
 	SnapID *string `json:"snap_id,omitempty"`
   
    // Snapshot name.
    
 	Name *string `json:"name,omitempty"`
   
    // Snapshot name.
    
 	SnapName *string `json:"snap_name,omitempty"`
}
