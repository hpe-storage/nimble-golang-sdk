/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// Subscription - Subscriptions represent the list of object types or alerts that a websocket client is interested in getting notifications for. Each subscription belongs to a single notification client.
// Export SubscriptionFields for advance operations like search filter etc.
var SubscriptionFields *Subscription

func init(){
	IDfield:= "id"
	SubscriberIDfield:= "subscriber_id"
	ObjectIDfield:= "object_id"
		
	SubscriptionFields= &Subscription{
		ID: &IDfield,
		SubscriberID: &SubscriberIDfield,
		ObjectID: &ObjectIDfield,
		
	}
}

type Subscription struct {
   
    // Identifier for subscription.
    
 	ID *string `json:"id,omitempty"`
   
    // Identifier for subscriber (notification client) that this subscription belongs to.
    
 	SubscriberID *string `json:"subscriber_id,omitempty"`
   
    // This indicates the type of notification being subscribed for.
    
   	NotificationType *NsNotificationType `json:"notification_type,omitempty"`
   
    // The object type that the notification subscriber is interested in. This is relevant for and required only for audit log based notifications.
    
   	ObjectType *NsObjectType `json:"object_type,omitempty"`
   
    // The object that the notification subscriber is interested in. Applies only to audit log based notifications.
    
 	ObjectID *string `json:"object_id,omitempty"`
   
    // The operation that the notification subscriber is interested in. Applies only to audit log based notifications.
    
   	Operation *NsAuditOperationTypeOrAll `json:"operation,omitempty"`
   
    // The kind of events or alerts that the notification subscriber is interested in. Applies only to events based notifications.
    
   	EventTargetType *NsEventTargetTypeOrAll `json:"event_target_type,omitempty"`
   
    // The severity of events that the notification subscriber is interested in. Applies only to events based notifications.
    
   	EventSeverity *NsSeverityLevel `json:"event_severity,omitempty"`
}
