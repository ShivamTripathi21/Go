package dockermanager

import (
	"github.com/docker/docker/client"
)

var DCluster *DockerCluster

type DockerCluster struct {
	dockerClusterUP      bool
	dockerClientAvilable bool
	dockerClient         *client.Client
}

type dockermanager interface {
	GetMyCluster() DockerCluster
	IsDockerClusterUP() bool
	NewDockerCluster() bool
	CreateDockerClient() (*client.Client, error)
}

func (param *DockerCluster) GetMyCluster() *DockerCluster {
	return DCluster
}
func (param *DockerCluster) IsDockerClusterUP() bool {
	if param == nil {
		return false
	} else {
		return DCluster.dockerClusterUP
	}
}
func (param *DockerCluster) NewDockerCluster() bool {
	if DCluster == nil {
		//prepare param with correct configuration
		param.CreateDockerClient()
		//assign param to dd
		DCluster = param
		return true
	} else {
		//dd is already there
		return false
	}
}
func (param *DockerCluster) CreateDockerClient() (bool, error) {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv)

	param.dockerClient = dockerClient

	if err != nil {
		param.dockerClientAvilable = false
		param.dockerClusterUP = false
		return false, err
	} else {
		param.dockerClientAvilable = true
		param.dockerClusterUP = true
		return true, nil
	}
}
