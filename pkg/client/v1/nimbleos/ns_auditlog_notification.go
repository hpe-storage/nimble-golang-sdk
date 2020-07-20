// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsAuditlogNotification - Represents an auditlog notification message.
// Export NsAuditlogNotificationFields for advance operations like search filter etc.
var NsAuditlogNotificationFields *NsAuditlogNotification

func init() {

	NsAuditlogNotificationFields = &NsAuditlogNotification{
		Activity:   "activity",
		ObjectId:   "object_id",
		ObjectName: "object_name",
	}
}

type NsAuditlogNotification struct {
	// SequenceNumber - Notification Sequence Number.
	SequenceNumber int64 `json:"sequence_number,omitempty"`
	// NotificationType - Represents the type of the notification.
	NotificationType *NsNotificationType `json:"notification_type,omitempty"`
	// Activity - Represents CUD message of auditlog notification.
	Activity string `json:"activity,omitempty"`
	// ObjectType - Represents the object type of an auditlog based notification.
	ObjectType *NsObjectType `json:"object_type,omitempty"`
	// ObjectId - Represents the object of an auditlog based notification.
	ObjectId string `json:"object_id,omitempty"`
	// ObjectName - Represents the object name of an auditlog based notification.
	ObjectName string `json:"object_name,omitempty"`
	// Timestamp - The timestamp when the activity happened in seconds since last epoch.
	Timestamp int64 `json:"timestamp,omitempty"`
}

// sdk internal struct
type nsAuditlogNotification struct {
	SequenceNumber   *int64              `json:"sequence_number,omitempty"`
	NotificationType *NsNotificationType `json:"notification_type,omitempty"`
	Activity         *string             `json:"activity,omitempty"`
	ObjectType       *NsObjectType       `json:"object_type,omitempty"`
	ObjectId         *string             `json:"object_id,omitempty"`
	ObjectName       *string             `json:"object_name,omitempty"`
	Timestamp        *int64              `json:"timestamp,omitempty"`
}

// EncodeNsAuditlogNotification - Transform NsAuditlogNotification to nsAuditlogNotification type
func EncodeNsAuditlogNotification(request interface{}) (*nsAuditlogNotification, error) {
	reqNsAuditlogNotification := request.(*NsAuditlogNotification)
	byte, err := json.Marshal(reqNsAuditlogNotification)
	if err != nil {
		return nil, err
	}
	respNsAuditlogNotificationPtr := &nsAuditlogNotification{}
	err = json.Unmarshal(byte, respNsAuditlogNotificationPtr)
	return respNsAuditlogNotificationPtr, err
}

// DecodeNsAuditlogNotification Transform nsAuditlogNotification to NsAuditlogNotification type
func DecodeNsAuditlogNotification(request interface{}) (*NsAuditlogNotification, error) {
	reqNsAuditlogNotification := request.(*nsAuditlogNotification)
	byte, err := json.Marshal(reqNsAuditlogNotification)
	if err != nil {
		return nil, err
	}
	respNsAuditlogNotification := &NsAuditlogNotification{}
	err = json.Unmarshal(byte, respNsAuditlogNotification)
	return respNsAuditlogNotification, err
}
