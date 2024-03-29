// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsOfflineReason Enum.

type NsOfflineReason string

const (
	cNsOfflineReasonReplica                  NsOfflineReason = "replica"
	cNsOfflineReasonRecovery                 NsOfflineReason = "recovery"
	cNsOfflineReasonNvramLossRecovery        NsOfflineReason = "nvram_loss_recovery"
	cNsOfflineReasonSerialNumberCollision    NsOfflineReason = "serial_number_collision"
	cNsOfflineReasonEncryptionKeyDeleted     NsOfflineReason = "encryption_key_deleted"
	cNsOfflineReasonOverVolumeUsageLimit     NsOfflineReason = "over_volume_usage_limit"
	cNsOfflineReasonOverFolderOverdraftLimit NsOfflineReason = "over_folder_overdraft_limit"
	cNsOfflineReasonCacheUnpinInProgress     NsOfflineReason = "cache_unpin_in_progress"
	cNsOfflineReasonVvolUnbind               NsOfflineReason = "vvol_unbind"
	cNsOfflineReasonOverVolumeLimit          NsOfflineReason = "over_volume_limit"
	cNsOfflineReasonOverSnapshotLimit        NsOfflineReason = "over_snapshot_limit"
	cNsOfflineReasonEncryptionInactive       NsOfflineReason = "encryption_inactive"
	cNsOfflineReasonPoolFreeSpaceExhausted   NsOfflineReason = "pool_free_space_exhausted"
	cNsOfflineReasonSrepUnconfigured         NsOfflineReason = "srep_unconfigured"
	cNsOfflineReasonUser                     NsOfflineReason = "user"
	cNsOfflineReasonOverVolumeReserve        NsOfflineReason = "over_volume_reserve"
	cNsOfflineReasonOverSnapshotReserve      NsOfflineReason = "over_snapshot_reserve"
)

var pNsOfflineReasonReplica NsOfflineReason
var pNsOfflineReasonRecovery NsOfflineReason
var pNsOfflineReasonNvramLossRecovery NsOfflineReason
var pNsOfflineReasonSerialNumberCollision NsOfflineReason
var pNsOfflineReasonEncryptionKeyDeleted NsOfflineReason
var pNsOfflineReasonOverVolumeUsageLimit NsOfflineReason
var pNsOfflineReasonOverFolderOverdraftLimit NsOfflineReason
var pNsOfflineReasonCacheUnpinInProgress NsOfflineReason
var pNsOfflineReasonVvolUnbind NsOfflineReason
var pNsOfflineReasonOverVolumeLimit NsOfflineReason
var pNsOfflineReasonOverSnapshotLimit NsOfflineReason
var pNsOfflineReasonEncryptionInactive NsOfflineReason
var pNsOfflineReasonPoolFreeSpaceExhausted NsOfflineReason
var pNsOfflineReasonSrepUnconfigured NsOfflineReason
var pNsOfflineReasonUser NsOfflineReason
var pNsOfflineReasonOverVolumeReserve NsOfflineReason
var pNsOfflineReasonOverSnapshotReserve NsOfflineReason

// NsOfflineReasonReplica enum exports
var NsOfflineReasonReplica *NsOfflineReason

// NsOfflineReasonRecovery enum exports
var NsOfflineReasonRecovery *NsOfflineReason

// NsOfflineReasonNvramLossRecovery enum exports
var NsOfflineReasonNvramLossRecovery *NsOfflineReason

// NsOfflineReasonSerialNumberCollision enum exports
var NsOfflineReasonSerialNumberCollision *NsOfflineReason

// NsOfflineReasonEncryptionKeyDeleted enum exports
var NsOfflineReasonEncryptionKeyDeleted *NsOfflineReason

// NsOfflineReasonOverVolumeUsageLimit enum exports
var NsOfflineReasonOverVolumeUsageLimit *NsOfflineReason

// NsOfflineReasonOverFolderOverdraftLimit enum exports
var NsOfflineReasonOverFolderOverdraftLimit *NsOfflineReason

// NsOfflineReasonCacheUnpinInProgress enum exports
var NsOfflineReasonCacheUnpinInProgress *NsOfflineReason

// NsOfflineReasonVvolUnbind enum exports
var NsOfflineReasonVvolUnbind *NsOfflineReason

// NsOfflineReasonOverVolumeLimit enum exports
var NsOfflineReasonOverVolumeLimit *NsOfflineReason

// NsOfflineReasonOverSnapshotLimit enum exports
var NsOfflineReasonOverSnapshotLimit *NsOfflineReason

// NsOfflineReasonEncryptionInactive enum exports
var NsOfflineReasonEncryptionInactive *NsOfflineReason

// NsOfflineReasonPoolFreeSpaceExhausted enum exports
var NsOfflineReasonPoolFreeSpaceExhausted *NsOfflineReason

// NsOfflineReasonSrepUnconfigured enum exports
var NsOfflineReasonSrepUnconfigured *NsOfflineReason

// NsOfflineReasonUser enum exports
var NsOfflineReasonUser *NsOfflineReason

// NsOfflineReasonOverVolumeReserve enum exports
var NsOfflineReasonOverVolumeReserve *NsOfflineReason

// NsOfflineReasonOverSnapshotReserve enum exports
var NsOfflineReasonOverSnapshotReserve *NsOfflineReason

func init() {
	pNsOfflineReasonReplica = cNsOfflineReasonReplica
	NsOfflineReasonReplica = &pNsOfflineReasonReplica

	pNsOfflineReasonRecovery = cNsOfflineReasonRecovery
	NsOfflineReasonRecovery = &pNsOfflineReasonRecovery

	pNsOfflineReasonNvramLossRecovery = cNsOfflineReasonNvramLossRecovery
	NsOfflineReasonNvramLossRecovery = &pNsOfflineReasonNvramLossRecovery

	pNsOfflineReasonSerialNumberCollision = cNsOfflineReasonSerialNumberCollision
	NsOfflineReasonSerialNumberCollision = &pNsOfflineReasonSerialNumberCollision

	pNsOfflineReasonEncryptionKeyDeleted = cNsOfflineReasonEncryptionKeyDeleted
	NsOfflineReasonEncryptionKeyDeleted = &pNsOfflineReasonEncryptionKeyDeleted

	pNsOfflineReasonOverVolumeUsageLimit = cNsOfflineReasonOverVolumeUsageLimit
	NsOfflineReasonOverVolumeUsageLimit = &pNsOfflineReasonOverVolumeUsageLimit

	pNsOfflineReasonOverFolderOverdraftLimit = cNsOfflineReasonOverFolderOverdraftLimit
	NsOfflineReasonOverFolderOverdraftLimit = &pNsOfflineReasonOverFolderOverdraftLimit

	pNsOfflineReasonCacheUnpinInProgress = cNsOfflineReasonCacheUnpinInProgress
	NsOfflineReasonCacheUnpinInProgress = &pNsOfflineReasonCacheUnpinInProgress

	pNsOfflineReasonVvolUnbind = cNsOfflineReasonVvolUnbind
	NsOfflineReasonVvolUnbind = &pNsOfflineReasonVvolUnbind

	pNsOfflineReasonOverVolumeLimit = cNsOfflineReasonOverVolumeLimit
	NsOfflineReasonOverVolumeLimit = &pNsOfflineReasonOverVolumeLimit

	pNsOfflineReasonOverSnapshotLimit = cNsOfflineReasonOverSnapshotLimit
	NsOfflineReasonOverSnapshotLimit = &pNsOfflineReasonOverSnapshotLimit

	pNsOfflineReasonEncryptionInactive = cNsOfflineReasonEncryptionInactive
	NsOfflineReasonEncryptionInactive = &pNsOfflineReasonEncryptionInactive

	pNsOfflineReasonPoolFreeSpaceExhausted = cNsOfflineReasonPoolFreeSpaceExhausted
	NsOfflineReasonPoolFreeSpaceExhausted = &pNsOfflineReasonPoolFreeSpaceExhausted

	pNsOfflineReasonSrepUnconfigured = cNsOfflineReasonSrepUnconfigured
	NsOfflineReasonSrepUnconfigured = &pNsOfflineReasonSrepUnconfigured

	pNsOfflineReasonUser = cNsOfflineReasonUser
	NsOfflineReasonUser = &pNsOfflineReasonUser

	pNsOfflineReasonOverVolumeReserve = cNsOfflineReasonOverVolumeReserve
	NsOfflineReasonOverVolumeReserve = &pNsOfflineReasonOverVolumeReserve

	pNsOfflineReasonOverSnapshotReserve = cNsOfflineReasonOverSnapshotReserve
	NsOfflineReasonOverSnapshotReserve = &pNsOfflineReasonOverSnapshotReserve

}
