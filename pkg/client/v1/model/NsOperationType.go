/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsOperationType.</p>
 */

type NsOperationType string 

const (
	NSOPERATIONTYPE_CREATE_INITIATOR_GRP NsOperationType = "create_initiator_grp"
	NSOPERATIONTYPE_VOLUME_COLLECTIONS_HANDOVER NsOperationType = "volume_collections_handover"
	NSOPERATIONTYPE_FOLDERS_DELETE NsOperationType = "folders_delete"
	NSOPERATIONTYPE_SNAPSHOTS_DELETE NsOperationType = "snapshots_delete"
	NSOPERATIONTYPE_ADD_VOL_ACL NsOperationType = "add_vol_acl"
	NSOPERATIONTYPE_ALARMS_UPDATE NsOperationType = "alarms_update"
	NSOPERATIONTYPE_ALARMS_ACKNOWLEDGE NsOperationType = "alarms_acknowledge"
	NSOPERATIONTYPE_EDIT_INITIATOR_GRP NsOperationType = "edit_initiator_grp"
	NSOPERATIONTYPE_VOLUME_COLLECTIONS_DEMOTE NsOperationType = "volume_collections_demote"
	NSOPERATIONTYPE_VOLUMES_BULK_DELETE NsOperationType = "volumes_bulk_delete"
	NSOPERATIONTYPE_EDIT_FOLDER NsOperationType = "edit_folder"
	NSOPERATIONTYPE_REMOVE_SNAPSHOT NsOperationType = "remove_snapshot"
	NSOPERATIONTYPE_VOLUME_COLLECTIONS_BULK_HANDOVER NsOperationType = "volume_collections_bulk_handover"
	NSOPERATIONTYPE_VOLUMES_DELETE NsOperationType = "volumes_delete"
	NSOPERATIONTYPE_ADD_FC_VOL_ACL NsOperationType = "add_fc_vol_acl"
	NSOPERATIONTYPE_SNAPSHOT_COLLECTIONS_DELETE NsOperationType = "snapshot_collections_delete"
	NSOPERATIONTYPE_ADD_INITIATOR NsOperationType = "add_initiator"
	NSOPERATIONTYPE_EDIT_SNAP NsOperationType = "edit_snap"
	NSOPERATIONTYPE_ALARMS_UNACKNOWLEDGE NsOperationType = "alarms_unacknowledge"
	NSOPERATIONTYPE_ALARMS_BULK_UNACKNOWLEDGE NsOperationType = "alarms_bulk_unacknowledge"
	NSOPERATIONTYPE_SNAP_VOLCOLL NsOperationType = "snap_volcoll"
	NSOPERATIONTYPE_EDIT_VOL NsOperationType = "edit_vol"
	NSOPERATIONTYPE_VOLUME_COLLECTIONS_BULK_DEMOTE NsOperationType = "volume_collections_bulk_demote"
	NSOPERATIONTYPE_ALARMS_BULK_ACKNOWLEDGE NsOperationType = "alarms_bulk_acknowledge"
	NSOPERATIONTYPE_SNAPSHOTS_BULK_DELETE NsOperationType = "snapshots_bulk_delete"
	NSOPERATIONTYPE_FOLDERS_BULK_DELETE NsOperationType = "folders_bulk_delete"
	NSOPERATIONTYPE_VOLUME_COLLECTIONS_BULK_PROMOTE NsOperationType = "volume_collections_bulk_promote"
	NSOPERATIONTYPE_VOLUMES_BULK_UPDATE NsOperationType = "volumes_bulk_update"
	NSOPERATIONTYPE_SNAP_VOL NsOperationType = "snap_vol"
	NSOPERATIONTYPE_REMOVE_VOL_ACL NsOperationType = "remove_vol_acl"
	NSOPERATIONTYPE_VOLUMES_UPDATE NsOperationType = "volumes_update"
	NSOPERATIONTYPE_ALARMS_BULK_UPDATE NsOperationType = "alarms_bulk_update"
	NSOPERATIONTYPE_DELETE_INITIATOR_GRP NsOperationType = "delete_initiator_grp"
	NSOPERATIONTYPE_SNAPSHOTS_BULK_CREATE NsOperationType = "snapshots_bulk_create"
	NSOPERATIONTYPE_SNAPSHOT_COLLECTIONS_BULK_DELETE NsOperationType = "snapshot_collections_bulk_delete"
	NSOPERATIONTYPE_SNAPSHOTS_CREATE NsOperationType = "snapshots_create"
	NSOPERATIONTYPE_VOLUME_COLLECTIONS_PROMOTE NsOperationType = "volume_collections_promote"
	NSOPERATIONTYPE_REMOVE_INITIATOR NsOperationType = "remove_initiator"

) 
