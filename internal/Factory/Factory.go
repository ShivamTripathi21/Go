package factory

import (
	"runtime"

	"github.com/ShivamTripathi21/Go/internal/model"
)

// FactoryConfig config information of run and build
type FactoryConfig struct {
	pieInfo       *model.Pie
	containerInfo []*model.ContainerInfo
}

// scdFunForOSInfo Information about os
func (FConfig *FactoryConfig) scdFunForClusterType() {
	os := &model.ClusterType{}

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
