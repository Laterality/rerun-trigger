package main

import "testing"

func Test_read_spec_from_file(t *testing.T) {
	expectedImageName := "nginx"
	expectedContainerName := "webserver"
	conf, _ := NewConfigFromFile("./config_example.yml")
	if conf.Spec.ImageName != expectedImageName {
		t.Errorf("Image name not matched. expected '%s' but actual was '%s'\n", expectedImageName, conf.Spec.ImageName)
	}
	if conf.Spec.ContainerName != expectedContainerName {
		t.Errorf("Container name not matched. expected '%s' but actual was '%s'\n", expectedContainerName, conf.Spec.ContainerName)
	}
}

func Test_read_server_from_file(t *testing.T) {
	expectedPort := 8000
	conf, _ := NewConfigFromFile("./config_example.yml")
	if conf.Server.Port != expectedPort {
		t.Errorf("Server port not matched. expected %d but actual was '%d'\n", expectedPort, conf.Server.Port)
	}
}
