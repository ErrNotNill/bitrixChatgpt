package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "POST" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("error reading response body:", err)
			http.Error(w, "Error reading request", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var feedback Feedback
		err = json.Unmarshal(body, &feedback)
		if err != nil {
			log.Println("error parsing JSON:", err)
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}

		log.Printf("Parsed feedback data: %+v\n", feedback)
		log.Printf("feedback.Rating: %s, feedback.Comment: %s", feedback.Rating, feedback.Comment)

		// Lookup for Rating
		ratingMap := map[string]int{
			"1":  857,
			"2":  859,
			"3":  861,
			"4":  863,
			"5":  865,
			"6":  867,
			"7":  869,
			"8":  871,
			"9":  873,
			"10": 875,
		}
		numericRating, exists := ratingMap[feedback.Rating]
		if !exists {
			log.Printf("Invalid rating value: %s", feedback.Rating)
		}

		// Get Deal information and branch mapping
		apiResponse, err := GetDealById(DealGlobalId)
		if err != nil {
			log.Println("Error getting deal info")
			http.Error(w, "Failed to get deal info", http.StatusInternalServerError)
			return
		}
		//C17:NEW НОВАЯ
		//C17:UC_LRSZUX ОТРИЦ (1-5)
		//C17:1 НЕЙТР (6-8)
		//C17:UC_NBQJD0 ПОЛОЖ (9-10)
		stageMap := map[string]string{
			"1":  "C17:UC_LRSZUX",
			"2":  "C17:UC_LRSZUX",
			"3":  "C17:UC_LRSZUX",
			"4":  "C17:UC_LRSZUX",
			"5":  "C17:UC_LRSZUX",
			"6":  "C17:1",
			"7":  "C17:1",
			"8":  "C17:1",
			"9":  "C17:UC_NBQJD0",
			"10": "C17:UC_NBQJD0",
		}
		stageValue, exists := stageMap[feedback.Rating]
		if !exists {
			log.Printf("Invalid rating value: %s", feedback.Rating)
		}

		// Assuming CreateDeal handles the error internally and logs as needed
		err = CreateDeal(feedback.Comment, "17", fmt.Sprintf("https://harizma.bitrix24.ru/crm/deal/details/%s/", DealGlobalId),
			apiResponse.Result.ContactID, apiResponse.Result.Branch, numericRating, apiResponse.Result.DateCreate, stageValue)
		if err != nil {
			log.Println("CreateDeal failed")
			http.Error(w, "Failed to create deal", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Feedback received successfully"))
	} else {
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
