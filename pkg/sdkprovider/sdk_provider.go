package sdkprovider

import (
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

// VolumeService defines the interface to Volume related operations
type VolumeService interface {
	GetVolumes(params *param.GetParams) ([]*nimbleos.Volume, error)
	CreateVolume(obj *nimbleos.Volume) (*nimbleos.Volume, error)
	UpdateVolume(id string, obj *nimbleos.Volume) (*nimbleos.Volume, error)
	GetVolumeById(id string) (*nimbleos.Volume, error)
	GetVolumeByName(name string) (*nimbleos.Volume, error)
	GetVolumeBySerialNumber(serialNumber string) (*nimbleos.Volume, error)
	OnlineVolume(id string, force bool) (*nimbleos.Volume, error)
	OfflineVolume(id string, force bool) (*nimbleos.Volume, error)
	DeleteVolume(id string) error
	AssociateVolume(volId string, volcollId string) error
	DisassociateVolume(volId string) error
	RestoreVolume(id string, baseSnapId string) error
	MoveVolume(id string, destPoolId string, forceVvol *bool) (*nimbleos.NsVolumeListReturn, error)
	BulkMoveVolumes(volIds []*string, destPoolId string, forceVvol *bool) (*nimbleos.NsVolumeListReturn, error)
	AbortMoveVolume(id string) error
	BulkSetDedupeVolumes(volIds []*string, dedupeEnabled bool) error
	BulkSetOnlineAndOfflineVolumes(volIds []*string, online bool) error
	DestroyVolume(id string) error
}

// PoolService defines the interface to Pool related operations
type PoolService interface {
	GetPools(params *param.GetParams) ([]*nimbleos.Pool, error)
	UpdatePool(id string, obj *nimbleos.Pool) (*nimbleos.Pool, error)
	GetPoolById(id string) (*nimbleos.Pool, error)
	GetPoolByName(name string) (*nimbleos.Pool, error)
	DeletePool(id string) error
	MergePool(id string, targetPoolId string, force *bool) (*nimbleos.NsPoolMergeReturn, error)
}

// NsGroupService defines the interface to any sdk related operations
type NsGroupService interface {
	GetVolumeService() (vs VolumeService)
	GetPoolService() (vs PoolService)
	LogoutService() error
	SetDebug()
}
