// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// HealthCheck - View the health of the group of arrays.
// Export HealthCheckFields for advance operations like search filter etc.
var HealthCheckFields *HealthCheckStringFields

func init() {
	IDfield := "id"
	Scopefield := "scope"
	Contextfield := "context"
	OnDemandfield := "on_demand"
	ArrayIdfield := "array_id"
	CtrlrIdfield := "ctrlr_id"
	ElementResultfield := "element_result"

	HealthCheckFields = &HealthCheckStringFields{
		ID:            &IDfield,
		Scope:         &Scopefield,
		Context:       &Contextfield,
		OnDemand:      &OnDemandfield,
		ArrayId:       &ArrayIdfield,
		CtrlrId:       &CtrlrIdfield,
		ElementResult: &ElementResultfield,
	}
}

type HealthCheck struct {
	// ID - Identifier for the health check.
	ID *string `json:"id,omitempty"`
	// Scope - Scope at which the health check is to be run.
	Scope *NsHcfScope `json:"scope,omitempty"`
	// Context - Context for aggregating health check results.
	Context *NsHcfContext `json:"context,omitempty"`
	// OnDemand - Flag to indicate running the health checks and then report results.
	OnDemand *bool `json:"on_demand,omitempty"`
	// ArrayId - Identifier of the array to which this result belongs.
	ArrayId *string `json:"array_id,omitempty"`
	// CtrlrId - Identifier of the controller to which this result belongs.
	CtrlrId *string `json:"ctrlr_id,omitempty"`
	// ElementResult - List of health check errors for a particular element.
	ElementResult *NsHcfResult `json:"element_result,omitempty"`
}

// Struct for HealthCheckFields
type HealthCheckStringFields struct {
	ID            *string
	Scope         *string
	Context       *string
	OnDemand      *string
	ArrayId       *string
	CtrlrId       *string
	ElementResult *string
}
