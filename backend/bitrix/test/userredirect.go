package test

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Feedback struct {
	Rating  string `json:"rating"` // Assuming rating is sent as a string; change to int if it's sent as a number
	Comment string `json:"comment"`
}

var UserGlobalId string

func UserForm(w http.ResponseWriter, r *http.Request) {
	// Read the entire request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading response body:", err)
		http.Error(w, "Error reading request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	log.Println("Received raw data:", string(body))

	// Parse the JSON body into the Feedback struct
	var feedback Feedback
	err = json.Unmarshal(body, &feedback)
	if err != nil {
		log.Println("error parsing JSON:", err)
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	// Log the parsed data
	log.Printf("Parsed feedback data: %+v\n", feedback)
	log.Printf("feedback.Rating: %s, feedback.Comment: %s", feedback.Rating, feedback.Comment)

	// Respond to the client to indicate success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Feedback received successfully"))
}

func UserRedirect(w http.ResponseWriter, r *http.Request) {
	// Since we're dealing with query parameters, reading the body might not be necessary unless you expect a POST request with additional data.
	queryParams := r.URL.Query()
	id := queryParams.Get("id")
	date := queryParams.Get("date")
	phoneNumber := queryParams.Get("phone_number")
	branch := queryParams.Get("branch")

	// Use these values as needed
	log.Printf("Received - ID: %s, Date: %s, Phone Number: %s, Branch: %s\n", id, date, phoneNumber, branch)

	// Redirect or process further as required
	redirectURL := "https://b24-yeth0y.bitrix24site.ru/empty/"
	http.Redirect(w, r, redirectURL, http.StatusFound)
}
