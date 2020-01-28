package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server ServerConfig  `yaml: "server"`
	Spec   ContainerSpec `yaml:"spec"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type ContainerSpec struct {
	ImageName     string        `yaml:"image"`
	ImagePath     string        `yaml:"imagePath"`
	ContainerName string        `yaml:"containerName"`
	InitialTag    string        `yaml:"initialTag"`
	Ports         []PortBinding `yaml:"ports"`
}

type PortBinding struct {
	ContainerPort int `yaml:"containerPort"`
	HostPort      int `yaml:"hostPort"`
}

func NewConfigFromFile(path string) (*Config, error) {
	var err error = nil
	config := new(Config)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	yaml.Unmarshal(bytes, config)

	return config, err
}
