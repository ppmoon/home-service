package client

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/ppmoon/home-service/log"
	"io"
	"net"
	"net/http"
	"strings"
)

const PodmanUnixSocket = "/run/podman/podman.sock"
const HostURL = "http://d/v1.0.0"

type PodmanClient struct {
	*resty.Client
}

// If combine PodmanClient Struct must be new this function
func NewPodmanClient() *PodmanClient {
	// Create a Go's http.Transport so we can set it in resty.
	transport := http.Transport{
		DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
			return net.Dial("unix", PodmanUnixSocket)
		},
	}
	// Create a Resty Client
	client := resty.New()

	// Set the previous transport that we created, set the scheme of the communication to the
	// socket and set the unixSocket as the HostURL.
	client.SetTransport(&transport).SetScheme("http").SetHostURL(HostURL)
	return &PodmanClient{
		client,
	}
}

//
// Ping podman
func (p *PodmanClient) Ping() (pong string, err error) {
	resp, err := p.R().Get("/libpod/_ping")
	if err != nil {
		log.Error("Podman ping error ", err)
		log.Info("Please check podman.service status.Type  systemctl status podman.socket  in terminal for check.Or Does home-service run in root user?")
		return
	}
	pong = resp.String()
	return
}

type PullImagesResp struct {
	Error  string   `json:"error"`
	ID     string   `json:"id"`
	Images []string `json:"images"`
	Stream string   `json:"stream"`
}

// podman pull image
func (p *PodmanClient) PullImages(reference string) error {
	resp, err := p.R().
		SetQueryParams(map[string]string{
			"reference": reference,
		}).
		SetDoNotParseResponse(true).
		Post("/libpod/images/pull")
	if err != nil {
		log.Error("Podman pull images error ", err)
		return err
	}
	defer resp.RawBody().Close()
	reader := bufio.NewReader(resp.RawBody())
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		var pullImagesResp PullImagesResp
		err = json.Unmarshal(line, &pullImagesResp)
		if err != nil {
			fmt.Println("error", err)
			return err
		}
		if pullImagesResp.Error != "" {
			err = errors.New(pullImagesResp.Error)
			fmt.Println("error", err)
			return err
		}
		if pullImagesResp.ID != "" {
			fmt.Println("ID:", strings.Replace(pullImagesResp.ID, "\n", "", -1))
		}
		if len(pullImagesResp.Images) != 0 {
			for _, image := range pullImagesResp.Images {
				fmt.Println("Image:", strings.Replace(image, "\n", "", -1))
			}
		}
		if pullImagesResp.Stream != "" {
			fmt.Println(strings.Replace(pullImagesResp.Stream, "\n", "", -1))
		}
	}
	return nil
}
