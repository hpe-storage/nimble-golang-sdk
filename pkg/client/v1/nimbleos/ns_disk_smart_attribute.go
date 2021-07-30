// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsDiskSmartAttribute - One Smart attribute of a disk.
// Export NsDiskSmartAttributeFields for advance operations like search filter etc.
var NsDiskSmartAttributeFields *NsDiskSmartAttributeStringFields

func init() {
	Namefield := "name"
	SmartIdfield := "smart_id"
	CurValuefield := "cur_value"
	WorstValuefield := "worst_value"
	Thresholdfield := "threshold"
	Flagsfield := "flags"
	RawValuefield := "raw_value"
	LastUpdatedEpochSecsfield := "last_updated_epoch_secs"

	NsDiskSmartAttributeFields = &NsDiskSmartAttributeStringFields{
		Name:                 &Namefield,
		SmartId:              &SmartIdfield,
		CurValue:             &CurValuefield,
		WorstValue:           &WorstValuefield,
		Threshold:            &Thresholdfield,
		Flags:                &Flagsfield,
		RawValue:             &RawValuefield,
		LastUpdatedEpochSecs: &LastUpdatedEpochSecsfield,
	}
}

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

// Struct for NsDiskSmartAttributeFields
type NsDiskSmartAttributeStringFields struct {
	Name                 *string
	SmartId              *string
	CurValue             *string
	WorstValue           *string
	Threshold            *string
	Flags                *string
	RawValue             *string
	LastUpdatedEpochSecs *string
}
