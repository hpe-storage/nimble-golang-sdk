// Copyright 2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAlertNotification - Represents an Alert notification message.
// Export NsAlertNotificationFields for advance operations like search filter etc.
var NsAlertNotificationFields *NsAlertNotification

func init() {
	Activityfield := "activity"

	NsAlertNotificationFields = &NsAlertNotification{
		Activity: &Activityfield,
	}
}

type NsAlertNotification struct {
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
