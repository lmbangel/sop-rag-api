package services

import "fmt"

type Chunk struct {
	ID     string
	Text   string
	Source string
}

func ChunkText(text string, source string, chunkSize int, overlap int) []Chunk {
	var cs []Chunk
	r := []rune(text)
	len := len(r)
	if len == 0 {
		return cs
	}
	start := 0
	chunkIndex := 0

	for start < len {
		end := start + chunkSize
		if end > len {
			end = len
		}

		c := string(r[start:end])
		cs = append(cs, Chunk{
			ID:     fmt.Sprintf("%s-chunk-%d", source, chunkIndex),
			Text:   c,
			Source: source,
		})

		chunkIndex++
		start = start + (chunkSize - overlap)

		if start >= len {
			break
		}
	}

	return cs
}
