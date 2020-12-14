package client

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/ppmoon/home-service/log"
	"net"
	"net/http"
)

const PodmanUnixSocket = "/run/podman/podman.sock"

type PodmanClient struct {
	*resty.Client
}

func NewPodmanClient() *PodmanClient {
	// Create a Go's http.Transport so we can set it in resty.
	transport := http.Transport{
		DialContext: func(_ context.Context,_,_ string) (net.Conn,error){
			return net.Dial("unix",PodmanUnixSocket)
		},
	}
	// Create a Resty Client
	client := resty.New()

	// Set the previous transport that we created, set the scheme of the communication to the
	// socket and set the unixSocket as the HostURL.
	client.SetTransport(&transport).SetScheme("http").SetHostURL("http://d/v1.0.0")
	return &PodmanClient{
		client,
	}
}
//
// Ping podman
func (p *PodmanClient) Ping() (pong string,err error) {
	resp,err := p.R().Get("/libpod/_ping")
	if err != nil {
		log.Error("Podman ping error ",err)
		log.Info("Please check podman.service status.Type  systemctl status podman.socket  in terminal for check.Or Does home-service run in root user?")
		return
	}
	pong = resp.String()
	return
}
// podman pull image
//func (p *PodmanClient) PullImages()