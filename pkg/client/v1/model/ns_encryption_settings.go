/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsEncryptionSettings


// NsEncryptionSettings :
type NsEncryptionSettings struct {
   // MasterKeySet
   MasterKeySet bool `json:"master_key_set,omitempty"`
   // EncryptionActive
   EncryptionActive bool `json:"encryption_active,omitempty"`
}
