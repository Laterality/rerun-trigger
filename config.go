package serviceConfig

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	}
	Spec struct {
		ImageName     string `yaml:"image"`
		ContainerName string `yaml:"containerName"`
	} `yaml:"spec"`
}

func NewConfigFromFile(path string) (*Config, error) {
	var err error = nil
	config := new(Config)
	bytes, err := ioutil.ReadFile(path)
	log.Println("read: ", string(bytes))
	if err != nil {
		return nil, err
	}
	yaml.Unmarshal(bytes, config)

	return config, err
}
