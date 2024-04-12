package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Feedback struct {
	Rating  string `json:"rating"` // Make sure the JSON tags match the keys in your JSON object.
	Comment string `json:"comment"`
}

var DealGlobalId string
var RatingGlobalText string
var CommentGlobalText string
var DateGlobal string
var PhoneNumberGlobal string
var BranchGlobal string

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

		RatingGlobalText = feedback.Rating
		CommentGlobalText = feedback.Comment
		// Respond to the client to indicate success
		w.WriteHeader(http.StatusOK) // This is now only set once in this block
		w.Write([]byte("Feedback received successfully"))

		category := os.Getenv("CATEGORY_DEAL_HARIZMA_DEAL_ADD")
		link := fmt.Sprintf(`https://harizma.bitrix24.ru/crm/deal/details/%s/`, DealGlobalId)

		apiResponse, err := GetDealById(DealGlobalId)
		if err != nil {
			log.Println("Error getting deal info")
		}

		var contactId, branch string
		var dateCreate time.Time
		for _, v := range apiResponse.Result {
			contactId, branch = v.ContactID, v.Branch
			dateCreate = v.DateCreate
		}
		err = CreateDeal(feedback.Comment, category, link, contactId, branch, feedback.Rating, dateCreate)
		if err != nil {
			log.Println("CreateDeal failed")
		}

	} else {
		// If not OPTIONS or POST, inform the client that the method is not allowed
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

}

type ResponseData struct {
	DealID      string `json:"deal_id"`
	Rating      string `json:"rating"`
	Comment     string `json:"comment"`
	Date        string `json:"date"`
	PhoneNumber string `json:"phone_number"`
	Branch      string `json:"branch"`
}

func SendJsonInGoogle(w http.ResponseWriter, r *http.Request) {
	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Populate data structure with your global variables

	data := ResponseData{
		DealID:      "STR",
		Rating:      "SSSSS",
		Comment:     "ZXCZXC",
		Date:        "ASDADSAD",
		PhoneNumber: "QWEQEQEWQE",
		Branch:      "GDSGDSGDSGSGDSG",
	}

	// Encode data to JSON and send as response
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

func UserRedirect(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameters from the URL
	query := r.URL.Query()

	// Retrieve individual query parameter values
	id := query.Get("id")

	// Log the values for debugging (or use them as needed)
	log.Printf("Received ID: %s", id)

	DealGlobalId = id

	// Redirect or process further as required
	redirectURL := "https://b24-yeth0y.bitrix24site.ru/empty/"
	http.Redirect(w, r, redirectURL, http.StatusFound)
}
