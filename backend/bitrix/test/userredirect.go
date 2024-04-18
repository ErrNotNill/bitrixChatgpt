package test

import (
	"bitrix_app/backend/bitrix/test/spreadsheets"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Feedback struct {
	Id      string `json:"id"` // Add this line
	Rating  string `json:"rating"`
	Comment string `json:"comment"`
}

var CountGetUrl = 68

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

		log.Printf("Parsed feedback ID: %+v\n", feedback.Id)
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
		apiResponse, err := GetDealById(feedback.Id)
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
		urlDeal := fmt.Sprintf("https://harizma.bitrix24.ru/crm/deal/details/%s/", feedback.Id)

		// Assuming CreateDeal handles the error internally and logs as needed
		err = CreateDeal(feedback.Comment, "17", urlDeal,
			apiResponse.Result.ContactID, apiResponse.Result.Branch, numericRating, apiResponse.Result.DateCreate, stageValue)
		if err != nil {
			log.Println("CreateDeal failed")
			http.Error(w, "Failed to create deal", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		var branchConvertedToText string
		//sheet := spreadsheets.GoogleSheetsUpdate()
		if apiResponse.Result.Branch == "471" {
			branchConvertedToText = "м. Бауманская, ул. Бакунинская 32/36 с1"
		} else if apiResponse.Result.Branch == "473" {
			branchConvertedToText = "м. Лубянка, ул. Сретенский переулок, 4"
		} else if apiResponse.Result.Branch == "475" {
			branchConvertedToText = "м. Молодежная, Рублёвское шоссе, 28к1"
		} else if apiResponse.Result.Branch == "477" {
			branchConvertedToText = "м. Сухаревская, Москва, ул. Сретенка, 30"
		} else if apiResponse.Result.Branch == "764" {
			branchConvertedToText = "м. Новослободская,  ул. Новослободская, 20с6"
		}

		CountUserRequests++
		spreadsheets.GoogleSheetsUpdate(CountGetUrl, 1, feedback.Rating)        //оценка
		spreadsheets.GoogleSheetsUpdate(CountGetUrl, 2, feedback.Comment)       //комментарий
		spreadsheets.GoogleSheetsUpdate(CountGetUrl, 3, urlDeal)                //ссылка на сделку
		spreadsheets.GoogleSheetsUpdate(CountGetUrl, 4, branchConvertedToText)  //филиал
		spreadsheets.GoogleSheetsUpdate(1, 10, strconv.Itoa(CountUserRequests)) //ответов по ссылке

		//w.Write([]byte("Feedback received successfully"))
		//http.Redirect(w, r, "https://b24-yeth0y.bitrix24site.ru/empty_jekf/", http.StatusFound)
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

var CountUserRedirect = 7  //переходов по ссылке
var CountUserRequests = 1  //ответов по ссылке
var CountSendedSms = 15    //сообщение отправлено
var CountSendedDoneSms = 3 //сообщение не отправилось

var store = sessions.NewCookieStore([]byte(GenerateSecretKey(32)))

func UserRedirect(w http.ResponseWriter, r *http.Request) {
	CountGetUrl++
	CountUserRedirect++

	session, _ := store.Get(r, "session-name")
	query := r.URL.Query()
	id := query.Get("id")
	session.Values["dealId"] = id
	session.Save(r, w)

	// Include the ID in the redirect URL as a query parameter
	redirectURL := fmt.Sprintf("https://harizma-service.ru/form?id=%s", id)
	http.Redirect(w, r, redirectURL, http.StatusFound)

	go func() {
		spreadsheets.GoogleSheetsUpdate(1, 9, strconv.Itoa(CountUserRedirect))
		spreadsheets.GoogleSheetsUpdate(CountGetUrl, 0, id)
	}()
}

func SendedSms(w http.ResponseWriter, r *http.Request) {
	CountSendedSms++
	spreadsheets.GoogleSheetsUpdate(1, 7, strconv.Itoa(CountSendedSms))
}

func SendedDoneSms(w http.ResponseWriter, r *http.Request) {
	CountSendedDoneSms++
	spreadsheets.GoogleSheetsUpdate(1, 8, strconv.Itoa(CountSendedDoneSms))
}

func GenerateSecretKey(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err) // Handle error appropriately in production code
	}
	return base64.StdEncoding.EncodeToString(b)
}
