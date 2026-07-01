package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lmbangel/sop-rag-api/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	sText := "Go is a statically typed, compiled programming language designed at Google. It is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency. The language was designed by Robert Griesemer, Rob Pike, and Ken Thompson. Go was publicly announced in November 2009, and version 1.0 was released in March 2012. Go is widely used in production at Google and many other organizations and open-source projects. The language provides excellent support for concurrent programming through goroutines and channels, making it ideal for building scalable network services and distributed systems that handle thousands of simultaneous connections."

	chunks := services.ChunkText(sText, "rag-test", 200, 100)

	// for i, c := range chunks {
	// 	fmt.Printf("--- Chunk %d ---\n", i)
	// 	fmt.Printf("ID: %s\n", c.ID)
	// 	fmt.Printf("Source: %s\n", c.Source)
	// 	fmt.Printf("Length: %d characters\n", len([]rune(c.Text)))
	// 	fmt.Printf("Text: %s\n\n", c.Text)
	// }

	ctx := context.TODO()
	s := services.NewEmbeddingService(os.Getenv("OPENAI_API_KEY"))

	res, err := s.GenerateNewEmbeddings(chunks, ctx)
	if err != err {
		fmt.Println("Error: ", err.Error())
		return
	}
	fmt.Println(res)
}
