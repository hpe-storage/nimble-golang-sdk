// Copyright 2020 Hewlett Packard Enterprise Development LP
package model


/**
 * <p>Golang package for NsOfflineReason Enum.</p>
 */

type NsOfflineReason string 

const (
	NSOFFLINEREASON_REPLICA NsOfflineReason = "replica"
	NSOFFLINEREASON_RECOVERY NsOfflineReason = "recovery"
	NSOFFLINEREASON_NVRAM_LOSS_RECOVERY NsOfflineReason = "nvram_loss_recovery"
	NSOFFLINEREASON_SERIAL_NUMBER_COLLISION NsOfflineReason = "serial_number_collision"
	NSOFFLINEREASON_ENCRYPTION_KEY_DELETED NsOfflineReason = "encryption_key_deleted"
	NSOFFLINEREASON_OVER_VOLUME_USAGE_LIMIT NsOfflineReason = "over_volume_usage_limit"
	NSOFFLINEREASON_OVER_FOLDER_OVERDRAFT_LIMIT NsOfflineReason = "over_folder_overdraft_limit"
	NSOFFLINEREASON_CACHE_UNPIN_IN_PROGRESS NsOfflineReason = "cache_unpin_in_progress"
	NSOFFLINEREASON_VVOL_UNBIND NsOfflineReason = "vvol_unbind"
	NSOFFLINEREASON_OVER_VOLUME_LIMIT NsOfflineReason = "over_volume_limit"
	NSOFFLINEREASON_OVER_SNAPSHOT_LIMIT NsOfflineReason = "over_snapshot_limit"
	NSOFFLINEREASON_ENCRYPTION_INACTIVE NsOfflineReason = "encryption_inactive"
	NSOFFLINEREASON_POOL_FREE_SPACE_EXHAUSTED NsOfflineReason = "pool_free_space_exhausted"
	NSOFFLINEREASON_SREP_UNCONFIGURED NsOfflineReason = "srep_unconfigured"
	NSOFFLINEREASON_USER NsOfflineReason = "user"
	NSOFFLINEREASON_OVER_VOLUME_RESERVE NsOfflineReason = "over_volume_reserve"
	NSOFFLINEREASON_OVER_SNAPSHOT_RESERVE NsOfflineReason = "over_snapshot_reserve"

) 
