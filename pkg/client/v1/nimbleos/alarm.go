// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Alarm - View alarms.
// Export AlarmFields for advance operations like search filter etc.
var AlarmFields *AlarmStringFields

func init() {
	IDfield := "id"
	Typefield := "type"
	Arrayfield := "array"
	CurrOnsetEventIdfield := "curr_onset_event_id"
	ObjectIdfield := "object_id"
	ObjectNamefield := "object_name"
	ObjectTypefield := "object_type"
	OnsetTimefield := "onset_time"
	AckTimefield := "ack_time"
	Statusfield := "status"
	UserIdfield := "user_id"
	UserNamefield := "user_name"
	UserFullNamefield := "user_full_name"
	Categoryfield := "category"
	Severityfield := "severity"
	RemindEveryfield := "remind_every"
	RemindEveryUnitfield := "remind_every_unit"
	Activityfield := "activity"
	NextNotificationTimefield := "next_notification_time"

	AlarmFields = &AlarmStringFields{
		ID:                   &IDfield,
		Type:                 &Typefield,
		Array:                &Arrayfield,
		CurrOnsetEventId:     &CurrOnsetEventIdfield,
		ObjectId:             &ObjectIdfield,
		ObjectName:           &ObjectNamefield,
		ObjectType:           &ObjectTypefield,
		OnsetTime:            &OnsetTimefield,
		AckTime:              &AckTimefield,
		Status:               &Statusfield,
		UserId:               &UserIdfield,
		UserName:             &UserNamefield,
		UserFullName:         &UserFullNamefield,
		Category:             &Categoryfield,
		Severity:             &Severityfield,
		RemindEvery:          &RemindEveryfield,
		RemindEveryUnit:      &RemindEveryUnitfield,
		Activity:             &Activityfield,
		NextNotificationTime: &NextNotificationTimefield,
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
