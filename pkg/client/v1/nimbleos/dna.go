// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Dna - Get list of DNAs.
// Export DnaFields for advance operations like search filter etc.
var DnaFields *DnaStringFields

func init() {
	IDfield := "id"
	ArrayNamefield := "array_name"
	CreationTimefield := "creation_time"
	FileSizefield := "file_size"
	DnaNamefield := "dna_name"
	ControllerNamefield := "controller_name"

	DnaFields = &DnaStringFields{
		ID:             &IDfield,
		ArrayName:      &ArrayNamefield,
		CreationTime:   &CreationTimefield,
		FileSize:       &FileSizefield,
		DnaName:        &DnaNamefield,
		ControllerName: &ControllerNamefield,
	}
}

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

// Struct for DnaFields
type DnaStringFields struct {
	ID             *string
	ArrayName      *string
	CreationTime   *string
	FileSize       *string
	DnaName        *string
	ControllerName *string
}
