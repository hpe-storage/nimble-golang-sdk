// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsArrayUnassignMigStatusFields provides field names to use in filter parameters, for example.
var NsArrayUnassignMigStatusFields *NsArrayUnassignMigStatusFieldHandles

func init() {
	fieldID := "id"
	fieldName := "name"
	fieldDestinationArrays := "destination_arrays"
	fieldBytesMigrated := "bytes_migrated"
	fieldBytesRemaining := "bytes_remaining"
	fieldStartTime := "start_time"
	fieldEstimatedCompletionTime := "estimated_completion_time"

	NsArrayUnassignMigStatusFields = &NsArrayUnassignMigStatusFieldHandles{
		ID:                      &fieldID,
		Name:                    &fieldName,
		DestinationArrays:       &fieldDestinationArrays,
		BytesMigrated:           &fieldBytesMigrated,
		BytesRemaining:          &fieldBytesRemaining,
		StartTime:               &fieldStartTime,
		EstimatedCompletionTime: &fieldEstimatedCompletionTime,
	}
}

// NsArrayUnassignMigStatus - Data migration status for array being unassigned from its pool.
type NsArrayUnassignMigStatus struct {
	// ID - Unique identifier of the array being unassigned.
	ID *string `json:"id,omitempty"`
	// Name - Name of the array being unassigned.
	Name *string `json:"name,omitempty"`
	// DestinationArrays - List of arrays to which data is being migrated.
	DestinationArrays []*NsArraySummary `json:"destination_arrays,omitempty"`
	// BytesMigrated - Number of bytes already migrated to the destination arrays.
	BytesMigrated *int64 `json:"bytes_migrated,omitempty"`
	// BytesRemaining - Number of bytes yet to be migrated to the destination arrays.
	BytesRemaining *int64 `json:"bytes_remaining,omitempty"`
	// StartTime - Time when array unassign operation started.
	StartTime *int64 `json:"start_time,omitempty"`
	// EstimatedCompletionTime - Estimated completion time. 0 if unknown.
	EstimatedCompletionTime *int64 `json:"estimated_completion_time,omitempty"`
}

// NsArrayUnassignMigStatusFieldHandles provides a string representation for each AccessControlRecord field.
type NsArrayUnassignMigStatusFieldHandles struct {
	ID                      *string
	Name                    *string
	DestinationArrays       *string
	BytesMigrated           *string
	BytesRemaining          *string
	StartTime               *string
	EstimatedCompletionTime *string
}
