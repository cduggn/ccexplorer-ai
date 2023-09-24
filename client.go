package ccexplorer_ai

import (
	"context"
	embedings "github.com/tmc/langchaingo/embeddings/openai"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/pinecone"
	"log/slog"
	"os"
)

type Client struct {
	config Config
	store  pinecone.Store
	logger *slog.Logger
}

func NewClient(opts ...Option) (*Client, error) {
	config, err := ClientConfig(opts...)
	if err != nil {
		return nil, err
	}

	return &Client{
		config: *config,
		logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}, nil

}

func (c *Client) LoadVectorStoreContext(ctx context.Context) {
	embedder, err := embedings.NewOpenAI()
	if err != nil {
		c.logger.Error(err.Error())
	}

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
		c.logger.Error(err.Error())
	}

	c.store = store
}

func (c *Client) Search(ctx context.Context, q string) {
	if q == "" {
		c.logger.Error("query is empty")
	}
	// Search for similar documents using score threshold.
	docs, err := c.store.SimilaritySearch(ctx, q, 10, vectorstores.WithScoreThreshold(0.80))
	if err != nil {
		c.logger.Error(err.Error(), "query", q, "docs", docs)
	}

	if docs != nil {
		c.logger.Info("docs returned", "docs", docs)
	} else {
		c.logger.Info("no docs returned")
	}
}
