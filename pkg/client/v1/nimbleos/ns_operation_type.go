// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsOperationType Enum.

type NsOperationType string

const (
	cNsOperationTypeCreateInitiatorGrp            NsOperationType = "create_initiator_grp"
	cNsOperationTypeVolumeCollectionsHandover     NsOperationType = "volume_collections_handover"
	cNsOperationTypeFoldersDelete                 NsOperationType = "folders_delete"
	cNsOperationTypeSnapshotsDelete               NsOperationType = "snapshots_delete"
	cNsOperationTypeAddVolAcl                     NsOperationType = "add_vol_acl"
	cNsOperationTypeAlarmsUpdate                  NsOperationType = "alarms_update"
	cNsOperationTypeAlarmsAcknowledge             NsOperationType = "alarms_acknowledge"
	cNsOperationTypeEditInitiatorGrp              NsOperationType = "edit_initiator_grp"
	cNsOperationTypeVolumeCollectionsDemote       NsOperationType = "volume_collections_demote"
	cNsOperationTypeVolumesBulkDelete             NsOperationType = "volumes_bulk_delete"
	cNsOperationTypeEditFolder                    NsOperationType = "edit_folder"
	cNsOperationTypeRemoveSnapshot                NsOperationType = "remove_snapshot"
	cNsOperationTypeVolumeCollectionsBulkHandover NsOperationType = "volume_collections_bulk_handover"
	cNsOperationTypeVolumesDelete                 NsOperationType = "volumes_delete"
	cNsOperationTypeAddFcVolAcl                   NsOperationType = "add_fc_vol_acl"
	cNsOperationTypeSnapshotCollectionsDelete     NsOperationType = "snapshot_collections_delete"
	cNsOperationTypeAddInitiator                  NsOperationType = "add_initiator"
	cNsOperationTypeEditSnap                      NsOperationType = "edit_snap"
	cNsOperationTypeAlarmsUnacknowledge           NsOperationType = "alarms_unacknowledge"
	cNsOperationTypeAlarmsBulkUnacknowledge       NsOperationType = "alarms_bulk_unacknowledge"
	cNsOperationTypeSnapVolcoll                   NsOperationType = "snap_volcoll"
	cNsOperationTypeEditVol                       NsOperationType = "edit_vol"
	cNsOperationTypeVolumeCollectionsBulkDemote   NsOperationType = "volume_collections_bulk_demote"
	cNsOperationTypeAlarmsBulkAcknowledge         NsOperationType = "alarms_bulk_acknowledge"
	cNsOperationTypeSnapshotsBulkDelete           NsOperationType = "snapshots_bulk_delete"
	cNsOperationTypeFoldersBulkDelete             NsOperationType = "folders_bulk_delete"
	cNsOperationTypeVolumeCollectionsBulkPromote  NsOperationType = "volume_collections_bulk_promote"
	cNsOperationTypeVolumesBulkUpdate             NsOperationType = "volumes_bulk_update"
	cNsOperationTypeSnapVol                       NsOperationType = "snap_vol"
	cNsOperationTypeRemoveVolAcl                  NsOperationType = "remove_vol_acl"
	cNsOperationTypeVolumesUpdate                 NsOperationType = "volumes_update"
	cNsOperationTypeAlarmsBulkUpdate              NsOperationType = "alarms_bulk_update"
	cNsOperationTypeDeleteInitiatorGrp            NsOperationType = "delete_initiator_grp"
	cNsOperationTypeSnapshotsBulkCreate           NsOperationType = "snapshots_bulk_create"
	cNsOperationTypeSnapshotCollectionsBulkDelete NsOperationType = "snapshot_collections_bulk_delete"
	cNsOperationTypeSnapshotsCreate               NsOperationType = "snapshots_create"
	cNsOperationTypeVolumeCollectionsPromote      NsOperationType = "volume_collections_promote"
	cNsOperationTypeRemoveInitiator               NsOperationType = "remove_initiator"
)

var pNsOperationTypeCreateInitiatorGrp NsOperationType
var pNsOperationTypeVolumeCollectionsHandover NsOperationType
var pNsOperationTypeFoldersDelete NsOperationType
var pNsOperationTypeSnapshotsDelete NsOperationType
var pNsOperationTypeAddVolAcl NsOperationType
var pNsOperationTypeAlarmsUpdate NsOperationType
var pNsOperationTypeAlarmsAcknowledge NsOperationType
var pNsOperationTypeEditInitiatorGrp NsOperationType
var pNsOperationTypeVolumeCollectionsDemote NsOperationType
var pNsOperationTypeVolumesBulkDelete NsOperationType
var pNsOperationTypeEditFolder NsOperationType
var pNsOperationTypeRemoveSnapshot NsOperationType
var pNsOperationTypeVolumeCollectionsBulkHandover NsOperationType
var pNsOperationTypeVolumesDelete NsOperationType
var pNsOperationTypeAddFcVolAcl NsOperationType
var pNsOperationTypeSnapshotCollectionsDelete NsOperationType
var pNsOperationTypeAddInitiator NsOperationType
var pNsOperationTypeEditSnap NsOperationType
var pNsOperationTypeAlarmsUnacknowledge NsOperationType
var pNsOperationTypeAlarmsBulkUnacknowledge NsOperationType
var pNsOperationTypeSnapVolcoll NsOperationType
var pNsOperationTypeEditVol NsOperationType
var pNsOperationTypeVolumeCollectionsBulkDemote NsOperationType
var pNsOperationTypeAlarmsBulkAcknowledge NsOperationType
var pNsOperationTypeSnapshotsBulkDelete NsOperationType
var pNsOperationTypeFoldersBulkDelete NsOperationType
var pNsOperationTypeVolumeCollectionsBulkPromote NsOperationType
var pNsOperationTypeVolumesBulkUpdate NsOperationType
var pNsOperationTypeSnapVol NsOperationType
var pNsOperationTypeRemoveVolAcl NsOperationType
var pNsOperationTypeVolumesUpdate NsOperationType
var pNsOperationTypeAlarmsBulkUpdate NsOperationType
var pNsOperationTypeDeleteInitiatorGrp NsOperationType
var pNsOperationTypeSnapshotsBulkCreate NsOperationType
var pNsOperationTypeSnapshotCollectionsBulkDelete NsOperationType
var pNsOperationTypeSnapshotsCreate NsOperationType
var pNsOperationTypeVolumeCollectionsPromote NsOperationType
var pNsOperationTypeRemoveInitiator NsOperationType

// NsOperationTypeCreateInitiatorGrp enum exports
var NsOperationTypeCreateInitiatorGrp *NsOperationType

// NsOperationTypeVolumeCollectionsHandover enum exports
var NsOperationTypeVolumeCollectionsHandover *NsOperationType

// NsOperationTypeFoldersDelete enum exports
var NsOperationTypeFoldersDelete *NsOperationType

// NsOperationTypeSnapshotsDelete enum exports
var NsOperationTypeSnapshotsDelete *NsOperationType

// NsOperationTypeAddVolAcl enum exports
var NsOperationTypeAddVolAcl *NsOperationType

// NsOperationTypeAlarmsUpdate enum exports
var NsOperationTypeAlarmsUpdate *NsOperationType

// NsOperationTypeAlarmsAcknowledge enum exports
var NsOperationTypeAlarmsAcknowledge *NsOperationType

// NsOperationTypeEditInitiatorGrp enum exports
var NsOperationTypeEditInitiatorGrp *NsOperationType

// NsOperationTypeVolumeCollectionsDemote enum exports
var NsOperationTypeVolumeCollectionsDemote *NsOperationType

// NsOperationTypeVolumesBulkDelete enum exports
var NsOperationTypeVolumesBulkDelete *NsOperationType

// NsOperationTypeEditFolder enum exports
var NsOperationTypeEditFolder *NsOperationType

// NsOperationTypeRemoveSnapshot enum exports
var NsOperationTypeRemoveSnapshot *NsOperationType

// NsOperationTypeVolumeCollectionsBulkHandover enum exports
var NsOperationTypeVolumeCollectionsBulkHandover *NsOperationType

// NsOperationTypeVolumesDelete enum exports
var NsOperationTypeVolumesDelete *NsOperationType

// NsOperationTypeAddFcVolAcl enum exports
var NsOperationTypeAddFcVolAcl *NsOperationType

// NsOperationTypeSnapshotCollectionsDelete enum exports
var NsOperationTypeSnapshotCollectionsDelete *NsOperationType

// NsOperationTypeAddInitiator enum exports
var NsOperationTypeAddInitiator *NsOperationType

// NsOperationTypeEditSnap enum exports
var NsOperationTypeEditSnap *NsOperationType

// NsOperationTypeAlarmsUnacknowledge enum exports
var NsOperationTypeAlarmsUnacknowledge *NsOperationType

// NsOperationTypeAlarmsBulkUnacknowledge enum exports
var NsOperationTypeAlarmsBulkUnacknowledge *NsOperationType

// NsOperationTypeSnapVolcoll enum exports
var NsOperationTypeSnapVolcoll *NsOperationType

// NsOperationTypeEditVol enum exports
var NsOperationTypeEditVol *NsOperationType

// NsOperationTypeVolumeCollectionsBulkDemote enum exports
var NsOperationTypeVolumeCollectionsBulkDemote *NsOperationType

// NsOperationTypeAlarmsBulkAcknowledge enum exports
var NsOperationTypeAlarmsBulkAcknowledge *NsOperationType

// NsOperationTypeSnapshotsBulkDelete enum exports
var NsOperationTypeSnapshotsBulkDelete *NsOperationType

// NsOperationTypeFoldersBulkDelete enum exports
var NsOperationTypeFoldersBulkDelete *NsOperationType

// NsOperationTypeVolumeCollectionsBulkPromote enum exports
var NsOperationTypeVolumeCollectionsBulkPromote *NsOperationType

// NsOperationTypeVolumesBulkUpdate enum exports
var NsOperationTypeVolumesBulkUpdate *NsOperationType

// NsOperationTypeSnapVol enum exports
var NsOperationTypeSnapVol *NsOperationType

// NsOperationTypeRemoveVolAcl enum exports
var NsOperationTypeRemoveVolAcl *NsOperationType

// NsOperationTypeVolumesUpdate enum exports
var NsOperationTypeVolumesUpdate *NsOperationType

// NsOperationTypeAlarmsBulkUpdate enum exports
var NsOperationTypeAlarmsBulkUpdate *NsOperationType

// NsOperationTypeDeleteInitiatorGrp enum exports
var NsOperationTypeDeleteInitiatorGrp *NsOperationType

// NsOperationTypeSnapshotsBulkCreate enum exports
var NsOperationTypeSnapshotsBulkCreate *NsOperationType

// NsOperationTypeSnapshotCollectionsBulkDelete enum exports
var NsOperationTypeSnapshotCollectionsBulkDelete *NsOperationType

// NsOperationTypeSnapshotsCreate enum exports
var NsOperationTypeSnapshotsCreate *NsOperationType

// NsOperationTypeVolumeCollectionsPromote enum exports
var NsOperationTypeVolumeCollectionsPromote *NsOperationType

// NsOperationTypeRemoveInitiator enum exports
var NsOperationTypeRemoveInitiator *NsOperationType

func init() {
	pNsOperationTypeCreateInitiatorGrp = cNsOperationTypeCreateInitiatorGrp
	NsOperationTypeCreateInitiatorGrp = &pNsOperationTypeCreateInitiatorGrp

	pNsOperationTypeVolumeCollectionsHandover = cNsOperationTypeVolumeCollectionsHandover
	NsOperationTypeVolumeCollectionsHandover = &pNsOperationTypeVolumeCollectionsHandover

	pNsOperationTypeFoldersDelete = cNsOperationTypeFoldersDelete
	NsOperationTypeFoldersDelete = &pNsOperationTypeFoldersDelete

	pNsOperationTypeSnapshotsDelete = cNsOperationTypeSnapshotsDelete
	NsOperationTypeSnapshotsDelete = &pNsOperationTypeSnapshotsDelete

	pNsOperationTypeAddVolAcl = cNsOperationTypeAddVolAcl
	NsOperationTypeAddVolAcl = &pNsOperationTypeAddVolAcl

	pNsOperationTypeAlarmsUpdate = cNsOperationTypeAlarmsUpdate
	NsOperationTypeAlarmsUpdate = &pNsOperationTypeAlarmsUpdate

	pNsOperationTypeAlarmsAcknowledge = cNsOperationTypeAlarmsAcknowledge
	NsOperationTypeAlarmsAcknowledge = &pNsOperationTypeAlarmsAcknowledge

	pNsOperationTypeEditInitiatorGrp = cNsOperationTypeEditInitiatorGrp
	NsOperationTypeEditInitiatorGrp = &pNsOperationTypeEditInitiatorGrp

	pNsOperationTypeVolumeCollectionsDemote = cNsOperationTypeVolumeCollectionsDemote
	NsOperationTypeVolumeCollectionsDemote = &pNsOperationTypeVolumeCollectionsDemote

	pNsOperationTypeVolumesBulkDelete = cNsOperationTypeVolumesBulkDelete
	NsOperationTypeVolumesBulkDelete = &pNsOperationTypeVolumesBulkDelete

	pNsOperationTypeEditFolder = cNsOperationTypeEditFolder
	NsOperationTypeEditFolder = &pNsOperationTypeEditFolder

	pNsOperationTypeRemoveSnapshot = cNsOperationTypeRemoveSnapshot
	NsOperationTypeRemoveSnapshot = &pNsOperationTypeRemoveSnapshot

	pNsOperationTypeVolumeCollectionsBulkHandover = cNsOperationTypeVolumeCollectionsBulkHandover
	NsOperationTypeVolumeCollectionsBulkHandover = &pNsOperationTypeVolumeCollectionsBulkHandover

	pNsOperationTypeVolumesDelete = cNsOperationTypeVolumesDelete
	NsOperationTypeVolumesDelete = &pNsOperationTypeVolumesDelete

	pNsOperationTypeAddFcVolAcl = cNsOperationTypeAddFcVolAcl
	NsOperationTypeAddFcVolAcl = &pNsOperationTypeAddFcVolAcl

	pNsOperationTypeSnapshotCollectionsDelete = cNsOperationTypeSnapshotCollectionsDelete
	NsOperationTypeSnapshotCollectionsDelete = &pNsOperationTypeSnapshotCollectionsDelete

	pNsOperationTypeAddInitiator = cNsOperationTypeAddInitiator
	NsOperationTypeAddInitiator = &pNsOperationTypeAddInitiator

	pNsOperationTypeEditSnap = cNsOperationTypeEditSnap
	NsOperationTypeEditSnap = &pNsOperationTypeEditSnap

	pNsOperationTypeAlarmsUnacknowledge = cNsOperationTypeAlarmsUnacknowledge
	NsOperationTypeAlarmsUnacknowledge = &pNsOperationTypeAlarmsUnacknowledge

	pNsOperationTypeAlarmsBulkUnacknowledge = cNsOperationTypeAlarmsBulkUnacknowledge
	NsOperationTypeAlarmsBulkUnacknowledge = &pNsOperationTypeAlarmsBulkUnacknowledge

	pNsOperationTypeSnapVolcoll = cNsOperationTypeSnapVolcoll
	NsOperationTypeSnapVolcoll = &pNsOperationTypeSnapVolcoll

	pNsOperationTypeEditVol = cNsOperationTypeEditVol
	NsOperationTypeEditVol = &pNsOperationTypeEditVol

	pNsOperationTypeVolumeCollectionsBulkDemote = cNsOperationTypeVolumeCollectionsBulkDemote
	NsOperationTypeVolumeCollectionsBulkDemote = &pNsOperationTypeVolumeCollectionsBulkDemote

	pNsOperationTypeAlarmsBulkAcknowledge = cNsOperationTypeAlarmsBulkAcknowledge
	NsOperationTypeAlarmsBulkAcknowledge = &pNsOperationTypeAlarmsBulkAcknowledge

	pNsOperationTypeSnapshotsBulkDelete = cNsOperationTypeSnapshotsBulkDelete
	NsOperationTypeSnapshotsBulkDelete = &pNsOperationTypeSnapshotsBulkDelete

	pNsOperationTypeFoldersBulkDelete = cNsOperationTypeFoldersBulkDelete
	NsOperationTypeFoldersBulkDelete = &pNsOperationTypeFoldersBulkDelete

	pNsOperationTypeVolumeCollectionsBulkPromote = cNsOperationTypeVolumeCollectionsBulkPromote
	NsOperationTypeVolumeCollectionsBulkPromote = &pNsOperationTypeVolumeCollectionsBulkPromote

	pNsOperationTypeVolumesBulkUpdate = cNsOperationTypeVolumesBulkUpdate
	NsOperationTypeVolumesBulkUpdate = &pNsOperationTypeVolumesBulkUpdate

	pNsOperationTypeSnapVol = cNsOperationTypeSnapVol
	NsOperationTypeSnapVol = &pNsOperationTypeSnapVol

	pNsOperationTypeRemoveVolAcl = cNsOperationTypeRemoveVolAcl
	NsOperationTypeRemoveVolAcl = &pNsOperationTypeRemoveVolAcl

	pNsOperationTypeVolumesUpdate = cNsOperationTypeVolumesUpdate
	NsOperationTypeVolumesUpdate = &pNsOperationTypeVolumesUpdate

	pNsOperationTypeAlarmsBulkUpdate = cNsOperationTypeAlarmsBulkUpdate
	NsOperationTypeAlarmsBulkUpdate = &pNsOperationTypeAlarmsBulkUpdate

	pNsOperationTypeDeleteInitiatorGrp = cNsOperationTypeDeleteInitiatorGrp
	NsOperationTypeDeleteInitiatorGrp = &pNsOperationTypeDeleteInitiatorGrp

	pNsOperationTypeSnapshotsBulkCreate = cNsOperationTypeSnapshotsBulkCreate
	NsOperationTypeSnapshotsBulkCreate = &pNsOperationTypeSnapshotsBulkCreate

	pNsOperationTypeSnapshotCollectionsBulkDelete = cNsOperationTypeSnapshotCollectionsBulkDelete
	NsOperationTypeSnapshotCollectionsBulkDelete = &pNsOperationTypeSnapshotCollectionsBulkDelete

	pNsOperationTypeSnapshotsCreate = cNsOperationTypeSnapshotsCreate
	NsOperationTypeSnapshotsCreate = &pNsOperationTypeSnapshotsCreate

	pNsOperationTypeVolumeCollectionsPromote = cNsOperationTypeVolumeCollectionsPromote
	NsOperationTypeVolumeCollectionsPromote = &pNsOperationTypeVolumeCollectionsPromote

	pNsOperationTypeRemoveInitiator = cNsOperationTypeRemoveInitiator
	NsOperationTypeRemoveInitiator = &pNsOperationTypeRemoveInitiator

}
