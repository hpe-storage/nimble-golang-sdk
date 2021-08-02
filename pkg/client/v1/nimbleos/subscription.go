// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// SubscriptionFields provides field names to use in filter parameters, for example.
var SubscriptionFields *SubscriptionFieldHandles

func init() {
	SubscriptionFields = &SubscriptionFieldHandles{
		ID:               "id",
		SubscriberId:     "subscriber_id",
		NotificationType: "notification_type",
		ObjectType:       "object_type",
		ObjectId:         "object_id",
		Operation:        "operation",
		EventTargetType:  "event_target_type",
		EventSeverity:    "event_severity",
	}
}

// Subscription - Subscriptions represent the list of object types or alerts that a websocket client is interested in getting notifications for. Each subscription belongs to a single notification client.
type Subscription struct {
	// ID - Identifier for subscription.
	ID *string `json:"id,omitempty"`
	// SubscriberId - Identifier for subscriber (notification client) that this subscription belongs to.
	SubscriberId *string `json:"subscriber_id,omitempty"`
	// NotificationType - This indicates the type of notification being subscribed for.
	NotificationType *NsNotificationType `json:"notification_type,omitempty"`
	// ObjectType - The object type that the notification subscriber is interested in. This is relevant for and required only for audit log based notifications.
	ObjectType *NsObjectType `json:"object_type,omitempty"`
	// ObjectId - The object that the notification subscriber is interested in. Applies only to audit log based notifications.
	ObjectId *string `json:"object_id,omitempty"`
	// Operation - The operation that the notification subscriber is interested in. Applies only to audit log based notifications.
	Operation *NsAuditOperationTypeOrAll `json:"operation,omitempty"`
	// EventTargetType - The kind of events or alerts that the notification subscriber is interested in. Applies only to events based notifications.
	EventTargetType *NsEventTargetTypeOrAll `json:"event_target_type,omitempty"`
	// EventSeverity - The severity of events that the notification subscriber is interested in. Applies only to events based notifications.
	EventSeverity *NsSeverityLevel `json:"event_severity,omitempty"`
}

// SubscriptionFieldHandles provides a string representation for each Subscription field.
type SubscriptionFieldHandles struct {
	ID               string
	SubscriberId     string
	NotificationType string
	ObjectType       string
	ObjectId         string
	Operation        string
	EventTargetType  string
	EventSeverity    string
}
