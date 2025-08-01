package hr

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"math"
	"time"

	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
	"github.com/shopspring/decimal"
)

type CreatePayrollReqModel struct {
	EmployeeID                 int64   `json:"emp_id"`
	Date                       string  `json:"date"`
	SalaryType                 string  `json:"salary_type"`
	Amount                     string  `json:"amount"`
	SalaryAmountType           string  `json:"salary_amount_type"`
	TotalOfSalaryAllowances    string  `json:"total_of_salary_allowances"`
	TotalAllowancesType        string  `json:"total_allowances_type"`
	Pension                    bool    `json:"pension"`
	PensionEmployer            string  `json:"pension_employer"`
	PensionEmployerType        string  `json:"pension_employer_type"`
	PensionEmployee            string  `json:"pension_employee"`
	PensionEmployeeType        string  `json:"pension_employee_type"`
	TotalNetSalary             string  `json:"total_net_salary"`
	TotalNetSalaryType         string  `json:"total_net_salary_type"`
	Tax                        bool    `json:"tax"`
	TrainerCom                 float64 `json:"trainer_com"`
	TaxPercentage              string  `json:"tax_percentage"`
	TotalNetSalaryAfterTax     string  `json:"total_net_salary_after_tax"`
	TotalNetSalaryAfterTaxType string  `json:"total_net_salary_after_tax_type"`
	ERID                       int64   `json:"er_id"`
	UpdatedBy                  *int64  `json:"updated_by"`
}

func (A *CreatePayrollReqModel) ToCreatePayrollParams(admin_id int64, ex_rate float64) (db.CreatePayrollParams, error) {

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

	var emp_id sql.NullInt64
	if A.EmployeeID != 0 {
		emp_id.Int64 = A.EmployeeID
		emp_id.Valid = true
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

	var pension_employee_type sql.NullString
	if A.PensionEmployeeType != "" {
		pension_employee_type.String = A.PensionEmployeeType
		pension_employee_type.Valid = true
	}

	var pension_employer_type sql.NullString
	if A.PensionEmployerType != "" {
		pension_employer_type.String = A.PensionEmployerType
		pension_employer_type.Valid = true
	}

	var er_id sql.NullInt64
	er_id.Int64 = A.ERID
	er_id.Valid = true

	var amount_float float64

	if A.SalaryAmountType == "USD" {
		amount_float = amount.InexactFloat64() * ex_rate
	} else {
		amount_float = amount.InexactFloat64()
	}

	total_salary_allowances_float := total_of_salary_allowances.InexactFloat64()

	Total_Gross_Salary := amount_float + total_salary_allowances_float

	var total_pension_float float64
	var employee_pension_float float64
	var pension_employer_float float64

	if A.Pension {
		employee_pension_float = 0.0
		pension_employer_float = 0.0

		if pension_employee.Valid && pension_employee.String != "" {
			var err error
			employee_pension_float, err = strconv.ParseFloat(pension_employee.String, 64)
			if err != nil {
				return db.CreatePayrollParams{}, fmt.Errorf("invalid pension_employee format: %v", err)
			}
		}

		if pension_employer.Valid && pension_employer.String != "" {
			var err error
			pension_employer_float, err = strconv.ParseFloat(pension_employer.String, 64)
			if err != nil {
				return db.CreatePayrollParams{}, fmt.Errorf("invalid pension_employer format: %v", err)
			}
		}

		total_pension_float = employee_pension_float + pension_employer_float
	} else {
		total_pension_float = 0.0
	}

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
	var tax_percentage_float float64
	if A.Tax {
		tax_percentage_float = tax_float / 100
	} else {
		tax_percentage_float = 0.0
	}

	total_net_salary_after_tax_float_req := total_net_salary_float- (total_net_salary_float * tax_percentage_float)

	total_net_salary_after_tax_float := total_net_salary_after_tax.InexactFloat64()

	total_net_salary_float_req = math.Round(total_net_salary_float_req*100) / 100
	total_net_salary_float = math.Round(total_net_salary_float*100) / 100
	total_net_salary_after_tax_float_req = math.Round(total_net_salary_after_tax_float_req*100) / 100
	total_net_salary_after_tax_float = math.Round(total_net_salary_after_tax_float*100) / 100

	if total_net_salary_float != total_net_salary_float_req || total_net_salary_after_tax_float != total_net_salary_after_tax_float_req {
		log.Printf("Salary calculation mismatch detected:")
		log.Printf("  Calculated net salary: %f, Requested net salary: %f", total_net_salary_float_req, total_net_salary_float)
		log.Printf("  Calculated net salary after tax: %f, Requested net salary after tax: %f", total_net_salary_after_tax_float_req, total_net_salary_after_tax_float)
		log.Printf("  Calculation inputs:")
		log.Printf("exchange rate: %f", ex_rate)
		log.Printf("    Gross salary: %f (base: %f, allowances: %f)", Total_Gross_Salary, amount_float, total_salary_allowances_float)
		log.Printf("    Pension: %f (employer: %f, employee: %f)", total_pension_float, pension_employer_float, employee_pension_float)
		log.Printf("    Trainer commission: %f", A.TrainerCom)
		log.Printf("    Tax percentage: %f%%", tax_float)
		log.Printf("Tax Perecentage float : %f", tax_percentage_float)
		return db.CreatePayrollParams{}, fmt.Errorf("total net salary does not match the calculated value")
	}

	return db.CreatePayrollParams{
		EmpID:                      emp_id,
		Date:                       date,
		SalaryType:                 A.SalaryType,
		Amount:                     amount,
		SalaryAmountType:           A.SalaryAmountType,
		TotalOfSalaryAllowances:    total_of_salary_allowances,
		TotalAllowancesType:        A.TotalAllowancesType,
		Pension:                    A.Pension,
		PensionEmployer:            pension_employer,
		PensionEmployerType:        pension_employer_type,
		PensionEmployee:            pension_employee,
		PensionEmployeeType:        pension_employee_type,
		TotalNetSalary:             total_net_salary,
		TotalNetSalaryType:         A.TotalNetSalaryType,
		Tax:                        A.Tax,
		TaxPercentage:              tax_percentage,
		TotalNetSalaryAfterTax:     total_net_salary_after_tax,
		TotalNetSalaryAfterTaxType: A.TotalNetSalaryAfterTaxType,
		ErID:                       er_id,
		UpdatedBy:                  updated_by,
	}, nil
}

func (A *CreatePayrollReqModel) ToUpdatePayrollParams(id int64, admin_id int64) (db.UpdatePayrollParams, error) {

	var updated_by sql.NullInt64
	if A.UpdatedBy != nil {
		updated_by.Int64 = admin_id
		updated_by.Valid = true
	}
	var pension_employee_type sql.NullString
	if A.PensionEmployeeType != "" {
		pension_employee_type.String = A.PensionEmployeeType
		pension_employee_type.Valid = true
	}

	var pension_employer_type sql.NullString
	if A.PensionEmployerType != "" {
		pension_employer_type.String = A.PensionEmployerType
		pension_employer_type.Valid = true
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
		Date:                       date,
		SalaryType:                 A.SalaryType,
		Amount:                     amount,
		SalaryAmountType:           A.SalaryAmountType,
		TotalOfSalaryAllowances:    total_of_salary_allowances,
		TotalAllowancesType:        A.TotalAllowancesType,
		Pension:                    A.Pension,
		PensionEmployer:            pension_employer,
		PensionEmployerType:        pension_employer_type,
		PensionEmployee:            pension_employee,
		PensionEmployeeType:        pension_employee_type,
		TotalNetSalary:             total_net_salary,
		TotalNetSalaryType:         A.TotalNetSalaryType,
		Tax:                        A.Tax,
		TaxPercentage:              tax_percentage,
		TotalNetSalaryAfterTax:     total_net_salary_after_tax,
		TotalNetSalaryAfterTaxType: A.TotalNetSalaryAfterTaxType,
		UpdatedBy:                  updated_by,
		ID:                         id,
	}, nil
}

type CreatePayrollAllowancesParams struct {
	Name       string `json:"name"`
	Amount     string `json:"amount"`
	AmountType string `json:"amount_type"`
	PayrollID  int64  `json:"payroll_id"`
	UpdatedBy  *int64 `json:"updated_by"`
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
		Name:       A.Name,
		Amount:     amount,
		AmountType: A.AmountType,
		PayrollID:  A.PayrollID,
		UpdatedBy:  updated_by,
	}, nil
}

type HRTrainerComReqParams struct {
	IsTrainer      bool                        `json:"is_trainer"`
	TrainerComData CreateHRTrainerComReqParams `json:"trainer_com_data"`
}

type CreateHRTrainerComReqParams struct {
	PayrollID     int64   `json:"payroll_id"`
	TrainerID     int64   `json:"trainer_id"`
	EmployeeID    int64   `json:"employee_id"`
	Commission    float64 `json:"commission"`
	AssignedCount int64   `json:"assigned_count"`
	Total         float64 `json:"total"`
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
	Payroll    CreatePayrollReqModel           `json:"payroll"`
	Allowances []CreatePayrollAllowancesParams `json:"allowances"`
	TrainerCom HRTrainerComReqParams           `json:"trainer_com"`
}

type GetPayrollsReqModel struct {
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (A *GetPayrollsReqModel) ToGetPayrollsParams(id int64) (db.GetPayrollsParams, error) {

	return db.GetPayrollsParams{
		BranchID: id,
		Limit:    A.Limit,
		Offset:   A.Offset,
	}, nil
}

type TrainerComRow struct {
	Istrainer        bool    `json:"is_trainer"`
	TrainerID        int64   `json:"trainer_id"`
	Com_amount       float64 `json:"com_amount"`
	Assign_count     int64   `json:"assign_count"`
	Total_commission float64 `json:"total_commission"`
}
