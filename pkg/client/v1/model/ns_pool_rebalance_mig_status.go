/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsPoolRebalanceMigStatus - Status of data rebalance operations in a pool.
// Export NsPoolRebalanceMigStatusFields for advance operations like search filter etc.
var NsPoolRebalanceMigStatusFields *NsPoolRebalanceMigStatus

func init(){
	IDfield:= "id"
	Namefield:= "name"
		
	NsPoolRebalanceMigStatusFields= &NsPoolRebalanceMigStatus{
		ID: &IDfield,
		Name: &Namefield,
		
	}
}

type NsPoolRebalanceMigStatus struct {
   
    // Unique ID of the pool.
    
 	ID *string `json:"id,omitempty"`
   
    // Name of the Pool.
    
 	Name *string `json:"name,omitempty"`
   
    // Average space utilization for the arrays in the pool.
    
   	PoolAvgSpaceUtilization *int64 `json:"pool_avg_space_utilization,omitempty"`
   
    // Data migration status for a list of arrays in the pool.
    
   	ArrayDataMigrationStatus []*NsArrayMigStatus `json:"array_data_migration_status,omitempty"`
}
