// Copyright 2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsReplVolPartnerStatus Enum.

type NsReplVolPartnerStatus string

const (
	cNsReplVolPartnerStatusDataTransferTerminated             NsReplVolPartnerStatus = "Data transfer terminated"
	cNsReplVolPartnerStatusInitiatingContentBasedReplication  NsReplVolPartnerStatus = "Initiating content based replication"
	cNsReplVolPartnerStatusContentBasedReplicationComplete    NsReplVolPartnerStatus = "Content based replication complete"
	cNsReplVolPartnerStatusDataTransferComplete               NsReplVolPartnerStatus = "Data transfer complete"
	cNsReplVolPartnerStatusCreatingSnapshotOnPartner          NsReplVolPartnerStatus = "Creating snapshot on partner"
	cNsReplVolPartnerStatusSynchronizingPartnerConfiguration  NsReplVolPartnerStatus = "Synchronizing partner configuration"
	cNsReplVolPartnerStatusDataTransferInProgress             NsReplVolPartnerStatus = "Data transfer in progress"
	cNsReplVolPartnerStatusIdle                               NsReplVolPartnerStatus = "Idle"
	cNsReplVolPartnerStatusIdentifyingCommonAncestryOnPartner NsReplVolPartnerStatus = "Identifying common ancestry on partner"
	cNsReplVolPartnerStatusContentBasedReplicationInProgress  NsReplVolPartnerStatus = "Content based replication in progress"
	cNsReplVolPartnerStatusTerminatingDataTransfer            NsReplVolPartnerStatus = "Terminating data transfer"
	cNsReplVolPartnerStatusPaused                             NsReplVolPartnerStatus = "Paused"
	cNsReplVolPartnerStatusInitiatingDataTransfer             NsReplVolPartnerStatus = "Initiating data transfer"
)

var pNsReplVolPartnerStatusDataTransferTerminated NsReplVolPartnerStatus
var pNsReplVolPartnerStatusInitiatingContentBasedReplication NsReplVolPartnerStatus
var pNsReplVolPartnerStatusContentBasedReplicationComplete NsReplVolPartnerStatus
var pNsReplVolPartnerStatusDataTransferComplete NsReplVolPartnerStatus
var pNsReplVolPartnerStatusCreatingSnapshotOnPartner NsReplVolPartnerStatus
var pNsReplVolPartnerStatusSynchronizingPartnerConfiguration NsReplVolPartnerStatus
var pNsReplVolPartnerStatusDataTransferInProgress NsReplVolPartnerStatus
var pNsReplVolPartnerStatusIdle NsReplVolPartnerStatus
var pNsReplVolPartnerStatusIdentifyingCommonAncestryOnPartner NsReplVolPartnerStatus
var pNsReplVolPartnerStatusContentBasedReplicationInProgress NsReplVolPartnerStatus
var pNsReplVolPartnerStatusTerminatingDataTransfer NsReplVolPartnerStatus
var pNsReplVolPartnerStatusPaused NsReplVolPartnerStatus
var pNsReplVolPartnerStatusInitiatingDataTransfer NsReplVolPartnerStatus

// NsReplVolPartnerStatusDataTransferTerminated enum exports
var NsReplVolPartnerStatusDataTransferTerminated *NsReplVolPartnerStatus

// NsReplVolPartnerStatusInitiatingContentBasedReplication enum exports
var NsReplVolPartnerStatusInitiatingContentBasedReplication *NsReplVolPartnerStatus

// NsReplVolPartnerStatusContentBasedReplicationComplete enum exports
var NsReplVolPartnerStatusContentBasedReplicationComplete *NsReplVolPartnerStatus

// NsReplVolPartnerStatusDataTransferComplete enum exports
var NsReplVolPartnerStatusDataTransferComplete *NsReplVolPartnerStatus

// NsReplVolPartnerStatusCreatingSnapshotOnPartner enum exports
var NsReplVolPartnerStatusCreatingSnapshotOnPartner *NsReplVolPartnerStatus

// NsReplVolPartnerStatusSynchronizingPartnerConfiguration enum exports
var NsReplVolPartnerStatusSynchronizingPartnerConfiguration *NsReplVolPartnerStatus

// NsReplVolPartnerStatusDataTransferInProgress enum exports
var NsReplVolPartnerStatusDataTransferInProgress *NsReplVolPartnerStatus

// NsReplVolPartnerStatusIdle enum exports
var NsReplVolPartnerStatusIdle *NsReplVolPartnerStatus

// NsReplVolPartnerStatusIdentifyingCommonAncestryOnPartner enum exports
var NsReplVolPartnerStatusIdentifyingCommonAncestryOnPartner *NsReplVolPartnerStatus

// NsReplVolPartnerStatusContentBasedReplicationInProgress enum exports
var NsReplVolPartnerStatusContentBasedReplicationInProgress *NsReplVolPartnerStatus

// NsReplVolPartnerStatusTerminatingDataTransfer enum exports
var NsReplVolPartnerStatusTerminatingDataTransfer *NsReplVolPartnerStatus

// NsReplVolPartnerStatusPaused enum exports
var NsReplVolPartnerStatusPaused *NsReplVolPartnerStatus

// NsReplVolPartnerStatusInitiatingDataTransfer enum exports
var NsReplVolPartnerStatusInitiatingDataTransfer *NsReplVolPartnerStatus

func init() {
	pNsReplVolPartnerStatusDataTransferTerminated = cNsReplVolPartnerStatusDataTransferTerminated
	NsReplVolPartnerStatusDataTransferTerminated = &pNsReplVolPartnerStatusDataTransferTerminated

	pNsReplVolPartnerStatusInitiatingContentBasedReplication = cNsReplVolPartnerStatusInitiatingContentBasedReplication
	NsReplVolPartnerStatusInitiatingContentBasedReplication = &pNsReplVolPartnerStatusInitiatingContentBasedReplication

	pNsReplVolPartnerStatusContentBasedReplicationComplete = cNsReplVolPartnerStatusContentBasedReplicationComplete
	NsReplVolPartnerStatusContentBasedReplicationComplete = &pNsReplVolPartnerStatusContentBasedReplicationComplete

	pNsReplVolPartnerStatusDataTransferComplete = cNsReplVolPartnerStatusDataTransferComplete
	NsReplVolPartnerStatusDataTransferComplete = &pNsReplVolPartnerStatusDataTransferComplete

	pNsReplVolPartnerStatusCreatingSnapshotOnPartner = cNsReplVolPartnerStatusCreatingSnapshotOnPartner
	NsReplVolPartnerStatusCreatingSnapshotOnPartner = &pNsReplVolPartnerStatusCreatingSnapshotOnPartner

	pNsReplVolPartnerStatusSynchronizingPartnerConfiguration = cNsReplVolPartnerStatusSynchronizingPartnerConfiguration
	NsReplVolPartnerStatusSynchronizingPartnerConfiguration = &pNsReplVolPartnerStatusSynchronizingPartnerConfiguration

	pNsReplVolPartnerStatusDataTransferInProgress = cNsReplVolPartnerStatusDataTransferInProgress
	NsReplVolPartnerStatusDataTransferInProgress = &pNsReplVolPartnerStatusDataTransferInProgress

	pNsReplVolPartnerStatusIdle = cNsReplVolPartnerStatusIdle
	NsReplVolPartnerStatusIdle = &pNsReplVolPartnerStatusIdle

	pNsReplVolPartnerStatusIdentifyingCommonAncestryOnPartner = cNsReplVolPartnerStatusIdentifyingCommonAncestryOnPartner
	NsReplVolPartnerStatusIdentifyingCommonAncestryOnPartner = &pNsReplVolPartnerStatusIdentifyingCommonAncestryOnPartner

	pNsReplVolPartnerStatusContentBasedReplicationInProgress = cNsReplVolPartnerStatusContentBasedReplicationInProgress
	NsReplVolPartnerStatusContentBasedReplicationInProgress = &pNsReplVolPartnerStatusContentBasedReplicationInProgress

	pNsReplVolPartnerStatusTerminatingDataTransfer = cNsReplVolPartnerStatusTerminatingDataTransfer
	NsReplVolPartnerStatusTerminatingDataTransfer = &pNsReplVolPartnerStatusTerminatingDataTransfer

	pNsReplVolPartnerStatusPaused = cNsReplVolPartnerStatusPaused
	NsReplVolPartnerStatusPaused = &pNsReplVolPartnerStatusPaused

	pNsReplVolPartnerStatusInitiatingDataTransfer = cNsReplVolPartnerStatusInitiatingDataTransfer
	NsReplVolPartnerStatusInitiatingDataTransfer = &pNsReplVolPartnerStatusInitiatingDataTransfer

}
