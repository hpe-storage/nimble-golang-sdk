// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsVolFamMigStatusFields provides field names to use in filter parameters, for example.
var NsVolFamMigStatusFields *NsVolFamMigStatusFieldHandles

func init() {
	NsVolFamMigStatusFields = &NsVolFamMigStatusFieldHandles{
		RootVolId:          "root_vol_id",
		RootVolName:        "root_vol_name",
		SourcePoolId:       "source_pool_id",
		SourcePoolName:     "source_pool_name",
		DestPoolId:         "dest_pool_id",
		DestPoolName:       "dest_pool_name",
		MoveBytesMigrated:  "move_bytes_migrated",
		MoveBytesRemaining: "move_bytes_remaining",
		MoveStartTime:      "move_start_time",
		MoveEstComplTime:   "move_est_compl_time",
		ArrayList:          "array_list",
	}
}

// NsVolFamMigStatus - Data migration status for a group of related volumes.
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

// NsVolFamMigStatusFieldHandles provides a string representation for each AccessControlRecord field.
type NsVolFamMigStatusFieldHandles struct {
	RootVolId          string
	RootVolName        string
	SourcePoolId       string
	SourcePoolName     string
	DestPoolId         string
	DestPoolName       string
	MoveBytesMigrated  string
	MoveBytesRemaining string
	MoveStartTime      string
	MoveEstComplTime   string
	ArrayList          string
}
