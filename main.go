package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	openai "github.com/openai/openai-go/v1"
)

func main() {
	// Retrieve OpenAI API token from environment variable
	apiKey := os.Getenv("OPENAI_TOKEN")
	if apiKey == "" {
		fmt.Println("Error: OPENAI_TOKEN environment variable is not set.")
		return
	}

	// Initialize OpenAI client
	client := openai.NewClient(apiKey)

	// Prompt the user for input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your query: ")
	query, _ := reader.ReadString('\n')
	query = strings.TrimSpace(query)

	// Call OpenAI API to get response
	response, err := queryOpenAI(client, query)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the response
	fmt.Println("Response:")
	fmt.Println(response)
}

func queryOpenAI(client *openai.Client, query string) (string, error) {
	// Set up the completion parameters
	completionRequest := &openai.CompletionRequest{
		Model:       "text-davinci-002", // GPT-3.5 model
		MaxTokens:   50,                 // Limit response to 50 tokens (adjust as needed)
		Temperature: 0.7,
		Prompt:      "Terminal command: " + query + "\nDescription:", // Prompt with the query
	}

	// Call the completion endpoint
	completion, err := client.CreateCompletion(completionRequest)
	if err != nil {
		return "", err
	}

	// Extract and return the response text
	return completion.Choices[0].Text, nil
}
