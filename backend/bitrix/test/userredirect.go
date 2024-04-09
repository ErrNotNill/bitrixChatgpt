package test

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Feedback struct {
	Rating  string `json:"rating"` // Make sure the JSON tags match the keys in your JSON object.
	Comment string `json:"comment"`
}

var UserGlobalId string

func UserForm(w http.ResponseWriter, r *http.Request) {
	// Handle preflight request for CORS
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")                   // Allow any origin
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS") // Allowed methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       // Allow Content-Type header
		w.WriteHeader(http.StatusOK)
		return
	}

	// Main handling for POST request
	if r.Method == "POST" {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow any origin
		// Reading the request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("error reading response body:", err)
			http.Error(w, "Error reading request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		log.Println("Received raw data:", string(body))

		// Parsing the JSON body into the Feedback struct
		var feedback Feedback
		err = json.Unmarshal(body, &feedback)
		if err != nil {
			log.Println("error parsing JSON:", err)
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}

		// Logging the parsed data
		log.Printf("Parsed feedback data: %+v\n", feedback)
		log.Printf("feedback.Rating: %s, feedback.Comment: %s", feedback.Rating, feedback.Comment)

		// Respond to the client to indicate success
		w.WriteHeader(http.StatusOK) // This is now only set once in this block
		w.Write([]byte("Feedback received successfully"))
	} else {
		// If not OPTIONS or POST, inform the client that the method is not allowed
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
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

	// Redirect or process further as required
	redirectURL := "https://b24-yeth0y.bitrix24site.ru/empty/"
	http.Redirect(w, r, redirectURL, http.StatusFound)
}
