package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type GeneralConfig struct {
	Server Server `yaml:"server"`
}

type Server struct {
	Port    string  `yaml:"port"`
	Host    string  `yaml:"host,omitempty"`
	Service Service `yaml:"service"`
}

type Service struct {
	URL string `json:"url,omitempty"`
}

func NewGeneralConfig(fPath string) GeneralConfig {
	var fName, _ = filepath.Abs(fPath)
	var yamlFile, err = ioutil.ReadFile(fName)

	if err != nil {
		panic(err)
	}

	var c GeneralConfig
	err = yaml.Unmarshal(yamlFile, &c)

	return c
}
