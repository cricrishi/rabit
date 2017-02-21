package config

import (
	"github.com/jacobstr/confer"
)

var ConfigData map[string]*confer.Config

func LoadConfig(filePath string, aliasName string) *confer.Config {
	if len(ConfigData) == 0 {
		ConfigData = make(map[string]*confer.Config, 0)
	}
	ConfigData[aliasName] = confer.NewConfig()
	ConfigData[aliasName].ReadPaths(filePath)
	return ConfigData[filePath]
}
