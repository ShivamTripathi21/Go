package model

import (
	"runtime"
	"time"

	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/docker/docker/api/types"
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
	ClusterID                     guid.GUID
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
	PieID                    guid.GUID
	PieOwnerID               guid.GUID
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
	PieCreatedByID           guid.GUID
	PieCreatedTime           time.Time
}

type ContainerInfo struct {
	ContainerID          guid.GUID
	ContainerOwnerIDs    []int32
	DockerInfo           *DockerInfo
	PieID                guid.GUID
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

// FactoryConfig config information of run and build
type FactoryConfig struct {
	pieInfo       *Pie
	containerInfo []*ContainerInfo
}

// scdFunForOSInfo Information about os
func (FConfig *FactoryConfig) scdFunForClusterType() {
	os := &ClusterType{}

	gos := runtime.GOOS
	switch gos {
	case "windows":
		os.OsName = "windows"
	case "darwin":
		os.OsName = "darwin"
	case "linux":
		os.OsName = "linux"
	default:
		os.OsName = "nil"
	}
	FConfig.pieInfo.ClusterInoframtion.ClusterType = os
}

// scdFunForDockerInfo Information about docker
// func (FConfig *FactoryConfig) scdFunForDockerInfo() {
// 	di := &DockerInfo{}

// 	cmdd := exec.Command("docker", "--version")
// 	cmddc := exec.Command("docker-compose", "--version")

// 	cmddOutput := &bytes.Buffer{}
// 	cmddcOutput := &bytes.Buffer{}

// 	cmdd.Stdout = cmddOutput
// 	cmddc.Stdout = cmddcOutput

// 	errd := cmdd.Run()
// 	//for docker
// 	if errd != nil {
// 		di.Version = "nil"
// 	} else {
// 		di.Version = string(cmddOutput.Bytes())
// 	}

// 	errdc := cmddc.Run()
// 	//for docker-compose
// 	if errdc != nil {
// 		di.ComposeVersion = "nil"
// 	} else {
// 		di.ComposeVersion = string(cmddcOutput.Bytes())
// 	}

// 	FConfig.containerInfo.DockerInfo = di
// }

// NewFactory intiate the factory for build or run
func NewFactory() (*FactoryConfig, error) {
	fconf := &FactoryConfig{}
	return fconf, nil
}
