package hr

import (
	"database/sql"
	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
)

type CreateServicesReqModel struct {
	Category  string        `json:"category"`
	Value     string        `json:"value"`
	UpdatedBy *int64        `json:"updated_by"`
}

func (M *CreateServicesReqModel) ToCreateServicesParams() (db.CreateServicesParams, error) {

	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = *M.UpdatedBy
		updated_by.Valid = true
	}

	return db.CreateServicesParams{
		Category:  M.Category,
		Value:     M.Value,
		UpdatedBy: updated_by,
	}, nil
}

func (M *CreateServicesReqModel) ToUpdateServicesParams(id int64) db.UpdateServiceParams {

	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = *M.UpdatedBy
		updated_by.Valid = true
	}

	return db.UpdateServiceParams{
		Category:  M.Category,
		Value:     M.Value,
		UpdatedBy: updated_by,
		ID:        id,
	}
}