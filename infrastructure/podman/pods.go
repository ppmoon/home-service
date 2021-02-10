package podman

import (
	"encoding/json"
	"errors"
	"github.com/ppmoon/home-service/log"
	"net/http"
)

// Play a k8s YAML file response
type PlayK8sYamlResp struct {
	Cause    string `json:"cause"`
	Message  string `json:"message"`
	Response int    `json:"response"`
}

// Play a Kubernetes YAML file.
func (c *Client) PlayK8sYaml(yamlFileString string) error {
	resp, err := c.R().
		SetBody(yamlFileString).
		Post("/libpod/play/kube")
	if err != nil {
		return err
	}
	var playK8sYamlResp PlayK8sYamlResp
	err = json.Unmarshal(resp.Body(),&playK8sYamlResp)
	if err != nil {
		return err
	}
	if resp.StatusCode() == http.StatusInternalServerError {
		err = errors.New(playK8sYamlResp.Message)
		log.Error(playK8sYamlResp.Message)
		return err
	}
	return nil
}
