package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func LoadConfig() {
	viper.SetConfigName(".env")

	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	viper.SetDefault("VAR_NAME", "default_value")

	// Getting a string from .env file
	loadedVar := viper.GetString("VAR_NAME")
	fmt.Printf("Loaded variable from .env: %s\n", loadedVar)
}
