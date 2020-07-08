// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for SmFcInitiatorAliasSource Enum.

type SmFcInitiatorAliasSource string

const (
	cSmFcInitiatorAliasSourceFabric  SmFcInitiatorAliasSource = "fabric"
	cSmFcInitiatorAliasSourceInvalid SmFcInitiatorAliasSource = "invalid"
	cSmFcInitiatorAliasSourceUser    SmFcInitiatorAliasSource = "user"
)

var pSmFcInitiatorAliasSourceFabric SmFcInitiatorAliasSource
var pSmFcInitiatorAliasSourceInvalid SmFcInitiatorAliasSource
var pSmFcInitiatorAliasSourceUser SmFcInitiatorAliasSource

// Export
var SmFcInitiatorAliasSourceFabric *SmFcInitiatorAliasSource
var SmFcInitiatorAliasSourceInvalid *SmFcInitiatorAliasSource
var SmFcInitiatorAliasSourceUser *SmFcInitiatorAliasSource

func init() {
	pSmFcInitiatorAliasSourceFabric = cSmFcInitiatorAliasSourceFabric
	SmFcInitiatorAliasSourceFabric = &pSmFcInitiatorAliasSourceFabric

	pSmFcInitiatorAliasSourceInvalid = cSmFcInitiatorAliasSourceInvalid
	SmFcInitiatorAliasSourceInvalid = &pSmFcInitiatorAliasSourceInvalid

	pSmFcInitiatorAliasSourceUser = cSmFcInitiatorAliasSourceUser
	SmFcInitiatorAliasSourceUser = &pSmFcInitiatorAliasSourceUser

}
