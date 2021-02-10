package repo

import (
	"context"

	db "github.com/CezarGarrido/crypto-index/api/db"
	"github.com/CezarGarrido/crypto-index/api/entity"
)

type CurrencyRepoInterface interface {
	Update(ctx context.Context, currency *entity.Currency) (*entity.Currency, error)
}

func NewCurrencyInMemoryRepo(db *db.DB) CurrencyRepoInterface {
	return &currencyInMemoryRepo{
		db: db,
	}
}

type currencyInMemoryRepo struct {
	db *db.DB
}

func (c *currencyInMemoryRepo) Update(ctx context.Context, currency *entity.Currency) (*entity.Currency, error) {
	return currency, c.db.MemoryDB.WriteFile(currency)
}
