// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsPoolRebalanceMigStatusFields provides field names to use in filter parameters, for example.
var NsPoolRebalanceMigStatusFields *NsPoolRebalanceMigStatusFieldHandles

func init() {
	fieldID := "id"
	fieldName := "name"
	fieldPoolAvgSpaceUtilization := "pool_avg_space_utilization"
	fieldArrayDataMigrationStatus := "array_data_migration_status"

	NsPoolRebalanceMigStatusFields = &NsPoolRebalanceMigStatusFieldHandles{
		ID:                       &fieldID,
		Name:                     &fieldName,
		PoolAvgSpaceUtilization:  &fieldPoolAvgSpaceUtilization,
		ArrayDataMigrationStatus: &fieldArrayDataMigrationStatus,
	}
}

// NsPoolRebalanceMigStatus - Status of data rebalance operations in a pool.
type NsPoolRebalanceMigStatus struct {
	// ID - Unique ID of the pool.
	ID *string `json:"id,omitempty"`
	// Name - Name of the Pool.
	Name *string `json:"name,omitempty"`
	// PoolAvgSpaceUtilization - Average space utilization for the arrays in the pool.
	PoolAvgSpaceUtilization *int64 `json:"pool_avg_space_utilization,omitempty"`
	// ArrayDataMigrationStatus - Data migration status for a list of arrays in the pool.
	ArrayDataMigrationStatus []*NsArrayMigStatus `json:"array_data_migration_status,omitempty"`
}

// NsPoolRebalanceMigStatusFieldHandles provides a string representation for each AccessControlRecord field.
type NsPoolRebalanceMigStatusFieldHandles struct {
	ID                       *string
	Name                     *string
	PoolAvgSpaceUtilization  *string
	ArrayDataMigrationStatus *string
}
