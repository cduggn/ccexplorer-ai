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
	PineconeEmbedder    string
	PineconeAPIKey      string
	PineconeNameSpace   string
	OpenAIKey           string
}

func LoadConfig() Config {
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	cfg := Config{
		PineconeProjectName: viper.GetString("PINECONE_PROJECT_NAME"),
		PineconeIndexName:   viper.GetString("PINECONE_INDEX_NAME"),
		PineconeEnvironment: viper.GetString("PINECONE_ENVIRONMENT"),
		PineconeEmbedder:    viper.GetString("PINECONE_EMBEDDER"),
		PineconeAPIKey:      viper.GetString("PINECONE_API_KEY"),
		PineconeNameSpace:   viper.GetString("PINECONE_NAMESPACE"),
		OpenAIKey:           viper.GetString("OPENAI_KEY"),
	}

	return cfg

}
