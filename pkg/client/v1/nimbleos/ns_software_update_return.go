// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSoftwareUpdateReturn - The status returned by the software update precheck and start actions.

// Export NsSoftwareUpdateReturnFields provides field names to use in filter parameters, for example.
var NsSoftwareUpdateReturnFields *NsSoftwareUpdateReturnStringFields

func init() {
	fieldError := "error"
	fieldArrayResponseList := "array_response_list"

	NsSoftwareUpdateReturnFields = &NsSoftwareUpdateReturnStringFields{
		Error:             &fieldError,
		ArrayResponseList: &fieldArrayResponseList,
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
