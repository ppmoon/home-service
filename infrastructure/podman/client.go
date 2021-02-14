package podman

import (
	"context"
	"github.com/containers/podman/v2/pkg/bindings"
	"github.com/ppmoon/home-service/infrastructure/log"
	"os"
)

var connText context.Context

type Client struct {
	Image Image
}

func NewPodmanClient() (err error) {
	// Get Podman socket location
	sockDir := os.Getenv("XDG_RUNTIME_DIR")
	socket := "unix:" + sockDir + "/podman/podman.sock"
	// Connect to Podman socket
	connText, err = bindings.NewConnection(context.Background(), socket)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	return
}
