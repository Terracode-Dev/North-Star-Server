package hr

import (
	"fmt"
	"encoding/json"
	"database/sql"
	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
)

type EmpDataModel struct {
	Employee       CreateEmployeeReqModel            `json:"employee"`
	Emergency      CreateEmpEmergencyDetailsReqModel `json:"emergency"`
	Bank           CreateEmpBankDetailsReqModel      `json:"bank"`
	User           CreateEmpUserReqModel             `json:"user"`
	Expatriate     CreateEmpExpatriateReqModel       `json:"expatriate"`
	FileSubmit     []CreateFileSubmitReqModel		 `json:"file_submit"`
}

type CreateEmpLinkReqParams struct {
	EmpData    EmpDataModel    `json:"emp_data"`
	PresetID   int64           `json:"preset_id"`
	IsApproved bool            `json:"is_approved"`
	Email      string          `json:"email"`
}

func (c *CreateEmpLinkReqParams) ToCreateEmpLinkParams() (db.CreateEmpLinkParams, error) {
	empData, err := json.Marshal(c.EmpData)
	if err != nil {
		return db.CreateEmpLinkParams{}, fmt.Errorf("failed to marshal emp data: %w", err)
	}
	return db.CreateEmpLinkParams{
		EmpData:    empData,
		PresetID:   c.PresetID,
		IsApproved: c.IsApproved,
		Email:      c.Email,
	}, nil
}

type UpdateEmpLinkApprovalReqParams struct {
	IsApproved bool          `json:"is_approved"`
	ID         int64         `json:"id"`
}

func (u *UpdateEmpLinkApprovalReqParams) ToUpdateEmpLinkApprovalParams(adminID int64) (db.UpdateEmpLinkApprovalParams, error) {
	var updatedBy sql.NullInt64
	updatedBy.Int64 = adminID
	updatedBy.Valid = true
	return db.UpdateEmpLinkApprovalParams{
		IsApproved: u.IsApproved,
		UpdatedBy:  updatedBy,
		ID:         u.ID,
	},nil
}




