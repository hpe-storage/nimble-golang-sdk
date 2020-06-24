/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsVolumeCollectionHandoverAttr - Arguments to handover a volume collection.
// Export NsVolumeCollectionHandoverAttrFields for advance operations like search filter etc.
var NsVolumeCollectionHandoverAttrFields *NsVolumeCollectionHandoverAttr

func init(){
	IDfield:= "id"
	ReplicationPartnerIDfield:= "replication_partner_id"
		
	NsVolumeCollectionHandoverAttrFields= &NsVolumeCollectionHandoverAttr{
		ID: &IDfield,
		ReplicationPartnerID: &ReplicationPartnerIDfield,
		
	}
}

type NsVolumeCollectionHandoverAttr struct {
   
    // ID of the volume collection to be handed over to the downstream replication partner.
    
 	ID *string `json:"id,omitempty"`
   
    // ID of the new owner.
    
 	ReplicationPartnerID *string `json:"replication_partner_id,omitempty"`
   
    // Do not automatically reverse direction of replication. Using this argument will prevent the new owner from automatically replicating the volume collection to this node when the handover completes. The default behavior is to enable replication back to this node. Default: 'false'.
    
 	NoReverse *bool `json:"no_reverse,omitempty"`
}
