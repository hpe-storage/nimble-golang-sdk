// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFCInitiatorFields provides field names to use in filter parameters, for example.
var NsFCInitiatorFields *NsFCInitiatorFieldHandles

func init() {
	NsFCInitiatorFields = &NsFCInitiatorFieldHandles{
		ID:          "id",
		InitiatorId: "initiator_id",
		Wwpn:        "wwpn",
		Alias:       "alias",
	}
}

// NsFCInitiator - Fibre Channel initiator.
type NsFCInitiator struct {
	// ID - Unique identifier of the Fibre Channel initiator.
	ID *string `json:"id,omitempty"`
	// InitiatorId - Unique identifier of the Fibre Channel initiator.
	InitiatorId *string `json:"initiator_id,omitempty"`
	// Wwpn - WWPN (World Wide Port Name) of the Fibre Channel initiator.
	Wwpn *string `json:"wwpn,omitempty"`
	// Alias - Alias of the Fibre Channel initiator.
	Alias *string `json:"alias,omitempty"`
}

// NsFCInitiatorFieldHandles provides a string representation for each NsFCInitiator field.
type NsFCInitiatorFieldHandles struct {
	ID          string
	InitiatorId string
	Wwpn        string
	Alias       string
}
