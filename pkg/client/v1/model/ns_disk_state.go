// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsDiskState Enum.

type NsDiskState string

const (
	cNsDiskStateValid   NsDiskState = "valid"
	cNsDiskStateVoid    NsDiskState = "void"
	cNsDiskStateRemoved NsDiskState = "removed"
	cNsDiskStateTFail   NsDiskState = "t_fail"
	cNsDiskStateInUse   NsDiskState = "in use"
	cNsDiskStateAbsent  NsDiskState = "absent"
	cNsDiskStateFailed  NsDiskState = "failed"
	cNsDiskStateForeign NsDiskState = "foreign"
)

var pNsDiskStateValid NsDiskState
var pNsDiskStateVoid NsDiskState
var pNsDiskStateRemoved NsDiskState
var pNsDiskStateTFail NsDiskState
var pNsDiskStateInUse NsDiskState
var pNsDiskStateAbsent NsDiskState
var pNsDiskStateFailed NsDiskState
var pNsDiskStateForeign NsDiskState

// Export
var NsDiskStateValid *NsDiskState
var NsDiskStateVoid *NsDiskState
var NsDiskStateRemoved *NsDiskState
var NsDiskStateTFail *NsDiskState
var NsDiskStateInUse *NsDiskState
var NsDiskStateAbsent *NsDiskState
var NsDiskStateFailed *NsDiskState
var NsDiskStateForeign *NsDiskState

func init() {
	pNsDiskStateValid = cNsDiskStateValid
	NsDiskStateValid = &pNsDiskStateValid

	pNsDiskStateVoid = cNsDiskStateVoid
	NsDiskStateVoid = &pNsDiskStateVoid

	pNsDiskStateRemoved = cNsDiskStateRemoved
	NsDiskStateRemoved = &pNsDiskStateRemoved

	pNsDiskStateTFail = cNsDiskStateTFail
	NsDiskStateTFail = &pNsDiskStateTFail

	pNsDiskStateInUse = cNsDiskStateInUse
	NsDiskStateInUse = &pNsDiskStateInUse

	pNsDiskStateAbsent = cNsDiskStateAbsent
	NsDiskStateAbsent = &pNsDiskStateAbsent

	pNsDiskStateFailed = cNsDiskStateFailed
	NsDiskStateFailed = &pNsDiskStateFailed

	pNsDiskStateForeign = cNsDiskStateForeign
	NsDiskStateForeign = &pNsDiskStateForeign

}
