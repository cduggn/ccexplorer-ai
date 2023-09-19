package ccexplorer_ai

const (
	opeanAPIURLV1 = "https://api.openai.com/v1"
	openAIModel   = "text-embedding-ada-002"
)

type Pinecone struct {
	ProjectName string
	IndexName   string
	Environment string
	Embedder    string
	APIKey      string
	NameSpace   string
}

type ClientConfig struct {
	VectorStore Pinecone
}

func DefaultClientConfig() ClientConfig {
	return ClientConfig{
		VectorStore: Pinecone{
			ProjectName: "YOUR_PROJECT_NAME",
			IndexName:   "YOUR_INDEX_NAME",
			Environment: "",
			Embedder:    "",
			APIKey:      "",
			NameSpace:   "",
		},
	}
}
