package repository

import (
	"github.com/ppmoon/home-service/domain/entity"
)

const ConfigBoltDBName = "config.db"
const ConfigBoltBucketName = "config"

type ConfigRepo struct {
}

// Get Read from boltDB sync.Onec
func (c ConfigRepo) Get() (*entity.Config, error) {
	panic("")
}

func NewConfigRepo() Config {
	return &ConfigRepo{}
}
