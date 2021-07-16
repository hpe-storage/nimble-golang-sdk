// Copyright 2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsPerfPolicySummary - Select fields containing performance policy.
// Export NsPerfPolicySummaryFields for advance operations like search filter etc.
var NsPerfPolicySummaryFields *NsPerfPolicySummary

func init() {
	Namefield := "name"

	NsPerfPolicySummaryFields = &NsPerfPolicySummary{
		Name: &Namefield,
	}
}

type NsPerfPolicySummary struct {
	// Name - Name of performance policy.
	Name *string `json:"name,omitempty"`
}
