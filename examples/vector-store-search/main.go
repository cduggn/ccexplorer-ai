package main

import (
	"context"
	ccexplorerai "github.com/cduggn/ccexplorerai"
)

func main() {
	cfg := LoadConfig()

	client, err := ccexplorerai.NewClient(
		ccexplorerai.WiitPineconeProjectName(cfg.PineconeProjectName),
		ccexplorerai.WithPineconeIndexName(cfg.PineconeIndexName),
		ccexplorerai.WithPineconeEnvironment(cfg.PineconeEnvironment),
		ccexplorerai.WithPineconeAPIKey(cfg.PineconeAPIKey),
		ccexplorerai.WithOpenAIKey(cfg.OpenAIKey),
	)
	if err != nil {
		panic(err)
	}

	client.LoadVectorStoreContext(context.Background(), "page_content")

	var scoreThreshold float32 = 0.75
	client.Search(context.Background(), "UnblendedCost for August 2023", 10, scoreThreshold)
}
