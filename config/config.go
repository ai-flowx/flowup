package config

import (
	_ "embed"
)

var (
	Build   string
	Commit  string
	Version string
)

//go:embed config.yml
var ConfigData string

//go:embed env
var EnvData string

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
	Artifact Artifact `yaml:"artifact"`
}

type Artifact struct {
	Url  string `yaml:"url"`
	User int    `yaml:"user"`
	Pass int    `yaml:"pass"`
}
