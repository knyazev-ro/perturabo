package utils

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Settings struct {
	Migrations string `yaml:"migrations"`
}

func DefaultSettings() *Settings {
	return &Settings{
		Migrations: "./migrations",
	}
}

func LoadSettings() *Settings {
	data, err := os.ReadFile("settings.yaml")
	if err != nil {
		println("Error during reading settings file.")
		println("Load default settings.")
		return DefaultSettings()
	}

	var settings Settings
	err = yaml.Unmarshal(data, &settings)
	if err != nil {
		log.Fatal("Error when unmarshaling yaml file:", err)
		println("Load default settings.")
		return DefaultSettings()
	}

	return &settings
}
