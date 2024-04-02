package description

import (
	"bitrix_app/backend/bitrix/authorize"
	"bitrix_app/backend/bitrix/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func DescriptionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	pathSegments := strings.Split(r.URL.Path, "/")

	// Assuming the URL format is /api/documents/{ID}
	// The ID is expected to be the fourth segment, hence index 3 (0-based index)
	if len(pathSegments) < 4 {
		http.Error(w, "Invalid request path", http.StatusBadRequest)
		return
	}
	entityId := pathSegments[3]

	// Use the extracted ID as needed, for now, we'll just print it
	fmt.Println("Extracted ID DescriptionHandler: ", entityId)

	docs, err := GetDescription(authorize.GlobalAuthId, entityId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Directly write the mock JSON as the response
	_, err = w.Write(docs) // docs already contains the JSON
	if err != nil {
		log.Println("Error writing the mock response:", err)
	}
}

func GetDescription(authID string, dealId string) ([]byte, error) {
	method := "POST"
	bitrixMethod := "crm.deal.get"
	body := fmt.Sprintf(`{"id":"%s"}`, dealId)
	// Format the URL with the provided authID parameter
	requestUrl := fmt.Sprintf("https://b24-9f7fvg.bitrix24.ru/rest/%s?auth=%s", bitrixMethod, authID)
	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Println("error marshaling request body:", err)
		return nil, err
	}

	req, err := http.NewRequest(method, requestUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("error creating new request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading response body:", err)
		return nil, err
	}

	log.Println("resp_at_last_GetDescription:", string(bz))

	var apiResponse models.ApiResponse
	if err := json.Unmarshal(bz, &apiResponse); err != nil {
		log.Printf("error unmarshalling response to ApiResponse: %v", err)
		return nil, err
	}

	// If you need to return the data as JSON (for example, to send to another system or client),
	// you can re-marshal the ApiResponse struct back into JSON.
	jsonResponse, err := json.Marshal(apiResponse)
	if err != nil {
		log.Printf("error marshalling ApiResponse back to JSON: %v", err)
		return nil, err
	}

	return jsonResponse, nil
}
