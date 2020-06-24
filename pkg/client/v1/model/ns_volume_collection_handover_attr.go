// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsVolumeCollectionHandoverAttr - Arguments to handover a volume collection.
// Export NsVolumeCollectionHandoverAttrFields for advance operations like search filter etc.
var NsVolumeCollectionHandoverAttrFields *NsVolumeCollectionHandoverAttr

func init(){
	IDfield:= "id"
	ReplicationPartnerIdfield:= "replication_partner_id"
		
	NsVolumeCollectionHandoverAttrFields= &NsVolumeCollectionHandoverAttr{
	ID: &IDfield,
	ReplicationPartnerId: &ReplicationPartnerIdfield,
		
	}
}

type NsVolumeCollectionHandoverAttr struct {
	// ID - ID of the volume collection to be handed over to the downstream replication partner.
 	ID *string `json:"id,omitempty"`
	// ReplicationPartnerId - ID of the new owner.
 	ReplicationPartnerId *string `json:"replication_partner_id,omitempty"`
	// NoReverse - Do not automatically reverse direction of replication. Using this argument will prevent the new owner from automatically replicating the volume collection to this node when the handover completes. The default behavior is to enable replication back to this node. Default: 'false'.
 	NoReverse *bool `json:"no_reverse,omitempty"`
}
