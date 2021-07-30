// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAuditlogNotification - Represents an auditlog notification message.

// Export NsAuditlogNotificationFields provides field names to use in filter parameters, for example.
var NsAuditlogNotificationFields *NsAuditlogNotificationStringFields

func init() {
	fieldSequenceNumber := "sequence_number"
	fieldNotificationType := "notification_type"
	fieldActivity := "activity"
	fieldObjectType := "object_type"
	fieldObjectId := "object_id"
	fieldObjectName := "object_name"
	fieldTimestamp := "timestamp"

	NsAuditlogNotificationFields = &NsAuditlogNotificationStringFields{
		SequenceNumber:   &fieldSequenceNumber,
		NotificationType: &fieldNotificationType,
		Activity:         &fieldActivity,
		ObjectType:       &fieldObjectType,
		ObjectId:         &fieldObjectId,
		ObjectName:       &fieldObjectName,
		Timestamp:        &fieldTimestamp,
	}
}

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

// Struct for NsAuditlogNotificationFields
type NsAuditlogNotificationStringFields struct {
	SequenceNumber   *string
	NotificationType *string
	Activity         *string
	ObjectType       *string
	ObjectId         *string
	ObjectName       *string
	Timestamp        *string
}
