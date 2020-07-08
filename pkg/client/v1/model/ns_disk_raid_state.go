// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsDiskRaidState Enum.

type NsDiskRaidState string

const (
	cNsDiskRaidStateOkay            NsDiskRaidState = "okay"
	cNsDiskRaidStateNA              NsDiskRaidState = "N/A"
	cNsDiskRaidStateResynchronizing NsDiskRaidState = "resynchronizing"
	cNsDiskRaidStateFaulty          NsDiskRaidState = "faulty"
	cNsDiskRaidStateSpare           NsDiskRaidState = "spare"
)

var pNsDiskRaidStateOkay NsDiskRaidState
var pNsDiskRaidStateNA NsDiskRaidState
var pNsDiskRaidStateResynchronizing NsDiskRaidState
var pNsDiskRaidStateFaulty NsDiskRaidState
var pNsDiskRaidStateSpare NsDiskRaidState

// Export
var NsDiskRaidStateOkay *NsDiskRaidState
var NsDiskRaidStateNA *NsDiskRaidState
var NsDiskRaidStateResynchronizing *NsDiskRaidState
var NsDiskRaidStateFaulty *NsDiskRaidState
var NsDiskRaidStateSpare *NsDiskRaidState

func init() {
	pNsDiskRaidStateOkay = cNsDiskRaidStateOkay
	NsDiskRaidStateOkay = &pNsDiskRaidStateOkay

	pNsDiskRaidStateNA = cNsDiskRaidStateNA
	NsDiskRaidStateNA = &pNsDiskRaidStateNA

	pNsDiskRaidStateResynchronizing = cNsDiskRaidStateResynchronizing
	NsDiskRaidStateResynchronizing = &pNsDiskRaidStateResynchronizing

	pNsDiskRaidStateFaulty = cNsDiskRaidStateFaulty
	NsDiskRaidStateFaulty = &pNsDiskRaidStateFaulty

	pNsDiskRaidStateSpare = cNsDiskRaidStateSpare
	NsDiskRaidStateSpare = &pNsDiskRaidStateSpare

}
