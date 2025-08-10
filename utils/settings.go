package utils

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Settings struct {
	Migrations    string `yaml:"migrations"`
	SysMigrations string `yaml:"sysmigrations"`

	TemplateCreate string `yaml:"template-create"`
	TemplateAlter  string `yaml:"template-alter"`
}

func DefaultSettings() *Settings {
	return &Settings{
		Migrations:    "./migrations",
		SysMigrations: "./sysmigrations",

		TemplateCreate: "./templates/create.tmpl",
		TemplateAlter:  "./templates/alter.tmpl",
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
