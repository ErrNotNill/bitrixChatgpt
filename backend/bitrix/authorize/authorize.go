package authorize

import (
	"encoding/json"
	"fmt"
	goBX24 "github.com/whatcrm/go-bitrix24"
	"io"
	"log"
	"net/http"
	"os"
)

func AuthorizeBitrix() error {
	clientID := os.Getenv("BITRIX_CLIENT_ID")
	clientSecret := os.Getenv("BITRIX_CLIENT_SECRET")
	domain := os.Getenv("BITRIX_DOOMAIN")
	auth := "auth"

	b24 := goBX24.NewAPI(clientID, clientSecret)

	if err := b24.SetOptions(domain, auth, true); err != nil {
		return err
	}

	admin, _ := b24.IsAdmin()
	log.Println(admin.Result)

	dealId := "43"
	res, err := b24.Get().Deals(dealId)
	if err != nil {
		return err
	}
	log.Println("result: ", res)
	return nil
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
