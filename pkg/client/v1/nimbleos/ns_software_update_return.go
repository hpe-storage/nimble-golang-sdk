// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSoftwareUpdateReturn - The status returned by the software update precheck and start actions.
// Export NsSoftwareUpdateReturnFields for advance operations like search filter etc.
var NsSoftwareUpdateReturnFields *NsSoftwareUpdateReturnStringFields

func init() {
	Errorfield := "error"
	ArrayResponseListfield := "array_response_list"

	NsSoftwareUpdateReturnFields = &NsSoftwareUpdateReturnStringFields{
		Error:             &Errorfield,
		ArrayResponseList: &ArrayResponseListfield,
	}
}

type NsSoftwareUpdateReturn struct {
	// Error - Top level error.
	Error *string `json:"error,omitempty"`
	// ArrayResponseList - Errors from all the arrays in the group.
	ArrayResponseList []*NsArraySoftwareUpdateStatus `json:"array_response_list,omitempty"`
}

// Struct for NsSoftwareUpdateReturnFields
type NsSoftwareUpdateReturnStringFields struct {
	Error             *string
	ArrayResponseList *string
}
