package podman

import (
	"bufio"
	"encoding/json"
	"errors"
	"github.com/ppmoon/home-service/log"
	"io"
	"net/http"
	"strings"
)
// Podman pull images Response
type PullImagesResp struct {
	Error  string   `json:"error"`
	ID     string   `json:"id"`
	Images []string `json:"images"`
	Stream string   `json:"stream"`
}
// Podman pull image
func (c *Client) PullImages(reference string) error {
	resp, err := c.R().
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
			log.Error("error", err)
			return err
		}
		if pullImagesResp.Error != "" {
			err = errors.New(pullImagesResp.Error)
			log.Error("error", err)
			return err
		}
		if pullImagesResp.ID != "" {
			log.Info("ID:", strings.Replace(pullImagesResp.ID, "\n", "", -1))
		}
		if len(pullImagesResp.Images) != 0 {
			for _, image := range pullImagesResp.Images {
				log.Info("Image:", strings.Replace(image, "\n", "", -1))
			}
		}
		if pullImagesResp.Stream != "" {
			log.Info(strings.Replace(pullImagesResp.Stream, "\n", "", -1))
		}
	}
	return nil
}

// Podman image exists response
type ImageExistsResp struct {
	Cause    string `json:"cause"`
	Message  string `json:"message"`
	Response int    `json:"response"`
}

// Podman image exists
func (c *Client) ImageExists(name string) (isExist bool, err error) {
	resp, err := c.R().Get("/libpod/images/" + name + "/exists")
	if err != nil {
		log.Error("podman image exists error", err)
		return
	}
	if resp.StatusCode() == http.StatusNoContent {
		isExist = true
		return
	}
	return
}
