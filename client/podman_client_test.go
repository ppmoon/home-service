package client_test

import (
	"github.com/ppmoon/home-service/client"
	"testing"
)

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
	reference := "docker.io/library/alpine:3.12.2"
	err := pc.PullImages(reference)
	if err != nil {
		t.Error(err)
		return
	}
	// TODO complete this unit test
}
