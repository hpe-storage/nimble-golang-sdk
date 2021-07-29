// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsMergeState Enum.

type NsMergeState string

const (
	cNsMergeStateDestDbDone          NsMergeState = "dest_db_done"
	cNsMergeStateDestDb              NsMergeState = "dest_db"
	cNsMergeStateDestReassigned      NsMergeState = "dest_reassigned"
	cNsMergeStateSrcQuiesced         NsMergeState = "src_quiesced"
	cNsMergeStateNone                NsMergeState = "none"
	cNsMergeStateDestRelinquish      NsMergeState = "dest_relinquish"
	cNsMergeStateSrcReassigned       NsMergeState = "src_reassigned"
	cNsMergeStateDestStart           NsMergeState = "dest_start"
	cNsMergeStateSrcStart            NsMergeState = "src_start"
	cNsMergeStateSrcQuiesceFailed    NsMergeState = "src_quiesce_failed"
	cNsMergeStateDestMergeValidation NsMergeState = "dest_merge_validation"
	cNsMergeStateSrcReassignFailed   NsMergeState = "src_reassign_failed"
	cNsMergeStateDestDbFailed        NsMergeState = "dest_db_failed"
)

var pNsMergeStateDestDbDone NsMergeState
var pNsMergeStateDestDb NsMergeState
var pNsMergeStateDestReassigned NsMergeState
var pNsMergeStateSrcQuiesced NsMergeState
var pNsMergeStateNone NsMergeState
var pNsMergeStateDestRelinquish NsMergeState
var pNsMergeStateSrcReassigned NsMergeState
var pNsMergeStateDestStart NsMergeState
var pNsMergeStateSrcStart NsMergeState
var pNsMergeStateSrcQuiesceFailed NsMergeState
var pNsMergeStateDestMergeValidation NsMergeState
var pNsMergeStateSrcReassignFailed NsMergeState
var pNsMergeStateDestDbFailed NsMergeState

// NsMergeStateDestDbDone enum exports
var NsMergeStateDestDbDone *NsMergeState

// NsMergeStateDestDb enum exports
var NsMergeStateDestDb *NsMergeState

// NsMergeStateDestReassigned enum exports
var NsMergeStateDestReassigned *NsMergeState

// NsMergeStateSrcQuiesced enum exports
var NsMergeStateSrcQuiesced *NsMergeState

// NsMergeStateNone enum exports
var NsMergeStateNone *NsMergeState

// NsMergeStateDestRelinquish enum exports
var NsMergeStateDestRelinquish *NsMergeState

// NsMergeStateSrcReassigned enum exports
var NsMergeStateSrcReassigned *NsMergeState

// NsMergeStateDestStart enum exports
var NsMergeStateDestStart *NsMergeState

// NsMergeStateSrcStart enum exports
var NsMergeStateSrcStart *NsMergeState

// NsMergeStateSrcQuiesceFailed enum exports
var NsMergeStateSrcQuiesceFailed *NsMergeState

// NsMergeStateDestMergeValidation enum exports
var NsMergeStateDestMergeValidation *NsMergeState

// NsMergeStateSrcReassignFailed enum exports
var NsMergeStateSrcReassignFailed *NsMergeState

// NsMergeStateDestDbFailed enum exports
var NsMergeStateDestDbFailed *NsMergeState

func init() {
	pNsMergeStateDestDbDone = cNsMergeStateDestDbDone
	NsMergeStateDestDbDone = &pNsMergeStateDestDbDone

	pNsMergeStateDestDb = cNsMergeStateDestDb
	NsMergeStateDestDb = &pNsMergeStateDestDb

	pNsMergeStateDestReassigned = cNsMergeStateDestReassigned
	NsMergeStateDestReassigned = &pNsMergeStateDestReassigned

	pNsMergeStateSrcQuiesced = cNsMergeStateSrcQuiesced
	NsMergeStateSrcQuiesced = &pNsMergeStateSrcQuiesced

	pNsMergeStateNone = cNsMergeStateNone
	NsMergeStateNone = &pNsMergeStateNone

	pNsMergeStateDestRelinquish = cNsMergeStateDestRelinquish
	NsMergeStateDestRelinquish = &pNsMergeStateDestRelinquish

	pNsMergeStateSrcReassigned = cNsMergeStateSrcReassigned
	NsMergeStateSrcReassigned = &pNsMergeStateSrcReassigned

	pNsMergeStateDestStart = cNsMergeStateDestStart
	NsMergeStateDestStart = &pNsMergeStateDestStart

	pNsMergeStateSrcStart = cNsMergeStateSrcStart
	NsMergeStateSrcStart = &pNsMergeStateSrcStart

	pNsMergeStateSrcQuiesceFailed = cNsMergeStateSrcQuiesceFailed
	NsMergeStateSrcQuiesceFailed = &pNsMergeStateSrcQuiesceFailed

	pNsMergeStateDestMergeValidation = cNsMergeStateDestMergeValidation
	NsMergeStateDestMergeValidation = &pNsMergeStateDestMergeValidation

	pNsMergeStateSrcReassignFailed = cNsMergeStateSrcReassignFailed
	NsMergeStateSrcReassignFailed = &pNsMergeStateSrcReassignFailed

	pNsMergeStateDestDbFailed = cNsMergeStateDestDbFailed
	NsMergeStateDestDbFailed = &pNsMergeStateDestDbFailed

}
