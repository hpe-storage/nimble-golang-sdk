// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsSnapVol - Select fields containing volume info.
// Export NsSnapVolFields for advance operations like search filter etc.
var NsSnapVolFields *NsSnapVol

func init() {

	NsSnapVolFields = &NsSnapVol{
		VolId:           "vol_id",
		SnapName:        "snap_name",
		SnapDescription: "snap_description",
		Cookie:          "cookie",
		AppUuid:         "app_uuid",
	}
}

type NsSnapVol struct {
	// VolId - ID of volume.
	VolId string `json:"vol_id,omitempty"`
	// SnapName - Snapshot name.
	SnapName string `json:"snap_name,omitempty"`
	// SnapDescription - Snapshot description.
	SnapDescription string `json:"snap_description,omitempty"`
	// Cookie - A cookie.
	Cookie string `json:"cookie,omitempty"`
	// Online - Snapshot is online.
	Online *bool `json:"online,omitempty"`
	// Writable - Snapshot is writable.
	Writable *bool `json:"writable,omitempty"`
	// AppUuid - Application identifier of snapshots.
	AppUuid string `json:"app_uuid,omitempty"`
	// AgentType - External management agent type.
	AgentType *NsAgentType `json:"agent_type,omitempty"`
	// Metadata - Key-value pairs that augment a snapshot's attributes.
	Metadata []*NsKeyValue `json:"metadata,omitempty"`
}

// sdk internal struct
type nsSnapVol struct {
	VolId           *string       `json:"vol_id,omitempty"`
	SnapName        *string       `json:"snap_name,omitempty"`
	SnapDescription *string       `json:"snap_description,omitempty"`
	Cookie          *string       `json:"cookie,omitempty"`
	Online          *bool         `json:"online,omitempty"`
	Writable        *bool         `json:"writable,omitempty"`
	AppUuid         *string       `json:"app_uuid,omitempty"`
	AgentType       *NsAgentType  `json:"agent_type,omitempty"`
	Metadata        []*NsKeyValue `json:"metadata,omitempty"`
}

// EncodeNsSnapVol - Transform NsSnapVol to nsSnapVol type
func EncodeNsSnapVol(request interface{}) (*nsSnapVol, error) {
	reqNsSnapVol := request.(*NsSnapVol)
	byte, err := json.Marshal(reqNsSnapVol)
	if err != nil {
		return nil, err
	}
	respNsSnapVolPtr := &nsSnapVol{}
	err = json.Unmarshal(byte, respNsSnapVolPtr)
	return respNsSnapVolPtr, err
}

// DecodeNsSnapVol Transform nsSnapVol to NsSnapVol type
func DecodeNsSnapVol(request interface{}) (*NsSnapVol, error) {
	reqNsSnapVol := request.(*nsSnapVol)
	byte, err := json.Marshal(reqNsSnapVol)
	if err != nil {
		return nil, err
	}
	respNsSnapVol := &NsSnapVol{}
	err = json.Unmarshal(byte, respNsSnapVol)
	return respNsSnapVol, err
}
