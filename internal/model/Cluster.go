package model

import (
	"time"

	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/docker/docker/api/types"
	"github.com/google/uuid"
)

// OSInfo Operating system info
type ClusterType struct {
	OsName string
}

// DockerInfo user's docker information if it is there.
type DockerInfo struct {
	Version        string
	ComposeVersion string
}

type ClusterInfo struct {
	ClusterID                     uuid.UUID
	ClusterName                   string
	ClusterType                   *ClusterType
	ClusterROMSize                int
	ClusterRAMSize                int
	ClusterGPUSize                int
	ClusterSwapMemorySize         int
	IsGPUAvilable                 bool
	ClusterAvilableROMSize        int
	ClusterAvilableRAMSize        int
	ClusterAvilableGPUSize        int
	ClusterAvilableSwapMemorySize int
	TotalNumberOfPies             int
	ClusterCreatedByID            guid.GUID
	ClusterCreateTime             time.Time
	ClusterResourceProvider       string
	ClusterResourceURL            string
	IsLocalCluster                bool
	IsClusterUP                   bool
	IsClusterInUse                bool
	IsBackupClusterEnabled        bool
	BackupClusterID               guid.GUID
}

type Pie struct {
	PieID                    uuid.UUID
	PieOwnerID               uuid.UUID
	ClusterInoframtion       *ClusterInfo //pointer type because we can chage few Cluster info from the Pie, like Sizes.
	PieMaxROMSize            int
	PieMaxRAMSize            int
	PieMaxGPUSize            int
	PieSwapMemorySize        int
	IsGPUAvilable            bool
	IsPieUP                  bool
	IsPieInUse               bool
	PieROMPercentage         int
	PieRAMPercentage         int
	PieGPUPercentage         int
	PieClusterPercentage     int
	MaxNumberOfContainers    int
	TotalNumberOfContainer   int
	TotalNumberOfUPContainer int
	PieCreatedByID           uuid.UUID
	PieCreatedTime           time.Time
}

type ContainerInfo struct {
	ContainerID          uuid.UUID
	ContainerOwnerIDs    []int32
	DockerInfo           *DockerInfo
	PieID                uuid.UUID
	DockerImageInfo      *DockerImageInfo
	ContainerUrl         string
	ContainerServicePort int
	RAMUsagePercentage   int
	GPUUsagePercentage   int
	ROMUsagePercentage   int
	NetworkInuputBytes   int
	NetworkOutputBytes   int
	DiskReadBytes        int
	DiskWriteBytes       int
}

type DockerImageInfo struct {
	DockerImageID    string
	ImageBuildOption types.ImageBuildOptions
}
