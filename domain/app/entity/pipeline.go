package entity

import (
	"errors"
)

const (
	PipelineStatusNotRun  = 0
	PipelineStatusRunning = 1
	PipelineStatusFinish  = 2
)

type Pipeline struct {
	name     string
	stepList []*PipelineStep
	cursor   int
	status   int // 0 default 1 running 2 finish
}

func NewPipeline(name string) *Pipeline {
	return &Pipeline{
		name: name,
	}
}

func (p *Pipeline) AddStep(step *PipelineStep) error {
	if p.status != PipelineStatusNotRun {
		return errors.New("pipeline is running can't add step")
	}
	p.stepList = append(p.stepList, step)
	return nil
}

func (p *Pipeline) Run() error {
	if p.status != PipelineStatusNotRun {
		return errors.New("pipeline is running,don't repeat the call")
	}
	p.status = PipelineStatusRunning
	for _, step := range p.stepList {
		err := step.Fn()
		if err != nil {
			return err
		}
		p.cursor = p.cursor + 1
	}
	p.status = PipelineStatusFinish
	return nil
}

func (p *Pipeline) GetPipelineStatus() *PipelineStatus {
	return &PipelineStatus{
		Cursor:         p.cursor,
		Total:          len(p.stepList),
		RunningStep:    p.stepList[p.cursor],
		RateOfProgress: (p.cursor + 1) / len(p.stepList),
	}
}
