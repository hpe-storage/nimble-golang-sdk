// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFCInitiator - Fibre Channel initiator.

// Export NsFCInitiatorFields provides field names to use in filter parameters, for example.
var NsFCInitiatorFields *NsFCInitiatorStringFields

func init() {
	fieldID := "id"
	fieldInitiatorId := "initiator_id"
	fieldWwpn := "wwpn"
	fieldAlias := "alias"

	NsFCInitiatorFields = &NsFCInitiatorStringFields{
		ID:          &fieldID,
		InitiatorId: &fieldInitiatorId,
		Wwpn:        &fieldWwpn,
		Alias:       &fieldAlias,
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

// Struct for NsFCInitiatorFields
type NsFCInitiatorStringFields struct {
	ID          *string
	InitiatorId *string
	Wwpn        *string
	Alias       *string
}
