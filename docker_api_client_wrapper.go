package main

import (
	"context"
	"errors"
	"io"
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

type envWrapper struct {
}

func (w *envWrapper) getContext() context.Context {
	return context.Background()
}

func (w *envWrapper) getClient() *client.Client {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	return cli
}

func (w *envWrapper) Pull(imagePath string) error {
	out, err := w.getClient().ImagePull(w.getContext(), imagePath, types.ImagePullOptions{})
	if out == nil {
		return errors.New("Output of pulling image is nil, Please check if image path is valid")
	}
	// Output must be handled, do nothing
	io.Copy(ioutil.Discard, out)
	return err
}

func (w *envWrapper) Run(option ContainerStartOption) (ContainerID, error) {
	ctx := w.getContext()
	cli := w.getClient()
	portMap, err := w.newPortMap(option.ports)
	if err != nil {
		log.Println("Error occurred while create port binding")
		return "", err
	}
	con, err := cli.ContainerCreate(ctx, &container.Config{
		Image: option.image,
	}, &container.HostConfig{
		PortBindings: portMap,
	}, nil, option.name)
	if err != nil {
		return "", err
	}

	cli.ContainerStart(ctx, con.ID, types.ContainerStartOptions{})
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
	return w.getClient().ContainerStop(w.getContext(), id, nil)
}

func (w *envWrapper) Remove(id string) error {
	return w.getClient().ContainerRemove(w.getContext(), id, types.ContainerRemoveOptions{})
}

// NewEnvClientWrapper constructs docker client object based environment variables
func NewEnvClientWrapper() DockerClientWrapper {
	return new(envWrapper)
}
