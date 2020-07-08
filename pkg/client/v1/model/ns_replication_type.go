// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsReplicationType Enum.

type NsReplicationType string

const (
	cNsReplicationTypeSynchronous      NsReplicationType = "synchronous"
	cNsReplicationTypePeriodicSnapshot NsReplicationType = "periodic_snapshot"
)

var pNsReplicationTypeSynchronous NsReplicationType
var pNsReplicationTypePeriodicSnapshot NsReplicationType

// Export
var NsReplicationTypeSynchronous *NsReplicationType
var NsReplicationTypePeriodicSnapshot *NsReplicationType

func init() {
	pNsReplicationTypeSynchronous = cNsReplicationTypeSynchronous
	NsReplicationTypeSynchronous = &pNsReplicationTypeSynchronous

	pNsReplicationTypePeriodicSnapshot = cNsReplicationTypePeriodicSnapshot
	NsReplicationTypePeriodicSnapshot = &pNsReplicationTypePeriodicSnapshot

}
