package entity

import (
	"errors"
	"github.com/ppmoon/home-service/domain/app/repository"
	"github.com/ppmoon/home-service/infrastructure/memsql"
	"gorm.io/gorm"
	"time"
)

const (
	PipelineStatusNotRun  = 0
	PipelineStatusRunning = 1
	PipelineStatusFinish  = 2
)

type Pipeline struct {
	ID       int
	name     string
	stepList []*PipelineStep
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

func (p *Pipeline) Run(expireTime time.Duration) error {
	pipelineRepo := repository.NewPipeline()
	p.status = pipelineRepo.GetLastPipelineStatus(p.name)
	if p.status != PipelineStatusNotRun {
		return errors.New("pipeline is running,don't repeat the call")
	}
	stepRepo := repository.NewStep()
	db := memsql.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		var err error
		p.ID, err = pipelineRepo.Create(tx, &repository.Pipeline{
			Name:       p.name,
			Status:     PipelineStatusRunning,
			ExpireTime: time.Now().Add(expireTime),
			TotalTask:  len(p.stepList),
		})
		if err != nil {
			return err
		}
		// create step
		for _, step := range p.stepList {
			stepID, err := stepRepo.Create(tx, &repository.Step{
				Name:       step.Name,
				Status:     StepStatusNotRun,
				PipelineID: p.ID,
			})
			if err != nil {
				return err
			}
			step.ID = stepID
		}
		return err
	})
	if err != nil {
		return err
	}
	for _, step := range p.stepList {
		err = stepRepo.UpdateStatusFromXToY(step.ID, StepStatusNotRun, StepStatusRunning)
		if err != nil {
			return err
		}
		err = step.Fn()
		if err != nil {
			errStatus := stepRepo.UpdateStatusFromXToY(step.ID, StepStatusRunning, StepStatusFail)
			err = errors.New(err.Error() + ";" + errStatus.Error())
			return err
		}
		err = stepRepo.UpdateStatusFromXToY(step.ID, StepStatusRunning, StepStatusSuccess)
		if err != nil {
			return err
		}
	}
	p.status = PipelineStatusFinish
	return pipelineRepo.UpdateStatusFromXToY(p.ID, PipelineStatusRunning, PipelineStatusFinish)
}

func (p *Pipeline) GetPipelineStatus() *PipelineStatus {
	pipelineRepo := repository.NewPipeline()
	pipeline := pipelineRepo.GetLastPipeline(p.name)
	if pipeline == nil {
		return nil
	}
	stepRepo := repository.NewStep()
	steps := stepRepo.GetAllPipelineSteps(pipeline.ID)
	stepList := make([]*PipelineStep, 0)
	currentTask := 0
	for _, step := range steps {
		if step.Status == StepStatusSuccess {
			currentTask++
		}
		stepList = append(stepList, &PipelineStep{
			ID:     step.ID,
			Name:   step.Name,
			Status: step.Status,
		})
	}
	return &PipelineStatus{
		Pipeline: Pipeline{
			ID:       pipeline.ID,
			name:     pipeline.Name,
			stepList: stepList,
			status:   pipeline.Status,
		},
		RateOfProgress: (currentTask * 100) / (len(stepList) * 100),
	}
}
