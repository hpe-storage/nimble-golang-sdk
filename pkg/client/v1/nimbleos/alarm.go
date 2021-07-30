// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Alarm - View alarms.

// Export AlarmFields provides field names to use in filter parameters, for example.
var AlarmFields *AlarmStringFields

func init() {
	fieldID := "id"
	fieldType := "type"
	fieldArray := "array"
	fieldCurrOnsetEventId := "curr_onset_event_id"
	fieldObjectId := "object_id"
	fieldObjectName := "object_name"
	fieldObjectType := "object_type"
	fieldOnsetTime := "onset_time"
	fieldAckTime := "ack_time"
	fieldStatus := "status"
	fieldUserId := "user_id"
	fieldUserName := "user_name"
	fieldUserFullName := "user_full_name"
	fieldCategory := "category"
	fieldSeverity := "severity"
	fieldRemindEvery := "remind_every"
	fieldRemindEveryUnit := "remind_every_unit"
	fieldActivity := "activity"
	fieldNextNotificationTime := "next_notification_time"

	AlarmFields = &AlarmStringFields{
		ID:                   &fieldID,
		Type:                 &fieldType,
		Array:                &fieldArray,
		CurrOnsetEventId:     &fieldCurrOnsetEventId,
		ObjectId:             &fieldObjectId,
		ObjectName:           &fieldObjectName,
		ObjectType:           &fieldObjectType,
		OnsetTime:            &fieldOnsetTime,
		AckTime:              &fieldAckTime,
		Status:               &fieldStatus,
		UserId:               &fieldUserId,
		UserName:             &fieldUserName,
		UserFullName:         &fieldUserFullName,
		Category:             &fieldCategory,
		Severity:             &fieldSeverity,
		RemindEvery:          &fieldRemindEvery,
		RemindEveryUnit:      &fieldRemindEveryUnit,
		Activity:             &fieldActivity,
		NextNotificationTime: &fieldNextNotificationTime,
	}
}

type Alarm struct {
	// ID - Identifier for the alarm.
	ID *string `json:"id,omitempty"`
	// Type - Identifier for type of alarm.
	Type *int64 `json:"type,omitempty"`
	// Array - The array name where the alarm is generated.
	Array *string `json:"array,omitempty"`
	// CurrOnsetEventId - Identifier for the current onset event.
	CurrOnsetEventId *string `json:"curr_onset_event_id,omitempty"`
	// ObjectId - Identifier of object operated upon.
	ObjectId *string `json:"object_id,omitempty"`
	// ObjectName - Name of object operated upon.
	ObjectName *string `json:"object_name,omitempty"`
	// ObjectType - Type of the object being operated upon.
	ObjectType *NsObjectType `json:"object_type,omitempty"`
	// OnsetTime - Time when this alarm was triggered.
	OnsetTime *int64 `json:"onset_time,omitempty"`
	// AckTime - Time when this alarm was acknowledged.
	AckTime *int64 `json:"ack_time,omitempty"`
	// Status - Status of the operation -- open or acknowledged.
	Status *NsAlarmStatus `json:"status,omitempty"`
	// UserId - Identifier of the user who acknowledged the alarm.
	UserId *string `json:"user_id,omitempty"`
	// UserName - Username of the user who acknowledged the alarm.
	UserName *string `json:"user_name,omitempty"`
	// UserFullName - Full name of the user who acknowledged the alarm.
	UserFullName *string `json:"user_full_name,omitempty"`
	// Category - Category of the alarm.
	Category *NsEventCategory `json:"category,omitempty"`
	// Severity - Severity level of the event.
	Severity *NsAlarmSeverityLevel `json:"severity,omitempty"`
	// RemindEvery - Frequency of notification. This number and the remind_every_unit define how frequent one alarm notification is sent. For example, a value of 1 with the 'remind_every_unit' of 'days' results in one notification every day.
	RemindEvery *int64 `json:"remind_every,omitempty"`
	// RemindEveryUnit - Time unit over which to send the number of notification specified in 'remind_every'. For example, a value of 'days' with a 'remind_every' of '1' results in one notification every day.
	RemindEveryUnit *NsPeriodUnit `json:"remind_every_unit,omitempty"`
	// Activity - Description of activity performed and recorded in alarm.
	Activity *string `json:"activity,omitempty"`
	// NextNotificationTime - Time when next reminder for the alarm will be sent.
	NextNotificationTime *int64 `json:"next_notification_time,omitempty"`
}

// Struct for AlarmFields
type AlarmStringFields struct {
	ID                   *string
	Type                 *string
	Array                *string
	CurrOnsetEventId     *string
	ObjectId             *string
	ObjectName           *string
	ObjectType           *string
	OnsetTime            *string
	AckTime              *string
	Status               *string
	UserId               *string
	UserName             *string
	UserFullName         *string
	Category             *string
	Severity             *string
	RemindEvery          *string
	RemindEveryUnit      *string
	Activity             *string
	NextNotificationTime *string
}
