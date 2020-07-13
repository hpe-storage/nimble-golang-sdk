// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsBitMapReturn - Return bitmap under certain request.
// Export NsBitMapReturnFields for advance operations like search filter etc.
var NsBitMapReturnFields *NsBitMapReturn

func init() {

	NsBitMapReturnFields = &NsBitMapReturn{
		Bitmap: "bitmap",
	}
}

type NsBitMapReturn struct {
	// Bitmap - Returned bitmap.
	Bitmap string `json:"bitmap,omitempty"`
}

// sdk internal struct
type nsBitMapReturn struct {
	// Bitmap - Returned bitmap.
	Bitmap *string `json:"bitmap,omitempty"`
}

// EncodeNsBitMapReturn - Transform NsBitMapReturn to nsBitMapReturn type
func EncodeNsBitMapReturn(request interface{}) (*nsBitMapReturn, error) {
	reqNsBitMapReturn := request.(*NsBitMapReturn)
	byte, err := json.Marshal(reqNsBitMapReturn)
	objPtr := &nsBitMapReturn{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsBitMapReturn Transform nsBitMapReturn to NsBitMapReturn type
func DecodeNsBitMapReturn(request interface{}) (*NsBitMapReturn, error) {
	reqNsBitMapReturn := request.(*nsBitMapReturn)
	byte, err := json.Marshal(reqNsBitMapReturn)
	obj := &NsBitMapReturn{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
