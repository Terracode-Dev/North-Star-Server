package hr

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"strconv"

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
	TrainerCom              float64         `json:"trainer_com"`
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

	amount_float := amount.InexactFloat64()
	total_salary_allowances_float := total_of_salary_allowances.InexactFloat64()

	Total_Gross_Salary := amount_float + total_salary_allowances_float

	employee_pension_float, err := strconv.ParseFloat(pension_employee.String, 64)
	if err != nil {
		return db.CreatePayrollParams{}, fmt.Errorf("invalid pension_employee format: %v", err)
	}
	pension_employer_float, err := strconv.ParseFloat(pension_employer.String, 64)
	if err != nil {
		return db.CreatePayrollParams{}, fmt.Errorf("invalid pension_employer format: %v", err)
	}
	total_pension_float := employee_pension_float + pension_employer_float
	var total_net_salary_float_req float64

	if A.TrainerCom != 0 {
	total_net_salary_float_req = (Total_Gross_Salary + A.TrainerCom) - total_pension_float
	} else {
	total_net_salary_float_req = Total_Gross_Salary - total_pension_float
	}
	
	total_net_salary_float := total_net_salary.InexactFloat64()

	tax_float, err := strconv.ParseFloat(tax_percentage.String, 64)
	if err != nil {
		return db.CreatePayrollParams{}, fmt.Errorf("invalid tax_percentage format: %v", err)
	}
	tax_percentage_float := tax_float / 100.0
	total_net_salary_after_tax_float_req := total_net_salary_float*tax_percentage_float

	total_net_salary_after_tax_float := total_net_salary_after_tax.InexactFloat64()

	if total_net_salary_float != total_net_salary_float_req || total_net_salary_after_tax_float != total_net_salary_after_tax_float_req {
		return db.CreatePayrollParams{}, fmt.Errorf("total net salary does not match the calculated value")
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

type HRTrainerComReqParams struct {
	IsTrainer   bool            `json:"is_trainer"`
	TrainerComData CreateHRTrainerComReqParams `json:"trainer_com_data"`
}

type CreateHRTrainerComReqParams struct {
	PayrollID     int64           `json:"payroll_id"`
	TrainerID     int64           `json:"trainer_id"`
	EmployeeID    int64           `json:"employee_id"`
	Commission    float64         `json:"commission"`
	AssignedCount int64           `json:"assigned_count"`
	Total         float64         `json:"total"`
}

func (A *CreateHRTrainerComReqParams) ToCreateHRTrainerComParams() (db.CreateHRTrainerComParams, error) {
	var commission decimal.Decimal
	if A.Commission != 0 {
		commission = decimal.NewFromFloat(A.Commission)
	} else {
		commission = decimal.Zero
	}

	return db.CreateHRTrainerComParams{
		PayrollID:     A.PayrollID,
		TrainerID:     A.TrainerID,
		EmployeeID:    A.EmployeeID,
		Commission:    commission,
		AssignedCount: A.AssignedCount,
		Total:         decimal.NewFromFloat(A.Total),
	}, nil
}

type PayrollAllowances struct {
	Payroll     CreatePayrollReqModel           `json:"payroll"`
	Allowances  []CreatePayrollAllowancesParams `json:"allowances"`
	TrainerCom  HRTrainerComReqParams        `json:"trainer_com"`
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

type TrainerComRow struct {
	Istrainer bool   `json:"is_trainer"`
	Com_amount float64 `json:"com_amount"`
	Assign_count int64  `json:"assign_count"`
	Total_commission float64 `json:"total_commission"`
}