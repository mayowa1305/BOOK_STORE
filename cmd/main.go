package main

import (
	"github.com/mayowa1305/adx_bookstore/config"
	"github.com/mayowa1305/adx_bookstore/database"
)

func main() {
	cfg := config.InitConfig()

	database.InitDB(cfg)
}
