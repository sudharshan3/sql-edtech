package main

import (
	"log"
	"net/http"

	"github.com/sudharshan3/sql-edtech/routes"
)

func main() {
	r := routes.Router()
	err := http.ListenAndServe(":3000", r) // setting listening port
	if err != nil {
		log.Println("ListenAndServe: ", err)
	} else {
		log.Println("Backend Started at 3000...")
	}
}
