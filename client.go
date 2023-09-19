package ccexplorer_ai

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/pinecone"
	"log"
)

type Client struct {
	cfg *ClientConfig
}

func NewClient() *Client {
	return &Client{
		cfg: &ClientConfig{},
	}
}

func (c *Client) LoadContext() {

	// Create a new Pinecone vector store.
	store, err := pinecone.New(
		ctx,
		pinecone.WithProjectName("YOUR_PROJECT_NAME"),
		pinecone.WithIndexName("YOUR_INDEX_NAME"),
		pinecone.WithEnvironment("YOUR_ENVIRONMENT"),
		pinecone.WithEmbedder(e),
		pinecone.WithAPIKey("YOUR_API_KEY"),
		pinecone.WithNameSpace(uuid.New().String()),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Search for similar documents using score threshold.
	docs, err = store.SimilaritySearch(ctx, "only cities in south america", 10, vectorstores.WithScoreThreshold(0.80))
	fmt.Println(docs)

}
