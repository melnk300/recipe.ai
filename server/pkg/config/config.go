package config

import (
	"github.com/melnk300/recipe.ai/server/internal/configs"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var Config configs.Config

func Configure(path string) {
	configFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(configFile, &Config)
	if err != nil {
		log.Fatal(err)
	}
}
