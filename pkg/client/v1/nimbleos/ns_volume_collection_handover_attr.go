// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsVolumeCollectionHandoverAttr - Arguments to handover a volume collection.
// Export NsVolumeCollectionHandoverAttrFields for advance operations like search filter etc.
var NsVolumeCollectionHandoverAttrFields *NsVolumeCollectionHandoverAttr

func init() {

	NsVolumeCollectionHandoverAttrFields = &NsVolumeCollectionHandoverAttr{
		ID:                   "id",
		ReplicationPartnerId: "replication_partner_id",
	}
}

type NsVolumeCollectionHandoverAttr struct {
	// ID - ID of the volume collection to be handed over to the downstream replication partner.
	ID string `json:"id,omitempty"`
	// ReplicationPartnerId - ID of the new owner.
	ReplicationPartnerId string `json:"replication_partner_id,omitempty"`
	// NoReverse - Do not automatically reverse direction of replication. Using this argument will prevent the new owner from automatically replicating the volume collection to this node when the handover completes. The default behavior is to enable replication back to this node. Default: 'false'.
	NoReverse *bool `json:"no_reverse,omitempty"`
}

// sdk internal struct
type nsVolumeCollectionHandoverAttr struct {
	ID                   *string `json:"id,omitempty"`
	ReplicationPartnerId *string `json:"replication_partner_id,omitempty"`
	NoReverse            *bool   `json:"no_reverse,omitempty"`
}

// EncodeNsVolumeCollectionHandoverAttr - Transform NsVolumeCollectionHandoverAttr to nsVolumeCollectionHandoverAttr type
func EncodeNsVolumeCollectionHandoverAttr(request interface{}) (*nsVolumeCollectionHandoverAttr, error) {
	reqNsVolumeCollectionHandoverAttr := request.(*NsVolumeCollectionHandoverAttr)
	byte, err := json.Marshal(reqNsVolumeCollectionHandoverAttr)
	if err != nil {
		return nil, err
	}
	respNsVolumeCollectionHandoverAttrPtr := &nsVolumeCollectionHandoverAttr{}
	err = json.Unmarshal(byte, respNsVolumeCollectionHandoverAttrPtr)
	return respNsVolumeCollectionHandoverAttrPtr, err
}

// DecodeNsVolumeCollectionHandoverAttr Transform nsVolumeCollectionHandoverAttr to NsVolumeCollectionHandoverAttr type
func DecodeNsVolumeCollectionHandoverAttr(request interface{}) (*NsVolumeCollectionHandoverAttr, error) {
	reqNsVolumeCollectionHandoverAttr := request.(*nsVolumeCollectionHandoverAttr)
	byte, err := json.Marshal(reqNsVolumeCollectionHandoverAttr)
	if err != nil {
		return nil, err
	}
	respNsVolumeCollectionHandoverAttr := &NsVolumeCollectionHandoverAttr{}
	err = json.Unmarshal(byte, respNsVolumeCollectionHandoverAttr)
	return respNsVolumeCollectionHandoverAttr, err
}
