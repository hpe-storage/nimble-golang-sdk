// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsPowerSupplyStatus Enum.

type NsPowerSupplyStatus string

const (
	cNsPowerSupplyStatusPsAlerted NsPowerSupplyStatus = "ps_alerted"
	cNsPowerSupplyStatusPsOkay    NsPowerSupplyStatus = "ps_okay"
	cNsPowerSupplyStatusPsFailed  NsPowerSupplyStatus = "ps_failed"
	cNsPowerSupplyStatusPsUnknown NsPowerSupplyStatus = "ps_unknown"
)

var pNsPowerSupplyStatusPsAlerted NsPowerSupplyStatus
var pNsPowerSupplyStatusPsOkay NsPowerSupplyStatus
var pNsPowerSupplyStatusPsFailed NsPowerSupplyStatus
var pNsPowerSupplyStatusPsUnknown NsPowerSupplyStatus

// NsPowerSupplyStatusPsAlerted enum exports
var NsPowerSupplyStatusPsAlerted *NsPowerSupplyStatus

// NsPowerSupplyStatusPsOkay enum exports
var NsPowerSupplyStatusPsOkay *NsPowerSupplyStatus

// NsPowerSupplyStatusPsFailed enum exports
var NsPowerSupplyStatusPsFailed *NsPowerSupplyStatus

// NsPowerSupplyStatusPsUnknown enum exports
var NsPowerSupplyStatusPsUnknown *NsPowerSupplyStatus

func init() {
	pNsPowerSupplyStatusPsAlerted = cNsPowerSupplyStatusPsAlerted
	NsPowerSupplyStatusPsAlerted = &pNsPowerSupplyStatusPsAlerted

	pNsPowerSupplyStatusPsOkay = cNsPowerSupplyStatusPsOkay
	NsPowerSupplyStatusPsOkay = &pNsPowerSupplyStatusPsOkay

	pNsPowerSupplyStatusPsFailed = cNsPowerSupplyStatusPsFailed
	NsPowerSupplyStatusPsFailed = &pNsPowerSupplyStatusPsFailed

	pNsPowerSupplyStatusPsUnknown = cNsPowerSupplyStatusPsUnknown
	NsPowerSupplyStatusPsUnknown = &pNsPowerSupplyStatusPsUnknown

}
