package entity

type Pipeline struct {
	Name     string
	StepList []PipelineStep
	Cursor   int
}

func NewPipeline(name string) *Pipeline {
	return nil
}

func (p *Pipeline) AddStep() {

}

func (p *Pipeline) Run() {

}

func (p *Pipeline) GetProcess() {

}
