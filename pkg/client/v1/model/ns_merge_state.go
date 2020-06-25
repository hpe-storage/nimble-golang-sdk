// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsMergeState Enum.
 
type NsMergeState string 

const (
	NSMERGESTATE_DEST_DB_DONE NsMergeState = "dest_db_done"
	NSMERGESTATE_DEST_DB NsMergeState = "dest_db"
	NSMERGESTATE_DEST_REASSIGNED NsMergeState = "dest_reassigned"
	NSMERGESTATE_SRC_QUIESCED NsMergeState = "src_quiesced"
	NSMERGESTATE_NONE NsMergeState = "none"
	NSMERGESTATE_DEST_RELINQUISH NsMergeState = "dest_relinquish"
	NSMERGESTATE_SRC_REASSIGNED NsMergeState = "src_reassigned"
	NSMERGESTATE_DEST_START NsMergeState = "dest_start"
	NSMERGESTATE_SRC_START NsMergeState = "src_start"
	NSMERGESTATE_SRC_QUIESCE_FAILED NsMergeState = "src_quiesce_failed"
	NSMERGESTATE_DEST_MERGE_VALIDATION NsMergeState = "dest_merge_validation"
	NSMERGESTATE_SRC_REASSIGN_FAILED NsMergeState = "src_reassign_failed"
	NSMERGESTATE_DEST_DB_FAILED NsMergeState = "dest_db_failed"

) 
