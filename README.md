# ccexplorer-ai
This is a repository that provides AI augmented search and discovery capabilities for AWS cost and usage data surfaced through [`ccExplorer`](https://github.com/cduggn/ccExplorer) CLI . 

## Description
`ccExplorer-ai` provides a simple interface to query cost explorer data stored in [Pinecone](https://www.pinecone.io/) using natural language queries. It relies on the [langchaingo](https://github.com/tmc/langchaingo) to load
and process the cost explorer data and the [OpenAI](https://openai.com/) API to generate natural language descriptions of the cost explorer data.

## Usage 

```go
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

	client.Search(context.Background(), "firewall costs", 10, scoreThreshold, filter)
}

```

See [`examples`](./examples) for usage examples.