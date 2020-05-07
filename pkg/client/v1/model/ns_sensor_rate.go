/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsSensorRate


// NsSensorRate :
type NsSensorRate struct {
   // Name
   Name string `json:"name,omitempty"`
   // Rate
   Rate float32 `json:"rate,omitempty"`
}
