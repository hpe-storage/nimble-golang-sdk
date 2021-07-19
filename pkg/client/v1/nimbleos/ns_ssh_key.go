// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSshKey - SSH key.
// Export NsSshKeyFields for advance operations like search filter etc.
var NsSshKeyFields *NsSshKey

func init() {
	KeyNamefield := "key_name"
	KeyTypefield := "key_type"
	Keyfield := "key"

	NsSshKeyFields = &NsSshKey{
		KeyName: &KeyNamefield,
		KeyType: &KeyTypefield,
		Key:     &Keyfield,
	}
}

type NsSshKey struct {
	// KeyName - The user that owns the key.
	KeyName *string `json:"key_name,omitempty"`
	// KeyType - The key type.
	KeyType *string `json:"key_type,omitempty"`
	// Key - The key.
	Key *string `json:"key,omitempty"`
}
