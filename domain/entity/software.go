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
	Status       string                   `json:"status" yaml:"status"`
}

// Software service
type SoftwareService interface {
	Install(category, name, version string) (err error)
}

// Software use case
type SoftwareUseCase interface {
	Install(category, name, version string) (err error)
}
type SoftwareRepository interface {
	IsSoftwareInstalled()
}

// Software BluePrint repository
type SoftwareBlueprintRepository interface {
	Get(category, name, version string) (software *Software, err error)
	Search(category, name, version string) (software *Software, err error)
	ReadConfigParam(sourceName string) (configParam map[string]interface{}, err error)
	CreateRunConfigFile(content []byte) (err error)
}

// Software Environments repository
type EnvRepository interface {
}
