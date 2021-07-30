// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// FibreChannelInitiatorAlias - This API provides the alias information for Fibre Channel initiators.
// Export FibreChannelInitiatorAliasFields for advance operations like search filter etc.
var FibreChannelInitiatorAliasFields *FibreChannelInitiatorAliasStringFields

func init() {
	IDfield := "id"
	Aliasfield := "alias"
	Wwpnfield := "wwpn"
	Sourcefield := "source"

	FibreChannelInitiatorAliasFields = &FibreChannelInitiatorAliasStringFields{
		ID:     &IDfield,
		Alias:  &Aliasfield,
		Wwpn:   &Wwpnfield,
		Source: &Sourcefield,
	}
}

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

// Struct for FibreChannelInitiatorAliasFields
type FibreChannelInitiatorAliasStringFields struct {
	ID     *string
	Alias  *string
	Wwpn   *string
	Source *string
}
