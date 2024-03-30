package service

import (
	"bitrix_app/backend/bitrix/authorize"
	"fmt"
	"log"
	"net/http"
)

func TransferDealsOnVue(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello"))

	deals, err := GetDeals(authorize.GlobalAuthId)
	if err != nil {
		log.Println("error getting service: ", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set content type to application/json before writing the response
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("TransferDealsOnVue deals: ", string(deals))

	// Write the JSON data to the response
	_, writeErr := w.Write(deals)
	if writeErr != nil {
		log.Println("error writing service to response: ", writeErr.Error())
		// Note: In real scenarios, consider handling this error more gracefully,
		// since part of the HTTP response might have already been written,
		// making it tricky to send a proper HTTP status code at this point.
	}
}
