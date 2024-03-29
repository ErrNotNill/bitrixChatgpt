package routes

import (
	"bitrix_app/backend/bitrix/authorize"
	"net/http"
)

func Router() {

	//http.HandleFunc("/api/check", authorize.BotBitrix)
	http.HandleFunc("/api/connect", authorize.ConnectionBitrix)
	http.HandleFunc("/api/deals_get", authorize.TransferDealsOnVue)
	/*c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Change this to the specific origin of your Vue.js app in a production environment.
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})

	http.Handle("/api/auth_page", c.Handler(http.HandlerFunc(repository.AuthPage)))
	http.Handle("/api/login_page", c.Handler(http.HandlerFunc(repository.LoginPage)))
	http.HandleFunc("/api/redirect", repository.RedirectPage) //here user redirects from login page*/

	http.HandleFunc("/api/redirect", authorize.ConnectionBitrixLogger)

}
