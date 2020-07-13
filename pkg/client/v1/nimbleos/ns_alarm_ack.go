// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsAlarmAck - Arguments acknowledge alarm.
// Export NsAlarmAckFields for advance operations like search filter etc.
var NsAlarmAckFields *NsAlarmAck

func init() {

	NsAlarmAckFields = &NsAlarmAck{
		ID: "id",
	}
}

type NsAlarmAck struct {
	// ID - ID of the alarm.
	ID string `json:"id,omitempty"`
	// RemindEvery - Notification frequency unit.
	RemindEvery int64 `json:"remind_every,omitempty"`
	// RemindEveryUnit - Period unit.
	RemindEveryUnit *NsPeriodUnit `json:"remind_every_unit,omitempty"`
}

// sdk internal struct
type nsAlarmAck struct {
	// ID - ID of the alarm.
	ID *string `json:"id,omitempty"`
	// RemindEvery - Notification frequency unit.
	RemindEvery *int64 `json:"remind_every,omitempty"`
	// RemindEveryUnit - Period unit.
	RemindEveryUnit *NsPeriodUnit `json:"remind_every_unit,omitempty"`
}

// EncodeNsAlarmAck - Transform NsAlarmAck to nsAlarmAck type
func EncodeNsAlarmAck(request interface{}) (*nsAlarmAck, error) {
	reqNsAlarmAck := request.(*NsAlarmAck)
	byte, err := json.Marshal(reqNsAlarmAck)
	objPtr := &nsAlarmAck{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsAlarmAck Transform nsAlarmAck to NsAlarmAck type
func DecodeNsAlarmAck(request interface{}) (*NsAlarmAck, error) {
	reqNsAlarmAck := request.(*nsAlarmAck)
	byte, err := json.Marshal(reqNsAlarmAck)
	obj := &NsAlarmAck{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
