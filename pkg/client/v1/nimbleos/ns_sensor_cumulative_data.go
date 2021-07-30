// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSensorCumulativeData - Stat sensor cumulative data.
// Export NsSensorCumulativeDataFields for advance operations like search filter etc.
var NsSensorCumulativeDataFields *NsSensorCumulativeDataStringFields

func init() {
	Namefield := "name"
	Indexfield := "index"
	Msecfield := "msec"
	PrevUnavailfield := "prev_unavail"
	CurrUnavailfield := "curr_unavail"
	Currfield := "curr"
	Prevfield := "prev"

	NsSensorCumulativeDataFields = &NsSensorCumulativeDataStringFields{
		Name:        &Namefield,
		Index:       &Indexfield,
		Msec:        &Msecfield,
		PrevUnavail: &PrevUnavailfield,
		CurrUnavail: &CurrUnavailfield,
		Curr:        &Currfield,
		Prev:        &Prevfield,
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
