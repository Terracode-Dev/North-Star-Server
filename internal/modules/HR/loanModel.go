package hr

import (
	"database/sql"
	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
	// "golang.org/x/crypto/bcrypt"
)

type CreateRequestReqParams struct {
	EmpID         int64          `json:"emp_id"`
	Reason        string         `json:"reason"`
	Amount        string         `json:"amount"`
}

func (l *CreateRequestReqParams) ToDbParams() (db.CreateRequestParams, error){
	return db.CreateRequestParams{
		EmpID: l.EmpID,
		Reason: sql.NullString{String: l.Reason, Valid: true},
		Amount: sql.NullString{String: l.Amount, Valid: true},
		Status: sql.NullString{String: "", Valid: false},
		DeclinedBy: sql.NullInt64{Int64: 0, Valid: false},
		DeclineReason: sql.NullString{String: "", Valid: false},
	},nil
}

type UpdateRequestReqParams struct {
	Reason string         `json:"reason"`
	Amount string         `json:"amount"`
	ID     int64          `json:"id"`
}

func (u *UpdateRequestReqParams) ToDbParams() (db.UpdateRequestParams, error) {
	return db.UpdateRequestParams{
		Reason: sql.NullString{String: u.Reason, Valid: true},
		Amount: sql.NullString{String: u.Amount, Valid: true},
		ID: u.ID,
	},nil
}