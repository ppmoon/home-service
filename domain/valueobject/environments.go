package valueobject

type Environments []Environment

type Environment struct {
	Key         string `yaml:"key"`
	Value       string `yaml:"value"`
	Placeholder string `yaml:"placeholder"`
}
