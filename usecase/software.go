package usecase

import (
	"github.com/coreos/go-systemd/v22/dbus"
	"github.com/ppmoon/home-service/infrastructure/podman"
)

type SoftwareUseCase struct {
	*dbus.Conn
	*podman.Client
}

func NewSoftwareUseCase(conn *dbus.Conn, client *podman.Client) *SoftwareUseCase {
	return &SoftwareUseCase{
		Conn:   conn,
		Client: client,
	}
}

// Install software
func (s *SoftwareUseCase) Install(name, version string) (err error) {
	// get software by name and version
	return
}

// Get software info
