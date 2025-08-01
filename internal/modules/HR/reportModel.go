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

type GetExpiredVisaOrReportsReqParams struct {
	BranchID int64       `json:"branch_id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Department string     `json:"department"`
	PassportID int64     `json:"passport_id"`
	VisaNumber string     `json:"visanumber"`
	Limit    int64      `json:"limit"`
	Offset   int64       `json:"offset"`
}

func (s *GetExpiredVisaOrReportsReqParams) ConvertToDbParams() (db.GetExpiredVisaOrReportsParams, error) {
	return db.GetExpiredVisaOrReportsParams{
		BranchID:   s.BranchID,
		Column2:  s.FirstName,
		CONCAT:   s.FirstName,
		Column4: s.LastName,
		CONCAT_2: s.LastName,
		Column6: s.Department,
		CONCAT_3: s.Department,
		Column8: s.PassportID,
		CONCAT_4: s.PassportID,
		Column10: s.VisaNumber,
		CONCAT_5: s.VisaNumber,
		Limit:      int32(s.Limit),
		Offset:     int32(s.Offset),
	}, nil
}

func (s *GetExpiredVisaOrReportsReqParams) ConvertToDbParamsForSoonExpiring() (db.GetVisaOrPassportExpiringSoonParams, error) {
	return db.GetVisaOrPassportExpiringSoonParams{
		BranchID:   s.BranchID,
		Column2:  s.FirstName,
		CONCAT:   s.FirstName,
		Column4: s.LastName,
		CONCAT_2: s.LastName,
		Column6: s.Department,
		CONCAT_3: s.Department,
		Column8: s.PassportID,
		CONCAT_4: s.PassportID,
		Column10: s.VisaNumber,
		CONCAT_5: s.VisaNumber,
		Limit:      int32(s.Limit),
		Offset:     int32(s.Offset),
	}, nil
}

type GetEmployeeInsuranceReqParams struct {
	BranchID int64       `json:"branch_id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Department string     `json:"department"`
	Limit    int32       `json:"limit"`
	Offset   int32       `json:"offset"`
}

func (s *GetEmployeeInsuranceReqParams) ConvertToDbParams() (db.GetempployeeInsuranceParams, error) {
	return db.GetempployeeInsuranceParams{
		BranchID:   s.BranchID,
		Column2:  s.FirstName,
		CONCAT:   s.FirstName,
		Column4: s.LastName,
		CONCAT_2: s.LastName,
		Column6: s.Department,
		CONCAT_3: s.Department,
		Limit:      s.Limit,
		Offset:     s.Offset,
	}, nil
}