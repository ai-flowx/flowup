package config

import (
	_ "embed"
)

type Config struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	MetaData   MetaData `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

type MetaData struct {
	Name string `yaml:"name"`
}

type Spec struct {
	Drive Drive `yaml:"drive"`
}

type Drive struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

var (
	Build   string
	Commit  string
	Version string
)

//go:embed config.yml
var ConfigData string
