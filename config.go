package ccexplorer_ai

type Config struct {
	PineconeProjectName string
	PineconeIndexName   string
	PineconeEnvironment string
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
