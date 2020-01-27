package main

import (
	"context"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

// ContainerID represents container's id
type ContainerID = string

// ContainerPort represents container's port to bind with host
type ContainerPort = int

// HostPort represents host's port to bind with container
type HostPort = int

type DockerClientWrapper interface {
	Pull(imagePath string) error
	Run(option ContainerStartOption) (ContainerID, error)
	Stop(id string) error
	Remove(id string) error
}

type ContainerStartOption struct {
	name      string // Name of container
	imagePath string // Path to pull image
	image     string // Name with tag of image
	ports     map[ContainerPort]HostPort
}

type envWrapper struct {
	ctx    context.Context
	client *client.Client
}

func (w *envWrapper) Pull(imagePath string) error {
	out, err := w.client.ImagePull(w.ctx, imagePath, types.ImagePullOptions{})
	// Output must be handled
	ioutil.ReadAll(out)
	return err
}

func (w *envWrapper) Run(option ContainerStartOption) (ContainerID, error) {
	portMap, err := w.newPortMap(option.ports)
	if err != nil {
		log.Println("Error occurred while create port binding")
		return "", err
	}
	con, err := w.client.ContainerCreate(w.ctx, &container.Config{
		Image: option.image,
	}, &container.HostConfig{
		PortBindings: portMap,
	}, nil, option.name)
	if err != nil {
		return "", err
	}

	w.client.ContainerStart(w.ctx, con.ID, types.ContainerStartOptions{})
	return con.ID, nil
}

func (w *envWrapper) newPortMap(ports map[ContainerPort]HostPort) (nat.PortMap, error) {
	portBindings := nat.PortMap{}
	for containerPort, hostPort := range ports {
		containerBinding, err := nat.NewPort("tcp", strconv.Itoa(containerPort))
		if err != nil {
			return nil, err
		}
		hostBinding := nat.PortBinding{HostIP: "0.0.0.0", HostPort: strconv.Itoa(hostPort)}
		portBindings[containerBinding] = []nat.PortBinding{hostBinding}
	}
	return portBindings, nil
}

func (w *envWrapper) Stop(id string) error {
	return w.client.ContainerStop(w.ctx, id, nil)
}

func (w *envWrapper) Remove(id string) error {
	return w.client.ContainerRemove(w.ctx, id, types.ContainerRemoveOptions{})
}

// NewEnvClientWrapper constructs docker client object based environment variables
func NewEnvClientWrapper() DockerClientWrapper {
	ctx := context.Background()
	cli, _ := client.NewEnvClient()
	wrapper := new(envWrapper)
	wrapper.client = cli
	wrapper.ctx = ctx

	return wrapper
}
