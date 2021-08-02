// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSensorCumulativeDataFields provides field names to use in filter parameters, for example.
var NsSensorCumulativeDataFields *NsSensorCumulativeDataFieldHandles

func init() {
	NsSensorCumulativeDataFields = &NsSensorCumulativeDataFieldHandles{
		Name:        "name",
		Index:       "index",
		Msec:        "msec",
		PrevUnavail: "prev_unavail",
		CurrUnavail: "curr_unavail",
		Curr:        "curr",
		Prev:        "prev",
	}
}

// NsSensorCumulativeData - Stat sensor cumulative data.
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

// NsSensorCumulativeDataFieldHandles provides a string representation for each AccessControlRecord field.
type NsSensorCumulativeDataFieldHandles struct {
	Name        string
	Index       string
	Msec        string
	PrevUnavail string
	CurrUnavail string
	Curr        string
	Prev        string
}
