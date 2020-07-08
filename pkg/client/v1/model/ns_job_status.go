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

// Export
var NsJobStatusCanceled *NsJobStatus
var NsJobStatusPending *NsJobStatus
var NsJobStatusInprogress *NsJobStatus
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
