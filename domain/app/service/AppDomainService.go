package service

import (
	"github.com/ppmoon/home-service/domain/app/entity"
	"github.com/ppmoon/home-service/infrastructure/lock"
	"time"
)

const InstallTimeout = 3600
const (
	StepNameParseAppBlueprint = "parse app blueprint"
	StepNameDownloadApp       = "download app"
	StepNameRun               = "run"
)

type AppDomainService struct {
	appEngine entity.IAppEngine
	lock      lock.IMultiRoutineLock
}

func NewAppDomainService() *AppDomainService {
	return &AppDomainService{
		appEngine: entity.NewAppEngine(),
		lock:      lock.NewMultiRoutineLock(),
	}
}

// Install software
func (a *AppDomainService) Install(blueprint entity.Blueprint) (err error) {
	// lock
	a.lock.Lock(blueprint.Name, time.Second*InstallTimeout)
	defer a.lock.Unlock(blueprint.Name)
	// new a pipeline
	pipeline := entity.NewPipeline(blueprint.Name)
	a.appEngine.LoadBlueprint(blueprint)
	err = pipeline.AddStep(&entity.PipelineStep{
		Name: StepNameParseAppBlueprint,
		Fn: func() error {
			return a.appEngine.ParseAppBlueprint()
		},
	})
	err = pipeline.AddStep(&entity.PipelineStep{
		Name: StepNameDownloadApp,
		Fn: func() error {
			return a.appEngine.DownloadApp()
		},
	})
	err = pipeline.AddStep(&entity.PipelineStep{
		Name: StepNameRun,
		Fn: func() error {
			return a.appEngine.RunApp()
		},
	})
	err = pipeline.Run()
	return err
}

func (a *AppDomainService) Update() {

}
func (a *AppDomainService) Delete() {

}
