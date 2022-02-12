package repository

import (
	"errors"
	"github.com/ppmoon/home-service/infrastructure/memsql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var pipelineMigrateOnce sync.Once

type Pipeline struct {
	Base
	Name       string
	Status     int
	ExpireTime time.Time
	TotalTask  int
}

func NewPipeline() *Pipeline {
	pipelineMigrateOnce.Do(func() {
		db := memsql.GetDB()
		err := db.AutoMigrate(&Pipeline{})
		if err != nil {
			panic(err)
		}
	})
	return &Pipeline{}
}

func (p *Pipeline) Create(tx *gorm.DB, pipeline *Pipeline) (id int, err error) {
	result := tx.Create(pipeline)
	return pipeline.ID, result.Error
}
func (p *Pipeline) GetLastPipeline(name string) *Pipeline {
	db := memsql.GetDB()
	var pipeline *Pipeline
	db.Last(&pipeline, "name = ? and expire_time > ?", name, time.Now())
	// if not find anything status will be zero its means not run
	return pipeline
}
func (p *Pipeline) GetLastPipelineStatus(name string) int {
	db := memsql.GetDB()
	var pipeline Pipeline
	db.Last(&pipeline, "name = ? and expire_time > ?", name, time.Now())
	// if not find anything status will be zero its means not run
	return pipeline.Status
}
func (p *Pipeline) UpdateStatusFromXToY(id, x, y int) (err error) {
	db := memsql.GetDB()
	result := db.Model(&Pipeline{}).Where("id = ? and status = ?", id, x).Update("status", y)
	err = result.Error
	if result.RowsAffected == 0 {
		err = errors.New("RowsAffected is zero")
	}
	return err
}
