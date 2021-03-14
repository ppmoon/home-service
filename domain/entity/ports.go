package entity

type Ports []Port

type Port struct {
	HostPort      int    `yaml:"host_port"`
	ContainerPort int    `yaml:"container_port"`
	Protocol      string `yaml:"protocol"`
}
