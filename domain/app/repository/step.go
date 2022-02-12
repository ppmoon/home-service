package repository

import (
	"errors"
	"github.com/ppmoon/home-service/infrastructure/memsql"
	"gorm.io/gorm"
	"sync"
)

var stepMigrateOnce sync.Once

type Step struct {
	Base
	Name       string
	Status     int
	PipelineID int
}

func NewStep() *Step {
	stepMigrateOnce.Do(func() {
		db := memsql.GetDB()
		err := db.AutoMigrate(&Step{})
		panic(err)
	})
	return &Step{}
}

func (s *Step) Create(tx *gorm.DB, step *Step) (id int, err error) {
	result := tx.Create(step)
	return step.ID, result.Error
}
func (s *Step) UpdateStatusFromXToY(id, x, y int) (err error) {
	db := memsql.GetDB()
	result := db.Model(&Step{}).Where("id = ? and status = ?", id, x).Update("status", y)
	err = result.Error
	if result.RowsAffected == 0 {
		err = errors.New("RowsAffected is zero")
	}
	return err
}
func (s *Step) GetAllPipelineSteps(pipelineID int) (steps []*Step) {
	db := memsql.GetDB()
	db.Where("pipeline_id = ?", pipelineID).Find(&steps)
	return steps
}
