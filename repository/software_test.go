package repository

import (
	"github.com/ppmoon/home-service/infrastructure/config"
	"testing"
)

func TestSoftwareRepository_Get(t *testing.T) {
	config.InitConfig("./../config")
	repo := NewSoftwareRepository("./../repo/")
	sList, err := repo.Get("", "", "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(sList)
}
