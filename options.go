package ccexplorer_ai

import "errors"

type Option func(p *Config)

func WiitPineconeProjectName(projectName string) Option {
	return func(p *Config) {
		p.PineconeProjectName = projectName
	}
}

func WithPineconeIndexName(indexName string) Option {
	return func(p *Config) {
		p.PineconeIndexName = indexName
	}
}

func WithPineconeEnvironment(env string) Option {
	return func(p *Config) {
		p.PineconeEnvironment = env
	}
}

func WithPineconeAPIKey(apiKey string) Option {
	return func(p *Config) {
		p.PineconeAPIKey = apiKey
	}
}

func WithPineconeNameSpace(nameSpace string) Option {
	return func(p *Config) {
		p.PineconeNameSpace = nameSpace
	}
}

func WithOpenAIKey(openAIKey string) Option {
	return func(p *Config) {
		p.OpenAIKey = openAIKey
	}
}

func applyOptions(opts ...Option) (*Config, error) {
	options := &Config{}
	for _, opt := range opts {
		opt(options)
	}

	if options.PineconeProjectName == "" {
		return nil, errors.New("pinecone project name is required")
	}

	if options.PineconeIndexName == "" {
		return nil, errors.New("pinecone index name is required")
	}

	if options.PineconeEnvironment == "" {
		return nil, errors.New("pinecone environment is required")
	}

	if options.PineconeAPIKey == "" {
		return nil, errors.New("pinecone api key is required")
	}

	if options.PineconeNameSpace == "" {
		return nil, errors.New("pinecone name space is required")
	}

	if options.OpenAIKey == "" {
		return nil, errors.New("openai key is required")
	}

	return options, nil
}
