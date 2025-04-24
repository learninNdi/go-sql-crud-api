package main

import (
	"go-sql-crud-api/app"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app.Run()
}
