// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsALUA Enum.

type NsALUA string

const (
	cNsALUALogicalBlockDependent NsALUA = "logical_block_dependent"
	cNsALUAOffline               NsALUA = "offline"
	cNsALUAActiveOptimized       NsALUA = "active_optimized"
	cNsALUAStandby               NsALUA = "standby"
	cNsALUATransitioning         NsALUA = "transitioning"
	cNsALUAUnavailable           NsALUA = "unavailable"
	cNsALUAInvalid               NsALUA = "invalid"
	cNsALUAActiveNonOptimized    NsALUA = "active_non_optimized"
)

var pNsALUALogicalBlockDependent NsALUA
var pNsALUAOffline NsALUA
var pNsALUAActiveOptimized NsALUA
var pNsALUAStandby NsALUA
var pNsALUATransitioning NsALUA
var pNsALUAUnavailable NsALUA
var pNsALUAInvalid NsALUA
var pNsALUAActiveNonOptimized NsALUA

// Export
var NsALUALogicalBlockDependent *NsALUA
var NsALUAOffline *NsALUA
var NsALUAActiveOptimized *NsALUA
var NsALUAStandby *NsALUA
var NsALUATransitioning *NsALUA
var NsALUAUnavailable *NsALUA
var NsALUAInvalid *NsALUA
var NsALUAActiveNonOptimized *NsALUA

func init() {
	pNsALUALogicalBlockDependent = cNsALUALogicalBlockDependent
	NsALUALogicalBlockDependent = &pNsALUALogicalBlockDependent

	pNsALUAOffline = cNsALUAOffline
	NsALUAOffline = &pNsALUAOffline

	pNsALUAActiveOptimized = cNsALUAActiveOptimized
	NsALUAActiveOptimized = &pNsALUAActiveOptimized

	pNsALUAStandby = cNsALUAStandby
	NsALUAStandby = &pNsALUAStandby

	pNsALUATransitioning = cNsALUATransitioning
	NsALUATransitioning = &pNsALUATransitioning

	pNsALUAUnavailable = cNsALUAUnavailable
	NsALUAUnavailable = &pNsALUAUnavailable

	pNsALUAInvalid = cNsALUAInvalid
	NsALUAInvalid = &pNsALUAInvalid

	pNsALUAActiveNonOptimized = cNsALUAActiveNonOptimized
	NsALUAActiveNonOptimized = &pNsALUAActiveNonOptimized

}
