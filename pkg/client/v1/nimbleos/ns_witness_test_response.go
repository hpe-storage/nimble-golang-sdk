// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsWitnessTestResponseFields provides field names to use in filter parameters, for example.
var NsWitnessTestResponseFields *NsWitnessTestResponseFieldHandles

func init() {
	NsWitnessTestResponseFields = &NsWitnessTestResponseFieldHandles{
		ArrayName:                  "array_name",
		Role:                       "role",
		WitnessConnectivityState:   "witness_connectivity_state",
		WitnessConnectivityMessage: "witness_connectivity_message",
	}
}

// NsWitnessTestResponse - Results of witness connection test.
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

// NsWitnessTestResponseFieldHandles provides a string representation for each AccessControlRecord field.
type NsWitnessTestResponseFieldHandles struct {
	ArrayName                  string
	Role                       string
	WitnessConnectivityState   string
	WitnessConnectivityMessage string
}
