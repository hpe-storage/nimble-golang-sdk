// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// FibreChannelInitiatorAliasFields provides field names to use in filter parameters, for example.
var FibreChannelInitiatorAliasFields *FibreChannelInitiatorAliasFieldHandles

func init() {
	FibreChannelInitiatorAliasFields = &FibreChannelInitiatorAliasFieldHandles{
		ID:     "id",
		Alias:  "alias",
		Wwpn:   "wwpn",
		Source: "source",
	}
}

// FibreChannelInitiatorAlias - This API provides the alias information for Fibre Channel initiators.
type FibreChannelInitiatorAlias struct {
	// ID - Unique identifier for the Fibre Channel initiator alias.
	ID *string `json:"id,omitempty"`
	// Alias - Alias of the Fibre Channel initiator.
	Alias *string `json:"alias,omitempty"`
	// Wwpn - WWPN (World Wide Port Name) of the Fibre Channel initiator.
	Wwpn *string `json:"wwpn,omitempty"`
	// Source - Source of the Fibre Channel initiator alias.
	Source *SmFcInitiatorAliasSource `json:"source,omitempty"`
}

// FibreChannelInitiatorAliasFieldHandles provides a string representation for each FibreChannelInitiatorAlias field.
type FibreChannelInitiatorAliasFieldHandles struct {
	ID     string
	Alias  string
	Wwpn   string
	Source string
}
