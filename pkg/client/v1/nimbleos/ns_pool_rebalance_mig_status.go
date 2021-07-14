// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos


// NsPoolRebalanceMigStatus - Status of data rebalance operations in a pool.
// Export NsPoolRebalanceMigStatusFields for advance operations like search filter etc.
var NsPoolRebalanceMigStatusFields *NsPoolRebalanceMigStatus

func init(){
 IDfield:= "id"
 Namefield:= "name"

 NsPoolRebalanceMigStatusFields= &NsPoolRebalanceMigStatus{
  ID:                         &IDfield,
  Name:                       &Namefield,
 }
}


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
