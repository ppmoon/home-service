package services

import (
	"github.com/ppmoon/home-service/domain/aggregates"
)

type ConfigInterface interface {
	Read() (*aggregates.Config, error)
}
