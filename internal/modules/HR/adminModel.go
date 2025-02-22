package hr

import (
	"database/sql"

	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
)

type CreateHrAdminReqModel struct {
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Status    string `json:"status"`
	BranchID  int64  `json:"branch_id"`
	CreatedBy *int64 `json:"created_by"`
	UpdatedBy *int64 `json:"updated_by"`
}

func (A *CreateHrAdminReqModel) convertToDbStruct() (db.CreateHrAdminParams, error) {
	var created_by sql.NullInt64
	if A.CreatedBy != nil {
		created_by.Int64 = *A.CreatedBy
		created_by.Valid = true
	}

	var updated_by sql.NullInt64
	if A.UpdatedBy != nil {
		updated_by.Int64 = *A.UpdatedBy
		updated_by.Valid = true
	}

	return db.CreateHrAdminParams{
		UserName:  A.UserName,
		Email:     A.Email,
		Password:  A.Password,
		Role:      A.Role,
		Status:    A.Status,
		BranchID:  A.BranchID,
		CreatedBy: created_by,
		UpdatedBy: updated_by,
	}, nil
}

type GetAdminReqModel struct {
	Search     string `json:"search"`
	Limit      int32  `json:"limit"`
	PageNumber int32  `json:"page"`
}
