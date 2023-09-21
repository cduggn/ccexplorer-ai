package main

import (
	"context"
	"github.com/cduggn/ccexplorerai"
)

func main() {

	client, err := ccexplorer_ai.NewClient(
		ccexplorer_ai.WiitPineconeProjectName("ccexplorer"),
		ccexplorer_ai.WithPineconeIndexName("ccexplorer"),
		ccexplorer_ai.WithPineconeEnvironment("production"),
		ccexplorer_ai.WithPineconeEmbedder("tfhub/universal-sentence-encoder-multilingual-large/3"),
		ccexplorer_ai.WithPineconeAPIKey(""),
		ccexplorer_ai.WithPineconeNameSpace("ccexplorer"),
		ccexplorer_ai.WithOpenAIKey(""),
	)

	client.LoadContext(context.Background())

}
