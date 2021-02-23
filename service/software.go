package service

import "github.com/ppmoon/home-service/domain/entity"

type SoftwareService struct {
	SoftwareRepo entity.SoftwareRepository
}

func NewSoftwareService(softwareRepo entity.SoftwareRepository) *SoftwareService {
	return &SoftwareService{
		SoftwareRepo: softwareRepo,
	}
}

func (s *SoftwareService) Install(category, name, version string) (err error) {
	// is software installed
}
