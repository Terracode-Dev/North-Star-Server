package hr

import (
	"database/sql"
	"fmt"

	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
	"github.com/shopspring/decimal"
)

type CreateAllowancesReqModel struct {
	AllowanceType string `json:"allowance_type"`
	Amount        string `json:"amount"`
	UpdatedBy     *int64 `json:"updated_by"`
	ID            int64  `json:"id"`
}

func (M *CreateAllowancesReqModel) ToCreateAllowancesParams(admin_id int64) (db.CreateAllowancesParams, error) {
	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true

	amount, err := decimal.NewFromString(M.Amount)
	if err != nil {
		return db.CreateAllowancesParams{}, fmt.Errorf("invalid amount format: %v", err)
	}

	return db.CreateAllowancesParams{
		AllowanceType: M.AllowanceType,
		Amount:        amount,
		UpdatedBy:     updated_by,
	}, nil
}

func (M *CreateAllowancesReqModel) ToUpdateAllowancesParams(id int64, admin_id int64) db.UpdateAllowanceParams {
	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = admin_id
		updated_by.Valid = true
	}

	amount, err := decimal.NewFromString(M.Amount)
	if err != nil {
		return db.UpdateAllowanceParams{}
	}

	return db.UpdateAllowanceParams{
		AllowanceType: M.AllowanceType,
		Amount:        amount,
		UpdatedBy:     updated_by,
		ID:            id,
	}
}

