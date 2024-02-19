package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mayowa1305/adx_bookstore/config"
	"github.com/spf13/viper"
)

var DB *sql.DB

// initializes db connection
func InitDB(cnfg *config.Config) {

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", viper.GetString("db.user"), viper.GetString("db.password"), viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.name"))

	log.Printf("DB URL: %s", dbURL)
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	DB = db

	err = DB.Ping()
	if err != nil {
		log.Fatal("error pinging the database:", err)
	}

	log.Println("connected to database")
}
