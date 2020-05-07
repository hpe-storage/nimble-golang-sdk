/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/Disk


// Disk :
type Disk struct {
   // ID
   ID string `json:"id,omitempty"`
   // IsDfc
   IsDfc bool `json:"is_dfc,omitempty"`
   // Serial
   Serial string `json:"serial,omitempty"`
   // Path
   Path string `json:"path,omitempty"`
   // ShelfSerial
   ShelfSerial string `json:"shelf_serial,omitempty"`
   // ShelfLocation
   ShelfLocation string `json:"shelf_location,omitempty"`
   // ShelfID
   ShelfID string `json:"shelf_id,omitempty"`
   // ShelfLocationID
   ShelfLocationID float64 `json:"shelf_location_id,omitempty"`
   // VshelfID
   VshelfID float64 `json:"vshelf_id,omitempty"`
   // Slot
   Slot float64 `json:"slot,omitempty"`
   // Bank
   Bank float64 `json:"bank,omitempty"`
   // Model
   Model string `json:"model,omitempty"`
   // Vendor
   Vendor string `json:"vendor,omitempty"`
   // FirmwareVersion
   FirmwareVersion string `json:"firmware_version,omitempty"`
   // Hba
   Hba float64 `json:"hba,omitempty"`
   // Port
   Port float64 `json:"port,omitempty"`
   // Size
   Size float64 `json:"size,omitempty"`
   // RaIDID
   RaIDID  int32 `json:"raid_id,omitempty"`
   // RaIDResyncPercent
   RaIDResyncPercent float64 `json:"raid_resync_percent,omitempty"`
   // RaIDResyncCurrentSpeed
   RaIDResyncCurrentSpeed float64 `json:"raid_resync_current_speed,omitempty"`
   // RaIDResyncAverageSpeed
   RaIDResyncAverageSpeed float64 `json:"raid_resync_average_speed,omitempty"`
   // DiskInternalStat1
   DiskInternalStat1 string `json:"disk_internal_stat1,omitempty"`
   // Force
   Force bool `json:"force,omitempty"`
   // ArrayName
   ArrayName string `json:"array_name,omitempty"`
   // ArrayID
   ArrayID string `json:"array_id,omitempty"`
   // PartialResponseOk
   PartialResponseOk bool `json:"partial_response_ok,omitempty"`
}
