// Copyright 2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsApiSeverityLevel Enum.

type NsApiSeverityLevel string

const (
	cNsApiSeverityLevelWarn  NsApiSeverityLevel = "warn"
	cNsApiSeverityLevelDebug NsApiSeverityLevel = "debug"
	cNsApiSeverityLevelError NsApiSeverityLevel = "error"
	cNsApiSeverityLevelFatal NsApiSeverityLevel = "fatal"
	cNsApiSeverityLevelInfo  NsApiSeverityLevel = "info"
)

var pNsApiSeverityLevelWarn NsApiSeverityLevel
var pNsApiSeverityLevelDebug NsApiSeverityLevel
var pNsApiSeverityLevelError NsApiSeverityLevel
var pNsApiSeverityLevelFatal NsApiSeverityLevel
var pNsApiSeverityLevelInfo NsApiSeverityLevel

// NsApiSeverityLevelWarn enum exports
var NsApiSeverityLevelWarn *NsApiSeverityLevel

// NsApiSeverityLevelDebug enum exports
var NsApiSeverityLevelDebug *NsApiSeverityLevel

// NsApiSeverityLevelError enum exports
var NsApiSeverityLevelError *NsApiSeverityLevel

// NsApiSeverityLevelFatal enum exports
var NsApiSeverityLevelFatal *NsApiSeverityLevel

// NsApiSeverityLevelInfo enum exports
var NsApiSeverityLevelInfo *NsApiSeverityLevel

func init() {
	pNsApiSeverityLevelWarn = cNsApiSeverityLevelWarn
	NsApiSeverityLevelWarn = &pNsApiSeverityLevelWarn

	pNsApiSeverityLevelDebug = cNsApiSeverityLevelDebug
	NsApiSeverityLevelDebug = &pNsApiSeverityLevelDebug

	pNsApiSeverityLevelError = cNsApiSeverityLevelError
	NsApiSeverityLevelError = &pNsApiSeverityLevelError

	pNsApiSeverityLevelFatal = cNsApiSeverityLevelFatal
	NsApiSeverityLevelFatal = &pNsApiSeverityLevelFatal

	pNsApiSeverityLevelInfo = cNsApiSeverityLevelInfo
	NsApiSeverityLevelInfo = &pNsApiSeverityLevelInfo

}
