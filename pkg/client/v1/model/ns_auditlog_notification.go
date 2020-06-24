/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsAuditlogNotification - Represents an auditlog notification message.
// Export NsAuditlogNotificationFields for advance operations like search filter etc.
var NsAuditlogNotificationFields *NsAuditlogNotification

func init(){
	Activityfield:= "activity"
	ObjectIDfield:= "object_id"
	ObjectNamefield:= "object_name"
		
	NsAuditlogNotificationFields= &NsAuditlogNotification{
		Activity: &Activityfield,
		ObjectID: &ObjectIDfield,
		ObjectName: &ObjectNamefield,
		
	}
}

type NsAuditlogNotification struct {
   
    // Notification Sequence Number.
    
   	SequenceNumber *int64 `json:"sequence_number,omitempty"`
   
    // Represents the type of the notification.
    
   	NotificationType *NsNotificationType `json:"notification_type,omitempty"`
   
    // Represents CUD message of auditlog notification.
    
 	Activity *string `json:"activity,omitempty"`
   
    // Represents the object type of an auditlog based notification.
    
   	ObjectType *NsObjectType `json:"object_type,omitempty"`
   
    // Represents the object of an auditlog based notification.
    
 	ObjectID *string `json:"object_id,omitempty"`
   
    // Represents the object name of an auditlog based notification.
    
 	ObjectName *string `json:"object_name,omitempty"`
   
    // The timestamp when the activity happened in seconds since last epoch.
    
   	Timestamp *int64 `json:"timestamp,omitempty"`
}
