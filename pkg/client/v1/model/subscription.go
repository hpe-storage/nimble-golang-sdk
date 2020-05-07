/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/Subscription


// Subscription :
type Subscription struct {
   // ID
   ID string `json:"id,omitempty"`
   // SubscriberID
   SubscriberID string `json:"subscriber_id,omitempty"`
   // ObjectID
   ObjectID string `json:"object_id,omitempty"`
}
