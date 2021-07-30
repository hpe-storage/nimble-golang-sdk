// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsSoftwareUpdateReturnFields provides field names to use in filter parameters, for example.
var NsSoftwareUpdateReturnFields *NsSoftwareUpdateReturnFieldHandles

func init() {
	NsSoftwareUpdateReturnFields = &NsSoftwareUpdateReturnFieldHandles{
		Error:             "error",
		ArrayResponseList: "array_response_list",
	}
}

// NsSoftwareUpdateReturn - The status returned by the software update precheck and start actions.
type NsSoftwareUpdateReturn struct {
	// Error - Top level error.
	Error *string `json:"error,omitempty"`
	// ArrayResponseList - Errors from all the arrays in the group.
	ArrayResponseList []*NsArraySoftwareUpdateStatus `json:"array_response_list,omitempty"`
}

// NsSoftwareUpdateReturnFieldHandles provides a string representation for each AccessControlRecord field.
type NsSoftwareUpdateReturnFieldHandles struct {
	Error             string
	ArrayResponseList string
}
