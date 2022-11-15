package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME")+"?charset=utf8mb4")
	if err != nil {
		log.Println("DB Connection Failed...")
		log.Println(err.Error())
	} else {
		log.Println("DB Connection Success...")
	}

	return db

}
