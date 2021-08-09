package lib

import (
	"encoding/json"
	"io/ioutil"
)

type PackageJson struct {
	Path         string
	Name         string
	Version      string
	Scu          Config
	Dependencies map[string]string
}

func NewPackageJson(path string) PackageJson {
	var p PackageJson
	var dat, _ = ioutil.ReadFile(path)
	json.Unmarshal(dat, &p)
	p.Path = path
	return p
}

func (p PackageJson) GetConfig() Config {
	return p.Scu
}
