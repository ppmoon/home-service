package usecase

import (
	"github.com/ppmoon/home-service/domain/entity"
	"github.com/ppmoon/home-service/infrastructure/podman"
	"github.com/ppmoon/home-service/infrastructure/systemd"
)

type SoftwareUseCase struct {
	systemd               *systemd.Client
	podman                *podman.Client
	SoftwareBlueprintRepo entity.SoftwareBlueprintRepository
	SoftwareRepo          entity.SoftwareRepository
}

func NewSoftwareUseCase(repo entity.SoftwareBlueprintRepository) (*SoftwareUseCase, error) {
	systemdClient, err := systemd.NewSystemdClient()
	if err != nil {
		return nil, err
	}
	podmanClient, err := podman.NewPodmanClient()
	if err != nil {
		return nil, err
	}
	return &SoftwareUseCase{
		systemd:               systemdClient,
		podman:                podmanClient,
		SoftwareBlueprintRepo: repo,
	}, nil
}

// Install software
func (s *SoftwareUseCase) Install(category, name, version string) (err error) {
	isInstalled, err := s.isSoftwareInstalled(category, name, version)
	if err != nil {
		return err
	}
	return
}

// Is Software Installed
func (s *SoftwareUseCase) isSoftwareInstalled(category, name, version string) (isInstalled bool, err error) {
	software, err := s.SoftwareRepo.Get(category, name, version)
	if err != nil {
		return false, err
	}
	if software != nil {
		return true, nil
	}
	return false, nil
}
