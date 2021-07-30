// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsObjectIDKV - A key value pair containing an object ID as a value.
// Export NsObjectIDKVFields for advance operations like search filter etc.
var NsObjectIDKVFields *NsObjectIDKVStringFields

func init() {
	Keyfield := "key"
	IDfield := "id"

	NsObjectIDKVFields = &NsObjectIDKVStringFields{
		Key: &Keyfield,
		ID:  &IDfield,
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
