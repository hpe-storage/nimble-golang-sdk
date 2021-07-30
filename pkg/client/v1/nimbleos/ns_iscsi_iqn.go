// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsISCSIIQN - ISCSI IQN.
// Export NsISCSIIQNFields for advance operations like search filter etc.
var NsISCSIIQNFields *NsISCSIIQNStringFields

func init() {
	Namefield := "name"

	NsISCSIIQNFields = &NsISCSIIQNStringFields{
		Name: &Namefield,
	}
}

type NsISCSIIQN struct {
	// Name - IQN name of the iSCSI initiator.
	Name *string `json:"name,omitempty"`
}

// Struct for NsISCSIIQNFields
type NsISCSIIQNStringFields struct {
	Name *string
}
