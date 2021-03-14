package entity

// Software
type Software struct {
	Name               string         `json:"name" yaml:"name"`
	Version            string         `json:"version" yaml:"version"`
	Category           string         `json:"category" yaml:"category"`
	Image              string         `json:"image" yaml:"image"`
	Ports              []Ports        `json:"ports" yaml:"ports"`
	Volumes            []Volumes      `json:"volumes" yaml:"volumes"`
	Environments       []Environments `json:"environments" yaml:"environments"`
	Dependence         []Software     `json:"dependence" yaml:"dependence"`
	Status             string         `json:"status" yaml:"status"`
	IsSetStartWithBoot bool           `json:"is_set_start_with_boot" yaml:"is_set_start_with_boot"`
}

// Software use case
type SoftwareUseCase interface {
	Install(category, name, version string) (err error)
}
type SoftwareRepository interface {
	Get(category, name, version string) (softwareList []*Software, err error)
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
