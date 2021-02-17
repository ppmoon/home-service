package valueobject

type Volumes []Volume

type Volume struct {
	HostVolume      string `yaml:"host_volume"`
	ContainerVolume string `yaml:"container_volume"`
}
