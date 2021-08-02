// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsVolumeCollectionDemoteAttrFields provides field names to use in filter parameters, for example.
var NsVolumeCollectionDemoteAttrFields *NsVolumeCollectionDemoteAttrFieldHandles

func init() {
	NsVolumeCollectionDemoteAttrFields = &NsVolumeCollectionDemoteAttrFieldHandles{
		ID:                   "id",
		ReplicationPartnerId: "replication_partner_id",
	}
}

// NsVolumeCollectionDemoteAttr - Arguments to demote a volume collection.
type NsVolumeCollectionDemoteAttr struct {
	// ID - ID of the demoted volume collection.
	ID *string `json:"id,omitempty"`
	// ReplicationPartnerId - ID of the new owner. If invoke_on_upstream_partner is provided, utilize the ID of the current owner i.e. upstream replication partner.
	ReplicationPartnerId *string `json:"replication_partner_id,omitempty"`
}

// NsVolumeCollectionDemoteAttrFieldHandles provides a string representation for each AccessControlRecord field.
type NsVolumeCollectionDemoteAttrFieldHandles struct {
	ID                   string
	ReplicationPartnerId string
}
