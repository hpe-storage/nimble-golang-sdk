// Copyright 2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsReplicationType Enum.

type NsReplicationType string

const (
	cNsReplicationTypeSynchronous      NsReplicationType = "synchronous"
	cNsReplicationTypePeriodicSnapshot NsReplicationType = "periodic_snapshot"
)

var pNsReplicationTypeSynchronous NsReplicationType
var pNsReplicationTypePeriodicSnapshot NsReplicationType

// NsReplicationTypeSynchronous enum exports
var NsReplicationTypeSynchronous *NsReplicationType

// NsReplicationTypePeriodicSnapshot enum exports
var NsReplicationTypePeriodicSnapshot *NsReplicationType

func init() {
	pNsReplicationTypeSynchronous = cNsReplicationTypeSynchronous
	NsReplicationTypeSynchronous = &pNsReplicationTypeSynchronous

	pNsReplicationTypePeriodicSnapshot = cNsReplicationTypePeriodicSnapshot
	NsReplicationTypePeriodicSnapshot = &pNsReplicationTypePeriodicSnapshot

}
