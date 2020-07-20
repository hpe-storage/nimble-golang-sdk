// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsAlarmCount - List of alarm count for each category.
// Export NsAlarmCountFields for advance operations like search filter etc.
var NsAlarmCountFields *NsAlarmCount

func init() {

	NsAlarmCountFields = &NsAlarmCount{}
}

type NsAlarmCount struct {
	// Category - Alert category.
	Category *NsEventCategory `json:"category,omitempty"`
	// Critical - Critical alarm count of a particular category.
	Critical int64 `json:"critical,omitempty"`
	// Warning - Warning alarm count of a particular category.
	Warning int64 `json:"warning,omitempty"`
}

// sdk internal struct
type nsAlarmCount struct {
	Category *NsEventCategory `json:"category,omitempty"`
	Critical *int64           `json:"critical,omitempty"`
	Warning  *int64           `json:"warning,omitempty"`
}

// EncodeNsAlarmCount - Transform NsAlarmCount to nsAlarmCount type
func EncodeNsAlarmCount(request interface{}) (*nsAlarmCount, error) {
	reqNsAlarmCount := request.(*NsAlarmCount)
	byte, err := json.Marshal(reqNsAlarmCount)
	if err != nil {
		return nil, err
	}
	respNsAlarmCountPtr := &nsAlarmCount{}
	err = json.Unmarshal(byte, respNsAlarmCountPtr)
	return respNsAlarmCountPtr, err
}

// DecodeNsAlarmCount Transform nsAlarmCount to NsAlarmCount type
func DecodeNsAlarmCount(request interface{}) (*NsAlarmCount, error) {
	reqNsAlarmCount := request.(*nsAlarmCount)
	byte, err := json.Marshal(reqNsAlarmCount)
	if err != nil {
		return nil, err
	}
	respNsAlarmCount := &NsAlarmCount{}
	err = json.Unmarshal(byte, respNsAlarmCount)
	return respNsAlarmCount, err
}
