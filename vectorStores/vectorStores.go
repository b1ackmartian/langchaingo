package vectorStores

import "github.com/tmc/langchaingo/schema"

type VectorStore interface {
	AddDocuments(documents []schema.Document) error
	SimilaritySearch(query string, numDocuments int) ([]schema.Document, error)
}
