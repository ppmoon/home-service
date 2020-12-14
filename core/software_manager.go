package core

type SoftwareManager struct {
}

func NewSoftwareManager() *SoftwareManager {
	return &SoftwareManager{}
}

// Download container image.
// Podman pull image
func (s *SoftwareManager) DownloadImage()  {

}
// Get software config repo cache
// install software
// remove software
// startup setting