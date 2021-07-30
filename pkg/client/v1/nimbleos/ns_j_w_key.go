// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsJWKey - Javascript Web Key.

// Export NsJWKeyFields provides field names to use in filter parameters, for example.
var NsJWKeyFields *NsJWKeyStringFields

func init() {
	fieldKid := "kid"
	fieldAlgorithm := "algorithm"
	fieldKey := "key"

	NsJWKeyFields = &NsJWKeyStringFields{
		Kid:       &fieldKid,
		Algorithm: &fieldAlgorithm,
		Key:       &fieldKey,
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
