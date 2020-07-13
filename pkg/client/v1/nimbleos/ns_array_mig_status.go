// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsArrayMigStatus - Data migration status for an array.
// Export NsArrayMigStatusFields for advance operations like search filter etc.
var NsArrayMigStatusFields *NsArrayMigStatus

func init() {

	NsArrayMigStatusFields = &NsArrayMigStatus{
		ID:   "id",
		Name: "name",
	}
}

type NsArrayMigStatus struct {
	// ID - Unique identifier of the array.
	ID string `json:"id,omitempty"`
	// Name - Name of the array.
	Name string `json:"name,omitempty"`
	// IsDataSource - Indicates whether the array is data source or not.
	IsDataSource *bool `json:"is_data_source,omitempty"`
	// SpaceUtilization - Space utilization as a percentage of array size.
	SpaceUtilization int64 `json:"space_utilization,omitempty"`
}

// sdk internal struct
type nsArrayMigStatus struct {
	// ID - Unique identifier of the array.
	ID *string `json:"id,omitempty"`
	// Name - Name of the array.
	Name *string `json:"name,omitempty"`
	// IsDataSource - Indicates whether the array is data source or not.
	IsDataSource *bool `json:"is_data_source,omitempty"`
	// SpaceUtilization - Space utilization as a percentage of array size.
	SpaceUtilization *int64 `json:"space_utilization,omitempty"`
}

// EncodeNsArrayMigStatus - Transform NsArrayMigStatus to nsArrayMigStatus type
func EncodeNsArrayMigStatus(request interface{}) (*nsArrayMigStatus, error) {
	reqNsArrayMigStatus := request.(*NsArrayMigStatus)
	byte, err := json.Marshal(reqNsArrayMigStatus)
	objPtr := &nsArrayMigStatus{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsArrayMigStatus Transform nsArrayMigStatus to NsArrayMigStatus type
func DecodeNsArrayMigStatus(request interface{}) (*NsArrayMigStatus, error) {
	reqNsArrayMigStatus := request.(*nsArrayMigStatus)
	byte, err := json.Marshal(reqNsArrayMigStatus)
	obj := &NsArrayMigStatus{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
