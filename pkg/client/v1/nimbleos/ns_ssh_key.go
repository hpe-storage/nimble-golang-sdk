// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsSshKeyFields provides field names to use in filter parameters, for example.
var NsSshKeyFields *NsSshKeyFieldHandles

func init() {
	fieldKeyName := "key_name"
	fieldKeyType := "key_type"
	fieldKey := "key"

	NsSshKeyFields = &NsSshKeyFieldHandles{
		KeyName: &fieldKeyName,
		KeyType: &fieldKeyType,
		Key:     &fieldKey,
	}
}

// NsSshKey - SSH key.
type NsSshKey struct {
	// KeyName - The user that owns the key.
	KeyName *string `json:"key_name,omitempty"`
	// KeyType - The key type.
	KeyType *string `json:"key_type,omitempty"`
	// Key - The key.
	Key *string `json:"key,omitempty"`
}

// NsSshKeyFieldHandles provides a string representation for each AccessControlRecord field.
type NsSshKeyFieldHandles struct {
	KeyName *string
	KeyType *string
	Key     *string
}
