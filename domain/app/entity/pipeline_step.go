package entity

type PipelineStep struct {
	Name string
	Fn   func() error
}
