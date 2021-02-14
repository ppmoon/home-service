package valueobject

type Volumes []Volume

type Volume struct {
	HostVolume      string
	ContainerVolume string
}
