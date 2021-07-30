// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsReapJobsReturnFields provides field names to use in filter parameters, for example.
var NsReapJobsReturnFields *NsReapJobsReturnFieldHandles

func init() {
	fieldReaped := "reaped"
	fieldRemaining := "remaining"

	NsReapJobsReturnFields = &NsReapJobsReturnFieldHandles{
		Reaped:    &fieldReaped,
		Remaining: &fieldRemaining,
	}
}

// NsReapJobsReturn - Response from reaping jobs.
type NsReapJobsReturn struct {
	// Reaped - Number of jobs reaped.
	Reaped *int64 `json:"reaped,omitempty"`
	// Remaining - Number of jobs remaining.
	Remaining *int64 `json:"remaining,omitempty"`
}

// NsReapJobsReturnFieldHandles provides a string representation for each AccessControlRecord field.
type NsReapJobsReturnFieldHandles struct {
	Reaped    *string
	Remaining *string
}
