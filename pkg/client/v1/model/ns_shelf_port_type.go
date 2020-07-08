// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsShelfPortType Enum.

type NsShelfPortType string

const (
	cNsShelfPortTypeUpstream   NsShelfPortType = "upstream"
	cNsShelfPortTypeDownstream NsShelfPortType = "downstream"
	cNsShelfPortTypeUnknown    NsShelfPortType = "unknown"
)

var pNsShelfPortTypeUpstream NsShelfPortType
var pNsShelfPortTypeDownstream NsShelfPortType
var pNsShelfPortTypeUnknown NsShelfPortType

// Export
var NsShelfPortTypeUpstream *NsShelfPortType
var NsShelfPortTypeDownstream *NsShelfPortType
var NsShelfPortTypeUnknown *NsShelfPortType

func init() {
	pNsShelfPortTypeUpstream = cNsShelfPortTypeUpstream
	NsShelfPortTypeUpstream = &pNsShelfPortTypeUpstream

	pNsShelfPortTypeDownstream = cNsShelfPortTypeDownstream
	NsShelfPortTypeDownstream = &pNsShelfPortTypeDownstream

	pNsShelfPortTypeUnknown = cNsShelfPortTypeUnknown
	NsShelfPortTypeUnknown = &pNsShelfPortTypeUnknown

}
