package valueobject

type Ports []Port

type Port struct {
	HostPort      int
	ContainerPort int
}
