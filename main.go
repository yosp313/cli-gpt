package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
) 

func main() {

  ctx := context.Background()
  // Access your API key as an environment variable (see "Set up your API key" above)
  client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyBji5cpv-O-osEne7l9LD3ozvV_GtyxxlA"))
  if err != nil {
    log.Fatal(err)
  }
  defer client.Close()

  model := client.GenerativeModel("gemini-1.5-flash")
  

  userPrompt := os.Args[1]
  // Generate a text response
  iter := model.GenerateContentStream(ctx, genai.Text(userPrompt))
  for {
    response, err := iter.Next()
    if err == iterator.Done{
      break
    }
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(response.Candidates[0].Content)
  }
}
