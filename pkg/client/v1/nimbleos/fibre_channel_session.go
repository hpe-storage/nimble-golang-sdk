// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// FibreChannelSession - Fibre Channel session is created when Fibre Channel initiator connects to this group.
// Export FibreChannelSessionFields for advance operations like search filter etc.
var FibreChannelSessionFields *FibreChannelSessionStringFields

func init() {
	IDfield := "id"
	InitiatorInfofield := "initiator_info"
	TargetInfofield := "target_info"

	FibreChannelSessionFields = &FibreChannelSessionStringFields{
		ID:            &IDfield,
		InitiatorInfo: &InitiatorInfofield,
		TargetInfo:    &TargetInfofield,
	}
}

type FibreChannelSession struct {
	// ID - Unique identifier of the Fibre Channel session.
	ID *string `json:"id,omitempty"`
	// InitiatorInfo - Information about the Fibre Channel initiator.
	InitiatorInfo *NsFcSessionInitiator `json:"initiator_info,omitempty"`
	// TargetInfo - Information about the Fibre Channel target.
	TargetInfo *NsFcSessionTarget `json:"target_info,omitempty"`
}

// Struct for FibreChannelSessionFields
type FibreChannelSessionStringFields struct {
	ID            *string
	InitiatorInfo *string
	TargetInfo    *string
}
