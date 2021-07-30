// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsISCSIInitiatorFields provides field names to use in filter parameters, for example.
var NsISCSIInitiatorFields *NsISCSIInitiatorFieldHandles

func init() {
	NsISCSIInitiatorFields = &NsISCSIInitiatorFieldHandles{
		ID:          "id",
		InitiatorId: "initiator_id",
		Label:       "label",
		Iqn:         "iqn",
		IpAddress:   "ip_address",
	}
}

// NsISCSIInitiator - ISCSI initiator.
type NsISCSIInitiator struct {
	// ID - Unique identifier of the iSCSI initiator.
	ID *string `json:"id,omitempty"`
	// InitiatorId - Unique identifier of the iSCSI initiator.
	InitiatorId *string `json:"initiator_id,omitempty"`
	// Label - Unique label of the iSCSI initiator.
	Label *string `json:"label,omitempty"`
	// Iqn - IQN name of the iSCSI initiator.
	Iqn *string `json:"iqn,omitempty"`
	// IpAddress - IP address of the iSCSI initiator.
	IpAddress *string `json:"ip_address,omitempty"`
}

// NsISCSIInitiatorFieldHandles provides a string representation for each AccessControlRecord field.
type NsISCSIInitiatorFieldHandles struct {
	ID          string
	InitiatorId string
	Label       string
	Iqn         string
	IpAddress   string
}
