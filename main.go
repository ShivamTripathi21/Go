package main

import (
	"context"
	"fmt"

	"github.com/ShivamTripathi21/Go/internal/dockermanager"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	//create cluster if there's no cluster
	dc := &dockermanager.DockerCluster{}

	dc.NewDockerCluster()
	println(dc.IsDockerClusterUP())

	// if dc.IsDockerClusterUP() {
	// 	println("Docker Cluster is UP")
	// } else {
	// 	println("Docker Cluster is NOT UP")
	// }

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	//configs, err := cli.Info(context.Background())
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s \n%s \n%s \n%s \n%s \n%s \n%s \n%s\n", configs, configs.ServerVersion, configs.DockerRootDir, configs.KernelVersion, configs.Architecture, strconv.Itoa(int(configs.MemTotal)), strconv.Itoa(int(configs.NCPU)), configs.OperatingSystem)
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
}
