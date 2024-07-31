package main

import (
	"crawler/db"
	"crawler/script"
)

func main() {
	pool := script.Connect()
	queries := db.New(pool)
	defer pool.Close()

	script.UpdateCompanies(queries) //Must run first

	script.UpdateOverview(queries)
	script.UpdateProfile(queries)
	script.UpdateShareholders(queries)
	script.UpdateInsiderDeals(queries)
}
