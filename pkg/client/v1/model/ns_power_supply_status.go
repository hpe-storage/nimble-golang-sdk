// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

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

// Export
var NsPowerSupplyStatusPsAlerted *NsPowerSupplyStatus
var NsPowerSupplyStatusPsOkay *NsPowerSupplyStatus
var NsPowerSupplyStatusPsFailed *NsPowerSupplyStatus
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
