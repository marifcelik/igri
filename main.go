package main

import (
	projectUtils "go-chat/utils"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

func init() {
	if projectUtils.GetAppEnv() != "prod" {
		if err := godotenv.Load(); err != nil {
			log.Error(err)
		}
	}
}

func main() {
	log.Info("main func")
	log.Fatal(nil)
}
