package entity

const (
	StepStatusNotRun  = 0
	StepStatusRunning = 1
	StepStatusSuccess = 2
	StepStatusFail    = 3
)

type PipelineStep struct {
	ID     int
	Name   string
	Fn     func() error
	Status int
}
