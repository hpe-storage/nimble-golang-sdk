// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// HealthCheck - View the health of the group of arrays.
// Export HealthCheckFields for advance operations like search filter etc.
var HealthCheckFields *HealthCheck

func init() {

	HealthCheckFields = &HealthCheck{
		ID:      "id",
		ArrayId: "array_id",
		CtrlrId: "ctrlr_id",
	}
}

type HealthCheck struct {
	// ID - Identifier for the health check.
	ID string `json:"id,omitempty"`
	// Scope - Scope at which the health check is to be run.
	Scope *NsHcfScope `json:"scope,omitempty"`
	// Context - Context for aggregating health check results.
	Context *NsHcfContext `json:"context,omitempty"`
	// OnDemand - Flag to indicate running the health checks and then report results.
	OnDemand *bool `json:"on_demand,omitempty"`
	// ArrayId - Identifier of the array to which this result belongs.
	ArrayId string `json:"array_id,omitempty"`
	// CtrlrId - Identifier of the controller to which this result belongs.
	CtrlrId string `json:"ctrlr_id,omitempty"`
	// ElementResult - List of health check errors for a particular element.
	ElementResult *NsHcfResult `json:"element_result,omitempty"`
}

// sdk internal struct
type healthCheck struct {
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

// EncodeHealthCheck - Transform HealthCheck to healthCheck type
func EncodeHealthCheck(request interface{}) (*healthCheck, error) {
	reqHealthCheck := request.(*HealthCheck)
	byte, err := json.Marshal(reqHealthCheck)
	objPtr := &healthCheck{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeHealthCheck Transform healthCheck to HealthCheck type
func DecodeHealthCheck(request interface{}) (*HealthCheck, error) {
	reqHealthCheck := request.(*healthCheck)
	byte, err := json.Marshal(reqHealthCheck)
	obj := &HealthCheck{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
