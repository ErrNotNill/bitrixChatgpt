package routes

import (
	"fmt"
	"log"
	"net/http"
)

func Router() {

	http.HandleFunc("/api/test", TestHandler)

	/*c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Change this to the specific origin of your Vue.js app in a production environment.
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})

	http.Handle("/api/auth_page", c.Handler(http.HandlerFunc(repository.AuthPage)))
	http.Handle("/api/login_page", c.Handler(http.HandlerFunc(repository.LoginPage)))
	http.HandleFunc("/api/redirect", repository.RedirectPage) //here user redirects from login page*/

}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(fmt.Sprintf("OK")))
	if err != nil {
		log.Println("Error redirect", err.Error())
		return
	}
	fmt.Println(r.Body)
}
