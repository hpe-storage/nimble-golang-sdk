/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// Disk - Disks are used for storing user data.
// Export DiskFields for advance operations like search filter etc.
var DiskFields *Disk

func init(){
	IDfield:= "id"
	Serialfield:= "serial"
	Pathfield:= "path"
	ShelfSerialfield:= "shelf_serial"
	ShelfLocationfield:= "shelf_location"
	ShelfIDfield:= "shelf_id"
	Modelfield:= "model"
	Vendorfield:= "vendor"
	FirmwareVersionfield:= "firmware_version"
	DiskInternalStat1field:= "disk_internal_stat1"
	ArrayNamefield:= "array_name"
	ArrayIDfield:= "array_id"
		
	DiskFields= &Disk{
		ID: &IDfield,
		Serial: &Serialfield,
		Path: &Pathfield,
		ShelfSerial: &ShelfSerialfield,
		ShelfLocation: &ShelfLocationfield,
		ShelfID: &ShelfIDfield,
		Model: &Modelfield,
		Vendor: &Vendorfield,
		FirmwareVersion: &FirmwareVersionfield,
		DiskInternalStat1: &DiskInternalStat1field,
		ArrayName: &ArrayNamefield,
		ArrayID: &ArrayIDfield,
		
	}
}

type Disk struct {
   
    // ID of disk.
    
 	ID *string `json:"id,omitempty"`
   
    // Is disk part of dual flash carrier.
    
 	IsDfc *bool `json:"is_dfc,omitempty"`
   
    // Disk serial number(N/A if empty).
    
 	Serial *string `json:"serial,omitempty"`
   
    // Disk SCSI device path.
    
 	Path *string `json:"path,omitempty"`
   
    // Serial number of the shelf the disk is attached to.
    
 	ShelfSerial *string `json:"shelf_serial,omitempty"`
   
    // Identifies the controller, port, and chain position of the shelf the disk belongs to.
    
 	ShelfLocation *string `json:"shelf_location,omitempty"`
   
    // Identifies the physical shelf the disk belongs to.
    
 	ShelfID *string `json:"shelf_id,omitempty"`
   
    // Identifies the position shelf the disk belongs to, as coded integer.
    
   	ShelfLocationID *int64 `json:"shelf_location_id,omitempty"`
   
    // Identifies the local shelf ID the disk belongs to.
    
   	VshelfID *int64 `json:"vshelf_id,omitempty"`
   
    // Disk slot number.
    
   	Slot *int64 `json:"slot,omitempty"`
   
    // Disk bank number.
    
   	Bank *int64 `json:"bank,omitempty"`
   
    // Disk model name.
    
 	Model *string `json:"model,omitempty"`
   
    // Vendor name of the disk manufacturer.
    
 	Vendor *string `json:"vendor,omitempty"`
   
    // Firmware version on the disk.
    
 	FirmwareVersion *string `json:"firmware_version,omitempty"`
   
    // HBA ID the disk is connected to.
    
   	Hba *int64 `json:"hba,omitempty"`
   
    // HBA port number the disk is connected to.
    
   	Port *int64 `json:"port,omitempty"`
   
    // Disk size in bytes.
    
   	Size *int64 `json:"size,omitempty"`
   
    // Disk hardware state.
    
   	State *NsDiskState `json:"state,omitempty"`
   
    // Type of disk (HDD, SSD, N/A).
    
   	Type *NsDiskType `json:"type,omitempty"`
   
    // Native block type of the disk.
    
   	BlockType *NsDiskBlockType `json:"block_type,omitempty"`
   
    // Raid ID.
    
  	RaIDID  *int64 `json:"raid_id,omitempty"`
   
    // Percentage RAID rebuild completed on this disk.
    
   	RaIDResyncPercent *int64 `json:"raid_resync_percent,omitempty"`
   
    // Current RAID rebuild speed (bytes/sec).
    
   	RaIDResyncCurrentSpeed *int64 `json:"raid_resync_current_speed,omitempty"`
   
    // Average RAID rebuild speed (bytes/sec).
    
   	RaIDResyncAverageSpeed *int64 `json:"raid_resync_average_speed,omitempty"`
   
    // RAID status for the disk (N/A, okay, resynchronizing, spare, faulty).
    
   	RaIDState *NsDiskRaIDState `json:"raid_state,omitempty"`
   
    // Internal disk statistic 1.
    
 	DiskInternalStat1 *string `json:"disk_internal_stat1,omitempty"`
   
    // S.M.A.R.T. attributes for the disk.
    
   	SmartAttributeList []*NsDiskSmartAttribute `json:"smart_attribute_list,omitempty"`
   
    // The intended operation to be performed on the specified disk.
    
   	DiskOp *NsDiskOp `json:"disk_op,omitempty"`
   
    // Forcibly add a disk.
    
 	Force *bool `json:"force,omitempty"`
   
    // Name of array the disk belongs to.
    
 	ArrayName *string `json:"array_name,omitempty"`
   
    // ID of array the disk belongs to.
    
 	ArrayID *string `json:"array_id,omitempty"`
   
    // Indicate that it is okay to provide partially available response.
    
 	PartialResponseOk *bool `json:"partial_response_ok,omitempty"`
}
