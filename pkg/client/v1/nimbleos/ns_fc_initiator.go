// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFCInitiator - Fibre Channel initiator.
// Export NsFCInitiatorFields for advance operations like search filter etc.
var NsFCInitiatorFields *NsFCInitiator

func init() {
	IDfield := "id"
	InitiatorIdfield := "initiator_id"
	Wwpnfield := "wwpn"
	Aliasfield := "alias"

	NsFCInitiatorFields = &NsFCInitiator{
		ID:          &IDfield,
		InitiatorId: &InitiatorIdfield,
		Wwpn:        &Wwpnfield,
		Alias:       &Aliasfield,
	}
}

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
