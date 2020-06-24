/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsVolFamMigStatus - Data migration status for a group of related volumes.
// Export NsVolFamMigStatusFields for advance operations like search filter etc.
var NsVolFamMigStatusFields *NsVolFamMigStatus

func init(){
	RootVolIDfield:= "root_vol_id"
	RootVolNamefield:= "root_vol_name"
	SourcePoolIDfield:= "source_pool_id"
	SourcePoolNamefield:= "source_pool_name"
	DestPoolIDfield:= "dest_pool_id"
	DestPoolNamefield:= "dest_pool_name"
		
	NsVolFamMigStatusFields= &NsVolFamMigStatus{
		RootVolID: &RootVolIDfield,
		RootVolName: &RootVolNamefield,
		SourcePoolID: &SourcePoolIDfield,
		SourcePoolName: &SourcePoolNamefield,
		DestPoolID: &DestPoolIDfield,
		DestPoolName: &DestPoolNamefield,
		
	}
}

type NsVolFamMigStatus struct {
   
    // ID of the root volume in the group.
    
 	RootVolID *string `json:"root_vol_id,omitempty"`
   
    // Name of the root volume in the group.
    
 	RootVolName *string `json:"root_vol_name,omitempty"`
   
    // ID of the source pool, where the volumes originally locate.
    
 	SourcePoolID *string `json:"source_pool_id,omitempty"`
   
    // Name of the source pool, where the volumes originally locate.
    
 	SourcePoolName *string `json:"source_pool_name,omitempty"`
   
    // ID of the destination pool, where the volumes are moved.
    
 	DestPoolID *string `json:"dest_pool_id,omitempty"`
   
    // Name of the destination pool, where the volumes are moved.
    
 	DestPoolName *string `json:"dest_pool_name,omitempty"`
   
    // The bytes of volumes which have been moved.
    
   	MoveBytesMigrated *int64 `json:"move_bytes_migrated,omitempty"`
   
    // The bytes of volumes which have not been moved.
    
   	MoveBytesRemaining *int64 `json:"move_bytes_remaining,omitempty"`
   
    // The start time when the volumes was moved.
    
   	MoveStartTime *int64 `json:"move_start_time,omitempty"`
   
    // The estimated time of completion of a move.
    
   	MoveEstComplTime *int64 `json:"move_est_compl_time,omitempty"`
   
    // Data migration status for the arrays that store the volumes.
    
   	ArrayList []*NsArrayMigStatus `json:"array_list,omitempty"`
}
