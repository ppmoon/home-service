package entity

type IPipeline interface {
	AddStep(step *PipelineStep) error
	Run() error
	GetPipelineStatus() *PipelineStatus
}
