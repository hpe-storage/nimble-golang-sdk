// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsPerfPolicySummary - Select fields containing performance policy.
// Export NsPerfPolicySummaryFields for advance operations like search filter etc.
var NsPerfPolicySummaryFields *NsPerfPolicySummaryStringFields

func init() {
	Namefield := "name"

	NsPerfPolicySummaryFields = &NsPerfPolicySummaryStringFields{
		Name: &Namefield,
	}
}

type NsPerfPolicySummary struct {
	// Name - Name of performance policy.
	Name *string `json:"name,omitempty"`
}

// Struct for NsPerfPolicySummaryFields
type NsPerfPolicySummaryStringFields struct {
	Name *string
}
