package main

import (
	"context"
	"github.com/cduggn/ccexplorerai"
)

func main() {

	cfg := LoadConfig()

	client, err := ccexplorer_ai.NewClient(
		ccexplorer_ai.WiitPineconeProjectName(cfg.PineconeProjectName),
		ccexplorer_ai.WithPineconeIndexName(cfg.PineconeIndexName),
		ccexplorer_ai.WithPineconeEnvironment(cfg.PineconeEnvironment),
		ccexplorer_ai.WithPineconeAPIKey(cfg.PineconeAPIKey),
		ccexplorer_ai.WithOpenAIKey(cfg.OpenAIKey),
	)
	if err != nil {
		panic(err)
	}

	client.LoadVectorStoreContext(context.Background())

	client.Search(context.Background(), "return any match", 0)
}
