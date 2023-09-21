package ccexplorer_ai

const (
	opeanAPIURLV1 = "https://api.openai.com/v1"
	openAIModel   = "text-embedding-ada-002"
)

type Config struct {
	PineconeProjectName string
	PineconeIndexName   string
	PineconeEnvironment string
	PineconeEmbedder    string
	PineconeAPIKey      string
	PineconeNameSpace   string
	OpenAIKey           string
}

func ClientConfig(opts ...Option) (*Config, error) {

	o, err := applyOptions(opts...)
	if err != nil {
		return &Config{}, err
	}

	return o, nil
}
