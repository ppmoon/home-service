package usecase

import "github.com/ppmoon/home-service/domain/entity"

type Config interface {
	Read() (entity.Config, error)
}
