// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// Alarm - View alarms.
// Export AlarmFields for advance operations like search filter etc.
var AlarmFields *Alarm

func init() {

	AlarmFields = &Alarm{
		ID:               "id",
		Array:            "array",
		CurrOnsetEventId: "curr_onset_event_id",
		ObjectId:         "object_id",
		ObjectName:       "object_name",
		UserId:           "user_id",
		UserName:         "user_name",
		UserFullName:     "user_full_name",
		Activity:         "activity",
	}
}

type Alarm struct {
	// ID - Identifier for the alarm.
	ID string `json:"id,omitempty"`
	// Type - Identifier for type of alarm.
	Type int64 `json:"type,omitempty"`
	// Array - The array name where the alarm is generated.
	Array string `json:"array,omitempty"`
	// CurrOnsetEventId - Identifier for the current onset event.
	CurrOnsetEventId string `json:"curr_onset_event_id,omitempty"`
	// ObjectId - Identifier of object operated upon.
	ObjectId string `json:"object_id,omitempty"`
	// ObjectName - Name of object operated upon.
	ObjectName string `json:"object_name,omitempty"`
	// ObjectType - Type of the object being operated upon.
	ObjectType *NsObjectType `json:"object_type,omitempty"`
	// OnsetTime - Time when this alarm was triggered.
	OnsetTime int64 `json:"onset_time,omitempty"`
	// AckTime - Time when this alarm was acknowledged.
	AckTime int64 `json:"ack_time,omitempty"`
	// Status - Status of the operation -- open or acknowledged.
	Status *NsAlarmStatus `json:"status,omitempty"`
	// UserId - Identifier of the user who acknowledged the alarm.
	UserId string `json:"user_id,omitempty"`
	// UserName - Username of the user who acknowledged the alarm.
	UserName string `json:"user_name,omitempty"`
	// UserFullName - Full name of the user who acknowledged the alarm.
	UserFullName string `json:"user_full_name,omitempty"`
	// Category - Category of the alarm.
	Category *NsEventCategory `json:"category,omitempty"`
	// Severity - Severity level of the event.
	Severity *NsAlarmSeverityLevel `json:"severity,omitempty"`
	// RemindEvery - Frequency of notification. This number and the remind_every_unit define how frequent one alarm notification is sent. For example, a value of 1 with the 'remind_every_unit' of 'days' results in one notification every day.
	RemindEvery int64 `json:"remind_every,omitempty"`
	// RemindEveryUnit - Time unit over which to send the number of notification specified in 'remind_every'. For example, a value of 'days' with a 'remind_every' of '1' results in one notification every day.
	RemindEveryUnit *NsPeriodUnit `json:"remind_every_unit,omitempty"`
	// Activity - Description of activity performed and recorded in alarm.
	Activity string `json:"activity,omitempty"`
	// NextNotificationTime - Time when next reminder for the alarm will be sent.
	NextNotificationTime int64 `json:"next_notification_time,omitempty"`
}

// sdk internal struct
type alarm struct {
	ID                   *string               `json:"id,omitempty"`
	Type                 *int64                `json:"type,omitempty"`
	Array                *string               `json:"array,omitempty"`
	CurrOnsetEventId     *string               `json:"curr_onset_event_id,omitempty"`
	ObjectId             *string               `json:"object_id,omitempty"`
	ObjectName           *string               `json:"object_name,omitempty"`
	ObjectType           *NsObjectType         `json:"object_type,omitempty"`
	OnsetTime            *int64                `json:"onset_time,omitempty"`
	AckTime              *int64                `json:"ack_time,omitempty"`
	Status               *NsAlarmStatus        `json:"status,omitempty"`
	UserId               *string               `json:"user_id,omitempty"`
	UserName             *string               `json:"user_name,omitempty"`
	UserFullName         *string               `json:"user_full_name,omitempty"`
	Category             *NsEventCategory      `json:"category,omitempty"`
	Severity             *NsAlarmSeverityLevel `json:"severity,omitempty"`
	RemindEvery          *int64                `json:"remind_every,omitempty"`
	RemindEveryUnit      *NsPeriodUnit         `json:"remind_every_unit,omitempty"`
	Activity             *string               `json:"activity,omitempty"`
	NextNotificationTime *int64                `json:"next_notification_time,omitempty"`
}

// EncodeAlarm - Transform Alarm to alarm type
func EncodeAlarm(request interface{}) (*alarm, error) {
	reqAlarm := request.(*Alarm)
	byte, err := json.Marshal(reqAlarm)
	if err != nil {
		return nil, err
	}
	respAlarmPtr := &alarm{}
	err = json.Unmarshal(byte, respAlarmPtr)
	return respAlarmPtr, err
}

// DecodeAlarm Transform alarm to Alarm type
func DecodeAlarm(request interface{}) (*Alarm, error) {
	reqAlarm := request.(*alarm)
	byte, err := json.Marshal(reqAlarm)
	if err != nil {
		return nil, err
	}
	respAlarm := &Alarm{}
	err = json.Unmarshal(byte, respAlarm)
	return respAlarm, err
}
