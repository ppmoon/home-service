package core

import (
	"github.com/coreos/go-systemd/v22/dbus"
	"github.com/ppmoon/home-service/client"
)

type SoftwareManager struct {
	*dbus.Conn
	*client.PodmanClient
}

func NewSoftwareManager() (*SoftwareManager, error) {
	dbusConn, err := dbus.NewSystemdConnection()
	if err != nil {
		return nil, err
	}
	podmanClient := client.NewPodmanClient()
	return &SoftwareManager{
		dbusConn,
		podmanClient,
	}, nil
}

// Download container image.
// Podman pull image
func (s *SoftwareManager) DownloadImage() {

}

// Get software config repo cache
// install software
// remove software
// startup setting
// Get Home service unit list
func (s *SoftwareManager) GetUnitList() ([]dbus.UnitStatus, error) {
	// TODO add home service prefix name for get unit list.Home service just control itself software.
	return s.ListUnitsByPatterns([]string{}, []string{"podman*"})
}
