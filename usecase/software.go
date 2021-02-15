package usecase

import (
	"github.com/ppmoon/home-service/domain/entity"
	"github.com/ppmoon/home-service/domain/valueobject"
	"github.com/ppmoon/home-service/infrastructure/podman"
	"github.com/ppmoon/home-service/infrastructure/systemd"
)

type SoftwareUseCase struct {
	systemd *systemd.Client
	podman  *podman.Client
	repo    entity.SoftwareRepository
}

func NewSoftwareUseCase(repo entity.SoftwareRepository) (*SoftwareUseCase, error) {
	systemdClient, err := systemd.NewSystemdClient()
	if err != nil {
		return nil, err
	}
	podmanClient, err := podman.NewPodmanClient()
	if err != nil {
		return nil, err
	}
	return &SoftwareUseCase{
		systemd: systemdClient,
		podman:  podmanClient,
		repo:    repo,
	}, nil
}

// Install software
func (s *SoftwareUseCase) Install(name, version, category string, environments valueobject.Environments) (err error) {
	// get software by name and version
	_, err = s.repo.Get(name, version, category)
	if err != nil {
		return
	}
	return
}
