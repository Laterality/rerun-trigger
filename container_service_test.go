package main

import (
	"reflect"
	"testing"
)

func TestInitiallyRun(t *testing.T) {
	wrapper := NewMockDockerClientWrapper()
	service, err := NewContainerService(wrapper, "config_example.yml")
	if err != nil {
		t.Error("Error while construct container service: ", err.Error())
	}
	service.Update("1.16")

	if wrapper.PullCount == 0 {
		t.Error("Image is not pulled")
	}

	if wrapper.RunCount == 0 {
		t.Error("Container.Run() haven't called")
	}
}

func TestCreateStartOptionFromConfig(t *testing.T) {
	actualOption := NewContainerStartOptionFromConfig(Config{
		Spec: ContainerSpec{
			ImageName:     "nginx",
			ImagePath:     "docker.io/library/nginx",
			ContainerName: "webserver",
			Ports: []PortBinding{
				PortBinding{
					ContainerPort: 80,
					HostPort:      8080,
				},
			}},
	})

	expectedOption := ContainerStartOption{
		name:      "nginx",
		imagePath: "docker.io/library/nginx",
		image:     "nginx",
		ports: map[ContainerPort]HostPort{
			80: 8080,
		},
	}

	if !reflect.DeepEqual(expectedOption, actualOption) {
		t.Error("Option not matched, expected: ", expectedOption, ", but actual: ", actualOption)
	}
}
