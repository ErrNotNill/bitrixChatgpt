package docs

import (
	"bitrix_app/backend/bitrix/authorize"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func DocumentHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Extract ID from the path
	pathSegments := strings.Split(r.URL.Path, "/")

	// Assuming the URL format is /api/documents/{ID}
	// The ID is expected to be the fourth segment, hence index 3 (0-based index)
	if len(pathSegments) < 4 {
		http.Error(w, "Invalid request path", http.StatusBadRequest)
		return
	}
	id := pathSegments[3]

	log.Println("id: ", id)
	// Use the extracted ID as needed, for now, we'll just print it

	docs, err := GetDocsByDeal(authorize.GlobalAuthId, id)
	if err != nil {
		log.Println("Error: ", err)
	}
	fmt.Println("docs: ", docs)
	fmt.Fprintf(w, "Requested documents for ID: %s", docs)
}

type requestBody struct {
	Select []string `json:"select"`
	Filter struct {
		EntityTypeId string `json:"entityTypeId"`
		EntityId     string `json:"entityId"`
	} `json:"filter"`
}

func GetDocsByDeal(authID string, entityId string) ([]byte, error) {
	bitrixMethod := "crm.documentgenerator.document.list"
	requestUrl := fmt.Sprintf("https://b24-9f7fvg.bitrix24.ru/rest/%s?auth=%s", bitrixMethod, authID)

	// Construct the request body
	body := requestBody{
		Select: []string{"*"},
		Filter: struct {
			EntityTypeId string "json:\"entityTypeId\""
			EntityId     string "json:\"entityId\""
		}{
			EntityTypeId: "2",
			EntityId:     entityId, // Use the variable value
		},
	}

	// Marshal the request body into JSON
	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Println("error marshaling request body:", err)
		return nil, err
	}

	// Create a new request with JSON body
	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(jsonData)) // Switch to POST if applicable
	if err != nil {
		log.Println("error creating new request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read and return the response body
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading response body:", err)
		return nil, err
	}

	log.Println("Response:", string(responseData))

	return responseData, nil
}
