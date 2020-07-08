// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsJobResult Enum.

type NsJobResult string

const (
	cNsJobResultFailed    NsJobResult = "failed"
	cNsJobResultPartial   NsJobResult = "partial"
	cNsJobResultUnknown   NsJobResult = "unknown"
	cNsJobResultSucceeded NsJobResult = "succeeded"
)

var pNsJobResultFailed NsJobResult
var pNsJobResultPartial NsJobResult
var pNsJobResultUnknown NsJobResult
var pNsJobResultSucceeded NsJobResult

// Export
var NsJobResultFailed *NsJobResult
var NsJobResultPartial *NsJobResult
var NsJobResultUnknown *NsJobResult
var NsJobResultSucceeded *NsJobResult

func init() {
	pNsJobResultFailed = cNsJobResultFailed
	NsJobResultFailed = &pNsJobResultFailed

	pNsJobResultPartial = cNsJobResultPartial
	NsJobResultPartial = &pNsJobResultPartial

	pNsJobResultUnknown = cNsJobResultUnknown
	NsJobResultUnknown = &pNsJobResultUnknown

	pNsJobResultSucceeded = cNsJobResultSucceeded
	NsJobResultSucceeded = &pNsJobResultSucceeded

}
