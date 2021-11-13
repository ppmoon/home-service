package repository

import (
	"github.com/ppmoon/home-service/domain/aggregates"
)

type Config interface {
	Get() (*aggregates.Config, error)
}
