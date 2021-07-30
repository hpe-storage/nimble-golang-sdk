// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsVolumeCollectionDemoteAttr - Arguments to demote a volume collection.

// Export NsVolumeCollectionDemoteAttrFields provides field names to use in filter parameters, for example.
var NsVolumeCollectionDemoteAttrFields *NsVolumeCollectionDemoteAttrStringFields

func init() {
	fieldID := "id"
	fieldReplicationPartnerId := "replication_partner_id"

	NsVolumeCollectionDemoteAttrFields = &NsVolumeCollectionDemoteAttrStringFields{
		ID:                   &fieldID,
		ReplicationPartnerId: &fieldReplicationPartnerId,
	}
}

type NsVolumeCollectionDemoteAttr struct {
	// ID - ID of the demoted volume collection.
	ID *string `json:"id,omitempty"`
	// ReplicationPartnerId - ID of the new owner. If invoke_on_upstream_partner is provided, utilize the ID of the current owner i.e. upstream replication partner.
	ReplicationPartnerId *string `json:"replication_partner_id,omitempty"`
}

// Struct for NsVolumeCollectionDemoteAttrFields
type NsVolumeCollectionDemoteAttrStringFields struct {
	ID                   *string
	ReplicationPartnerId *string
}
