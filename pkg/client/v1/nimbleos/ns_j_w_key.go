// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsJWKeyFields provides field names to use in filter parameters, for example.
var NsJWKeyFields *NsJWKeyFieldHandles

func init() {
	fieldKid := "kid"
	fieldAlgorithm := "algorithm"
	fieldKey := "key"

	NsJWKeyFields = &NsJWKeyFieldHandles{
		Kid:       &fieldKid,
		Algorithm: &fieldAlgorithm,
		Key:       &fieldKey,
	}
}

// NsJWKey - Javascript Web Key.
type NsJWKey struct {
	// Kid - Key ID.
	Kid *string `json:"kid,omitempty"`
	// Algorithm - The algorithm and key length of this key.
	Algorithm *NsJWKAlgorithm `json:"algorithm,omitempty"`
	// Key - Public Key.
	Key *string `json:"key,omitempty"`
}

// NsJWKeyFieldHandles provides a string representation for each AccessControlRecord field.
type NsJWKeyFieldHandles struct {
	Kid       *string
	Algorithm *string
	Key       *string
}
