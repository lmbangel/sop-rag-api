package services

import (
	"context"
	_ "context"
	_ "fmt"
	_ "os"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	_ "github.com/openai/openai-go/v3/responses"
)

type EmbeddingService struct {
	Client *openai.Client
}

func NewEmbeddingService(apiKey string) *EmbeddingService {
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)
	return &EmbeddingService{Client: &client}
}

func (s *EmbeddingService) GenerateNewEmbeddings(chunks []Chunk, ctx context.Context) ([][]float64, error) {
	var ems [][]float64

	for _, c := range chunks {
		resp, err := s.Client.Embeddings.New(ctx, openai.EmbeddingNewParams{
			Model: openai.EmbeddingModelTextEmbedding3Small,
			Input: openai.EmbeddingNewParamsInputUnion{OfString: openai.String(c.Text)},
		})
		if err != nil {
			return nil, err
		}
		ems = append(ems, resp.Data[0].Embedding)
	}

	return ems, nil
}
