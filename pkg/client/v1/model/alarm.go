/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// Alarm - View alarms.
// Export AlarmFields for advance operations like search filter etc.
var AlarmFields *Alarm

func init(){
	IDfield:= "id"
	Arrayfield:= "array"
	CurrOnsetEventIDfield:= "curr_onset_event_id"
	ObjectIDfield:= "object_id"
	ObjectNamefield:= "object_name"
	UserIDfield:= "user_id"
	UserNamefield:= "user_name"
	UserFullNamefield:= "user_full_name"
	Activityfield:= "activity"
		
	AlarmFields= &Alarm{
		ID: &IDfield,
		Array: &Arrayfield,
		CurrOnsetEventID: &CurrOnsetEventIDfield,
		ObjectID: &ObjectIDfield,
		ObjectName: &ObjectNamefield,
		UserID: &UserIDfield,
		UserName: &UserNamefield,
		UserFullName: &UserFullNamefield,
		Activity: &Activityfield,
		
	}
}

type Alarm struct {
   
    // Identifier for the alarm.
    
 	ID *string `json:"id,omitempty"`
   
    // Identifier for type of alarm.
    
   	Type *int64 `json:"type,omitempty"`
   
    // The array name where the alarm is generated.
    
 	Array *string `json:"array,omitempty"`
   
    // Identifier for the current onset event.
    
 	CurrOnsetEventID *string `json:"curr_onset_event_id,omitempty"`
   
    // Identifier of object operated upon.
    
 	ObjectID *string `json:"object_id,omitempty"`
   
    // Name of object operated upon.
    
 	ObjectName *string `json:"object_name,omitempty"`
   
    // Type of the object being operated upon.
    
   	ObjectType *NsObjectType `json:"object_type,omitempty"`
   
    // Time when this alarm was triggered.
    
   	OnsetTime *int64 `json:"onset_time,omitempty"`
   
    // Time when this alarm was acknowledged.
    
   	AckTime *int64 `json:"ack_time,omitempty"`
   
    // Status of the operation -- open or acknowledged.
    
   	Status *NsAlarmStatus `json:"status,omitempty"`
   
    // Identifier of the user who acknowledged the alarm.
    
 	UserID *string `json:"user_id,omitempty"`
   
    // Username of the user who acknowledged the alarm.
    
 	UserName *string `json:"user_name,omitempty"`
   
    // Full name of the user who acknowledged the alarm.
    
 	UserFullName *string `json:"user_full_name,omitempty"`
   
    // Category of the alarm.
    
   	Category *NsEventCategory `json:"category,omitempty"`
   
    // Severity level of the event.
    
   	Severity *NsAlarmSeverityLevel `json:"severity,omitempty"`
   
    // Frequency of notification. This number and the remind_every_unit define how frequent one alarm notification is sent. For example, a value of 1 with the 'remind_every_unit' of 'days' results in one notification every day.
    
   	RemindEvery *int64 `json:"remind_every,omitempty"`
   
    // Time unit over which to send the number of notification specified in 'remind_every'. For example, a value of 'days' with a 'remind_every' of '1' results in one notification every day.
    
   	RemindEveryUnit *NsPeriodUnit `json:"remind_every_unit,omitempty"`
   
    // Description of activity performed and recorded in alarm.
    
 	Activity *string `json:"activity,omitempty"`
   
    // Time when next reminder for the alarm will be sent.
    
  	NextNotificationTime  *int64 `json:"next_notification_time,omitempty"`
}
