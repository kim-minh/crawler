package main

import (
	"crawler/db"
	"crawler/script"
)

func main() {
	pool := script.Connect()
	queries := db.New(pool)
	defer pool.Close()

	script.UpdateCompanies(queries)
	defer script.UpdateOverview(queries)
	defer script.UpdateProfile(queries)
}
