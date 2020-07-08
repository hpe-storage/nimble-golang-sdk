// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsArrayUpgradeState Enum.

type NsArrayUpgradeState string

const (
	cNsArrayUpgradeStateAbortFailed       NsArrayUpgradeState = "abort_failed"
	cNsArrayUpgradeStatePrecheck          NsArrayUpgradeState = "precheck"
	cNsArrayUpgradeStateInProgress        NsArrayUpgradeState = "in_progress"
	cNsArrayUpgradeStatePaused            NsArrayUpgradeState = "paused"
	cNsArrayUpgradeStateAborted           NsArrayUpgradeState = "aborted"
	cNsArrayUpgradeStateAborting          NsArrayUpgradeState = "aborting"
	cNsArrayUpgradeStateStarted           NsArrayUpgradeState = "started"
	cNsArrayUpgradeStateNone              NsArrayUpgradeState = "none"
	cNsArrayUpgradeStateFailed            NsArrayUpgradeState = "failed"
	cNsArrayUpgradeStateAwaitingNextStage NsArrayUpgradeState = "awaiting_next_stage"
	cNsArrayUpgradeStateComplete          NsArrayUpgradeState = "complete"
)

var pNsArrayUpgradeStateAbortFailed NsArrayUpgradeState
var pNsArrayUpgradeStatePrecheck NsArrayUpgradeState
var pNsArrayUpgradeStateInProgress NsArrayUpgradeState
var pNsArrayUpgradeStatePaused NsArrayUpgradeState
var pNsArrayUpgradeStateAborted NsArrayUpgradeState
var pNsArrayUpgradeStateAborting NsArrayUpgradeState
var pNsArrayUpgradeStateStarted NsArrayUpgradeState
var pNsArrayUpgradeStateNone NsArrayUpgradeState
var pNsArrayUpgradeStateFailed NsArrayUpgradeState
var pNsArrayUpgradeStateAwaitingNextStage NsArrayUpgradeState
var pNsArrayUpgradeStateComplete NsArrayUpgradeState

// Export
var NsArrayUpgradeStateAbortFailed *NsArrayUpgradeState
var NsArrayUpgradeStatePrecheck *NsArrayUpgradeState
var NsArrayUpgradeStateInProgress *NsArrayUpgradeState
var NsArrayUpgradeStatePaused *NsArrayUpgradeState
var NsArrayUpgradeStateAborted *NsArrayUpgradeState
var NsArrayUpgradeStateAborting *NsArrayUpgradeState
var NsArrayUpgradeStateStarted *NsArrayUpgradeState
var NsArrayUpgradeStateNone *NsArrayUpgradeState
var NsArrayUpgradeStateFailed *NsArrayUpgradeState
var NsArrayUpgradeStateAwaitingNextStage *NsArrayUpgradeState
var NsArrayUpgradeStateComplete *NsArrayUpgradeState

func init() {
	pNsArrayUpgradeStateAbortFailed = cNsArrayUpgradeStateAbortFailed
	NsArrayUpgradeStateAbortFailed = &pNsArrayUpgradeStateAbortFailed

	pNsArrayUpgradeStatePrecheck = cNsArrayUpgradeStatePrecheck
	NsArrayUpgradeStatePrecheck = &pNsArrayUpgradeStatePrecheck

	pNsArrayUpgradeStateInProgress = cNsArrayUpgradeStateInProgress
	NsArrayUpgradeStateInProgress = &pNsArrayUpgradeStateInProgress

	pNsArrayUpgradeStatePaused = cNsArrayUpgradeStatePaused
	NsArrayUpgradeStatePaused = &pNsArrayUpgradeStatePaused

	pNsArrayUpgradeStateAborted = cNsArrayUpgradeStateAborted
	NsArrayUpgradeStateAborted = &pNsArrayUpgradeStateAborted

	pNsArrayUpgradeStateAborting = cNsArrayUpgradeStateAborting
	NsArrayUpgradeStateAborting = &pNsArrayUpgradeStateAborting

	pNsArrayUpgradeStateStarted = cNsArrayUpgradeStateStarted
	NsArrayUpgradeStateStarted = &pNsArrayUpgradeStateStarted

	pNsArrayUpgradeStateNone = cNsArrayUpgradeStateNone
	NsArrayUpgradeStateNone = &pNsArrayUpgradeStateNone

	pNsArrayUpgradeStateFailed = cNsArrayUpgradeStateFailed
	NsArrayUpgradeStateFailed = &pNsArrayUpgradeStateFailed

	pNsArrayUpgradeStateAwaitingNextStage = cNsArrayUpgradeStateAwaitingNextStage
	NsArrayUpgradeStateAwaitingNextStage = &pNsArrayUpgradeStateAwaitingNextStage

	pNsArrayUpgradeStateComplete = cNsArrayUpgradeStateComplete
	NsArrayUpgradeStateComplete = &pNsArrayUpgradeStateComplete

}
