package main

import (
  "context"
  "fmt"
  "log"
  "os"

  "github.com/google/generative-ai-go/genai"
  "github.com/joho/godotenv"
  "google.golang.org/api/iterator"
  "google.golang.org/api/option"
) 

func main() {

  err := godotenv.Load(".env")
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  ctx := context.Background()
  // Access your API key as an environment variable (see "Set up your API key" above)
  client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
  if err != nil {
    log.Fatal(err)
  }
  defer client.Close()

  model := client.GenerativeModel("gemini-1.5-flash")

  if len(os.Args) <= 2 {
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

    resp := response.Candidates[0].Content.Parts
    fmt.Println(resp)
  }
}
