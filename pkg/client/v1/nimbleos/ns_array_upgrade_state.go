// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

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

// NsArrayUpgradeStateAbortFailed enum exports
var NsArrayUpgradeStateAbortFailed *NsArrayUpgradeState

// NsArrayUpgradeStatePrecheck enum exports
var NsArrayUpgradeStatePrecheck *NsArrayUpgradeState

// NsArrayUpgradeStateInProgress enum exports
var NsArrayUpgradeStateInProgress *NsArrayUpgradeState

// NsArrayUpgradeStatePaused enum exports
var NsArrayUpgradeStatePaused *NsArrayUpgradeState

// NsArrayUpgradeStateAborted enum exports
var NsArrayUpgradeStateAborted *NsArrayUpgradeState

// NsArrayUpgradeStateAborting enum exports
var NsArrayUpgradeStateAborting *NsArrayUpgradeState

// NsArrayUpgradeStateStarted enum exports
var NsArrayUpgradeStateStarted *NsArrayUpgradeState

// NsArrayUpgradeStateNone enum exports
var NsArrayUpgradeStateNone *NsArrayUpgradeState

// NsArrayUpgradeStateFailed enum exports
var NsArrayUpgradeStateFailed *NsArrayUpgradeState

// NsArrayUpgradeStateAwaitingNextStage enum exports
var NsArrayUpgradeStateAwaitingNextStage *NsArrayUpgradeState

// NsArrayUpgradeStateComplete enum exports
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
