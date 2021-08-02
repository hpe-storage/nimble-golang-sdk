// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAlertNotificationFields provides field names to use in filter parameters, for example.
var NsAlertNotificationFields *NsAlertNotificationFieldHandles

func init() {
	NsAlertNotificationFields = &NsAlertNotificationFieldHandles{
		SequenceNumber:   "sequence_number",
		NotificationType: "notification_type",
		Activity:         "activity",
		EventTarget:      "event_target",
		Category:         "category",
		Severity:         "severity",
		AlertType:        "alert_type",
		Timestamp:        "timestamp",
	}
}

// NsAlertNotification - Represents an Alert notification message.
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

// NsAlertNotificationFieldHandles provides a string representation for each AccessControlRecord field.
type NsAlertNotificationFieldHandles struct {
	SequenceNumber   string
	NotificationType string
	Activity         string
	EventTarget      string
	Category         string
	Severity         string
	AlertType        string
	Timestamp        string
}
