// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsVolumeCollectionDemoteAttr - Arguments to demote a volume collection.
// Export NsVolumeCollectionDemoteAttrFields for advance operations like search filter etc.
var NsVolumeCollectionDemoteAttrFields *NsVolumeCollectionDemoteAttr

func init(){
	IDfield:= "id"
	ReplicationPartnerIdfield:= "replication_partner_id"
		
	NsVolumeCollectionDemoteAttrFields= &NsVolumeCollectionDemoteAttr{
		ID:                    &IDfield,
		ReplicationPartnerId:  &ReplicationPartnerIdfield,
	}
}

type NsVolumeCollectionDemoteAttr struct {
	// ID - ID of the demoted volume collection.
 	ID *string `json:"id,omitempty"`
	// ReplicationPartnerId - ID of the new owner. If invoke_on_upstream_partner is provided, utilize the ID of the current owner i.e. upstream replication partner.
 	ReplicationPartnerId *string `json:"replication_partner_id,omitempty"`
}
