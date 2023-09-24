package ccexplorer_ai

import (
	"context"
	"fmt"
	embedings "github.com/tmc/langchaingo/embeddings/openai"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/pinecone"
	"log"
)

type Client struct {
	config Config
	store  pinecone.Store
}

func NewClient(opts ...Option) (*Client, error) {
	config, err := ClientConfig(opts...)
	if err != nil {
		return nil, err
	}

	return &Client{
		config: *config,
	}, nil

}

func (c *Client) LoadVectorStoreContext(ctx context.Context) {

	openai.New()

	embedder, err := embedings.NewOpenAI()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Pinecone vector store.
	store, err := pinecone.New(
		ctx,
		pinecone.WithProjectName(c.config.PineconeProjectName),
		pinecone.WithIndexName(c.config.PineconeIndexName),
		pinecone.WithEnvironment(c.config.PineconeEnvironment),
		pinecone.WithEmbedder(embedder),
		pinecone.WithAPIKey(c.config.PineconeAPIKey),
		pinecone.WithNameSpace(c.config.PineconeIndexName),
	)
	if err != nil {
		log.Fatal(err)
	}

	c.store = store
}

func (c *Client) Search(ctx context.Context, q string) {
	// Search for similar documents using score threshold.
	docs, err := c.store.SimilaritySearch(ctx, q, 10, vectorstores.WithScoreThreshold(0.80))
	if err != nil {
		fmt.Println(err)
	}

	if docs != nil {
		fmt.Println(docs)
	} else {
		log.Println("No documents found")
	}
}
