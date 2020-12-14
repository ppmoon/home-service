package client_test

import (
	"github.com/ppmoon/home-service/client"
	"testing"
)

func TestPodmanClient_Ping(t *testing.T) {
	pc := client.NewPodmanClient()
	pong,err := pc.Ping()
	if err != nil {
		t.Error(err)
		return
	}
	if pong != "OK" {
		t.Error("Ping is not OK")
		return
	}
}
