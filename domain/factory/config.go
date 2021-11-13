package factory

import (
	"github.com/ppmoon/home-service/domain/repository"
	"github.com/ppmoon/home-service/domain/services"
)

func NewConfig(config repository.Config) services.ConfigInterface {
	return services.NewConfigService(config)
}
