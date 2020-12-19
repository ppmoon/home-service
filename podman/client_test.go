package podman_test

import (
	"github.com/ppmoon/home-service/podman"
	"testing"
)

const TestContainerName = "alpine:3.12.2"

func TestPodmanClient_Ping(t *testing.T) {
	pc := podman.NewPodmanClient()
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
