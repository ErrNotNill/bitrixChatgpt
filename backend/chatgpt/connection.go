package chatgpt

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"net/http"
	"os"
)

func SendRequest(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("CHATGPT_API_KEY")
	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)

}
