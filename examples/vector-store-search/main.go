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

	client.LoadVectorStoreContext(context.Background(), "page_content")

	client.Search(context.Background(), "lambda costs in february", 0.75)
}
