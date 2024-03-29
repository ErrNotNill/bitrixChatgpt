package routes

import (
	"bitrix_app/backend/bitrix/authorize"
	"bitrix_app/backend/bitrix/service"
	"net/http"
)

func Router() {

	//http.HandleFunc("/api/check", authorize.BotBitrix)
	http.HandleFunc("/api/connect", authorize.ConnectionBitrixLocalApp)
	http.HandleFunc("/api/deals_get", service.TransferDealsOnVue)

	//http.HandleFunc("/api/check_widget", widget.CheckWidget) //here we create widget in bitrix

	/*c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Change this to the specific origin of your Vue.js app in a production environment.
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})

	http.Handle("/api/auth_page", c.Handler(http.HandlerFunc(repository.AuthPage)))
	http.Handle("/api/login_page", c.Handler(http.HandlerFunc(repository.LoginPage)))
	http.HandleFunc("/api/redirect", repository.RedirectPage) //here user redirects from login page*/

	http.HandleFunc("/api/redirect", service.ConnectionBitrixLogger)

}
