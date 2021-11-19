package fileops

import "github.com/spf13/afero"

func NewOSFs() *afero.Afero {
	osFs := afero.NewOsFs()
	return &afero.Afero{Fs:osFs}
}