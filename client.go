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

func (c *Client) LoadVectorStoreContext(ctx context.Context, textKey string) {
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
		pinecone.WithNameSpace(c.config.PineconeNameSpace),
		pinecone.WithTextKey(textKey),
	)
	if err != nil {
		c.logger.Error(err.Error())
	}

	c.store = store
}

func (c *Client) Search(ctx context.Context, query string, numDocuments int, scoreThreshold float32) {
	if query == "" {
		c.logger.Error("query is empty")
	}
	// Search for similar documents using score threshold.
	docs, err := c.store.SimilaritySearch(ctx, query, numDocuments, vectorstores.WithScoreThreshold(scoreThreshold))
	if err != nil {
		c.logger.Error(err.Error(), "query", query, "docs", docs)
	}

	if docs != nil {
		for _, doc := range docs {
			c.logger.Info("doc", "doc", doc.PageContent)
		}
	} else {
		c.logger.Info("no docs returned")
	}
}
