# ccexplorer-ai
This is a repository that provides AI augmented search and discovery capabilities for AWS cost and usage data produced by the [`ccExplorer`](https://github.com/cduggn/ccExplorer) CLI . 

## Description
`ccExplorer-ai` relies on vectorised data produced by [`ccExplorer`](https://github.com/cduggn/ccExplorer) and pushed to [Pinecone](https://www.pinecone.io/). Search queries are augmented with context from Pinecone before being passed to an AI model for answer generation.
