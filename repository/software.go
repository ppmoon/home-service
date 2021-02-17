package repository

import (
	"github.com/ppmoon/home-service/domain/entity"
	"github.com/ppmoon/home-service/infrastructure/config"
	"github.com/ppmoon/home-service/infrastructure/git"
	"github.com/ppmoon/home-service/infrastructure/log"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type SoftwareRepository struct {
	git          *git.Client
	repoRootPath string
	programPath  string
}

const (
	VersionLatest       = "Latest"
	ConfigParamFileName = "config_param.yaml"
)

func NewSoftwareRepository(repoRootPath, programPath string) *SoftwareRepository {
	g := git.NewGitClient()
	return &SoftwareRepository{
		git:          g,
		repoRootPath: repoRootPath,
		programPath:  programPath,
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
		path := filepath.Join(s.repoRootPath, sourceName)
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

// Read Repo Config Param
func (s *SoftwareRepository) ReadConfigParam(sourceName string) (configParam map[string]interface{}, err error) {
	path := filepath.Join(s.repoRootPath, sourceName, ConfigParamFileName)
	configParamByte, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(configParamByte, &configParam)
	if err != nil {
		return nil, err
	}
	return configParam, err
}
