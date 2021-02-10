package repo

import (
	"context"

	db "github.com/CezarGarrido/crypto-index/api/db"
	"github.com/CezarGarrido/crypto-index/api/entity"
)
// CurrencyRepoInterface represents currency repo database
type CurrencyRepoInterface interface {
	Update(ctx context.Context, currency *entity.Currency) (*entity.Currency, error)
}

// NewCurrencyInMemoryRepo : Return new CurrencyRepoInterface with in memory database.
// Example:
// database, _ := db.NewWithMemoryDB("./api/storage/currencies.json")
// currencyRepo := NewCurrencyInMemoryRepo(database)
// currencyRepo.Update(...)

func NewCurrencyInMemoryRepo(db *db.DB) CurrencyRepoInterface {
	return &currencyInMemoryRepo{
		db: db,
	}
}

// currencyInMemoryRepo implements CurrencyRepoInterface
type currencyInMemoryRepo struct {
	db *db.DB // db: represents the structure that contains the database
}

// Update: Update the currency in the memory repository.
// Example: 
// database, _ := db.NewWithMemoryDB("./api/storage/currencies.json")
// currencyRepo := NewCurrencyInMemoryRepo(database)
// ctx := context.TODO()
// currencyToUpdate := &entity.Currency{BRL: "5.500",EUR: "0.620",CAD: "1.740"}
// currencyUpdated, err := currencyRepo.Update(ctx, currencyToUpdate)
// if err != nil { // Error handling return}
// Its work currencyUpdated

func (c *currencyInMemoryRepo) Update(ctx context.Context, currency *entity.Currency) (*entity.Currency, error) {
	return currency, c.db.MemoryDB.WriteFile(currency)
}
