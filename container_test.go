package main

import (
	"testing"
)

func TestCreate(t *testing.T) {
	container := NewContainer(NewMockDockerClientWrapper())
	if container == nil {
		t.Error("Created container is nil")
	}
}

func TestUp(t *testing.T) {
	clientWrapper := NewMockDockerClientWrapper()
	cont := NewContainer(clientWrapper)
	cont.Up(ContainerStartOption{
		name:      "webserver",
		imagePath: "docker.io/library/nginx:1.17",
		image:     "nginx:1.17",
		ports: map[ContainerPort]HostPort{
			80: 8080,
		},
	})

	if clientWrapper.PullCount == 0 {
		t.Error("Container haven't pulled image")
	}
}

func TestDown(t *testing.T) {
	clientWrapper := NewMockDockerClientWrapper()
	cont := NewContainer(clientWrapper)
	cont.Up(ContainerStartOption{
		name:      "webserver",
		imagePath: "docker.io/library/nginx:1.17",
		image:     "nginx:1.17",
		ports: map[ContainerPort]HostPort{
			80: 8080,
		},
	})
	cont.Down()

	if clientWrapper.RunCount == 0 || clientWrapper.StopCount == 0 || clientWrapper.RemoveCount == 0 {
		t.Error("Container isn't down correctly, Run: ", clientWrapper.RunCount, ", Stop: ", clientWrapper.StopCount, ", Remove: ", clientWrapper.RemoveCount)
	}
}

type MockDockerClientWrapper struct {
	PullCount   int
	RunCount    int
	StopCount   int
	RemoveCount int
}

func (w *MockDockerClientWrapper) Pull(imagePath string) error {
	w.PullCount++
	return nil
}

func (w *MockDockerClientWrapper) Run(option ContainerStartOption) (ContainerID, error) {
	w.RunCount++
	return "", nil
}

func (w *MockDockerClientWrapper) Stop(id string) error {
	w.StopCount++
	return nil
}

func (w *MockDockerClientWrapper) Remove(id string) error {
	w.RemoveCount++
	return nil
}

func NewMockDockerClientWrapper() *MockDockerClientWrapper {
	return new(MockDockerClientWrapper)
}
