package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type GeneralConfig struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

type Server struct {
	Port string `yaml:"port"`
	Host string `yaml:"host,omitempty"`
}

type Database struct {
	Name     string `yaml:"name,omitempty"`
	User     string `yaml:"user,omitempty"`
	Password string `yaml:"password,omitempty"`
	Addr     string `yaml:"addr,omitempty"`
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