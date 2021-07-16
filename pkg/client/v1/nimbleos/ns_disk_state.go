// Copyright 2021 Hewlett Packard Enterprise Development LP
package nimbleos

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

// NsDiskStateValid enum exports
var NsDiskStateValid *NsDiskState

// NsDiskStateVoid enum exports
var NsDiskStateVoid *NsDiskState

// NsDiskStateRemoved enum exports
var NsDiskStateRemoved *NsDiskState

// NsDiskStateTFail enum exports
var NsDiskStateTFail *NsDiskState

// NsDiskStateInUse enum exports
var NsDiskStateInUse *NsDiskState

// NsDiskStateAbsent enum exports
var NsDiskStateAbsent *NsDiskState

// NsDiskStateFailed enum exports
var NsDiskStateFailed *NsDiskState

// NsDiskStateForeign enum exports
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
