package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
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
	// Parse the query parameters from the URL
	query := r.URL.Query()

	// Retrieve individual query parameter values
	id := query.Get("id")
	date := query.Get("date")
	phoneNumber := query.Get("phone_number")
	branch := query.Get("branch")

	// URL decode values that may contain URL-encoded characters (e.g., Cyrillic text or spaces)
	decodedBranch, err := url.QueryUnescape(branch)
	if err != nil {
		log.Printf("Error decoding branch parameter: %v", err)
		http.Error(w, "Error processing request", http.StatusBadRequest)
		return
	}

	// Log the values for debugging (or use them as needed)
	log.Printf("Received ID: %s", id)
	log.Printf("Received Date: %s", date)
	log.Printf("Received Phone Number: %s", phoneNumber)
	log.Printf("Received Branch: %s", decodedBranch)

	// Respond back to the client
	fmt.Fprintf(w, "Received data - ID: %s, Date: %s, Phone Number: %s, Branch: %s", id, date, phoneNumber, decodedBranch)

	// Redirect or process further as required
	redirectURL := "https://b24-yeth0y.bitrix24site.ru/empty/"
	http.Redirect(w, r, redirectURL, http.StatusFound)
}
