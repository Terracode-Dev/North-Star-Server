package hr

import (
	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
	"github.com/shopspring/decimal"
)

type CreateTaxReqModel struct {
	TaxFrom       int   `json:"tax_from"`
	TaxTo         int   `json:"tax_to"`
	TaxPercentage int   `json:"tax_percentage"`
}

func (T *CreateTaxReqModel) ToCreateTaxParams() (db.CreateTaxParams, error) {

	tax_from := decimal.NewFromInt(int64(T.TaxFrom))

	tax_to := decimal.NewFromInt(int64(T.TaxTo))

	tax_percentage := decimal.NewFromInt(int64(T.TaxPercentage))
	
	return db.CreateTaxParams{
		TaxFrom:       tax_from,
		TaxTo:         tax_to,
		TaxPercentage: tax_percentage,
	}, nil
}