package hr

import (
	"fmt"

	"time"

	 db "github.com/Terracode-Dev/North-Star-Server/internal/database"

)

type GetAccountDetailsReqParams struct {
	MONTH    string    `json:"date"`
	BranchID int64     `json:"branch_id"`
}

func (s *GetAccountDetailsReqParams) ConvertToDbParams() (db.GetAccountDetailsParams, error) {
	month, err := time.Parse("2006-01", s.MONTH)
	if err != nil {
		return db.GetAccountDetailsParams{}, fmt.Errorf("invalid month format: %v", err)
	}
	return db.GetAccountDetailsParams{
		DATEFORMAT:     month,
		BranchID:  s.BranchID,
	}, nil
}

type SalaryTransferRes struct {
	AccountData []db.GetAccountDetailsRow `json:"accountdata"`
	Total      float64                   `json:"total"`
}

type ReportsReq struct {
	Date   string `json:"date"`
}

func (s *ReportsReq) ConvertToDbParams() (time.Time, error) {
	date, err := time.Parse("2006-01-02", s.Date)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format: %v", err)
	}
	return date, nil
}

type GetStaffPayrollReqParams struct {
	MONTH     string    `json:"date"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BranchID  int64     `json:"branch_id"`
}

func (s *GetStaffPayrollReqParams) ConvertToDbParams() (db.GetStaffPayrollParams, error) {
    month, err := time.Parse("2006-01", s.MONTH)
    if err != nil {
        return db.GetStaffPayrollParams{}, fmt.Errorf("invalid month format: %v", err)
    }
    
    return db.GetStaffPayrollParams{
        DATEFORMAT:     month,
        Column2: s.FirstName,
        CONCAT:  s.FirstName, // For the LIKE comparison
        Column4: s.LastName,
        CONCAT_2: s.LastName,   // For the LIKE comparison
        BranchID:  s.BranchID,
        Column7: s.BranchID,   // For the OR condition
    }, nil
}