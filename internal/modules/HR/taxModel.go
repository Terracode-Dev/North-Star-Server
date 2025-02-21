package hr

import (
	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
	"github.com/shopspring/decimal"
)

type CreateTaxReqModel struct {
	TaxFrom       string   `json:"tax_from"`
	TaxTo         string   `json:"tax_to"`
	TaxPercentage string   `json:"tax_percentage"`
}

func (T *CreateTaxReqModel) ToCreateTaxParams() (db.CreateTaxParams, error) {

	tax_from, err := decimal.NewFromString(T.TaxFrom)
	if err != nil {
		return db.CreateTaxParams{}, err
	}

	tax_to, err := decimal.NewFromString(T.TaxTo)
	if err != nil {
		return db.CreateTaxParams{}, err
	}

	tax_percentage, err := decimal.NewFromString(T.TaxPercentage)
	if err != nil {
		return db.CreateTaxParams{}, err
	}
	
	return db.CreateTaxParams{
		TaxFrom:       tax_from,
		TaxTo:         tax_to,
		TaxPercentage: tax_percentage,
	}, nil
}