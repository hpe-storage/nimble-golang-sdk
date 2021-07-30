// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsBitMapReturn - Return bitmap under certain request.
// Export NsBitMapReturnFields for advance operations like search filter etc.
var NsBitMapReturnFields *NsBitMapReturnStringFields

func init() {
	Bitmapfield := "bitmap"

	NsBitMapReturnFields = &NsBitMapReturnStringFields{
		Bitmap: &Bitmapfield,
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
