// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsVolumeCollectionDemoteAttr - Arguments to demote a volume collection.
// Export NsVolumeCollectionDemoteAttrFields for advance operations like search filter etc.
var NsVolumeCollectionDemoteAttrFields *NsVolumeCollectionDemoteAttr

func init() {

	NsVolumeCollectionDemoteAttrFields = &NsVolumeCollectionDemoteAttr{
		ID:                   "id",
		ReplicationPartnerId: "replication_partner_id",
	}
}

type NsVolumeCollectionDemoteAttr struct {
	// ID - ID of the demoted volume collection.
	ID string `json:"id,omitempty"`
	// ReplicationPartnerId - ID of the new owner. If invoke_on_upstream_partner is provided, utilize the ID of the current owner i.e. upstream replication partner.
	ReplicationPartnerId string `json:"replication_partner_id,omitempty"`
}

// sdk internal struct
type nsVolumeCollectionDemoteAttr struct {
	ID                   *string `json:"id,omitempty"`
	ReplicationPartnerId *string `json:"replication_partner_id,omitempty"`
}

// EncodeNsVolumeCollectionDemoteAttr - Transform NsVolumeCollectionDemoteAttr to nsVolumeCollectionDemoteAttr type
func EncodeNsVolumeCollectionDemoteAttr(request interface{}) (*nsVolumeCollectionDemoteAttr, error) {
	reqNsVolumeCollectionDemoteAttr := request.(*NsVolumeCollectionDemoteAttr)
	byte, err := json.Marshal(reqNsVolumeCollectionDemoteAttr)
	if err != nil {
		return nil, err
	}
	respNsVolumeCollectionDemoteAttrPtr := &nsVolumeCollectionDemoteAttr{}
	err = json.Unmarshal(byte, respNsVolumeCollectionDemoteAttrPtr)
	return respNsVolumeCollectionDemoteAttrPtr, err
}

// DecodeNsVolumeCollectionDemoteAttr Transform nsVolumeCollectionDemoteAttr to NsVolumeCollectionDemoteAttr type
func DecodeNsVolumeCollectionDemoteAttr(request interface{}) (*NsVolumeCollectionDemoteAttr, error) {
	reqNsVolumeCollectionDemoteAttr := request.(*nsVolumeCollectionDemoteAttr)
	byte, err := json.Marshal(reqNsVolumeCollectionDemoteAttr)
	if err != nil {
		return nil, err
	}
	respNsVolumeCollectionDemoteAttr := &NsVolumeCollectionDemoteAttr{}
	err = json.Unmarshal(byte, respNsVolumeCollectionDemoteAttr)
	return respNsVolumeCollectionDemoteAttr, err
}
