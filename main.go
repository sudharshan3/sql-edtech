package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sudharshan3/sql-edtech/routes"
)

func main() {
	r := routes.Router()
	err := http.ListenAndServe(":"+os.Getenv("API_PORT"), r) // setting listening port
	if err != nil {
		log.Println("ListenAndServe: ", err)
	} else {
		log.Println("Backend Started at 3000...")
	}
}
