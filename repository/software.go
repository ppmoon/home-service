package repository

import (
	"github.com/ppmoon/home-service/domain/entity"
	"github.com/ppmoon/home-service/infrastructure/bolt"
	"go.etcd.io/bbolt"
	"gopkg.in/yaml.v3"
)

type SoftwareRepository struct {
	boltDBPath string
}

func NewSoftwareRepository(boltDBPath string) *SoftwareRepository {
	return &SoftwareRepository{
		boltDBPath: boltDBPath,
	}
}

// Get Software from boltDB
func (s *SoftwareRepository) Get(category, name, version string) (softwareList []*entity.Software, err error) {
	if category == "" {
		category = DefaultRepo
	}
	boltDB, err := bolt.GetBoltDB(s.boltDBPath + "/" + category + BoltSuffix)
	if err != nil {
		return nil, err
	}
	err = boltDB.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(name))
		v := b.Get([]byte(version))
		if v != nil {
			var software *entity.Software
			err = yaml.Unmarshal(v, software)
			if err != nil {
				return err
			}
			softwareList = append(softwareList, software)
			return err
		}
		err = b.ForEach(func(k, v []byte) error {
			var software *entity.Software
			err = yaml.Unmarshal(v, software)
			if err != nil {
				return err
			}
			softwareList = append(softwareList, software)
			return err
		})
		return err
	})
	return
}
