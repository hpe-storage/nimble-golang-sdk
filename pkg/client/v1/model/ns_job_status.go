// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsJobStatus Enum.

type NsJobStatus string

const (
	cNsJobStatusCanceled   NsJobStatus = "canceled"
	cNsJobStatusPending    NsJobStatus = "pending"
	cNsJobStatusInprogress NsJobStatus = "inprogress"
	cNsJobStatusDone       NsJobStatus = "done"
)

var pNsJobStatusCanceled NsJobStatus
var pNsJobStatusPending NsJobStatus
var pNsJobStatusInprogress NsJobStatus
var pNsJobStatusDone NsJobStatus

// NsJobStatusCanceled enum exports
var NsJobStatusCanceled *NsJobStatus

// NsJobStatusPending enum exports
var NsJobStatusPending *NsJobStatus

// NsJobStatusInprogress enum exports
var NsJobStatusInprogress *NsJobStatus

// NsJobStatusDone enum exports
var NsJobStatusDone *NsJobStatus

func init() {
	pNsJobStatusCanceled = cNsJobStatusCanceled
	NsJobStatusCanceled = &pNsJobStatusCanceled

	pNsJobStatusPending = cNsJobStatusPending
	NsJobStatusPending = &pNsJobStatusPending

	pNsJobStatusInprogress = cNsJobStatusInprogress
	NsJobStatusInprogress = &pNsJobStatusInprogress

	pNsJobStatusDone = cNsJobStatusDone
	NsJobStatusDone = &pNsJobStatusDone

}
