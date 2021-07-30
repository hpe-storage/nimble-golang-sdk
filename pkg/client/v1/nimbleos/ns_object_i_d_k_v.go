// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsObjectIDKV - A key value pair containing an object ID as a value.

// Export NsObjectIDKVFields provides field names to use in filter parameters, for example.
var NsObjectIDKVFields *NsObjectIDKVStringFields

func init() {
	fieldKey := "key"
	fieldID := "id"

	NsObjectIDKVFields = &NsObjectIDKVStringFields{
		Key: &fieldKey,
		ID:  &fieldID,
	}
}

type NsObjectIDKV struct {
	// Key - Identifier of key-value pair.
	Key *string `json:"key,omitempty"`
	// ID - An object ID.
	ID *string `json:"id,omitempty"`
}

// Struct for NsObjectIDKVFields
type NsObjectIDKVStringFields struct {
	Key *string
	ID  *string
}
