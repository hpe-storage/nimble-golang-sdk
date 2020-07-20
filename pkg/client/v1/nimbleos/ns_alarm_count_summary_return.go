// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsAlarmCountSummaryReturn - List of alarm count for each category.
// Export NsAlarmCountSummaryReturnFields for advance operations like search filter etc.
var NsAlarmCountSummaryReturnFields *NsAlarmCountSummaryReturn

func init() {

	NsAlarmCountSummaryReturnFields = &NsAlarmCountSummaryReturn{}
}

type NsAlarmCountSummaryReturn struct {
	// AlarmSummary - List of alarm count for each category.
	AlarmSummary []*NsAlarmCount `json:"alarm_summary,omitempty"`
}

// sdk internal struct
type nsAlarmCountSummaryReturn struct {
	AlarmSummary []*NsAlarmCount `json:"alarm_summary,omitempty"`
}

// EncodeNsAlarmCountSummaryReturn - Transform NsAlarmCountSummaryReturn to nsAlarmCountSummaryReturn type
func EncodeNsAlarmCountSummaryReturn(request interface{}) (*nsAlarmCountSummaryReturn, error) {
	reqNsAlarmCountSummaryReturn := request.(*NsAlarmCountSummaryReturn)
	byte, err := json.Marshal(reqNsAlarmCountSummaryReturn)
	if err != nil {
		return nil, err
	}
	respNsAlarmCountSummaryReturnPtr := &nsAlarmCountSummaryReturn{}
	err = json.Unmarshal(byte, respNsAlarmCountSummaryReturnPtr)
	return respNsAlarmCountSummaryReturnPtr, err
}

// DecodeNsAlarmCountSummaryReturn Transform nsAlarmCountSummaryReturn to NsAlarmCountSummaryReturn type
func DecodeNsAlarmCountSummaryReturn(request interface{}) (*NsAlarmCountSummaryReturn, error) {
	reqNsAlarmCountSummaryReturn := request.(*nsAlarmCountSummaryReturn)
	byte, err := json.Marshal(reqNsAlarmCountSummaryReturn)
	if err != nil {
		return nil, err
	}
	respNsAlarmCountSummaryReturn := &NsAlarmCountSummaryReturn{}
	err = json.Unmarshal(byte, respNsAlarmCountSummaryReturn)
	return respNsAlarmCountSummaryReturn, err
}
