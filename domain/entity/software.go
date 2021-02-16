package entity

import "github.com/ppmoon/home-service/domain/valueobject"

// Software
type Software struct {
	Name         string                   `json:"name" yaml:"name"`
	Version      string                   `json:"version" yaml:"version"`
	Category     string                   `json:"category" yaml:"category"`
	Image        string                   `json:"image" yaml:"image"`
	Ports        valueobject.Ports        `json:"ports" yaml:"ports"`
	Volumes      valueobject.Volumes      `json:"volumes" yaml:"volumes"`
	Environments valueobject.Environments `json:"environments" yaml:"environments"`
	Dependence   []valueobject.Dependence `json:"dependence" yaml:"dependence"`
}

// Software use case
type SoftwareUseCase interface {
	Install(name, version, category string) (err error)
}

type SoftwareRepository interface {
	Get(name, version, category string) (software []*Software, err error)
	ReadConfigParam(sourceName string) (configParam map[string]interface{}, err error)
}
