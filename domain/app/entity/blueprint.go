package entity

// Blueprint is software blueprint
type Blueprint struct {
	Name         string            `toml:"name"`
	Version      string            `toml:"version"`
	Image        string            `toml:"image"`
	Volumes      map[string]string `toml:"volumes"`
	Ports        map[int]int       `toml:"ports"`
	Environment  map[string]string `toml:"environment"`
	Command      string            `toml:"command"`
	Dependencies []Blueprint       `toml:"dependencies"`
}
