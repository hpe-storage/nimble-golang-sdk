// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsObjectIDKVFields provides field names to use in filter parameters, for example.
var NsObjectIDKVFields *NsObjectIDKVFieldHandles

func init() {
	NsObjectIDKVFields = &NsObjectIDKVFieldHandles{
		Key: "key",
		ID:  "id",
	}
}

// NsObjectIDKV - A key value pair containing an object ID as a value.
type NsObjectIDKV struct {
	// Key - Identifier of key-value pair.
	Key *string `json:"key,omitempty"`
	// ID - An object ID.
	ID *string `json:"id,omitempty"`
}

// NsObjectIDKVFieldHandles provides a string representation for each AccessControlRecord field.
type NsObjectIDKVFieldHandles struct {
	Key string
	ID  string
}
