// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsReapJobsReturn - Response from reaping jobs.

// Export NsReapJobsReturnFields provides field names to use in filter parameters, for example.
var NsReapJobsReturnFields *NsReapJobsReturnStringFields

func init() {
	fieldReaped := "reaped"
	fieldRemaining := "remaining"

	NsReapJobsReturnFields = &NsReapJobsReturnStringFields{
		Reaped:    &fieldReaped,
		Remaining: &fieldRemaining,
	}
}

type NsReapJobsReturn struct {
	// Reaped - Number of jobs reaped.
	Reaped *int64 `json:"reaped,omitempty"`
	// Remaining - Number of jobs remaining.
	Remaining *int64 `json:"remaining,omitempty"`
}

// Struct for NsReapJobsReturnFields
type NsReapJobsReturnStringFields struct {
	Reaped    *string
	Remaining *string
}
