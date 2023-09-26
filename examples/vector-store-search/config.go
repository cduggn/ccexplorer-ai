package main

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	PineconeProjectName string
	PineconeIndexName   string
	PineconeEnvironment string
	PineconeAPIKey      string
	PineconeNameSpace   string
	OpenAIKey           string
}

func LoadConfig() Config {
	viper.AutomaticEnv()
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	cfg := Config{
		PineconeProjectName: viper.GetString("pinecone.project-name"),
		PineconeIndexName:   viper.GetString("pinecone.index-name"),
		PineconeEnvironment: viper.GetString("pinecone.environment"),
		PineconeAPIKey:      viper.GetString("pinecone.api-key"),
		PineconeNameSpace:   viper.GetString("pinecone.namespace"),
		OpenAIKey:           viper.GetString("OPENAI_API_KEY"),
	}

	return cfg

}
