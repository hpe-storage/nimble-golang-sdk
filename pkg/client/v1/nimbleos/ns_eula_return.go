// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsEulaReturn - Return end-user license information.
// Export NsEulaReturnFields for advance operations like search filter etc.
var NsEulaReturnFields *NsEulaReturnStringFields

func init() {
	Eulafield := "eula"

	NsEulaReturnFields = &NsEulaReturnStringFields{
		Eula: &Eulafield,
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
