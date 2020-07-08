// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

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

// Export
var NsMergeStateDestDbDone *NsMergeState
var NsMergeStateDestDb *NsMergeState
var NsMergeStateDestReassigned *NsMergeState
var NsMergeStateSrcQuiesced *NsMergeState
var NsMergeStateNone *NsMergeState
var NsMergeStateDestRelinquish *NsMergeState
var NsMergeStateSrcReassigned *NsMergeState
var NsMergeStateDestStart *NsMergeState
var NsMergeStateSrcStart *NsMergeState
var NsMergeStateSrcQuiesceFailed *NsMergeState
var NsMergeStateDestMergeValidation *NsMergeState
var NsMergeStateSrcReassignFailed *NsMergeState
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
