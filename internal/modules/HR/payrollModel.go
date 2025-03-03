package hr

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
	"github.com/shopspring/decimal"
)

type CreatePayrollReqModel struct {
	Employee                string          `json:"employee"`
	Date                    string          `json:"date"`
	SalaryType              string          `json:"salary_type"`
	Amount                  string          `json:"amount"`
	TotalOfSalaryAllowances string          `json:"total_of_salary_allowances"`
	Pension                 bool            `json:"pension"`
	PensionEmployer         string          `json:"pension_employer"`
	PensionEmployee         string          `json:"pension_employee"`
	TotalNetSalary          string          `json:"total_net_salary"`
	Tax                     bool            `json:"tax"`
	TaxPercentage           string          `json:"tax_percentage"`
	TotalNetSalaryAfterTax  string          `json:"total_net_salary_after_tax"`
	UpdatedBy               *int64          `json:"updated_by"`
}

func (A *CreatePayrollReqModel) ToCreatePayrollParams(admin_id int64) (db.CreatePayrollParams, error) {

	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true


	date, err := time.Parse(time.RFC3339, A.Date)
	if err != nil {
		log.Printf("Error parsing dob: %v", err)
		return db.CreatePayrollParams{}, fmt.Errorf("invalid date format: %v", err)
	}

	amount, err := decimal.NewFromString(A.Amount)
	if err != nil {
		return db.CreatePayrollParams{}, fmt.Errorf("invalid amount format: %v", err)
	}

	total_of_salary_allowances, err := decimal.NewFromString(A.TotalOfSalaryAllowances)
	if err != nil {
		return db.CreatePayrollParams{}, fmt.Errorf("invalid total_of_salary_allowances format: %v", err)
	}

	total_net_salary, err := decimal.NewFromString(A.TotalNetSalary)
	if err != nil {
		return db.CreatePayrollParams{}, fmt.Errorf("invalid total_net_salary format: %v", err)
	}

	total_net_salary_after_tax, err := decimal.NewFromString(A.TotalNetSalaryAfterTax)
	if err != nil {
		return db.CreatePayrollParams{}, fmt.Errorf("invalid total_net_salary_after_tax format: %v", err)
	}

	var tax_percentage sql.NullString
	if A.TaxPercentage != "" {
		tax_percentage.String = A.TaxPercentage
		tax_percentage.Valid = true
	}

	var pension_employer sql.NullString
	if A.PensionEmployer != "" {
		pension_employer.String = A.PensionEmployer
		pension_employer.Valid = true
	}

	var pension_employee sql.NullString
	if A.PensionEmployee != "" {
		pension_employee.String = A.PensionEmployee
		pension_employee.Valid = true
	}

	return db.CreatePayrollParams{
		Employee:                A.Employee,
		Date:                    date,
		SalaryType:              A.SalaryType,
		Amount:                  amount,
		TotalOfSalaryAllowances: total_of_salary_allowances,
		Pension:                 A.Pension,
		PensionEmployer:         pension_employer,
		PensionEmployee:         pension_employee,
		TotalNetSalary:          total_net_salary,
		Tax:                     A.Tax,
		TaxPercentage:           tax_percentage,
		TotalNetSalaryAfterTax:  total_net_salary_after_tax,
		UpdatedBy:               updated_by,
	}, nil
}

func (A *CreatePayrollReqModel) ToUpdatePayrollParams(id int64, admin_id int64) (db.UpdatePayrollParams, error) {

	var updated_by sql.NullInt64
	if A.UpdatedBy != nil {
		updated_by.Int64 = admin_id
		updated_by.Valid = true
	}

	date, err := time.Parse(time.RFC3339, A.Date)
	if err != nil {
		log.Printf("Error parsing dob: %v", err)
		return db.UpdatePayrollParams{}, err
	}

	amount, err := decimal.NewFromString(A.Amount)
	if err != nil {
		return db.UpdatePayrollParams{}, err
	}

	total_of_salary_allowances, err := decimal.NewFromString(A.TotalOfSalaryAllowances)
	if err != nil {
		return db.UpdatePayrollParams{}, err
	}

	total_net_salary, err := decimal.NewFromString(A.TotalNetSalary)
	if err != nil {
		return db.UpdatePayrollParams{}, err
	}

	total_net_salary_after_tax, err := decimal.NewFromString(A.TotalNetSalaryAfterTax)
	if err != nil {
		return db.UpdatePayrollParams{}, err
	}

	var tax_percentage sql.NullString
	if A.TaxPercentage != "" {
		tax_percentage.String = A.TaxPercentage
		tax_percentage.Valid = true
	}

	var pension_employer sql.NullString
	if A.PensionEmployer != "" {
		pension_employer.String = A.PensionEmployer
		pension_employer.Valid = true
	}

	var pension_employee sql.NullString
	if A.PensionEmployee != "" {
		pension_employee.String = A.PensionEmployee
		pension_employee.Valid = true
	}

	return db.UpdatePayrollParams{
		Employee:                A.Employee,
		Date:                    date,
		SalaryType:              A.SalaryType,
		Amount:                  amount,
		TotalOfSalaryAllowances: total_of_salary_allowances,
		Pension:                 A.Pension,
		PensionEmployer:         pension_employer,
		PensionEmployee:         pension_employee,
		TotalNetSalary:          total_net_salary,
		Tax:                     A.Tax,
		TaxPercentage:           tax_percentage,
		TotalNetSalaryAfterTax:  total_net_salary_after_tax,
		UpdatedBy:               updated_by,
		ID:                      id,
	},nil
}

type CreatePayrollAllowancesParams struct {
	Name      string          `json:"name"`
	Amount    string          `json:"amount"`
	PayrollID int64           `json:"payroll_id"`
	UpdatedBy *int64          `json:"updated_by"`
}

func (A *CreatePayrollAllowancesParams) ToCreatePayrollAllowancesParams(admin_id int64) (db.CreatePayrollAllowancesParams, error) {
	
	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true
	

	amount, err := decimal.NewFromString(A.Amount)
	if err != nil {
		return db.CreatePayrollAllowancesParams{}, fmt.Errorf("invalid amount format: %v", err)
	}

	return db.CreatePayrollAllowancesParams{
		Name:      A.Name,
		Amount:    amount,
		PayrollID: A.PayrollID,
		UpdatedBy: updated_by,
	}, nil
}

type PayrollAllowances struct {
	Payroll     CreatePayrollReqModel           `json:"payroll"`
	Allowances  []CreatePayrollAllowancesParams `json:"allowances"`
}

type GetPayrollsReqModel struct {
	Limit  int32 `json:"limit"`
	PageNumber int32 `json:"page"`
}

func (A *GetPayrollsReqModel) ToGetPayrollsParams() (db.GetPayrollsParams,error) {

	offset := (A.PageNumber - 1) * A.Limit

	return db.GetPayrollsParams{
		Limit:  A.Limit,
		Offset: offset,
	},nil
}