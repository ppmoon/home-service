package core_test

import (
	"github.com/ppmoon/home-service/core"
	"testing"
)

func TestSoftwareManager_GetUnitList(t *testing.T) {
	sm, err := core.NewSoftwareManager()
	if err != nil {
		t.Error(err)
		return
	}
	unitStatus, err := sm.GetUnitList()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(unitStatus)
}
