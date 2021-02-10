package repo_test

import (
	"context"
	"testing"

	db "github.com/CezarGarrido/crypto-index/api/db"
	"github.com/CezarGarrido/crypto-index/api/entity"
	"github.com/CezarGarrido/crypto-index/api/repo"
	"github.com/stretchr/testify/assert"

	
)

// TestUpdateCurrencyRepo: Creating test for new users
func TestUpdateCurrencyRepo(t *testing.T) {

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

	// test: Setup tests
	type test struct {
		currency entity.Currency
		expected error
	}

	currencyToTests := []test{
		{entity.Currency{
			BRL: "5.500",
			EUR: "0.620",
			CAD: "1.740",
		}, nil},
		{entity.Currency{
			BRL: "5.500",
			EUR: "0.620",
			CAD: "1.740",
		}, nil},
		{entity.Currency{
			BRL: "5.500",
			EUR: "0.620",
			CAD: "1.740",
		}, nil},
		{entity.Currency{
			BRL: "5.500",
			EUR: "0.620",
			CAD: "1.740",
		}, nil},
	}

	for _, tt := range currencyToTests {
		currencyUpdated, err := currencyRepo.Update(ctx, &tt.currency)
		if err != nil {
			t.Errorf("no error was expected but got: %q", err)
		}
		assert.Equal(t, tt.currency.BRL, currencyUpdated.BRL)
		assert.Equal(t, tt.currency.EUR, currencyUpdated.EUR)
		assert.Equal(t, tt.currency.CAD, currencyUpdated.CAD)
	}

}

