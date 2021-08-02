// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsJWKeyFields provides field names to use in filter parameters, for example.
var NsJWKeyFields *NsJWKeyFieldHandles

func init() {
	NsJWKeyFields = &NsJWKeyFieldHandles{
		Kid:       "kid",
		Algorithm: "algorithm",
		Key:       "key",
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
	Kid       string
	Algorithm string
	Key       string
}
