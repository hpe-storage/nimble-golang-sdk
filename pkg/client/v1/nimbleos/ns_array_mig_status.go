// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsArrayMigStatusFields provides field names to use in filter parameters, for example.
var NsArrayMigStatusFields *NsArrayMigStatusFieldHandles

func init() {
	fieldID := "id"
	fieldName := "name"
	fieldIsDataSource := "is_data_source"
	fieldSpaceUtilization := "space_utilization"

	NsArrayMigStatusFields = &NsArrayMigStatusFieldHandles{
		ID:               &fieldID,
		Name:             &fieldName,
		IsDataSource:     &fieldIsDataSource,
		SpaceUtilization: &fieldSpaceUtilization,
	}
}

// NsArrayMigStatus - Data migration status for an array.
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

// NsArrayMigStatusFieldHandles provides a string representation for each AccessControlRecord field.
type NsArrayMigStatusFieldHandles struct {
	ID               *string
	Name             *string
	IsDataSource     *string
	SpaceUtilization *string
}
