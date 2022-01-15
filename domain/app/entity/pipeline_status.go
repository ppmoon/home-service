package entity

type PipelineStatus struct {
	Cursor         int
	Total          int
	RateOfProgress int
	Status         int
	RunningStep    *PipelineStep
}
