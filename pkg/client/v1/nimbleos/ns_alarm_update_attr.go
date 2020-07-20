// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsAlarmUpdateAttr - Alarm object used for updating alarms.
// Export NsAlarmUpdateAttrFields for advance operations like search filter etc.
var NsAlarmUpdateAttrFields *NsAlarmUpdateAttr

func init() {

	NsAlarmUpdateAttrFields = &NsAlarmUpdateAttr{
		ID: "id",
	}
}

type NsAlarmUpdateAttr struct {
	// ID - Identifier for the alarm.
	ID string `json:"id,omitempty"`
	// RemindEvery - Frequency of notification.
	RemindEvery int64 `json:"remind_every,omitempty"`
	// RemindEveryUnit - Timeunit over which to send the number of notification specified in 'remind every'.
	RemindEveryUnit *NsPeriodUnit `json:"remind_every_unit,omitempty"`
}

// sdk internal struct
type nsAlarmUpdateAttr struct {
	ID              *string       `json:"id,omitempty"`
	RemindEvery     *int64        `json:"remind_every,omitempty"`
	RemindEveryUnit *NsPeriodUnit `json:"remind_every_unit,omitempty"`
}

// EncodeNsAlarmUpdateAttr - Transform NsAlarmUpdateAttr to nsAlarmUpdateAttr type
func EncodeNsAlarmUpdateAttr(request interface{}) (*nsAlarmUpdateAttr, error) {
	reqNsAlarmUpdateAttr := request.(*NsAlarmUpdateAttr)
	byte, err := json.Marshal(reqNsAlarmUpdateAttr)
	if err != nil {
		return nil, err
	}
	respNsAlarmUpdateAttrPtr := &nsAlarmUpdateAttr{}
	err = json.Unmarshal(byte, respNsAlarmUpdateAttrPtr)
	return respNsAlarmUpdateAttrPtr, err
}

// DecodeNsAlarmUpdateAttr Transform nsAlarmUpdateAttr to NsAlarmUpdateAttr type
func DecodeNsAlarmUpdateAttr(request interface{}) (*NsAlarmUpdateAttr, error) {
	reqNsAlarmUpdateAttr := request.(*nsAlarmUpdateAttr)
	byte, err := json.Marshal(reqNsAlarmUpdateAttr)
	if err != nil {
		return nil, err
	}
	respNsAlarmUpdateAttr := &NsAlarmUpdateAttr{}
	err = json.Unmarshal(byte, respNsAlarmUpdateAttr)
	return respNsAlarmUpdateAttr, err
}
