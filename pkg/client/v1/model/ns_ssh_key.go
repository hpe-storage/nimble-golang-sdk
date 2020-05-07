/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsSshKey


// NsSshKey :
type NsSshKey struct {
   // KeyName
   KeyName string `json:"key_name,omitempty"`
   // KeyType
   KeyType string `json:"key_type,omitempty"`
   // Key
   Key string `json:"key,omitempty"`
}
