package services

import (
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
