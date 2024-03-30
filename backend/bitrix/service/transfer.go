package service

import (
	"bitrix_app/backend/bitrix/authorize"
	"fmt"
	"log"
	"net/http"
)

func TransferDealsOnVue(w http.ResponseWriter, r *http.Request) {

	deals, err := GetDeals(authorize.GlobalAuthId)
	if err != nil {
		log.Println("error getting service: ", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Log the deals for debugging
	fmt.Println("TransferDealsOnVue deals: ", string(deals))

	// Write the JSON data to the response
	_, writeErr := w.Write(deals)
	if writeErr != nil {
		log.Println("error writing service to response: ", writeErr.Error())
	}
}
