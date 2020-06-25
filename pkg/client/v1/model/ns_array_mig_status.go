// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsArrayMigStatus - Data migration status for an array.
// Export NsArrayMigStatusFields for advance operations like search filter etc.
var NsArrayMigStatusFields *NsArrayMigStatus

func init(){
	IDfield:= "id"
	Namefield:= "name"
		
	NsArrayMigStatusFields= &NsArrayMigStatus{
		ID:               &IDfield,
		Name:             &Namefield,
	}
}

type NsArrayMigStatus struct {
	// ID - Unique identifier of the array.
 	ID *string `json:"id,omitempty"`
	// Name - Name of the array.
 	Name *string `json:"name,omitempty"`
	// IsDataSource - Indicates whether the array is data source or not.
    IsDataSource *bool `json:"is_data_source,omitempty"`
	// SpaceUtilization - Space utilization as a percentage of array size.
   	SpaceUtilization *int64 `json:"space_utilization,omitempty"`
}
