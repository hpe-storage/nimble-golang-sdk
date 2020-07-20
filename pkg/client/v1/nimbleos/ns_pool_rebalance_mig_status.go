// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsPoolRebalanceMigStatus - Status of data rebalance operations in a pool.
// Export NsPoolRebalanceMigStatusFields for advance operations like search filter etc.
var NsPoolRebalanceMigStatusFields *NsPoolRebalanceMigStatus

func init() {

	NsPoolRebalanceMigStatusFields = &NsPoolRebalanceMigStatus{
		ID:   "id",
		Name: "name",
	}
}

type NsPoolRebalanceMigStatus struct {
	// ID - Unique ID of the pool.
	ID string `json:"id,omitempty"`
	// Name - Name of the Pool.
	Name string `json:"name,omitempty"`
	// PoolAvgSpaceUtilization - Average space utilization for the arrays in the pool.
	PoolAvgSpaceUtilization int64 `json:"pool_avg_space_utilization,omitempty"`
	// ArrayDataMigrationStatus - Data migration status for a list of arrays in the pool.
	ArrayDataMigrationStatus []*NsArrayMigStatus `json:"array_data_migration_status,omitempty"`
}

// sdk internal struct
type nsPoolRebalanceMigStatus struct {
	ID                       *string             `json:"id,omitempty"`
	Name                     *string             `json:"name,omitempty"`
	PoolAvgSpaceUtilization  *int64              `json:"pool_avg_space_utilization,omitempty"`
	ArrayDataMigrationStatus []*NsArrayMigStatus `json:"array_data_migration_status,omitempty"`
}

// EncodeNsPoolRebalanceMigStatus - Transform NsPoolRebalanceMigStatus to nsPoolRebalanceMigStatus type
func EncodeNsPoolRebalanceMigStatus(request interface{}) (*nsPoolRebalanceMigStatus, error) {
	reqNsPoolRebalanceMigStatus := request.(*NsPoolRebalanceMigStatus)
	byte, err := json.Marshal(reqNsPoolRebalanceMigStatus)
	if err != nil {
		return nil, err
	}
	respNsPoolRebalanceMigStatusPtr := &nsPoolRebalanceMigStatus{}
	err = json.Unmarshal(byte, respNsPoolRebalanceMigStatusPtr)
	return respNsPoolRebalanceMigStatusPtr, err
}

// DecodeNsPoolRebalanceMigStatus Transform nsPoolRebalanceMigStatus to NsPoolRebalanceMigStatus type
func DecodeNsPoolRebalanceMigStatus(request interface{}) (*NsPoolRebalanceMigStatus, error) {
	reqNsPoolRebalanceMigStatus := request.(*nsPoolRebalanceMigStatus)
	byte, err := json.Marshal(reqNsPoolRebalanceMigStatus)
	if err != nil {
		return nil, err
	}
	respNsPoolRebalanceMigStatus := &NsPoolRebalanceMigStatus{}
	err = json.Unmarshal(byte, respNsPoolRebalanceMigStatus)
	return respNsPoolRebalanceMigStatus, err
}
