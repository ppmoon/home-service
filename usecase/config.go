package usecase

import (
	"github.com/ppmoon/home-service/domain/entity"
	"github.com/ppmoon/home-service/repository"
)

type ConfigUseCase struct {
	repo repository.Config
}

func NewConfigUseCase(repo repository.Config) Config {
	return &ConfigUseCase{
		repo: repo,
	}
}

// Read Config
func (c ConfigUseCase) Read() (entity.Config, error) {
	return c.repo.Get()
}
