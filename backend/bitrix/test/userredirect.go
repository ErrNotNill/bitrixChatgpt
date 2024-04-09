package test

import (
	"io"
	"log"
	"net/http"
)

func UserRedirect(w http.ResponseWriter, r *http.Request) {
	// Read the body content (if needed)
	bs, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading response body:", err)
	}
	log.Println("resp_UserRedirect:", string(bs))

	// Parse query parameters
	queryParams := r.URL.Query()
	id := queryParams.Get("id") // This will extract the value of "id" parameter

	// Check if "id" parameter exists and is not empty
	if id != "" {
		log.Println("Received ID:", id)
		// Here, you can use the "id" to perform further actions
	} else {
		log.Println("No ID provided")
	}

	// Example response to send back
	w.Write([]byte("Received ID: " + id))
}
