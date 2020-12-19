package podman

// Play a Kubernetes YAML file.
func (c *Client) PlayK8sYaml(yamlFileString string) (string, error) {
	resp, err := c.R().
		SetBody(yamlFileString).
		Post("/libpod/play/kube")
	if err != nil {
		return "", err
	}
	// TODO parse resp
	return resp.String(), nil
}
