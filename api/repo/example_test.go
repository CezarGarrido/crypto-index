package repo_test

import (
	"log"

	db "github.com/CezarGarrido/crypto-index/api/db"
	"github.com/CezarGarrido/crypto-index/api/repo"
)

// Examples: NewCurrencyInMemoryRepo

func ExampleNewCurrencyInMemoryRepo() {
	// Init in memory database
	database, err := db.NewWithMemoryDB("../../api/storage/currencies.json")
	if err != nil {
		log.Fatal(err)
	}

	// Return NewCurrencyInMemoryRepo
	currencyRepo := repo.NewCurrencyInMemoryRepo(database)
	// Use currencyRepo
	_ = currencyRepo
}
