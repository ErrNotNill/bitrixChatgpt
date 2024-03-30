package events

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func OnCrmDealAddEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OnCrmDealAddEvent started: ")
	bs, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading request body:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
	log.Println("resp_at_first:", string(bs))
	defer r.Body.Close()

	values, err := url.ParseQuery(string(bs))
	if err != nil {
		panic(err)
	}

	// Convert to a map (manually handling nested structures)
	result := make(map[string]interface{})
	for key, value := range values {
		// Assuming only one value per key for simplicity
		result[key] = value[0]
	}

	// Converting the map to JSON
	jsonBytes, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonBytes))

	// To extract the ID, you can unmarshal the relevant part of the JSON back into a Go data structure
	// or manipulate the JSON directly if you prefer
	type DataFields struct {
		ID string `json:"data[FIELDS][ID]"`
	}

	var dataFields DataFields
	if err := json.Unmarshal(jsonBytes, &dataFields); err != nil {
		panic(err)
	}

	fmt.Println("Extracted ID:", dataFields.ID)

}

func OnCrmDealAddEventRegistration(authId string) {

	method := "POST"
	event := "OnCrmDealAdd"
	handler := "https://b24app.rwp2.com/api/event_deal_add"

	// Format the URL with the provided authID parameter
	requestUrl := fmt.Sprintf("https://b24-9f7fvg.bitrix24.ru/rest/event.bind.json?auth=%s&auth_type=0&event=%s&handler=%s", authId, event, handler)

	req, err := http.NewRequest(method, requestUrl, nil)
	if err != nil {
		log.Println("error creating new request:", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("error sending request:", err)
	}
	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading response body:", err)
	}
	log.Println("resp_at_last_AddDeal:", string(bz))

}

//func OnCrmDealUpdateEvent(w http.ResponseWriter, r *http.Request) {}
//func OnCrmDealDeleteEvent(w http.ResponseWriter, r *http.Request) {}
