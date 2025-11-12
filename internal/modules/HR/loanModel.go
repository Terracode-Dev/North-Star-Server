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

type UpdateRequestStatusReqParams struct {
	Status        string         `json:"status"`
	DeclinedBy    int64          `json:"declined_by"`
	DeclineReason string         `json:"decline_reason"`
	ID            int64          `json:"id"`
}

func ( u *UpdateRequestStatusReqParams) ToDbParams() (db.UpdateRequestStatusParams, error) {
	return db.UpdateRequestStatusParams{
		Status: sql.NullString{String: u.Status, Valid: true},
		DeclinedBy: sql.NullInt64{Int64: u.DeclinedBy, Valid: true},
		DeclineReason: sql.NullString{String: u.DeclineReason, Valid: true},
		ID: u.ID,
	}, nil
}

type GetRequestsReqParams struct {
	FirstName string     `json:"firstname"`
	LastName string      `json:"lastname"`
	Limit    int32       `json:"limit"`
	Offset   int32       `json:"offset"`
}

func (u *GetRequestsReqParams) ToDbParams(emp_id int64) (db.GetRequestsParams, error) {
	return db.GetRequestsParams{
		Column1: u.FirstName,
		CONCAT: u.FirstName,
		Column3: u.LastName,
		CONCAT_2: u.LastName,
		EmpID: emp_id,
		Limit: u.Limit,
		Offset: u.Offset,
	},nil
}

type GetRequestsAdminReqParams struct {
	FirstName string     `json:"firstname"`
	LastName string      `json:"lastname"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (u *GetRequestsAdminReqParams) ToDbParams(branchId int64) (db.GetRequestsAdminParams, error) {
	return db.GetRequestsAdminParams{
		ID: branchId,
		Column2: u.FirstName,
		CONCAT: u.FirstName,
		Column4: u.LastName,
		CONCAT_2: u.LastName,
		Limit: u.Limit,
		Offset: u.Offset,
	},nil
}
