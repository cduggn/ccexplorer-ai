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

	filter := map[string]any{
		"$and": []map[string]interface{}{
			{
				"start": map[string]interface{}{
					"$eq": "2023-08-01",
				},
			},
			{
				"end": map[string]interface{}{
					"$eq": "2023-09-01",
				},
			},
		},
	}

	client.Search(context.Background(), "AWS WAF metrics", 10, scoreThreshold, filter)
}
