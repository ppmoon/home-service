package entity

import "github.com/ppmoon/home-service/domain/valueobject"

// Software
type Software struct {
	Name         string                   `json:"name"`
	Version      string                   `json:"version"`
	Image        string                   `json:"image"`
	Ports        valueobject.Ports        `json:"ports"`
	Volumes      valueobject.Volumes      `json:"volumes"`
	Environments valueobject.Environments `json:"environments"`
	Dependence   []Software               `json:"dependence"`
}

// Software use case
type SoftwareUseCase interface {
	Install(name, version string) (err error)
}

type SoftwareRepository interface {
	GetLastByName(name string) (err error)
	GetByNameAndVersion(name, version string)
}
