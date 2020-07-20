// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsReapJobsReturn - Response from reaping jobs.
// Export NsReapJobsReturnFields for advance operations like search filter etc.
var NsReapJobsReturnFields *NsReapJobsReturn

func init() {

	NsReapJobsReturnFields = &NsReapJobsReturn{}
}

type NsReapJobsReturn struct {
	// Reaped - Number of jobs reaped.
	Reaped int64 `json:"reaped,omitempty"`
	// Remaining - Number of jobs remaining.
	Remaining int64 `json:"remaining,omitempty"`
}

// sdk internal struct
type nsReapJobsReturn struct {
	Reaped    *int64 `json:"reaped,omitempty"`
	Remaining *int64 `json:"remaining,omitempty"`
}

// EncodeNsReapJobsReturn - Transform NsReapJobsReturn to nsReapJobsReturn type
func EncodeNsReapJobsReturn(request interface{}) (*nsReapJobsReturn, error) {
	reqNsReapJobsReturn := request.(*NsReapJobsReturn)
	byte, err := json.Marshal(reqNsReapJobsReturn)
	if err != nil {
		return nil, err
	}
	respNsReapJobsReturnPtr := &nsReapJobsReturn{}
	err = json.Unmarshal(byte, respNsReapJobsReturnPtr)
	return respNsReapJobsReturnPtr, err
}

// DecodeNsReapJobsReturn Transform nsReapJobsReturn to NsReapJobsReturn type
func DecodeNsReapJobsReturn(request interface{}) (*NsReapJobsReturn, error) {
	reqNsReapJobsReturn := request.(*nsReapJobsReturn)
	byte, err := json.Marshal(reqNsReapJobsReturn)
	if err != nil {
		return nil, err
	}
	respNsReapJobsReturn := &NsReapJobsReturn{}
	err = json.Unmarshal(byte, respNsReapJobsReturn)
	return respNsReapJobsReturn, err
}
