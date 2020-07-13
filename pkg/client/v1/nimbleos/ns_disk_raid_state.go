// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

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

// NsDiskRaidStateOkay enum exports
var NsDiskRaidStateOkay *NsDiskRaidState

// NsDiskRaidStateNA enum exports
var NsDiskRaidStateNA *NsDiskRaidState

// NsDiskRaidStateResynchronizing enum exports
var NsDiskRaidStateResynchronizing *NsDiskRaidState

// NsDiskRaidStateFaulty enum exports
var NsDiskRaidStateFaulty *NsDiskRaidState

// NsDiskRaidStateSpare enum exports
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
