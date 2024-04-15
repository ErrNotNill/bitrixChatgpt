package main

import (
	"bitrix_app/backend/bitrix/endpoints"
	"bitrix_app/backend/bitrix/repo/mysql"
	"bitrix_app/backend/bitrix/test/spreadsheets"
	"bitrix_app/backend/routes"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	fmt.Println("service starting...")

	if err := godotenv.Load(filepath.Join(".env")); err != nil {
		log.Print("No .env file found")
	} else {
		fmt.Println("Loaded .env file")
	}
	//authorize.StartB24()
	endpoints.BitrixDomain = os.Getenv("BITRIX_DOMAIN")
	//authorize.CustomAuthorizeBitrix()

	routes.Router()

	/*	err := authorize.AuthorizeBitrix()
		if err != nil {
			log.Println("Bitrix authorization error...")
		}*/

	server := &http.Server{
		Addr:              ":9090",
		ReadHeaderTimeout: 3 * time.Second,
	}

	spreadsheets.GoogleSheetsStart()

	return
	urlMysql := os.Getenv("URL_MYSQL")

	err := mysql.InitDB(urlMysql)
	if err != nil {
		fmt.Println("cant' connect to mysql")
		log.Fatal(err)
	} else {
		fmt.Println("db init accepted")
	}

	fmt.Printf("server started on addr: %s", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Server started with error")
		panic(err)
	}

}
