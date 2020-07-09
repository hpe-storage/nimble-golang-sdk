// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsTraceLevel Enum.

type NsTraceLevel string

const (
	cNsTraceLevelNote     NsTraceLevel = "note"
	cNsTraceLevelWarn     NsTraceLevel = "warn"
	cNsTraceLevelCritical NsTraceLevel = "critical"
	cNsTraceLevelError    NsTraceLevel = "error"
	cNsTraceLevelInfo     NsTraceLevel = "info"
)

var pNsTraceLevelNote NsTraceLevel
var pNsTraceLevelWarn NsTraceLevel
var pNsTraceLevelCritical NsTraceLevel
var pNsTraceLevelError NsTraceLevel
var pNsTraceLevelInfo NsTraceLevel

// NsTraceLevelNote enum exports
var NsTraceLevelNote *NsTraceLevel

// NsTraceLevelWarn enum exports
var NsTraceLevelWarn *NsTraceLevel

// NsTraceLevelCritical enum exports
var NsTraceLevelCritical *NsTraceLevel

// NsTraceLevelError enum exports
var NsTraceLevelError *NsTraceLevel

// NsTraceLevelInfo enum exports
var NsTraceLevelInfo *NsTraceLevel

func init() {
	pNsTraceLevelNote = cNsTraceLevelNote
	NsTraceLevelNote = &pNsTraceLevelNote

	pNsTraceLevelWarn = cNsTraceLevelWarn
	NsTraceLevelWarn = &pNsTraceLevelWarn

	pNsTraceLevelCritical = cNsTraceLevelCritical
	NsTraceLevelCritical = &pNsTraceLevelCritical

	pNsTraceLevelError = cNsTraceLevelError
	NsTraceLevelError = &pNsTraceLevelError

	pNsTraceLevelInfo = cNsTraceLevelInfo
	NsTraceLevelInfo = &pNsTraceLevelInfo

}
