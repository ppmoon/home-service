package valueobject

type Dependence struct {
	Name     string `json:"name" yaml:"name"`
	Version  string `json:"version" yaml:"version"`
	Category string `json:"category" yaml:"category"`
}
