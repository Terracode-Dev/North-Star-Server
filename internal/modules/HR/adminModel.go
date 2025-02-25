package hr

import (
	"database/sql"

	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type CreateHrAdminReqModel struct {
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Status    bool `json:"status"`
	BranchID  int64  `json:"branch_id"`
	CreatedBy *int64 `json:"created_by"`
	UpdatedBy *int64 `json:"updated_by"`
}

func (A *CreateHrAdminReqModel) convertToDbStruct(admin_id int64) (db.CreateHrAdminParams, error) {
	var created_by sql.NullInt64
	created_by.Int64 = admin_id
	created_by.Valid = true

	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(A.Password), bcrypt.DefaultCost)
	if err != nil {
		return db.CreateHrAdminParams{}, err
	}

	var Status_string string

	if A.Status {
		Status_string = "active"
	}else {
		Status_string = "suspended"
	}

	return db.CreateHrAdminParams{
		UserName:  A.UserName,
		Email:     A.Email,
		Password:  string(hashedPassword),
		Role:      A.Role,
		Status:    Status_string,
		BranchID:  A.BranchID,
		CreatedBy: created_by,
		UpdatedBy: updated_by,
	}, nil
}

func (A *CreateHrAdminReqModel) convertToDbStructForUpdate(id int64, admin_id int64) (db.UpdateHrAdminParams, error) {
	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true
	
	var Status_string string

	if A.Status {
		Status_string = "active"
	}else {
		Status_string = "suspended"
	}

	return db.UpdateHrAdminParams{
		UserName:  A.UserName,
		Email:     A.Email,
		Role:      A.Role,
		Status:    Status_string,
		BranchID:  A.BranchID,
		UpdatedBy: updated_by,
		ID:        id,
	}, nil
}

type GetAdminReqModel struct {
	Search     string `json:"search"`
	Limit      int32  `json:"limit"`
	PageNumber int32  `json:"page"`
}

type AdminLoginReqModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
