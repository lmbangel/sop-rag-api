package services

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
	return cs
}
