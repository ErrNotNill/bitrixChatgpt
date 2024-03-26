package authorize

import (
	"encoding/json"
	"fmt"
	"github.com/ikarpovich/go-bitrix/client"
	"github.com/ikarpovich/go-bitrix/types"
	"io"
	"log"
	"net/http"
	"os"
)

func BitrixAuthorize() {
	c, err := client.NewEnvClientWithWebhookAuth()

	if err != nil {
		log.Fatalf("Can't create client: %s", err)
	}

	c.SetInsecureSSL(true)
	c.SetDebug(true)

	resp, err := c.Methods(&types.MethodsRequest{
		Full:  true,
		Scope: "landing",
	})

	if err != nil {
		log.Fatalf("Request error: %s", err)
	}

	log.Print(resp)
}

func BotBitrix(w http.ResponseWriter, r *http.Request) {
	BitrixClientId := os.Getenv("BITRIX_CLIENT_ID")

	client := &http.Client{}

	tokenURL := fmt.Sprintf(`https://onviz.bitrix24.ru/oauth/authorize/?client_id=%s`,
		BitrixClientId)

	post, err := client.Post(tokenURL, "application/x-www-form-urlencoded", nil)
	if err != nil {
		fmt.Println("Failed to exchange authorization code for access token:", err)
		return
	}

	body, err := io.ReadAll(post.Body)
	fmt.Println("post.Body", post.Body)
	json.Unmarshal(body, &post.Body)
	fmt.Fprint(w, string(body))
}
