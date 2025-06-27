package hr

import (
	"fmt"

	"github.com/Terracode-Dev/North-Star-Server/internal/database"
	"github.com/shopspring/decimal"
)

type ExchangeRateReqModel struct {
	ExchangeRate string `json:"exchange_rate"`
	CurrencyType string          `json:"currency_type"`
}

func (e *ExchangeRateReqModel) convertToDbStruct() (database.CreateExchangeRateParams, error) {
	exchangeRate, err := decimal.NewFromString(e.ExchangeRate)
	if err != nil {
	return database.CreateExchangeRateParams{}, fmt.Errorf("invalid exchange rate: %w", err)
	}
	return database.CreateExchangeRateParams{
	ExchangeRate: exchangeRate,
	CurrencyType: e.CurrencyType,
	}, nil
}