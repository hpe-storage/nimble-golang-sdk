// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// Subscription - Subscriptions represent the list of object types or alerts that a websocket client is interested in getting notifications for. Each subscription belongs to a single notification client.
// Export SubscriptionFields for advance operations like search filter etc.
var SubscriptionFields *Subscription

func init() {

	SubscriptionFields = &Subscription{
		ID:           "id",
		SubscriberId: "subscriber_id",
		ObjectId:     "object_id",
	}
}

type Subscription struct {
	// ID - Identifier for subscription.
	ID string `json:"id,omitempty"`
	// SubscriberId - Identifier for subscriber (notification client) that this subscription belongs to.
	SubscriberId string `json:"subscriber_id,omitempty"`
	// NotificationType - This indicates the type of notification being subscribed for.
	NotificationType *NsNotificationType `json:"notification_type,omitempty"`
	// ObjectType - The object type that the notification subscriber is interested in. This is relevant for and required only for audit log based notifications.
	ObjectType *NsObjectType `json:"object_type,omitempty"`
	// ObjectId - The object that the notification subscriber is interested in. Applies only to audit log based notifications.
	ObjectId string `json:"object_id,omitempty"`
	// Operation - The operation that the notification subscriber is interested in. Applies only to audit log based notifications.
	Operation *NsAuditOperationTypeOrAll `json:"operation,omitempty"`
	// EventTargetType - The kind of events or alerts that the notification subscriber is interested in. Applies only to events based notifications.
	EventTargetType *NsEventTargetTypeOrAll `json:"event_target_type,omitempty"`
	// EventSeverity - The severity of events that the notification subscriber is interested in. Applies only to events based notifications.
	EventSeverity *NsSeverityLevel `json:"event_severity,omitempty"`
}

// sdk internal struct
type subscription struct {
	ID               *string                    `json:"id,omitempty"`
	SubscriberId     *string                    `json:"subscriber_id,omitempty"`
	NotificationType *NsNotificationType        `json:"notification_type,omitempty"`
	ObjectType       *NsObjectType              `json:"object_type,omitempty"`
	ObjectId         *string                    `json:"object_id,omitempty"`
	Operation        *NsAuditOperationTypeOrAll `json:"operation,omitempty"`
	EventTargetType  *NsEventTargetTypeOrAll    `json:"event_target_type,omitempty"`
	EventSeverity    *NsSeverityLevel           `json:"event_severity,omitempty"`
}

// EncodeSubscription - Transform Subscription to subscription type
func EncodeSubscription(request interface{}) (*subscription, error) {
	reqSubscription := request.(*Subscription)
	byte, err := json.Marshal(reqSubscription)
	if err != nil {
		return nil, err
	}
	respSubscriptionPtr := &subscription{}
	err = json.Unmarshal(byte, respSubscriptionPtr)
	return respSubscriptionPtr, err
}

// DecodeSubscription Transform subscription to Subscription type
func DecodeSubscription(request interface{}) (*Subscription, error) {
	reqSubscription := request.(*subscription)
	byte, err := json.Marshal(reqSubscription)
	if err != nil {
		return nil, err
	}
	respSubscription := &Subscription{}
	err = json.Unmarshal(byte, respSubscription)
	return respSubscription, err
}
