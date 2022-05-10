package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
)

const projectDirName = "DislinktXWS-back"

func LoadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentDirectory))

	err := godotenv.Load(string(rootPath) + `/.dev.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
