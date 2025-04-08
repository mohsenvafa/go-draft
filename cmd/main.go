package main

import (
	"cmd-test/config"
	"core-shared/config/global"
	"core-shared/config/loader"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello, World!")

	var shared global.GlobalConfig
	var local config.Config

	err := loader.LoadConfig(
		&shared,
		&local,
		loader.NewYAMLLoader("../common/config/global.local.yml"),
		loader.NewYAMLLoader("config/private.local.yml"),
		loader.NewEnvLoader(),
	)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Printf("Shared Client Secret: %+v\n", shared.ClientSecret)
	fmt.Printf("Shared CLient Id: %+v\n", shared.ClientID)
	fmt.Printf("Local Config:  %+v\n", local.SchemaPath)
}
