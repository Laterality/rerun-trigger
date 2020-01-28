package main

type ContainerService interface {
	Update(tag string) (string, error)
}

type ContainerServiceImpl struct {
	clientWrapper    DockerClientWrapper
	config           *Config
	currentContainer Container
}

func (s *ContainerServiceImpl) Update(tag string) (string, error) {
	if s.currentContainer != nil {
		s.currentContainer.Down()
	}
	s.currentContainer = NewContainer(s.clientWrapper)
	err := s.currentContainer.Up(NewContainerStartOptionFromConfig(*s.config))
	return s.currentContainer.ID(), err
}

// NewContainerService constructs new container service with config at specified path, client wrapper factory
func NewContainerService(clientWrapper DockerClientWrapper, configPath string) (ContainerService, error) {
	service := new(ContainerServiceImpl)
	config, err := NewConfigFromFile(configPath)
	if err != nil {
		return nil, err
	}
	service.config = config
	service.clientWrapper = clientWrapper
	return service, nil
}

func NewContainerStartOptionFromConfig(config Config) ContainerStartOption {
	return ContainerStartOption{
		name:      config.Spec.ImageName,
		imagePath: config.Spec.ImagePath,
		image:     config.Spec.ImageName,
		ports:     mapPort(config.Spec.Ports),
	}
}

func mapPort(bindings []PortBinding) map[ContainerPort]HostPort {
	ports := make(map[ContainerPort]HostPort)
	for _, binding := range bindings {
		ports[binding.ContainerPort] = binding.HostPort
	}
	return ports
}
