// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsJWKey - Javascript Web Key.
// Export NsJWKeyFields for advance operations like search filter etc.
var NsJWKeyFields *NsJWKeyStringFields

func init() {
	Kidfield := "kid"
	Algorithmfield := "algorithm"
	Keyfield := "key"

	NsJWKeyFields = &NsJWKeyStringFields{
		Kid:       &Kidfield,
		Algorithm: &Algorithmfield,
		Key:       &Keyfield,
	}
}

type NsJWKey struct {
	// Kid - Key ID.
	Kid *string `json:"kid,omitempty"`
	// Algorithm - The algorithm and key length of this key.
	Algorithm *NsJWKAlgorithm `json:"algorithm,omitempty"`
	// Key - Public Key.
	Key *string `json:"key,omitempty"`
}

// Struct for NsJWKeyFields
type NsJWKeyStringFields struct {
	Kid       *string
	Algorithm *string
	Key       *string
}
