package main

import (
	"crawler/db"
	"crawler/script"
)

func main() {
	pool := script.Connect()
	queries := db.New(pool)
	defer pool.Close()

	// script.UpdateCompanies(queries)
	script.UpdateOverview(queries)
}
