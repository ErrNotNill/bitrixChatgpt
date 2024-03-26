package main

import (
	"bitrix_app/backend/bitrix/repo/mysql"
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

	if err := godotenv.Load(filepath.Join(".env")); err != nil {
		log.Print("No .env file found")
	} else {
		fmt.Println("Loaded .env file")
	}

	routes.Router()

	server := &http.Server{
		Addr:              ":9090",
		ReadHeaderTimeout: 3 * time.Second,
	}

	urlMysql := os.Getenv("URL_MYSQL")
	err := mysql.InitDB(urlMysql)
	if err != nil {
		fmt.Println("cant' connect to mysql")
		log.Fatal(err)
	} else {
		fmt.Println("db init accepted")
	}

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Server started with error")
		panic(err)
	}

}
