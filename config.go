package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	}
	Spec struct {
		ImageName     string `yaml:"image"`
		ImagePath     string `yaml:"imagePath"`
		ContainerName string `yaml:"containerName"`
		InitialTag    string `yaml:"initialTag"`
		Ports         []struct {
			ContainerPort int `yaml:"containerPort"`
			HostPort      int `yaml:"hostPort"`
		} `yamle:"ports"`
	} `yaml:"spec"`
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
