package main

import "log"

type ContainerStartOption struct {
	name      string // Name of container
	imagePath string // Path to pull image
	image     string // Name with tag of image
	ports     map[ContainerPort]HostPort
}

type Container interface {
	Up(option ContainerStartOption) error
	Down() error
	ID() string
}

type containerObject struct {
	clientWraper DockerClientWrapper
	id           string
	option       ContainerStartOption
}

func (c *containerObject) Up(option ContainerStartOption) error {
	log.Println("Pull " + option.imagePath)
	err := c.clientWraper.Pull(option.imagePath)
	if err != nil {
		return err
	}
	id, err := c.clientWraper.Run(option)
	if err != nil {
		return err
	}
	c.id = id
	return nil
}

func (c *containerObject) Down() error {
	err := c.clientWraper.Stop(c.id)
	if err != nil {
		return err
	}

	return c.clientWraper.Remove(c.id)
}

func (c *containerObject) ID() string {
	return c.id
}

func NewContainer(client DockerClientWrapper) Container {
	container := new(containerObject)
	container.clientWraper = client
	return container
}
