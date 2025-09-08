package hr

import (
	"database/sql"

	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
	"github.com/shopspring/decimal"
)

type CreateTaxReqModel struct {
	TaxFrom       int `json:"tax_from"`
	TaxTo         int `json:"tax_to"`
	TaxPercentage float64 `json:"tax_percentage"`
}

func (T *CreateTaxReqModel) ToCreateTaxParams(admin_id int64) (db.CreateTaxParams, error) {
	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true
	tax_from := decimal.NewFromInt(int64(T.TaxFrom))

	tax_to := decimal.NewFromInt(int64(T.TaxTo))

	tax_percentage := decimal.NewFromFloat(T.TaxPercentage)

	return db.CreateTaxParams{
		TaxFrom:       tax_from,
		TaxTo:         tax_to,
		TaxPercentage: tax_percentage,
		UpdatedBy:     updated_by,
	}, nil
}

