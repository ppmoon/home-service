package entity

type IAppEngine interface {
	ParseAppBlueprint() error
	DownloadApp() error
	RunApp() error
	LoadBlueprint(blueprint Blueprint)
}
