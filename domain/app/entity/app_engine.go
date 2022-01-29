package entity

type AppEngine struct {
	blueprint Blueprint
}

func (a *AppEngine) LoadBlueprint(blueprint Blueprint) {
	//TODO implement me
	panic("implement me")
}

func NewAppEngine() *AppEngine {
	return &AppEngine{}
}

func (a *AppEngine) DownloadApp() error {
	//TODO implement me
	panic("implement me")
}

func (a *AppEngine) RunApp() error {
	//TODO implement me
	panic("implement me")
}

func (a *AppEngine) ParseAppBlueprint() error {
	//TODO implement me
	panic("implement me")
}
