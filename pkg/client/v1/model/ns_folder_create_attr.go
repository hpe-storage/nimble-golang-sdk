/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsFolderCreateAttr


// NsFolderCreateAttr :
type NsFolderCreateAttr struct {
   // Name
   Name string `json:"name,omitempty"`
   // PoolID
   PoolID string `json:"pool_id,omitempty"`
}
