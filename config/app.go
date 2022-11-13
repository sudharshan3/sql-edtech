package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() *sql.DB {
	dbDriver := "mysql"
	dbURL := "localhost:3306"
	dbName := "root"
	dbPass := "ShanSud3198#"
	dbTable := "edtech"
	db, err := sql.Open(dbDriver, dbName+":"+dbPass+"@tcp("+dbURL+")/"+dbTable)
	if err != nil {
		log.Println("DB Connection Failed...")
		log.Println(err.Error())
	}
	log.Println("DB Connection Success...")
	return db

}
