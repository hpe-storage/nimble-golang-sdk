// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsPerfPolicySummary - Select fields containing performance policy.

// Export NsPerfPolicySummaryFields provides field names to use in filter parameters, for example.
var NsPerfPolicySummaryFields *NsPerfPolicySummaryStringFields

func init() {
	fieldName := "name"

	NsPerfPolicySummaryFields = &NsPerfPolicySummaryStringFields{
		Name: &fieldName,
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
