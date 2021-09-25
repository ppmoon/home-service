package repository

import "github.com/ppmoon/home-service/domain/entity"

type Config interface {
	Get() (*entity.Config, error)
}
