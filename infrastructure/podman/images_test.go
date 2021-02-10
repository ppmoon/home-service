package podman_test

import (
	"github.com/ppmoon/home-service/infrastructure/podman"
	"testing"
)

func TestPodmanClient_PullImages(t *testing.T) {
	pc := podman.NewPodmanClient()
	reference := "docker.io/library/" + TestContainerName
	err := pc.PullImages(reference)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestPodmanClient_ImageExists(t *testing.T) {
	pc := podman.NewPodmanClient()
	isExist, err := pc.ImageExists(TestContainerName)
	if err != nil {
		t.Error(err)
		return
	}
	if !isExist {
		t.Error("podman image exists error.Is Exist ", isExist)
	}
}
