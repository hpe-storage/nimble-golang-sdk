// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// FibreChannelSessionFields provides field names to use in filter parameters, for example.
var FibreChannelSessionFields *FibreChannelSessionFieldHandles

func init() {
	FibreChannelSessionFields = &FibreChannelSessionFieldHandles{
		ID:            "id",
		InitiatorInfo: "initiator_info",
		TargetInfo:    "target_info",
	}
}

// FibreChannelSession - Fibre Channel session is created when Fibre Channel initiator connects to this group.
type FibreChannelSession struct {
	// ID - Unique identifier of the Fibre Channel session.
	ID *string `json:"id,omitempty"`
	// InitiatorInfo - Information about the Fibre Channel initiator.
	InitiatorInfo *NsFcSessionInitiator `json:"initiator_info,omitempty"`
	// TargetInfo - Information about the Fibre Channel target.
	TargetInfo *NsFcSessionTarget `json:"target_info,omitempty"`
}

// FibreChannelSessionFieldHandles provides a string representation for each AccessControlRecord field.
type FibreChannelSessionFieldHandles struct {
	ID            string
	InitiatorInfo string
	TargetInfo    string
}
