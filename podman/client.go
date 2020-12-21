package podman

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/ppmoon/home-service/log"
	"net"
	"net/http"
)

const UnixSocket = "/run/podman/podman.sock"
const HostURL = "http://d/v1.0.0"

// Podman Http Client
type Client struct {
	*resty.Client
}

// If combine PodmanClient Struct must be new this function
func NewPodmanClient() *Client {
	// Create a Go's http.Transport so we can set it in resty.
	transport := http.Transport{
		DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
			return net.Dial("unix", UnixSocket)
		},
	}
	// Create a Resty Client
	client := resty.New()
	// Set the previous transport that we created, set the scheme of the communication to the
	// socket and set the unixSocket as the HostURL.
	client.SetTransport(&transport).SetScheme("http").SetHostURL(HostURL)
	return &Client{
		client,
	}
}

//
// Ping podman
func (c *Client) Ping() (pong string, err error) {
	resp, err := c.R().Get("/libpod/_ping")
	if err != nil {
		log.Error("Podman ping error ", err, `\nPlease check podman.service status.Type  systemctl status podman.socket  in terminal for check Or Does home-service run in root user?`)
		return
	}
	pong = resp.String()
	return
}

// TODO add a network check and change mirror functon.
