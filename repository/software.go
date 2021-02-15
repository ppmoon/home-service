package repository

import (
	"github.com/ppmoon/home-service/domain/entity"
	"github.com/ppmoon/home-service/infrastructure/config"
	"github.com/ppmoon/home-service/infrastructure/git"
	"os"
)

type SoftwareRepository struct {
	git *git.Client
}

const (
	VersionLatest = "Latest"
	RepoFolder    = "./repo/"
)

func NewSoftwareRepository(git *git.Client) *SoftwareRepository {
	return &SoftwareRepository{
		git: git,
	}
}

// Get Software by name and version
func (s *SoftwareRepository) Get(name, version, category string) (softwareList []*entity.Software, err error) {
	// check is exist software repo
	err = s.checkSoftwareRepo()
	if err != nil {
		return nil, err
	}
	// check version
	if version == "" {
		version = VersionLatest
	}
	// find software

	return nil, err
}

// Check software repo.
// If not exist.git clone
// If exit git pull
func (s *SoftwareRepository) checkSoftwareRepo() error {
	// read source_list
	conf := config.GetConfig()
	for k, v := range conf.SourceList {
		path := RepoFolder + k
		var isExist bool
		isExist = s.isRepoFolderExist(path)
		if isExist {
			// git pull
			err := s.git.Pull(path)
			if err != nil {
				return err
			}
		} else {
			// git clone
			_, err := s.git.PlainClone(path, v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Check repo folder exist
func (s *SoftwareRepository) isRepoFolderExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
