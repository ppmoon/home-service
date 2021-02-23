package entity

import "github.com/ppmoon/home-service/domain/valueobject"

type Blueprint struct {
	Name         string                   `json:"name" yaml:"name"`
	Version      string                   `json:"version" yaml:"version"`
	Category     string                   `json:"category" yaml:"category"`
	Image        string                   `json:"image" yaml:"image"`
	Ports        valueobject.Ports        `json:"ports" yaml:"ports"`
	Volumes      valueobject.Volumes      `json:"volumes" yaml:"volumes"`
	Environments valueobject.Environments `json:"environments" yaml:"environments"`
	Dependence   []valueobject.Dependence `json:"dependence" yaml:"dependence"`
}

type BlueprintRepository interface {
}
