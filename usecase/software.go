package usecase

import (
	"github.com/coreos/go-systemd/v22/dbus"
	"github.com/ppmoon/home-service/infrastructure/podman"
)

type SoftwareUseCase struct {
	dbusConn     *dbus.Conn
	podmanClient *podman.Client
}

func NewSoftwareUseCase(dbusConn *dbus.Conn, podmanClient *podman.Client) *SoftwareUseCase {
	return &SoftwareUseCase{
		dbusConn:     dbusConn,
		podmanClient: podmanClient,
	}
}

// Install software
func (s *SoftwareUseCase) Install(name, version string) (err error) {
	// get software by name and version
	return
}

// Get software info
