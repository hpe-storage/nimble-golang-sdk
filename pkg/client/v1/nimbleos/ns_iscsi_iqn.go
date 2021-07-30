// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsISCSIIQN - ISCSI IQN.

// Export NsISCSIIQNFields provides field names to use in filter parameters, for example.
var NsISCSIIQNFields *NsISCSIIQNStringFields

func init() {
	fieldName := "name"

	NsISCSIIQNFields = &NsISCSIIQNStringFields{
		Name: &fieldName,
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
