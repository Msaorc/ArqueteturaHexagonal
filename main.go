package main

import (
	"database/sql"

	db2 "github.com/Msaorc/ArqueteturaHexagonal/adapters/db"
	"github.com/Msaorc/ArqueteturaHexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Noteboook", 8000)
	productService.Enable(product)
}
