package podman

import (
	"context"
	"github.com/containers/podman/v2/pkg/bindings"
	"os"
)

var connText context.Context

type Client struct {
	Image Image
}

func NewPodmanClient() (client *Client, err error) {
	// Get Podman socket location
	sockDir := os.Getenv("XDG_RUNTIME_DIR")
	socket := "unix:" + sockDir + "/podman/podman.sock"
	// Connect to Podman socket
	connText, err = bindings.NewConnection(context.Background(), socket)
	if err != nil {
		return nil, err
	}
	return &Client{}, nil
}
