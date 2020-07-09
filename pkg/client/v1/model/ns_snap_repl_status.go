// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsSnapReplStatus Enum.

type NsSnapReplStatus string

const (
	cNsSnapReplStatusFail       NsSnapReplStatus = "fail"
	cNsSnapReplStatusInProgress NsSnapReplStatus = "in_progress"
	cNsSnapReplStatusPending    NsSnapReplStatus = "pending"
	cNsSnapReplStatusComplete   NsSnapReplStatus = "complete"
)

var pNsSnapReplStatusFail NsSnapReplStatus
var pNsSnapReplStatusInProgress NsSnapReplStatus
var pNsSnapReplStatusPending NsSnapReplStatus
var pNsSnapReplStatusComplete NsSnapReplStatus

// NsSnapReplStatusFail enum exports
var NsSnapReplStatusFail *NsSnapReplStatus

// NsSnapReplStatusInProgress enum exports
var NsSnapReplStatusInProgress *NsSnapReplStatus

// NsSnapReplStatusPending enum exports
var NsSnapReplStatusPending *NsSnapReplStatus

// NsSnapReplStatusComplete enum exports
var NsSnapReplStatusComplete *NsSnapReplStatus

func init() {
	pNsSnapReplStatusFail = cNsSnapReplStatusFail
	NsSnapReplStatusFail = &pNsSnapReplStatusFail

	pNsSnapReplStatusInProgress = cNsSnapReplStatusInProgress
	NsSnapReplStatusInProgress = &pNsSnapReplStatusInProgress

	pNsSnapReplStatusPending = cNsSnapReplStatusPending
	NsSnapReplStatusPending = &pNsSnapReplStatusPending

	pNsSnapReplStatusComplete = cNsSnapReplStatusComplete
	NsSnapReplStatusComplete = &pNsSnapReplStatusComplete

}
