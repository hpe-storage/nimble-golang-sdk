// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsKeyValueFields provides field names to use in filter parameters, for example.
var NsKeyValueFields *NsKeyValueFieldHandles

func init() {
	NsKeyValueFields = &NsKeyValueFieldHandles{
		Key:   "key",
		Value: "value",
	}
}

// NsKeyValue - Key-value pair.
type NsKeyValue struct {
	// Key - Identifier of key-value pair.
	Key *string `json:"key,omitempty"`
	// Value - Value of key-value pair.
	Value *string `json:"value,omitempty"`
}

// NsKeyValueFieldHandles provides a string representation for each NsKeyValue field.
type NsKeyValueFieldHandles struct {
	Key   string
	Value string
}
