package client_test

import (
	"github.com/ppmoon/home-service/client"
	"testing"
)

const TestContainerName = "alpine:3.12.2"

func TestPodmanClient_Ping(t *testing.T) {
	pc := client.NewPodmanClient()
	pong, err := pc.Ping()
	if err != nil {
		t.Error(err)
		return
	}
	if pong != "OK" {
		t.Error("Ping is not OK")
		return
	}
}

func TestPodmanClient_PullImages(t *testing.T) {
	pc := client.NewPodmanClient()
	reference := "docker.io/library/" + TestContainerName
	err := pc.PullImages(reference)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestPodmanClient_ImageExists(t *testing.T) {
	pc := client.NewPodmanClient()
	isExist, err := pc.ImageExists(TestContainerName)
	if err != nil {
		t.Error(err)
		return
	}
	if !isExist {
		t.Error("podman image exists error.Is Exist ", isExist)
	}
}
