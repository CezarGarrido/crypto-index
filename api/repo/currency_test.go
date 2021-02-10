package repo_test

import (
	"context"
	"testing"

	db "github.com/CezarGarrido/crypto-index/api/db"
	"github.com/CezarGarrido/crypto-index/api/entity"
	"github.com/CezarGarrido/crypto-index/api/repo"
)

// TestNewCurrencyRepo: Creating test for new users
func TestNewCurrencyRepo(t *testing.T) {

	ctx := context.TODO()
	database, err := db.NewWithMemoryDB("../../api/storage/currencies.json")
	if err != nil {
		t.Error(err.Error())
	}

	if err := database.MemoryDB.WriteFile(entity.Currency{
		BRL: "5.400",
		EUR: "0.920",
		CAD: "1.440",
	}); err != nil {
		t.Error(err.Error())
	}

	currencyRepo := repo.NewCurrencyInMemoryRepo(database)
	currencyToUpdate := entity.Currency{
		BRL: "5.500",
		EUR: "0.620",
		CAD: "1.740",
	}
	currencyUpdated, err := currencyRepo.Update(ctx, &currencyToUpdate)
	if err != nil {
		t.Error(err.Error())
	}

	if currencyToUpdate.BRL != currencyUpdated.BRL {
		t.Error("currencyToUpdate.BRL != currencyUpdated.BRL")
	}

	if currencyToUpdate.EUR != currencyUpdated.EUR {
		t.Error("error: currencyToUpdate.EUR != currencyUpdated.EUR")
	}

	if currencyToUpdate.CAD != currencyUpdated.CAD {
		t.Error("currencyToUpdate.CAD != currencyUpdated.CADL")
	}
}
