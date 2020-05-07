/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/MasterKey


// MasterKey :
type MasterKey struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // Passphrase
   Passphrase string `json:"passphrase,omitempty"`
   // NewPassphrase
   NewPassphrase string `json:"new_passphrase,omitempty"`
   // Active
   Active bool `json:"active,omitempty"`
}
