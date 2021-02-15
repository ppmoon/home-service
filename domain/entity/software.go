package entity

import "github.com/ppmoon/home-service/domain/valueobject"

// Software
type Software struct {
	Name         string                   `json:"name"`
	Version      string                   `json:"version"`
	Category     string                   `json:"category"`
	Image        string                   `json:"image"`
	Ports        valueobject.Ports        `json:"ports"`
	Volumes      valueobject.Volumes      `json:"volumes"`
	Environments valueobject.Environments `json:"environments"`
	Dependence   []Software               `json:"dependence"`
}

// Software use case
type SoftwareUseCase interface {
	Install(name, version, category string, environments valueobject.Environments) (err error)
}

type SoftwareRepository interface {
	Get(name, version, category string) (software []*Software, err error)
}
