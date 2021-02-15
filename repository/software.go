package repository

import (
	"github.com/ppmoon/home-service/domain/entity"
	"github.com/ppmoon/home-service/infrastructure/config"
	"github.com/ppmoon/home-service/infrastructure/git"
	"github.com/ppmoon/home-service/infrastructure/log"
	"os"
)

type SoftwareRepository struct {
	git          *git.Client
	repoRootPath string
}

const (
	VersionLatest = "Latest"
)

func NewSoftwareRepository(repoRootPath string) *SoftwareRepository {
	g := git.NewGitClient()
	return &SoftwareRepository{
		git:          g,
		repoRootPath: repoRootPath,
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
	log.Infof("conf.SourceList %v", conf.SourceList)
	for sourceName, sourceGitUrl := range conf.SourceList {
		path := s.repoRootPath + sourceName
		var isExist bool
		isExist = s.isRepoFolderExist(path)
		log.Infof("check repo folder exist path:%s,result:%t", path, isExist)
		if isExist {
			// git pull
			err := s.git.Pull(path)
			if err != nil {
				return err
			}
		} else {
			// git clone
			_, err := s.git.PlainClone(path, sourceGitUrl)
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
