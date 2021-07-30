// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsKeyValue - Key-value pair.

// Export NsKeyValueFields provides field names to use in filter parameters, for example.
var NsKeyValueFields *NsKeyValueStringFields

func init() {
	fieldKey := "key"
	fieldValue := "value"

	NsKeyValueFields = &NsKeyValueStringFields{
		Key:   &fieldKey,
		Value: &fieldValue,
	}
}

type NsKeyValue struct {
	// Key - Identifier of key-value pair.
	Key *string `json:"key,omitempty"`
	// Value - Value of key-value pair.
	Value *string `json:"value,omitempty"`
}

// Struct for NsKeyValueFields
type NsKeyValueStringFields struct {
	Key   *string
	Value *string
}
