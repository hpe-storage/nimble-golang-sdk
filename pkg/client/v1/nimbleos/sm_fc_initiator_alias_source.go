// Copyright 2021 Hewlett Packard Enterprise Development LP
package nimbleos

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

// SmFcInitiatorAliasSourceFabric enum exports
var SmFcInitiatorAliasSourceFabric *SmFcInitiatorAliasSource

// SmFcInitiatorAliasSourceInvalid enum exports
var SmFcInitiatorAliasSourceInvalid *SmFcInitiatorAliasSource

// SmFcInitiatorAliasSourceUser enum exports
var SmFcInitiatorAliasSourceUser *SmFcInitiatorAliasSource

func init() {
	pSmFcInitiatorAliasSourceFabric = cSmFcInitiatorAliasSourceFabric
	SmFcInitiatorAliasSourceFabric = &pSmFcInitiatorAliasSourceFabric

	pSmFcInitiatorAliasSourceInvalid = cSmFcInitiatorAliasSourceInvalid
	SmFcInitiatorAliasSourceInvalid = &pSmFcInitiatorAliasSourceInvalid

	pSmFcInitiatorAliasSourceUser = cSmFcInitiatorAliasSourceUser
	SmFcInitiatorAliasSourceUser = &pSmFcInitiatorAliasSourceUser

}
