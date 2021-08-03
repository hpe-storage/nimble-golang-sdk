// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAuditlogNotificationFields provides field names to use in filter parameters, for example.
var NsAuditlogNotificationFields *NsAuditlogNotificationFieldHandles

func init() {
	NsAuditlogNotificationFields = &NsAuditlogNotificationFieldHandles{
		SequenceNumber:   "sequence_number",
		NotificationType: "notification_type",
		Activity:         "activity",
		ObjectType:       "object_type",
		ObjectId:         "object_id",
		ObjectName:       "object_name",
		Timestamp:        "timestamp",
	}
}

// NsAuditlogNotification - Represents an auditlog notification message.
type NsAuditlogNotification struct {
	// SequenceNumber - Notification Sequence Number.
	SequenceNumber *int64 `json:"sequence_number,omitempty"`
	// NotificationType - Represents the type of the notification.
	NotificationType *NsNotificationType `json:"notification_type,omitempty"`
	// Activity - Represents CUD message of auditlog notification.
	Activity *string `json:"activity,omitempty"`
	// ObjectType - Represents the object type of an auditlog based notification.
	ObjectType *NsObjectType `json:"object_type,omitempty"`
	// ObjectId - Represents the object of an auditlog based notification.
	ObjectId *string `json:"object_id,omitempty"`
	// ObjectName - Represents the object name of an auditlog based notification.
	ObjectName *string `json:"object_name,omitempty"`
	// Timestamp - The timestamp when the activity happened in seconds since last epoch.
	Timestamp *int64 `json:"timestamp,omitempty"`
}

// NsAuditlogNotificationFieldHandles provides a string representation for each NsAuditlogNotification field.
type NsAuditlogNotificationFieldHandles struct {
	SequenceNumber   string
	NotificationType string
	Activity         string
	ObjectType       string
	ObjectId         string
	ObjectName       string
	Timestamp        string
}
