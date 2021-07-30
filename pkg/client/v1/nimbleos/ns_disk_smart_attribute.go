// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsDiskSmartAttributeFields provides field names to use in filter parameters, for example.
var NsDiskSmartAttributeFields *NsDiskSmartAttributeFieldHandles

func init() {
	fieldName := "name"
	fieldSmartId := "smart_id"
	fieldCurValue := "cur_value"
	fieldWorstValue := "worst_value"
	fieldThreshold := "threshold"
	fieldFlags := "flags"
	fieldRawValue := "raw_value"
	fieldLastUpdatedEpochSecs := "last_updated_epoch_secs"

	NsDiskSmartAttributeFields = &NsDiskSmartAttributeFieldHandles{
		Name:                 &fieldName,
		SmartId:              &fieldSmartId,
		CurValue:             &fieldCurValue,
		WorstValue:           &fieldWorstValue,
		Threshold:            &fieldThreshold,
		Flags:                &fieldFlags,
		RawValue:             &fieldRawValue,
		LastUpdatedEpochSecs: &fieldLastUpdatedEpochSecs,
	}
}

// NsDiskSmartAttribute - One Smart attribute of a disk.
type NsDiskSmartAttribute struct {
	// Name - Name of the Smart attribute.
	Name *string `json:"name,omitempty"`
	// SmartId - Smart attribute ID.
	SmartId *int64 `json:"smart_id,omitempty"`
	// CurValue - Current value.
	CurValue *int64 `json:"cur_value,omitempty"`
	// WorstValue - Worst value.
	WorstValue *int64 `json:"worst_value,omitempty"`
	// Threshold - Smart threshold.
	Threshold *int64 `json:"threshold,omitempty"`
	// Flags - Smart flags.
	Flags *int64 `json:"flags,omitempty"`
	// RawValue - Raw value.
	RawValue *int64 `json:"raw_value,omitempty"`
	// LastUpdatedEpochSecs - Last update time in epoch seconds.
	LastUpdatedEpochSecs *int64 `json:"last_updated_epoch_secs,omitempty"`
}

// NsDiskSmartAttributeFieldHandles provides a string representation for each AccessControlRecord field.
type NsDiskSmartAttributeFieldHandles struct {
	Name                 *string
	SmartId              *string
	CurValue             *string
	WorstValue           *string
	Threshold            *string
	Flags                *string
	RawValue             *string
	LastUpdatedEpochSecs *string
}
