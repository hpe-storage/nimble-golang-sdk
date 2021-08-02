// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsArrayDetailFields provides field names to use in filter parameters, for example.
var NsArrayDetailFields *NsArrayDetailFieldHandles

func init() {
	NsArrayDetailFields = &NsArrayDetailFieldHandles{
		ID:                       "id",
		ArrayId:                  "array_id",
		Name:                     "name",
		ArrayName:                "array_name",
		EvacTime:                 "evac_time",
		EvacUsage:                "evac_usage",
		UsableCapacity:           "usable_capacity",
		Usage:                    "usage",
		VolUsageCompressedBytes:  "vol_usage_compressed_bytes",
		SnapUsageCompressedBytes: "snap_usage_compressed_bytes",
		UsageValid:               "usage_valid",
		Migrate:                  "migrate",
	}
}

// NsArrayDetail - Detailed array information.
type NsArrayDetail struct {
	// ID - Array API ID.
	ID *string `json:"id,omitempty"`
	// ArrayId - Array API ID.
	ArrayId *string `json:"array_id,omitempty"`
	// Name - Unique name of array.
	Name *string `json:"name,omitempty"`
	// ArrayName - Unique name of array.
	ArrayName *string `json:"array_name,omitempty"`
	// EvacTime - Start time of array evacuation.
	EvacTime *int64 `json:"evac_time,omitempty"`
	// EvacUsage - Initial data in the array.
	EvacUsage *int64 `json:"evac_usage,omitempty"`
	// UsableCapacity - Usable capacity of the array.
	UsableCapacity *int64 `json:"usable_capacity,omitempty"`
	// Usage - Usage of the array.
	Usage *int64 `json:"usage,omitempty"`
	// VolUsageCompressedBytes - Usage of volumes in the array.
	VolUsageCompressedBytes *int64 `json:"vol_usage_compressed_bytes,omitempty"`
	// SnapUsageCompressedBytes - Usage of snapshots in the array.
	SnapUsageCompressedBytes *int64 `json:"snap_usage_compressed_bytes,omitempty"`
	// UsageValid - Indicate whether usage of the array is valid.
	UsageValid *bool `json:"usage_valid,omitempty"`
	// Migrate - Migrate status of array.
	Migrate *NsPoolMigrate `json:"migrate,omitempty"`
}

// NsArrayDetailFieldHandles provides a string representation for each AccessControlRecord field.
type NsArrayDetailFieldHandles struct {
	ID                       string
	ArrayId                  string
	Name                     string
	ArrayName                string
	EvacTime                 string
	EvacUsage                string
	UsableCapacity           string
	Usage                    string
	VolUsageCompressedBytes  string
	SnapUsageCompressedBytes string
	UsageValid               string
	Migrate                  string
}
