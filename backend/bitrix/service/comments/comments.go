package comments

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

func CommentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Assuming the URL format is /api/documents/{ID}
	// The ID is expected to be the fourth segment, hence index 3 (0-based index)

	rdr, _ := io.ReadAll(r.Body)
	fmt.Println("rdr:____: ", rdr)

	pathSegments := strings.Split(r.URL.Path, "/")

	// Assuming the URL format is /api/documents/{ID}
	// The ID is expected to be the fourth segment, hence index 3 (0-based index)
	if len(pathSegments) < 4 {
		http.Error(w, "Invalid request path", http.StatusBadRequest)
		return
	}
	entityId := pathSegments[3]

	// Use the extracted ID as needed, for now, we'll just print it
	fmt.Println("Extracted ID CommentsHandler: ", entityId)

	comments, err := GetComments(authorize.GlobalAuthId, entityId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(string(comments))
	// Directly write the mock JSON as the response
	_, err = w.Write(comments) // comments already contains the JSON
	if err != nil {
		log.Println("Error writing the mock response:", err)
	}
}

func GetComments(authID string, entityId string) ([]byte, error) {
	bitrixMethod := "crm.timeline.comment.list"

	requestUrl := fmt.Sprintf("https://b24-9f7fvg.bitrix24.ru/rest/%s?auth=%s", bitrixMethod, authID)

	// Construct the request body
	body := fmt.Sprintf(`{"filter": {
			"ENTITY_ID": %s,
			"ENTITY_TYPE": "deal"
		},
		"select": [ "ID", "COMMENT"]
}`, entityId)

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

	log.Println("GetComments Response:", string(responseData))

	return responseData, nil
}
