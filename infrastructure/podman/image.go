package podman

import (
	"github.com/containers/podman/v2/pkg/bindings/images"
	"github.com/containers/podman/v2/pkg/domain/entities"
)

type Image struct {
}

func (i *Image) Pull(rawImage string, options entities.ImagePullOptions) ([]string, error) {
	return images.Pull(connText, rawImage, options)
}
