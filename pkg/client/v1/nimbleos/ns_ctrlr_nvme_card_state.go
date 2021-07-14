// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsCtrlrNvmeCardState Enum.

type NsCtrlrNvmeCardState string

const (
 cNsCtrlrNvmeCardStateValid NsCtrlrNvmeCardState = "valid"
 cNsCtrlrNvmeCardStateInUse NsCtrlrNvmeCardState = "in use"
 cNsCtrlrNvmeCardStateFailed NsCtrlrNvmeCardState = "failed"
)

var pNsCtrlrNvmeCardStateValid NsCtrlrNvmeCardState
var pNsCtrlrNvmeCardStateInUse NsCtrlrNvmeCardState
var pNsCtrlrNvmeCardStateFailed NsCtrlrNvmeCardState

// NsCtrlrNvmeCardStateValid enum exports
var NsCtrlrNvmeCardStateValid *NsCtrlrNvmeCardState

// NsCtrlrNvmeCardStateInUse enum exports
var NsCtrlrNvmeCardStateInUse *NsCtrlrNvmeCardState

// NsCtrlrNvmeCardStateFailed enum exports
var NsCtrlrNvmeCardStateFailed *NsCtrlrNvmeCardState

func init() {
 pNsCtrlrNvmeCardStateValid = cNsCtrlrNvmeCardStateValid
 NsCtrlrNvmeCardStateValid = &pNsCtrlrNvmeCardStateValid

 pNsCtrlrNvmeCardStateInUse = cNsCtrlrNvmeCardStateInUse
 NsCtrlrNvmeCardStateInUse = &pNsCtrlrNvmeCardStateInUse

 pNsCtrlrNvmeCardStateFailed = cNsCtrlrNvmeCardStateFailed
 NsCtrlrNvmeCardStateFailed = &pNsCtrlrNvmeCardStateFailed

}

