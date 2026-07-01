package services

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

type QdrantService struct {
	client         *qdrant.Client
	collectionName string
}

func NewQdrantService(host string, port int, colName string) (*QdrantService, error) {
	c, err := qdrant.NewClient(&qdrant.Config{
		Host: host,
		Port: port,
	})
	if err != nil {
		return nil, err
	}
	q := &QdrantService{
		client:         c,
		collectionName: colName,
	}
	return q, nil
}

func (q *QdrantService) CreateCollection(vectorSize uint64) error {
	return q.client.CreateCollection(context.Background(), &qdrant.CreateCollection{
		CollectionName: q.collectionName,
		VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
			Size:     vectorSize,
			Distance: qdrant.Distance_Cosine,
		}),
	})
}

func (s *QdrantService) Upset(ctx context.Context, chunks []Chunk, embeddings [][]float64) error {
	return nil
}
