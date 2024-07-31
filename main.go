package main

import (
	"crawler/db"
	"crawler/script"
	"crawler/utils"
)

func main() {
	pool := utils.ConnectDB()
	queries := db.New(pool)
	defer pool.Close()

	script.UpdateCompanies(queries) //Must run first

	script.UpdateOverview(queries)
	script.UpdateProfile(queries)
	script.UpdateShareholders(queries)
	script.UpdateInsiderDeals(queries)
	script.UpdateSubsidiaries(queries)
	script.UpdateOfficers(queries)
}
