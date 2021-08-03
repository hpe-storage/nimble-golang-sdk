// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsISCSIIQNFields provides field names to use in filter parameters, for example.
var NsISCSIIQNFields *NsISCSIIQNFieldHandles

func init() {
	NsISCSIIQNFields = &NsISCSIIQNFieldHandles{
		Name: "name",
	}
}

// NsISCSIIQN - ISCSI IQN.
type NsISCSIIQN struct {
	// Name - IQN name of the iSCSI initiator.
	Name *string `json:"name,omitempty"`
}

// NsISCSIIQNFieldHandles provides a string representation for each NsISCSIIQN field.
type NsISCSIIQNFieldHandles struct {
	Name string
}
