package main

import (
	"context"
	"io"
	"log"
	"os"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func TestListContainer(t *testing.T) {
	ctx := context.Background()
	cli, _ := client.NewEnvClient()
	containers, _ := cli.ContainerList(ctx, types.ContainerListOptions{})
	for _, container := range containers {
		log.Printf("%s\t%s", container.ID, container.Names[0])
	}
}

func TestPullImage(t *testing.T) {
	ctx := context.Background()
	cli, _ := client.NewEnvClient()

	out, err := cli.ImagePull(ctx, "docker.io/library/nginx:1.17", types.ImagePullOptions{})
	if err != nil {
		t.Error(err.Error())
		return
	}
	io.Copy(os.Stdout, out)
}

func TestRunContainer(t *testing.T) {
	ctx := context.Background()
	cli, _ := client.NewEnvClient()

	containerPort, _ := nat.NewPort("tcp", "80")
	hostBinding := nat.PortBinding{HostIP: "0.0.0.0", HostPort: "80"}
	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}
	container, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "nginx:1.17",
	}, &container.HostConfig{
		PortBindings: portBinding,
	}, nil, "nginx-test")
	if err != nil {
		t.Error(err.Error())
		return
	}
	log.Printf("Container created: %s\n", container.ID)
	cli.ContainerStart(ctx, container.ID, types.ContainerStartOptions{})
}
