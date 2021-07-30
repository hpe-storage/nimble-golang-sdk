// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsBitMapReturn - Return bitmap under certain request.

// Export NsBitMapReturnFields provides field names to use in filter parameters, for example.
var NsBitMapReturnFields *NsBitMapReturnStringFields

func init() {
	fieldBitmap := "bitmap"

	NsBitMapReturnFields = &NsBitMapReturnStringFields{
		Bitmap: &fieldBitmap,
	}
}

type NsBitMapReturn struct {
	// Bitmap - Returned bitmap.
	Bitmap *string `json:"bitmap,omitempty"`
}

// Struct for NsBitMapReturnFields
type NsBitMapReturnStringFields struct {
	Bitmap *string
}
