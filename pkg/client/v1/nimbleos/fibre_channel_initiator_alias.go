// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// FibreChannelInitiatorAlias - This API provides the alias information for Fibre Channel initiators.
// Export FibreChannelInitiatorAliasFields for advance operations like search filter etc.
var FibreChannelInitiatorAliasFields *FibreChannelInitiatorAlias

func init() {

	FibreChannelInitiatorAliasFields = &FibreChannelInitiatorAlias{
		ID:    "id",
		Alias: "alias",
		Wwpn:  "wwpn",
	}
}

type FibreChannelInitiatorAlias struct {
	// ID - Unique identifier for the Fibre Channel initiator alias.
	ID string `json:"id,omitempty"`
	// Alias - Alias of the Fibre Channel initiator.
	Alias string `json:"alias,omitempty"`
	// Wwpn - WWPN (World Wide Port Name) of the Fibre Channel initiator.
	Wwpn string `json:"wwpn,omitempty"`
	// Source - Source of the Fibre Channel initiator alias.
	Source *SmFcInitiatorAliasSource `json:"source,omitempty"`
}

// sdk internal struct
type fibreChannelInitiatorAlias struct {
	ID     *string                   `json:"id,omitempty"`
	Alias  *string                   `json:"alias,omitempty"`
	Wwpn   *string                   `json:"wwpn,omitempty"`
	Source *SmFcInitiatorAliasSource `json:"source,omitempty"`
}

// EncodeFibreChannelInitiatorAlias - Transform FibreChannelInitiatorAlias to fibreChannelInitiatorAlias type
func EncodeFibreChannelInitiatorAlias(request interface{}) (*fibreChannelInitiatorAlias, error) {
	reqFibreChannelInitiatorAlias := request.(*FibreChannelInitiatorAlias)
	byte, err := json.Marshal(reqFibreChannelInitiatorAlias)
	if err != nil {
		return nil, err
	}
	respFibreChannelInitiatorAliasPtr := &fibreChannelInitiatorAlias{}
	err = json.Unmarshal(byte, respFibreChannelInitiatorAliasPtr)
	return respFibreChannelInitiatorAliasPtr, err
}

// DecodeFibreChannelInitiatorAlias Transform fibreChannelInitiatorAlias to FibreChannelInitiatorAlias type
func DecodeFibreChannelInitiatorAlias(request interface{}) (*FibreChannelInitiatorAlias, error) {
	reqFibreChannelInitiatorAlias := request.(*fibreChannelInitiatorAlias)
	byte, err := json.Marshal(reqFibreChannelInitiatorAlias)
	if err != nil {
		return nil, err
	}
	respFibreChannelInitiatorAlias := &FibreChannelInitiatorAlias{}
	err = json.Unmarshal(byte, respFibreChannelInitiatorAlias)
	return respFibreChannelInitiatorAlias, err
}
