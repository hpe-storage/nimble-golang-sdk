// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsPerfPolicySummaryFields provides field names to use in filter parameters, for example.
var NsPerfPolicySummaryFields *NsPerfPolicySummaryFieldHandles

func init() {
	NsPerfPolicySummaryFields = &NsPerfPolicySummaryFieldHandles{
		Name: "name",
	}
}

// NsPerfPolicySummary - Select fields containing performance policy.
type NsPerfPolicySummary struct {
	// Name - Name of performance policy.
	Name *string `json:"name,omitempty"`
}

// NsPerfPolicySummaryFieldHandles provides a string representation for each NsPerfPolicySummary field.
type NsPerfPolicySummaryFieldHandles struct {
	Name string
}
