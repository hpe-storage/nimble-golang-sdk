// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

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

// Export
var NsApiSeverityLevelWarn *NsApiSeverityLevel
var NsApiSeverityLevelDebug *NsApiSeverityLevel
var NsApiSeverityLevelError *NsApiSeverityLevel
var NsApiSeverityLevelFatal *NsApiSeverityLevel
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
