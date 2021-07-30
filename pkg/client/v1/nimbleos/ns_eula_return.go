// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsEulaReturn - Return end-user license information.

// Export NsEulaReturnFields provides field names to use in filter parameters, for example.
var NsEulaReturnFields *NsEulaReturnStringFields

func init() {
	fieldEula := "eula"

	NsEulaReturnFields = &NsEulaReturnStringFields{
		Eula: &fieldEula,
	}
}

type NsEulaReturn struct {
	// Eula - License information.
	Eula *string `json:"eula,omitempty"`
}

// Struct for NsEulaReturnFields
type NsEulaReturnStringFields struct {
	Eula *string
}
