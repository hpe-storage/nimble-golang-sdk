// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export DnaFields provides field names to use in filter parameters, for example.
var DnaFields *DnaFieldHandles

func init() {
	fieldID := "id"
	fieldArrayName := "array_name"
	fieldCreationTime := "creation_time"
	fieldFileSize := "file_size"
	fieldDnaName := "dna_name"
	fieldControllerName := "controller_name"

	DnaFields = &DnaFieldHandles{
		ID:             &fieldID,
		ArrayName:      &fieldArrayName,
		CreationTime:   &fieldCreationTime,
		FileSize:       &fieldFileSize,
		DnaName:        &fieldDnaName,
		ControllerName: &fieldControllerName,
	}
}

// Dna - Get list of DNAs.
type Dna struct {
	// ID - Identifier for the DNA.
	ID *string `json:"id,omitempty"`
	// ArrayName - Name of array holding the DNA.
	ArrayName *string `json:"array_name,omitempty"`
	// CreationTime - Creation time for the DNA.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// FileSize - Size of the DNA in bytes.
	FileSize *int64 `json:"file_size,omitempty"`
	// DnaName - Name of the DNA, unique in the context of its controller.
	DnaName *string `json:"dna_name,omitempty"`
	// ControllerName - Name of the controller storing the DNA.
	ControllerName *NsControllerSide `json:"controller_name,omitempty"`
}

// DnaFieldHandles provides a string representation for each AccessControlRecord field.
type DnaFieldHandles struct {
	ID             *string
	ArrayName      *string
	CreationTime   *string
	FileSize       *string
	DnaName        *string
	ControllerName *string
}
