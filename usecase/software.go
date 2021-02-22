package usecase

import (
	"github.com/ppmoon/home-service/domain/entity"
	"github.com/ppmoon/home-service/infrastructure/podman"
	"github.com/ppmoon/home-service/infrastructure/systemd"
	"gopkg.in/yaml.v3"
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
func (s *SoftwareUseCase) Install(category, name, version string) (err error) {
	// check run config exist
	// check software is run
	// get software by name and version
	software, err := s.repo.Get(category, name, version)
	if err != nil {
		return
	}
	// create run config
	// read run config
	// create container by run config
	return
}
// Create software podman run config file
func (s *SoftwareUseCase) createRunConfig(software *entity.Software) (err error) {
	// process environments
	// save config file to software folder
	runConfig,err := yaml.Marshal(software)
	if err != nil {
		return
	}
	s.repo.CreateRunConfigFile(runConfig)
	return
}
