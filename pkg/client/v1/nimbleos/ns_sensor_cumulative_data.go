// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSensorCumulativeData - Stat sensor cumulative data.

// Export NsSensorCumulativeDataFields provides field names to use in filter parameters, for example.
var NsSensorCumulativeDataFields *NsSensorCumulativeDataStringFields

func init() {
	fieldName := "name"
	fieldIndex := "index"
	fieldMsec := "msec"
	fieldPrevUnavail := "prev_unavail"
	fieldCurrUnavail := "curr_unavail"
	fieldCurr := "curr"
	fieldPrev := "prev"

	NsSensorCumulativeDataFields = &NsSensorCumulativeDataStringFields{
		Name:        &fieldName,
		Index:       &fieldIndex,
		Msec:        &fieldMsec,
		PrevUnavail: &fieldPrevUnavail,
		CurrUnavail: &fieldCurrUnavail,
		Curr:        &fieldCurr,
		Prev:        &fieldPrev,
	}
}

type NsSensorCumulativeData struct {
	// Name - Name of sensor.
	Name *string `json:"name,omitempty"`
	// Index - Index of sensor in stat for versioning.
	Index *int64 `json:"index,omitempty"`
	// Msec - Creation time of stat sensor in millisecs.
	Msec *int64 `json:"msec,omitempty"`
	// PrevUnavail - Previous stat value is not valid.
	PrevUnavail *bool `json:"prev_unavail,omitempty"`
	// CurrUnavail - Current stat value is not valid.
	CurrUnavail *bool `json:"curr_unavail,omitempty"`
	// Curr - Current stat value.
	Curr *int64 `json:"curr,omitempty"`
	// Prev - Previous stat value from last sample period.
	Prev *int64 `json:"prev,omitempty"`
}

// Struct for NsSensorCumulativeDataFields
type NsSensorCumulativeDataStringFields struct {
	Name        *string
	Index       *string
	Msec        *string
	PrevUnavail *string
	CurrUnavail *string
	Curr        *string
	Prev        *string
}
