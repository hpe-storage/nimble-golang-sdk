// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// Array - Retrieve information of specified arrays. The array is the management and configuration for the underlying physical hardware array box.
// Export ArrayFields for advance operations like search filter etc.
var ArrayFields *Array

func init() {

	ArrayFields = &Array{
		ID:              "id",
		Name:            "name",
		FullName:        "full_name",
		SearchName:      "search_name",
		PoolName:        "pool_name",
		PoolId:          "pool_id",
		Model:           "model",
		Serial:          "serial",
		Version:         "version",
		ExtendedModel:   "extended_model",
		PoolDescription: "pool_description",
		CtrlrASupportIp: "ctrlr_a_support_ip",
		CtrlrBSupportIp: "ctrlr_b_support_ip",
		ModelSubType:    "model_sub_type",
		SecondaryMgmtIp: "secondary_mgmt_ip",
	}
}

type Array struct {
	// ID - Identifier for array.
	ID string `json:"id,omitempty"`
	// Name - The user provided name of the array. It is also the array's hostname.
	Name string `json:"name,omitempty"`
	// Force - Forcibly delete the specified array.
	Force *bool `json:"force,omitempty"`
	// FullName - The array's fully qualified name.
	FullName string `json:"full_name,omitempty"`
	// SearchName - The array name used for object search.
	SearchName string `json:"search_name,omitempty"`
	// Status - Reachability status of the array in the group.
	Status *NsArrayStatus `json:"status,omitempty"`
	// Role - Role of the array in the group.
	Role *NsArrayRole `json:"role,omitempty"`
	// GroupState - State of the array in the group.
	GroupState *NsArrayGroupState `json:"group_state,omitempty"`
	// PoolName - Name of pool to which this is a member.
	PoolName string `json:"pool_name,omitempty"`
	// PoolId - ID of pool to which this is a member.
	PoolId string `json:"pool_id,omitempty"`
	// Model - Array model.
	Model string `json:"model,omitempty"`
	// Serial - Serial number of the array.
	Serial string `json:"serial,omitempty"`
	// Version - Software version of the array.
	Version string `json:"version,omitempty"`
	// IsSfa - True if this array supports SFA; false otherwise.
	IsSfa *bool `json:"is_sfa,omitempty"`
	// CreationTime - Time when this array object was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this array object was last modified.
	LastModified int64 `json:"last_modified,omitempty"`
	// UsageValid - Indicates whether the usage of array is valid.
	UsageValid *bool `json:"usage_valid,omitempty"`
	// UsableCapacityBytes - The usable capacity of the array in bytes.
	UsableCapacityBytes int64 `json:"usable_capacity_bytes,omitempty"`
	// UsableCacheCapacityBytes - The usable cache capacity of the array in bytes.
	UsableCacheCapacityBytes int64 `json:"usable_cache_capacity_bytes,omitempty"`
	// RawCapacityBytes - The raw capacity of the array in bytes.
	RawCapacityBytes int64 `json:"raw_capacity_bytes,omitempty"`
	// VolUsageBytes - The compressed usage of volumes in array.
	VolUsageBytes int64 `json:"vol_usage_bytes,omitempty"`
	// VolUsageUncompressedBytes - The uncompressed usage of volumes in array. This is the pre-reduced usage.
	VolUsageUncompressedBytes int64 `json:"vol_usage_uncompressed_bytes,omitempty"`
	// VolCompression - The compression rate of volumes in array expressed as ratio.
	VolCompression float32 `json:"vol_compression,omitempty"`
	// VolSavedBytes - The saved space of volumes in array.
	VolSavedBytes int64 `json:"vol_saved_bytes,omitempty"`
	// SnapUsageBytes - The compressed usage of snapshots in array.
	SnapUsageBytes int64 `json:"snap_usage_bytes,omitempty"`
	// SnapUsageUncompressedBytes - The uncompressed usage of snapshots in array. This is the pre-reduced usage.
	SnapUsageUncompressedBytes int64 `json:"snap_usage_uncompressed_bytes,omitempty"`
	// SnapCompression - The compression rate of snapshots in array expressed as ratio.
	SnapCompression float32 `json:"snap_compression,omitempty"`
	// SnapSpaceReduction - The space reduction rate of snapshots in array expressed as ratio.
	SnapSpaceReduction float32 `json:"snap_space_reduction,omitempty"`
	// SnapSavedBytes - The saved space of snapshots in array.
	SnapSavedBytes int64 `json:"snap_saved_bytes,omitempty"`
	// PendingDeleteBytes - The pending delete bytes in array.
	PendingDeleteBytes int64 `json:"pending_delete_bytes,omitempty"`
	// AvailableBytes - The available space of array.
	AvailableBytes int64 `json:"available_bytes,omitempty"`
	// Usage - Used space of the array in bytes.
	Usage int64 `json:"usage,omitempty"`
	// AllFlash - Whether it is an all-flash array.
	AllFlash *bool `json:"all_flash,omitempty"`
	// DedupeCapacityBytes - The dedupe capacity of a hybrid array. Does not apply to all-flash arrays.
	DedupeCapacityBytes int64 `json:"dedupe_capacity_bytes,omitempty"`
	// DedupeUsageBytes - The dedupe usage of a hybrid array. Does not apply to all-flash arrays.
	DedupeUsageBytes int64 `json:"dedupe_usage_bytes,omitempty"`
	// IsFullyDedupeCapable - Is array fully capable to dedupe its usable capacity.
	IsFullyDedupeCapable *bool `json:"is_fully_dedupe_capable,omitempty"`
	// ExtendedModel - Extended model of the array.
	ExtendedModel string `json:"extended_model,omitempty"`
	// IsSupportedHwConfig - Whether it is a supported hardware config.
	IsSupportedHwConfig *bool `json:"is_supported_hw_config,omitempty"`
	// GigNicPortCount - Count of 1G NIC Ports installed on the array.
	GigNicPortCount int64 `json:"gig_nic_port_count,omitempty"`
	// TenGigSfpNicPortCount - Count of 10G SFP NIC Ports installed on the array.
	TenGigSfpNicPortCount int64 `json:"ten_gig_sfp_nic_port_count,omitempty"`
	// TenGigTNicPortCount - Count of 10G BaseT NIC Ports installed on the array.
	TenGigTNicPortCount int64 `json:"ten_gig_t_nic_port_count,omitempty"`
	// FcPortCount - Count of Fibre Channel Ports installed on the array.
	FcPortCount int64 `json:"fc_port_count,omitempty"`
	// PublicKey - Public key of the array.
	PublicKey *NsSshKey `json:"public_key,omitempty"`
	// Upgrade - The array upgrade data.
	Upgrade *NsArrayUpgrade `json:"upgrade,omitempty"`
	// CreatePool - Whether to create associated pool during array create.
	CreatePool *bool `json:"create_pool,omitempty"`
	// PoolDescription - Text description of the pool to be created during array creation.
	PoolDescription string `json:"pool_description,omitempty"`
	// AllowLowerLimits - A True setting will allow you to add an array with lower limits to a pool with higher limits.
	AllowLowerLimits *bool `json:"allow_lower_limits,omitempty"`
	// CtrlrASupportIp - Controller A Support IP Address.
	CtrlrASupportIp string `json:"ctrlr_a_support_ip,omitempty"`
	// CtrlrBSupportIp - Controller B Support IP Address.
	CtrlrBSupportIp string `json:"ctrlr_b_support_ip,omitempty"`
	// NicList - List NICs information. Used when creating an array.
	NicList []*NsNIC `json:"nic_list,omitempty"`
	// ModelSubType - Array model sub type.
	ModelSubType string `json:"model_sub_type,omitempty"`
	// ZconfIpaddrs - List of link-local zero-configuration addresses of the array.
	ZconfIpaddrs []*NsIPAddressObject `json:"zconf_ipaddrs,omitempty"`
	// SecondaryMgmtIp - Secondary management IP address for the Group.
	SecondaryMgmtIp string `json:"secondary_mgmt_ip,omitempty"`
}

// sdk internal struct
type array struct {
	ID                         *string              `json:"id,omitempty"`
	Name                       *string              `json:"name,omitempty"`
	Force                      *bool                `json:"force,omitempty"`
	FullName                   *string              `json:"full_name,omitempty"`
	SearchName                 *string              `json:"search_name,omitempty"`
	Status                     *NsArrayStatus       `json:"status,omitempty"`
	Role                       *NsArrayRole         `json:"role,omitempty"`
	GroupState                 *NsArrayGroupState   `json:"group_state,omitempty"`
	PoolName                   *string              `json:"pool_name,omitempty"`
	PoolId                     *string              `json:"pool_id,omitempty"`
	Model                      *string              `json:"model,omitempty"`
	Serial                     *string              `json:"serial,omitempty"`
	Version                    *string              `json:"version,omitempty"`
	IsSfa                      *bool                `json:"is_sfa,omitempty"`
	CreationTime               *int64               `json:"creation_time,omitempty"`
	LastModified               *int64               `json:"last_modified,omitempty"`
	UsageValid                 *bool                `json:"usage_valid,omitempty"`
	UsableCapacityBytes        *int64               `json:"usable_capacity_bytes,omitempty"`
	UsableCacheCapacityBytes   *int64               `json:"usable_cache_capacity_bytes,omitempty"`
	RawCapacityBytes           *int64               `json:"raw_capacity_bytes,omitempty"`
	VolUsageBytes              *int64               `json:"vol_usage_bytes,omitempty"`
	VolUsageUncompressedBytes  *int64               `json:"vol_usage_uncompressed_bytes,omitempty"`
	VolCompression             *float32             `json:"vol_compression,omitempty"`
	VolSavedBytes              *int64               `json:"vol_saved_bytes,omitempty"`
	SnapUsageBytes             *int64               `json:"snap_usage_bytes,omitempty"`
	SnapUsageUncompressedBytes *int64               `json:"snap_usage_uncompressed_bytes,omitempty"`
	SnapCompression            *float32             `json:"snap_compression,omitempty"`
	SnapSpaceReduction         *float32             `json:"snap_space_reduction,omitempty"`
	SnapSavedBytes             *int64               `json:"snap_saved_bytes,omitempty"`
	PendingDeleteBytes         *int64               `json:"pending_delete_bytes,omitempty"`
	AvailableBytes             *int64               `json:"available_bytes,omitempty"`
	Usage                      *int64               `json:"usage,omitempty"`
	AllFlash                   *bool                `json:"all_flash,omitempty"`
	DedupeCapacityBytes        *int64               `json:"dedupe_capacity_bytes,omitempty"`
	DedupeUsageBytes           *int64               `json:"dedupe_usage_bytes,omitempty"`
	IsFullyDedupeCapable       *bool                `json:"is_fully_dedupe_capable,omitempty"`
	ExtendedModel              *string              `json:"extended_model,omitempty"`
	IsSupportedHwConfig        *bool                `json:"is_supported_hw_config,omitempty"`
	GigNicPortCount            *int64               `json:"gig_nic_port_count,omitempty"`
	TenGigSfpNicPortCount      *int64               `json:"ten_gig_sfp_nic_port_count,omitempty"`
	TenGigTNicPortCount        *int64               `json:"ten_gig_t_nic_port_count,omitempty"`
	FcPortCount                *int64               `json:"fc_port_count,omitempty"`
	PublicKey                  *NsSshKey            `json:"public_key,omitempty"`
	Upgrade                    *NsArrayUpgrade      `json:"upgrade,omitempty"`
	CreatePool                 *bool                `json:"create_pool,omitempty"`
	PoolDescription            *string              `json:"pool_description,omitempty"`
	AllowLowerLimits           *bool                `json:"allow_lower_limits,omitempty"`
	CtrlrASupportIp            *string              `json:"ctrlr_a_support_ip,omitempty"`
	CtrlrBSupportIp            *string              `json:"ctrlr_b_support_ip,omitempty"`
	NicList                    []*NsNIC             `json:"nic_list,omitempty"`
	ModelSubType               *string              `json:"model_sub_type,omitempty"`
	ZconfIpaddrs               []*NsIPAddressObject `json:"zconf_ipaddrs,omitempty"`
	SecondaryMgmtIp            *string              `json:"secondary_mgmt_ip,omitempty"`
}

// EncodeArray - Transform Array to array type
func EncodeArray(request interface{}) (*array, error) {
	reqArray := request.(*Array)
	byte, err := json.Marshal(reqArray)
	if err != nil {
		return nil, err
	}
	respArrayPtr := &array{}
	err = json.Unmarshal(byte, respArrayPtr)
	return respArrayPtr, err
}

// DecodeArray Transform array to Array type
func DecodeArray(request interface{}) (*Array, error) {
	reqArray := request.(*array)
	byte, err := json.Marshal(reqArray)
	if err != nil {
		return nil, err
	}
	respArray := &Array{}
	err = json.Unmarshal(byte, respArray)
	return respArray, err
}
