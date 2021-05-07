// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// NsJWKey - Javascript Web Key.
// Export NsJWKeyFields for advance operations like search filter etc.
var NsJWKeyFields *NsJWKey

func init() {
	Kidfield := "kid"
	Keyfield := "key"

	NsJWKeyFields = &NsJWKey{
		Kid: &Kidfield,
		Key: &Keyfield,
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
