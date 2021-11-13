package services

import (
	"github.com/ppmoon/home-service/domain/aggregates"
	"github.com/ppmoon/home-service/domain/repository"
)

type ConfigService struct {
	repo repository.Config
}

func NewConfigService(config repository.Config) ConfigInterface {
	return &ConfigService{repo: config}
}

// Read Config
func (c *ConfigService) Read() (*aggregates.Config, error) {
	return c.repo.Get()
}
