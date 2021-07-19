// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsReapJobsReturn - Response from reaping jobs.
// Export NsReapJobsReturnFields for advance operations like search filter etc.
var NsReapJobsReturnFields *NsReapJobsReturn

func init() {

	NsReapJobsReturnFields = &NsReapJobsReturn{}
}

type NsReapJobsReturn struct {
	// Reaped - Number of jobs reaped.
	Reaped *int64 `json:"reaped,omitempty"`
	// Remaining - Number of jobs remaining.
	Remaining *int64 `json:"remaining,omitempty"`
}
