package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
) 

func main() {

  ctx := context.Background()
  // Access your API key as an environment variable (see "Set up your API key" above)
  client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
  if err != nil {
    log.Fatal(err)
  }
  defer client.Close()

  model := client.GenerativeModel("gemini-1.5-flash")

  if len(os.Args) <= 1 {
    fmt.Println("Please provide a prompt")
    os.Exit(1)
  }

  userPrompt := "You are a helpful terminal assistant that helps the user with anything related to the terminal and bash or zsh please provide the user with a helpful and very concise response. The user prompt is: " + os.Args[1]


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

    out, err := json.Marshal(response.Candidates[0].Content.Parts)
    if err != nil {
      log.Fatal(err)
    }
    cleanedResponse := strings.Replace(string(out), "\\n", "\n", -1)
    fmt.Println(strings.Trim(cleanedResponse, "[]`\"\\/*"))

  }
}
