/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsReplVolPartnerStatus.</p>
 */

type NsReplVolPartnerStatus string 

const (
	NSREPLVOLPARTNERSTATUS_DATA_TRANSFER_TERMINATED NsReplVolPartnerStatus = "Data transfer terminated"
	NSREPLVOLPARTNERSTATUS_INITIATING_CONTENT_BASED_REPLICATION NsReplVolPartnerStatus = "Initiating content based replication"
	NSREPLVOLPARTNERSTATUS_CONTENT_BASED_REPLICATION_COMPLETE NsReplVolPartnerStatus = "Content based replication complete"
	NSREPLVOLPARTNERSTATUS_DATA_TRANSFER_COMPLETE NsReplVolPartnerStatus = "Data transfer complete"
	NSREPLVOLPARTNERSTATUS_CREATING_SNAPSHOT_ON_PARTNER NsReplVolPartnerStatus = "Creating snapshot on partner"
	NSREPLVOLPARTNERSTATUS_SYNCHRONIZING_PARTNER_CONFIGURATION NsReplVolPartnerStatus = "Synchronizing partner configuration"
	NSREPLVOLPARTNERSTATUS_DATA_TRANSFER_IN_PROGRESS NsReplVolPartnerStatus = "Data transfer in progress"
	NSREPLVOLPARTNERSTATUS_IDLE NsReplVolPartnerStatus = "Idle"
	NSREPLVOLPARTNERSTATUS_IDENTIFYING_COMMON_ANCESTRY_ON_PARTNER NsReplVolPartnerStatus = "Identifying common ancestry on partner"
	NSREPLVOLPARTNERSTATUS_CONTENT_BASED_REPLICATION_IN_PROGRESS NsReplVolPartnerStatus = "Content based replication in progress"
	NSREPLVOLPARTNERSTATUS_TERMINATING_DATA_TRANSFER NsReplVolPartnerStatus = "Terminating data transfer"
	NSREPLVOLPARTNERSTATUS_PAUSED NsReplVolPartnerStatus = "Paused"
	NSREPLVOLPARTNERSTATUS_INITIATING_DATA_TRANSFER NsReplVolPartnerStatus = "Initiating data transfer"

) 
