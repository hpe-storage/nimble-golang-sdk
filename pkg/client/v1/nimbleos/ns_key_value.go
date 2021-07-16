// Copyright 2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsKeyValue - Key-value pair.
// Export NsKeyValueFields for advance operations like search filter etc.
var NsKeyValueFields *NsKeyValue

func init() {
	Keyfield := "key"
	Valuefield := "value"

	NsKeyValueFields = &NsKeyValue{
		Key:   &Keyfield,
		Value: &Valuefield,
	}
}

type NsKeyValue struct {
	// Key - Identifier of key-value pair.
	Key *string `json:"key,omitempty"`
	// Value - Value of key-value pair.
	Value *string `json:"value,omitempty"`
}
