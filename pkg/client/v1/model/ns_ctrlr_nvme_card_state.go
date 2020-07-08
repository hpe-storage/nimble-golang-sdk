// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsCtrlrNvmeCardState Enum.

type NsCtrlrNvmeCardState string

const (
	cNsCtrlrNvmeCardStateValid  NsCtrlrNvmeCardState = "valid"
	cNsCtrlrNvmeCardStateInUse  NsCtrlrNvmeCardState = "in use"
	cNsCtrlrNvmeCardStateFailed NsCtrlrNvmeCardState = "failed"
)

var pNsCtrlrNvmeCardStateValid NsCtrlrNvmeCardState
var pNsCtrlrNvmeCardStateInUse NsCtrlrNvmeCardState
var pNsCtrlrNvmeCardStateFailed NsCtrlrNvmeCardState

// Export
var NsCtrlrNvmeCardStateValid *NsCtrlrNvmeCardState
var NsCtrlrNvmeCardStateInUse *NsCtrlrNvmeCardState
var NsCtrlrNvmeCardStateFailed *NsCtrlrNvmeCardState

func init() {
	pNsCtrlrNvmeCardStateValid = cNsCtrlrNvmeCardStateValid
	NsCtrlrNvmeCardStateValid = &pNsCtrlrNvmeCardStateValid

	pNsCtrlrNvmeCardStateInUse = cNsCtrlrNvmeCardStateInUse
	NsCtrlrNvmeCardStateInUse = &pNsCtrlrNvmeCardStateInUse

	pNsCtrlrNvmeCardStateFailed = cNsCtrlrNvmeCardStateFailed
	NsCtrlrNvmeCardStateFailed = &pNsCtrlrNvmeCardStateFailed

}
