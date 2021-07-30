// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsWitnessTestResponse - Results of witness connection test.

// Export NsWitnessTestResponseFields provides field names to use in filter parameters, for example.
var NsWitnessTestResponseFields *NsWitnessTestResponseStringFields

func init() {
	fieldArrayName := "array_name"
	fieldRole := "role"
	fieldWitnessConnectivityState := "witness_connectivity_state"
	fieldWitnessConnectivityMessage := "witness_connectivity_message"

	NsWitnessTestResponseFields = &NsWitnessTestResponseStringFields{
		ArrayName:                  &fieldArrayName,
		Role:                       &fieldRole,
		WitnessConnectivityState:   &fieldWitnessConnectivityState,
		WitnessConnectivityMessage: &fieldWitnessConnectivityMessage,
	}
}

type NsWitnessTestResponse struct {
	// ArrayName - Name of an array.
	ArrayName *string `json:"array_name,omitempty"`
	// Role - Role of an array in the group.
	Role *NsArrayRole `json:"role,omitempty"`
	// WitnessConnectivityState - Reachability status of the witness.
	WitnessConnectivityState *string `json:"witness_connectivity_state,omitempty"`
	// WitnessConnectivityMessage - Reachability message of the witness.
	WitnessConnectivityMessage *string `json:"witness_connectivity_message,omitempty"`
}

// Struct for NsWitnessTestResponseFields
type NsWitnessTestResponseStringFields struct {
	ArrayName                  *string
	Role                       *string
	WitnessConnectivityState   *string
	WitnessConnectivityMessage *string
}
