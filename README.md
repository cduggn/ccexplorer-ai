# ccexplorer-ai
This is a repository that provides AI augmented search and discovery capabilities for AWS cost and usage data produced by the [`ccExplorer`](https://github.com/cduggn/ccExplorer) CLI . 

## Description
`ccExplorer-ai` provides a simple interface to query cost explorer data stored in [Pinecone](https://www.pinecone.io/) using natural language queries. It relies on the [langchaingo](https://github.com/tmc/langchaingo) to load
and process the cost explorer data and the [OpenAI](https://openai.com/) API to generate natural language descriptions of the cost explorer data.

## Usage 

```go
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

```

See [`examples`](./examples) for usage examples.