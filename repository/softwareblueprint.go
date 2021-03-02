package repository

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ppmoon/home-service/domain/entity"
	"github.com/ppmoon/home-service/infrastructure/config"
	"github.com/ppmoon/home-service/infrastructure/git"
	"github.com/ppmoon/home-service/infrastructure/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type SoftwareBlueprintRepository struct {
	git          *git.Client
	repoRootPath string
}

func NewSoftwareBlueprintRepository(repoRootPath string) *SoftwareBlueprintRepository {
	g := git.NewGitClient()
	return &SoftwareBlueprintRepository{
		git:          g,
		repoRootPath: repoRootPath,
	}
}

// Get software by name and version
func (s *SoftwareBlueprintRepository) Get(category, name, version string) (software *entity.Software, err error) {
	// check is exist software blueprint
	err = s.checkSoftwareRepo()
	if err != nil {
		return nil, err
	}
	// check category
	if category == "" {
		category = DefaultRepo
	}
	// check version
	if version == "" {
		version, err = s.getSoftwareLastVersion(category, name)
		if err != nil {
			return nil, err
		}
		log.Infof("Get Software category=%s||name=%s||version=%s", category, name, version)
	}

	// find software
	softwareFilePath := filepath.Join(s.repoRootPath, category, name, version+YAMLSuffix)
	softwareFileByte, err := ioutil.ReadFile(softwareFilePath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(softwareFileByte, &software)
	if err != nil {
		return nil, err
	}
	return software, err
}

// Get software last version
func (s *SoftwareBlueprintRepository) getSoftwareLastVersion(category, name string) (version string, err error) {
	softwareVersionPath := filepath.Join(s.repoRootPath, category, name)
	files, err := ioutil.ReadDir(softwareVersionPath)
	if err != nil {
		return "", err
	}
	if len(files) == 0 {
		return "", entity.ErrorSoftwareFolderIsEmpty
	}
	var versionList []*semver.Version
	versionFileNameMap := make(map[*semver.Version]string)
	for _, file := range files {
		fileName := strings.TrimSuffix(file.Name(), YAMLSuffix)
		var v *semver.Version
		v, err = semver.NewVersion(fileName)
		if err != nil {
			return "", err
		}
		versionList = append(versionList, v)
		versionFileNameMap[v] = fileName
	}
	sort.Sort(semver.Collection(versionList))
	if versionList == nil {
		return "", entity.ErrorSoftwareVersionListIsNil
	}
	return versionFileNameMap[versionList[len(versionList)-1]], nil
}

// Check software blueprint.
// If not exist.git clone
// If exit git pull
func (s *SoftwareBlueprintRepository) checkSoftwareRepo() error {
	// read source_list
	conf := config.GetConfig()
	log.Infof("conf.SourceList %v", conf.SourceList)
	for sourceName, sourceGitUrl := range conf.SourceList {
		path := filepath.Join(s.repoRootPath, sourceName)
		var isExist bool
		isExist = s.isRepoFolderExist(path)
		log.Infof("check blueprint folder exist path:%s,result:%t", path, isExist)
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

// Check blueprint folder exist
func (s *SoftwareBlueprintRepository) isRepoFolderExist(path string) bool {
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
func (s *SoftwareBlueprintRepository) ReadConfigParam(sourceName string) (configParam map[string]interface{}, err error) {
	path := filepath.Join(s.repoRootPath, sourceName, ConfigParamFileName)
	configParamByte, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(configParamByte, &configParam)
	if err != nil {
		return nil, err
	}
	return configParam, err
}
