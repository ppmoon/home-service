package entity

type Blueprint struct {
	Name         string       `json:"name" yaml:"name"`
	Version      string       `json:"version" yaml:"version"`
	Category     string       `json:"category" yaml:"category"`
	Image        string       `json:"image" yaml:"image"`
	Ports        []Ports        `json:"ports" yaml:"ports"`
	Volumes      []Volumes      `json:"volumes" yaml:"volumes"`
	Environments []Environments `json:"environments" yaml:"environments"`
	Dependence   []Dependence `json:"dependence" yaml:"dependence"`
}

type BlueprintRepository interface {
}
