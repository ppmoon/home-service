package entity

type Pipeline struct {
	name     string
	stepList []PipelineStep
	cursor   int
}

func NewPipeline(name string) *Pipeline {
	return nil
}

func (p *Pipeline) AddStep(step PipelineStep) {
	p.stepList = append(p.stepList, step)
}

func (p *Pipeline) Run() {

}

func (p *Pipeline) GetProcess() {

}
