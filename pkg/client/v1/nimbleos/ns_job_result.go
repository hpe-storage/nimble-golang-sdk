// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsJobResult Enum.

type NsJobResult string

const (
 cNsJobResultFailed NsJobResult = "failed"
 cNsJobResultPartial NsJobResult = "partial"
 cNsJobResultUnknown NsJobResult = "unknown"
 cNsJobResultSucceeded NsJobResult = "succeeded"
)

var pNsJobResultFailed NsJobResult
var pNsJobResultPartial NsJobResult
var pNsJobResultUnknown NsJobResult
var pNsJobResultSucceeded NsJobResult

// NsJobResultFailed enum exports
var NsJobResultFailed *NsJobResult

// NsJobResultPartial enum exports
var NsJobResultPartial *NsJobResult

// NsJobResultUnknown enum exports
var NsJobResultUnknown *NsJobResult

// NsJobResultSucceeded enum exports
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

