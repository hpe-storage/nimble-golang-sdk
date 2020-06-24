/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsVolumeCollectionDemoteAttr - Arguments to demote a volume collection.
// Export NsVolumeCollectionDemoteAttrFields for advance operations like search filter etc.
var NsVolumeCollectionDemoteAttrFields *NsVolumeCollectionDemoteAttr

func init(){
	IDfield:= "id"
	ReplicationPartnerIDfield:= "replication_partner_id"
		
	NsVolumeCollectionDemoteAttrFields= &NsVolumeCollectionDemoteAttr{
		ID: &IDfield,
		ReplicationPartnerID: &ReplicationPartnerIDfield,
		
	}
}

type NsVolumeCollectionDemoteAttr struct {
   
    // ID of the demoted volume collection.
    
 	ID *string `json:"id,omitempty"`
   
    // ID of the new owner. If invoke_on_upstream_partner is provided, utilize the ID of the current owner i.e. upstream replication partner.
    
 	ReplicationPartnerID *string `json:"replication_partner_id,omitempty"`
}
