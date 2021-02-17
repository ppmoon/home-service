package repository

import (
	"github.com/ppmoon/home-service/infrastructure/config"
	"testing"
)

func TestSoftwareRepository_Get(t *testing.T) {
	config.InitConfig("./../config")
	repo := NewSoftwareRepository("./../repo", "")
	sList, err := repo.Get("", "kms", "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(sList)
}

func TestSoftwareRepository_ReadConfigParam(t *testing.T) {
	config.InitConfig("./../config")
	repo := NewSoftwareRepository("./../repo", "./../program")
	cp, err := repo.ReadConfigParam("default_repo")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(cp)
}
