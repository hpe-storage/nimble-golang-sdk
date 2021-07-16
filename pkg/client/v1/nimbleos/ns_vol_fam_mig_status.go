// Copyright 2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsVolFamMigStatus - Data migration status for a group of related volumes.
// Export NsVolFamMigStatusFields for advance operations like search filter etc.
var NsVolFamMigStatusFields *NsVolFamMigStatus

func init() {
	RootVolIdfield := "root_vol_id"
	RootVolNamefield := "root_vol_name"
	SourcePoolIdfield := "source_pool_id"
	SourcePoolNamefield := "source_pool_name"
	DestPoolIdfield := "dest_pool_id"
	DestPoolNamefield := "dest_pool_name"

	NsVolFamMigStatusFields = &NsVolFamMigStatus{
		RootVolId:      &RootVolIdfield,
		RootVolName:    &RootVolNamefield,
		SourcePoolId:   &SourcePoolIdfield,
		SourcePoolName: &SourcePoolNamefield,
		DestPoolId:     &DestPoolIdfield,
		DestPoolName:   &DestPoolNamefield,
	}
}

type NsVolFamMigStatus struct {
	// RootVolId - ID of the root volume in the group.
	RootVolId *string `json:"root_vol_id,omitempty"`
	// RootVolName - Name of the root volume in the group.
	RootVolName *string `json:"root_vol_name,omitempty"`
	// SourcePoolId - ID of the source pool, where the volumes originally locate.
	SourcePoolId *string `json:"source_pool_id,omitempty"`
	// SourcePoolName - Name of the source pool, where the volumes originally locate.
	SourcePoolName *string `json:"source_pool_name,omitempty"`
	// DestPoolId - ID of the destination pool, where the volumes are moved.
	DestPoolId *string `json:"dest_pool_id,omitempty"`
	// DestPoolName - Name of the destination pool, where the volumes are moved.
	DestPoolName *string `json:"dest_pool_name,omitempty"`
	// MoveBytesMigrated - The bytes of volumes which have been moved.
	MoveBytesMigrated *int64 `json:"move_bytes_migrated,omitempty"`
	// MoveBytesRemaining - The bytes of volumes which have not been moved.
	MoveBytesRemaining *int64 `json:"move_bytes_remaining,omitempty"`
	// MoveStartTime - The start time when the volumes was moved.
	MoveStartTime *int64 `json:"move_start_time,omitempty"`
	// MoveEstComplTime - The estimated time of completion of a move.
	MoveEstComplTime *int64 `json:"move_est_compl_time,omitempty"`
	// ArrayList - Data migration status for the arrays that store the volumes.
	ArrayList []*NsArrayMigStatus `json:"array_list,omitempty"`
}
