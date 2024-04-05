package chatgpt

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
)

const (
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)

func SendRequest() {
	apiKey := os.Getenv("CHATGPT_API_KEY")

	fmt.Println("apiKey: ", apiKey)
	client := resty.New()
	response, err := client.R().
		SetAuthToken(apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model":      "gpt-3.5-turbo",
			"messages":   []interface{}{map[string]interface{}{"role": "system", "content": "Hi can you tell me what is the factorial of 10?"}},
			"max_tokens": 50,
		}).
		Post(apiEndpoint)

	if err != nil {
		log.Fatalf("Error while sending the request: %v", err)
	}

	// Log the raw response body for debugging
	log.Printf("Raw response body: %s\n", response)

	// Proceed with your existing unmarshalling and data handling logic...
}
