// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsAlertNotification - Represents an Alert notification message.
// Export NsAlertNotificationFields for advance operations like search filter etc.
var NsAlertNotificationFields *NsAlertNotification

func init() {

	NsAlertNotificationFields = &NsAlertNotification{
		Activity: "activity",
	}
}

type NsAlertNotification struct {
	// SequenceNumber - Notification Sequence Number.
	SequenceNumber int64 `json:"sequence_number,omitempty"`
	// NotificationType - Represents the type of the notification.
	NotificationType *NsNotificationType `json:"notification_type,omitempty"`
	// Activity - Represents the alert message details of the notification.
	Activity string `json:"activity,omitempty"`
	// EventTarget - The kind of events or alerts that the notification subscriber is interested in.
	EventTarget *NsEventTargetTypeOrAll `json:"event_target,omitempty"`
	// Category - The category of the events or alerts that the notification subscriber is interested in.
	Category *NsEventCategory `json:"category,omitempty"`
	// Severity - The severity of events that the notification subscriber is interested in.
	Severity *NsSeverityLevel `json:"severity,omitempty"`
	// AlertType - Identifier for type of alert.
	AlertType int64 `json:"alert_type,omitempty"`
	// Timestamp - The timestamp when the activity happened in seconds since last epoch.
	Timestamp int64 `json:"timestamp,omitempty"`
}

// sdk internal struct
type nsAlertNotification struct {
	// SequenceNumber - Notification Sequence Number.
	SequenceNumber *int64 `json:"sequence_number,omitempty"`
	// NotificationType - Represents the type of the notification.
	NotificationType *NsNotificationType `json:"notification_type,omitempty"`
	// Activity - Represents the alert message details of the notification.
	Activity *string `json:"activity,omitempty"`
	// EventTarget - The kind of events or alerts that the notification subscriber is interested in.
	EventTarget *NsEventTargetTypeOrAll `json:"event_target,omitempty"`
	// Category - The category of the events or alerts that the notification subscriber is interested in.
	Category *NsEventCategory `json:"category,omitempty"`
	// Severity - The severity of events that the notification subscriber is interested in.
	Severity *NsSeverityLevel `json:"severity,omitempty"`
	// AlertType - Identifier for type of alert.
	AlertType *int64 `json:"alert_type,omitempty"`
	// Timestamp - The timestamp when the activity happened in seconds since last epoch.
	Timestamp *int64 `json:"timestamp,omitempty"`
}

// EncodeNsAlertNotification - Transform NsAlertNotification to nsAlertNotification type
func EncodeNsAlertNotification(request interface{}) (*nsAlertNotification, error) {
	reqNsAlertNotification := request.(*NsAlertNotification)
	byte, err := json.Marshal(reqNsAlertNotification)
	objPtr := &nsAlertNotification{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsAlertNotification Transform nsAlertNotification to NsAlertNotification type
func DecodeNsAlertNotification(request interface{}) (*NsAlertNotification, error) {
	reqNsAlertNotification := request.(*nsAlertNotification)
	byte, err := json.Marshal(reqNsAlertNotification)
	obj := &NsAlertNotification{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
