package repository

import (
	"errors"
	"github.com/ppmoon/home-service/domain/entity"
)

var (
	SoftwareHasNotBeenInstalled = errors.New("software has not been installed")
	VersionHasNotBeenInstalled  = errors.New("version has not been installed")
)

type SoftwareRepository struct {
}

func NewSoftwareRepository() *SoftwareRepository {
	return &SoftwareRepository{}
}

// Get Software
func (s *SoftwareRepository) Get(category, name, version string) (softwareList []*entity.Software, err error) {

	return
}
