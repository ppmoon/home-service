package entity

type IPipeline interface {
	AddStep()
	Run()
	GetProcess()
}
