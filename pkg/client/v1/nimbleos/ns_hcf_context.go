// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsHcfContext Enum.

type NsHcfContext string

const (
	cNsHcfContextAll      NsHcfContext = "all"
	cNsHcfContextFailover NsHcfContext = "failover"
	cNsHcfContextSwUpdate NsHcfContext = "sw_update"
)

var pNsHcfContextAll NsHcfContext
var pNsHcfContextFailover NsHcfContext
var pNsHcfContextSwUpdate NsHcfContext

// NsHcfContextAll enum exports
var NsHcfContextAll *NsHcfContext

// NsHcfContextFailover enum exports
var NsHcfContextFailover *NsHcfContext

// NsHcfContextSwUpdate enum exports
var NsHcfContextSwUpdate *NsHcfContext

func init() {
	pNsHcfContextAll = cNsHcfContextAll
	NsHcfContextAll = &pNsHcfContextAll

	pNsHcfContextFailover = cNsHcfContextFailover
	NsHcfContextFailover = &pNsHcfContextFailover

	pNsHcfContextSwUpdate = cNsHcfContextSwUpdate
	NsHcfContextSwUpdate = &pNsHcfContextSwUpdate

}
