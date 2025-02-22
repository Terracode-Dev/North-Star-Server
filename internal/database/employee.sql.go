// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: employee.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/shopspring/decimal"
)

const createEmpAccessiability = `-- name: CreateEmpAccessiability :exec
INSERT INTO HR_EMP_Accessiability (
    accessibility, accessibility_from, accessibility_till, enable, updated_by, employee_id
) VALUES (
    ?, ?, ?, ?, ?, ?
)
`

type CreateEmpAccessiabilityParams struct {
	Accessibility     bool          `json:"accessibility"`
	AccessibilityFrom time.Time     `json:"accessibility_from"`
	AccessibilityTill time.Time     `json:"accessibility_till"`
	Enable            bool          `json:"enable"`
	UpdatedBy         sql.NullInt64 `json:"updated_by"`
	EmployeeID        int64         `json:"employee_id"`
}

func (q *Queries) CreateEmpAccessiability(ctx context.Context, arg CreateEmpAccessiabilityParams) error {
	_, err := q.db.ExecContext(ctx, createEmpAccessiability,
		arg.Accessibility,
		arg.AccessibilityFrom,
		arg.AccessibilityTill,
		arg.Enable,
		arg.UpdatedBy,
		arg.EmployeeID,
	)
	return err
}

const createEmpAllowances = `-- name: CreateEmpAllowances :exec
INSERT INTO HR_EMP_Allowances (
   name, amount, updated_by, employee_id 
) VALUES (
    ?, ?, ?, ?
)
`

type CreateEmpAllowancesParams struct {
	Name       string          `json:"name"`
	Amount     decimal.Decimal `json:"amount"`
	UpdatedBy  sql.NullInt64   `json:"updated_by"`
	EmployeeID int64           `json:"employee_id"`
}

func (q *Queries) CreateEmpAllowances(ctx context.Context, arg CreateEmpAllowancesParams) error {
	_, err := q.db.ExecContext(ctx, createEmpAllowances,
		arg.Name,
		arg.Amount,
		arg.UpdatedBy,
		arg.EmployeeID,
	)
	return err
}

const createEmpBankDetails = `-- name: CreateEmpBankDetails :exec
INSERT INTO HR_EMP_Bank_Details (
    bank_name, branch_name, account_number, account_holder, employee_id, updated_by
) VALUES (
    ?, ?, ?, ?, ?, ?
)
`

type CreateEmpBankDetailsParams struct {
	BankName      string        `json:"bank_name"`
	BranchName    string        `json:"branch_name"`
	AccountNumber string        `json:"account_number"`
	AccountHolder string        `json:"account_holder"`
	EmployeeID    int64         `json:"employee_id"`
	UpdatedBy     sql.NullInt64 `json:"updated_by"`
}

func (q *Queries) CreateEmpBankDetails(ctx context.Context, arg CreateEmpBankDetailsParams) error {
	_, err := q.db.ExecContext(ctx, createEmpBankDetails,
		arg.BankName,
		arg.BranchName,
		arg.AccountNumber,
		arg.AccountHolder,
		arg.EmployeeID,
		arg.UpdatedBy,
	)
	return err
}

const createEmpBenifits = `-- name: CreateEmpBenifits :exec
INSERT INTO HR_EMP_Benifits (
    leave_status, leave_type, leave_count, health_insurance, insurance_from, insurance_till, retainment_plan, retainment_plan_from, retainment_plan_till, benifits, benifits_from, benifits_till, updated_by, employee_id
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
`

type CreateEmpBenifitsParams struct {
	LeaveStatus        bool          `json:"leave_status"`
	LeaveType          string        `json:"leave_type"`
	LeaveCount         int32         `json:"leave_count"`
	HealthInsurance    string        `json:"health_insurance"`
	InsuranceFrom      time.Time     `json:"insurance_from"`
	InsuranceTill      time.Time     `json:"insurance_till"`
	RetainmentPlan     string        `json:"retainment_plan"`
	RetainmentPlanFrom time.Time     `json:"retainment_plan_from"`
	RetainmentPlanTill time.Time     `json:"retainment_plan_till"`
	Benifits           string        `json:"benifits"`
	BenifitsFrom       time.Time     `json:"benifits_from"`
	BenifitsTill       time.Time     `json:"benifits_till"`
	UpdatedBy          sql.NullInt64 `json:"updated_by"`
	EmployeeID         int64         `json:"employee_id"`
}

func (q *Queries) CreateEmpBenifits(ctx context.Context, arg CreateEmpBenifitsParams) error {
	_, err := q.db.ExecContext(ctx, createEmpBenifits,
		arg.LeaveStatus,
		arg.LeaveType,
		arg.LeaveCount,
		arg.HealthInsurance,
		arg.InsuranceFrom,
		arg.InsuranceTill,
		arg.RetainmentPlan,
		arg.RetainmentPlanFrom,
		arg.RetainmentPlanTill,
		arg.Benifits,
		arg.BenifitsFrom,
		arg.BenifitsTill,
		arg.UpdatedBy,
		arg.EmployeeID,
	)
	return err
}

const createEmpCertificates = `-- name: CreateEmpCertificates :exec
INSERT INTO HR_EMP_Certificates (
    date, name, image_path, updated_by, employee_id
) VALUES (
    ?, ?, ?, ?, ?
)
`

type CreateEmpCertificatesParams struct {
	Date       time.Time     `json:"date"`
	Name       string        `json:"name"`
	ImagePath  string        `json:"image_path"`
	UpdatedBy  sql.NullInt64 `json:"updated_by"`
	EmployeeID int64         `json:"employee_id"`
}

func (q *Queries) CreateEmpCertificates(ctx context.Context, arg CreateEmpCertificatesParams) error {
	_, err := q.db.ExecContext(ctx, createEmpCertificates,
		arg.Date,
		arg.Name,
		arg.ImagePath,
		arg.UpdatedBy,
		arg.EmployeeID,
	)
	return err
}

const createEmpEmergencyDetails = `-- name: CreateEmpEmergencyDetails :exec
INSERT INTO HR_EMP_Emergency_Details (
    first_name, last_name, relationship, contact, employee_id, updated_by
) VALUES (
    ?, ?, ?, ?, ?, ?
)
`

type CreateEmpEmergencyDetailsParams struct {
	FirstName    string        `json:"first_name"`
	LastName     string        `json:"last_name"`
	Relationship string        `json:"relationship"`
	Contact      string        `json:"contact"`
	EmployeeID   int64         `json:"employee_id"`
	UpdatedBy    sql.NullInt64 `json:"updated_by"`
}

func (q *Queries) CreateEmpEmergencyDetails(ctx context.Context, arg CreateEmpEmergencyDetailsParams) error {
	_, err := q.db.ExecContext(ctx, createEmpEmergencyDetails,
		arg.FirstName,
		arg.LastName,
		arg.Relationship,
		arg.Contact,
		arg.EmployeeID,
		arg.UpdatedBy,
	)
	return err
}

const createEmpExpatriate = `-- name: CreateEmpExpatriate :exec
INSERT INTO HR_EMP_Expatriate (
    expatriate, nationality, visa_type, visa_from, visa_till, visa_number, visa_fee, visa_image_path, updated_by, employee_id
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
`

type CreateEmpExpatriateParams struct {
	Expatriate    bool            `json:"expatriate"`
	Nationality   string          `json:"nationality"`
	VisaType      string          `json:"visa_type"`
	VisaFrom      time.Time       `json:"visa_from"`
	VisaTill      time.Time       `json:"visa_till"`
	VisaNumber    string          `json:"visa_number"`
	VisaFee       decimal.Decimal `json:"visa_fee"`
	VisaImagePath string          `json:"visa_image_path"`
	UpdatedBy     sql.NullInt64   `json:"updated_by"`
	EmployeeID    int64           `json:"employee_id"`
}

func (q *Queries) CreateEmpExpatriate(ctx context.Context, arg CreateEmpExpatriateParams) error {
	_, err := q.db.ExecContext(ctx, createEmpExpatriate,
		arg.Expatriate,
		arg.Nationality,
		arg.VisaType,
		arg.VisaFrom,
		arg.VisaTill,
		arg.VisaNumber,
		arg.VisaFee,
		arg.VisaImagePath,
		arg.UpdatedBy,
		arg.EmployeeID,
	)
	return err
}

const createEmpSalary = `-- name: CreateEmpSalary :exec
INSERT INTO HR_EMP_Salary (
    salary_type, amount, Total_of_salary_allowances, pension_employer, pension_employee, total_net_salary, employee_id, updated_by
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?
)
`

type CreateEmpSalaryParams struct {
	SalaryType              string          `json:"salary_type"`
	Amount                  decimal.Decimal `json:"amount"`
	TotalOfSalaryAllowances decimal.Decimal `json:"total_of_salary_allowances"`
	PensionEmployer         decimal.Decimal `json:"pension_employer"`
	PensionEmployee         decimal.Decimal `json:"pension_employee"`
	TotalNetSalary          decimal.Decimal `json:"total_net_salary"`
	EmployeeID              int64           `json:"employee_id"`
	UpdatedBy               sql.NullInt64   `json:"updated_by"`
}

func (q *Queries) CreateEmpSalary(ctx context.Context, arg CreateEmpSalaryParams) error {
	_, err := q.db.ExecContext(ctx, createEmpSalary,
		arg.SalaryType,
		arg.Amount,
		arg.TotalOfSalaryAllowances,
		arg.PensionEmployer,
		arg.PensionEmployee,
		arg.TotalNetSalary,
		arg.EmployeeID,
		arg.UpdatedBy,
	)
	return err
}

const createEmpStatus = `-- name: CreateEmpStatus :exec
INSERT INTO HR_EMP_Status (
    status, department, designation, valid_from, valid_till, updated_by, employee_id
) VALUES (
    ?, ?, ?, ?, ?, ?, ?
)
`

type CreateEmpStatusParams struct {
	Status      string        `json:"status"`
	Department  string        `json:"department"`
	Designation string        `json:"designation"`
	ValidFrom   time.Time     `json:"valid_from"`
	ValidTill   time.Time     `json:"valid_till"`
	UpdatedBy   sql.NullInt64 `json:"updated_by"`
	EmployeeID  int64         `json:"employee_id"`
}

func (q *Queries) CreateEmpStatus(ctx context.Context, arg CreateEmpStatusParams) error {
	_, err := q.db.ExecContext(ctx, createEmpStatus,
		arg.Status,
		arg.Department,
		arg.Designation,
		arg.ValidFrom,
		arg.ValidTill,
		arg.UpdatedBy,
		arg.EmployeeID,
	)
	return err
}

const createEmpUser = `-- name: CreateEmpUser :exec
INSERT INTO HR_EMP_User (
    email, password, updated_by, employee_id, branch_id 
) VALUES (
    ?, ?, ?, ?,?
)
`

type CreateEmpUserParams struct {
	Email      string        `json:"email"`
	Password   string        `json:"password"`
	UpdatedBy  sql.NullInt64 `json:"updated_by"`
	EmployeeID int64         `json:"employee_id"`
	BranchID   int64         `json:"branch_id"`
}

func (q *Queries) CreateEmpUser(ctx context.Context, arg CreateEmpUserParams) error {
	_, err := q.db.ExecContext(ctx, createEmpUser,
		arg.Email,
		arg.Password,
		arg.UpdatedBy,
		arg.EmployeeID,
		arg.BranchID,
	)
	return err
}

const createEmployee = `-- name: CreateEmployee :execresult
INSERT INTO HR_Employee (
    first_name, last_name, gender, dob, religion, primary_number, secondary_number, 
    passport_id, nationality, passport_valid_till, nic, country, nic_valid_till, 
    address, current_country, email, updated_by
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
`

type CreateEmployeeParams struct {
	FirstName         string        `json:"first_name"`
	LastName          string        `json:"last_name"`
	Gender            string        `json:"gender"`
	Dob               time.Time     `json:"dob"`
	Religion          string        `json:"religion"`
	PrimaryNumber     string        `json:"primary_number"`
	SecondaryNumber   string        `json:"secondary_number"`
	PassportID        string        `json:"passport_id"`
	Nationality       string        `json:"nationality"`
	PassportValidTill time.Time     `json:"passport_valid_till"`
	Nic               string        `json:"nic"`
	Country           string        `json:"country"`
	NicValidTill      time.Time     `json:"nic_valid_till"`
	Address           string        `json:"address"`
	CurrentCountry    string        `json:"current_country"`
	Email             string        `json:"email"`
	UpdatedBy         sql.NullInt64 `json:"updated_by"`
}

func (q *Queries) CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createEmployee,
		arg.FirstName,
		arg.LastName,
		arg.Gender,
		arg.Dob,
		arg.Religion,
		arg.PrimaryNumber,
		arg.SecondaryNumber,
		arg.PassportID,
		arg.Nationality,
		arg.PassportValidTill,
		arg.Nic,
		arg.Country,
		arg.NicValidTill,
		arg.Address,
		arg.CurrentCountry,
		arg.Email,
		arg.UpdatedBy,
	)
}

const deleteEmpAccessiability = `-- name: DeleteEmpAccessiability :exec
DELETE FROM HR_EMP_Accessiability WHERE employee_id = ?
`

func (q *Queries) DeleteEmpAccessiability(ctx context.Context, employeeID int64) error {
	_, err := q.db.ExecContext(ctx, deleteEmpAccessiability, employeeID)
	return err
}

const deleteEmpAllowances = `-- name: DeleteEmpAllowances :exec
DELETE FROM HR_EMP_Allowances WHERE employee_id = ?
`

func (q *Queries) DeleteEmpAllowances(ctx context.Context, employeeID int64) error {
	_, err := q.db.ExecContext(ctx, deleteEmpAllowances, employeeID)
	return err
}

const deleteEmpBankDetails = `-- name: DeleteEmpBankDetails :exec
DELETE FROM HR_EMP_Bank_Details WHERE employee_id = ?
`

func (q *Queries) DeleteEmpBankDetails(ctx context.Context, employeeID int64) error {
	_, err := q.db.ExecContext(ctx, deleteEmpBankDetails, employeeID)
	return err
}

const deleteEmpBenifits = `-- name: DeleteEmpBenifits :exec
DELETE FROM HR_EMP_Benifits WHERE employee_id = ?
`

func (q *Queries) DeleteEmpBenifits(ctx context.Context, employeeID int64) error {
	_, err := q.db.ExecContext(ctx, deleteEmpBenifits, employeeID)
	return err
}

const deleteEmpCertificates = `-- name: DeleteEmpCertificates :exec
DELETE FROM HR_EMP_Certificates WHERE employee_id = ?
`

func (q *Queries) DeleteEmpCertificates(ctx context.Context, employeeID int64) error {
	_, err := q.db.ExecContext(ctx, deleteEmpCertificates, employeeID)
	return err
}

const deleteEmpEmergencyDetails = `-- name: DeleteEmpEmergencyDetails :exec
DELETE FROM HR_EMP_Emergency_Details WHERE employee_id = ?
`

func (q *Queries) DeleteEmpEmergencyDetails(ctx context.Context, employeeID int64) error {
	_, err := q.db.ExecContext(ctx, deleteEmpEmergencyDetails, employeeID)
	return err
}

const deleteEmpExpatriate = `-- name: DeleteEmpExpatriate :exec
DELETE FROM HR_EMP_Expatriate WHERE employee_id = ?
`

func (q *Queries) DeleteEmpExpatriate(ctx context.Context, employeeID int64) error {
	_, err := q.db.ExecContext(ctx, deleteEmpExpatriate, employeeID)
	return err
}

const deleteEmpSalary = `-- name: DeleteEmpSalary :exec
DELETE FROM HR_EMP_Salary WHERE employee_id = ?
`

func (q *Queries) DeleteEmpSalary(ctx context.Context, employeeID int64) error {
	_, err := q.db.ExecContext(ctx, deleteEmpSalary, employeeID)
	return err
}

const deleteEmpStatus = `-- name: DeleteEmpStatus :exec
DELETE FROM HR_EMP_Status WHERE employee_id = ?
`

func (q *Queries) DeleteEmpStatus(ctx context.Context, employeeID int64) error {
	_, err := q.db.ExecContext(ctx, deleteEmpStatus, employeeID)
	return err
}

const deleteEmpUser = `-- name: DeleteEmpUser :exec
DELETE FROM HR_EMP_User WHERE employee_id = ?
`

func (q *Queries) DeleteEmpUser(ctx context.Context, employeeID int64) error {
	_, err := q.db.ExecContext(ctx, deleteEmpUser, employeeID)
	return err
}

const deleteEmployee = `-- name: DeleteEmployee :exec
DELETE FROM HR_Employee WHERE id = ?
`

func (q *Queries) DeleteEmployee(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteEmployee, id)
	return err
}

const employeeLogin = `-- name: EmployeeLogin :one
SELECT employee_id, password, email, branch_id
FROM HR_EMP_User
WHERE email = ?
`

type EmployeeLoginRow struct {
	EmployeeID int64  `json:"employee_id"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	BranchID   int64  `json:"branch_id"`
}

func (q *Queries) EmployeeLogin(ctx context.Context, email string) (EmployeeLoginRow, error) {
	row := q.db.QueryRowContext(ctx, employeeLogin, email)
	var i EmployeeLoginRow
	err := row.Scan(
		&i.EmployeeID,
		&i.Password,
		&i.Email,
		&i.BranchID,
	)
	return i, err
}

const getEmployee = `-- name: GetEmployee :many
SELECT
  e.id AS employee_id,
  e.first_name,
  e.last_name,
  usr.email AS user_email,
  br.name AS branch_name
FROM HR_Employee e
LEFT JOIN HR_EMP_User usr ON e.id = usr.employee_id
LEFT JOIN HR_Branch br ON usr.branch_id = br.id
WHERE 
  (
    CAST(e.id AS CHAR) LIKE CONCAT('%', ?, '%')
    OR e.first_name LIKE CONCAT('%', ?, '%')
    OR e.last_name  LIKE CONCAT('%', ?, '%')
    OR usr.email    LIKE CONCAT('%', ?, '%')
    OR br.name      LIKE CONCAT('%', ?, '%')
  )
  AND (? = '' OR br.id = ?)
ORDER BY e.id DESC
LIMIT ? OFFSET ?
`

type GetEmployeeParams struct {
	CONCAT   interface{} `json:"CONCAT"`
	CONCAT_2 interface{} `json:"CONCAT_2"`
	CONCAT_3 interface{} `json:"CONCAT_3"`
	CONCAT_4 interface{} `json:"CONCAT_4"`
	CONCAT_5 interface{} `json:"CONCAT_5"`
	Column6  interface{} `json:"column_6"`
	ID       int64       `json:"id"`
	Limit    int32       `json:"limit"`
	Offset   int32       `json:"offset"`
}

type GetEmployeeRow struct {
	EmployeeID int64          `json:"employee_id"`
	FirstName  string         `json:"first_name"`
	LastName   string         `json:"last_name"`
	UserEmail  sql.NullString `json:"user_email"`
	BranchName sql.NullString `json:"branch_name"`
}

func (q *Queries) GetEmployee(ctx context.Context, arg GetEmployeeParams) ([]GetEmployeeRow, error) {
	rows, err := q.db.QueryContext(ctx, getEmployee,
		arg.CONCAT,
		arg.CONCAT_2,
		arg.CONCAT_3,
		arg.CONCAT_4,
		arg.CONCAT_5,
		arg.Column6,
		arg.ID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetEmployeeRow
	for rows.Next() {
		var i GetEmployeeRow
		if err := rows.Scan(
			&i.EmployeeID,
			&i.FirstName,
			&i.LastName,
			&i.UserEmail,
			&i.BranchName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEmployeeByID = `-- name: GetEmployeeByID :many
SELECT 
    e.id AS employee_id, 
    e.first_name, 
    e.last_name, 
    e.gender, 
    e.dob, 
    e.religion, 
    e.primary_number, 
    e.secondary_number, 
    e.passport_id, 
    e.nationality, 
    e.passport_valid_till, 
    e.nic, 
    e.country, 
    e.nic_valid_till, 
    e.address,
    e.current_country, 
    e.email, 
    e.updated_by, 
    e.created_at, 
    e.updated_at,

    ed.first_name AS emergency_first_name, 
    ed.last_name AS emergency_last_name, 
    ed.relationship, 
    ed.contact AS emergency_contact, 

    bd.bank_name, 
    bd.branch_name, 
    bd.account_number, 
    bd.account_holder,

    s.salary_type, 
    s.amount, 
    s.Total_of_salary_allowances, 
    s.pension_employer, 
    s.pension_employee, 
    s.total_net_salary, 

    cert.date AS certificate_date, 
    cert.name AS certificate_name, 
    cert.image_path AS certificate_image, 

    stat.status, 
    stat.department, 
    stat.designation, 
    stat.valid_from AS status_valid_from, 
    stat.valid_till AS status_valid_till, 

    ben.leave_status, 
    ben.leave_type, 
    ben.leave_count, 
    ben.health_insurance, 
    ben.insurance_from, 
    ben.insurance_till, 
    ben.retainment_plan, 
    ben.retainment_plan_from, 
    ben.retainment_plan_till, 
    ben.benifits, 
    ben.benifits_from, 
    ben.benifits_till, 

    usr.email AS user_email, 
    usr.password AS user_password, 
    usr.branch_id AS user_branch_id,

    allw.name AS allowance_name, 
    allw.amount AS allowance_amount, 

    exp.expatriate, 
    exp.nationality AS exp_nationality, 
    exp.visa_type, 
    exp.visa_from, 
    exp.visa_till, 
    exp.visa_number, 
    exp.visa_fee, 
    exp.visa_image_path, 

    acc.accessibility, 
    acc.accessibility_from, 
    acc.accessibility_till, 
    acc.enable 

FROM HR_Employee e
LEFT JOIN HR_EMP_Emergency_Details ed ON e.id = ed.employee_id
LEFT JOIN HR_EMP_Bank_Details bd ON e.id = bd.employee_id
LEFT JOIN HR_EMP_Salary s ON e.id = s.employee_id
LEFT JOIN HR_EMP_Certificates cert ON e.id = cert.employee_id
LEFT JOIN HR_EMP_Status stat ON e.id = stat.employee_id
LEFT JOIN HR_EMP_Benifits ben ON e.id = ben.employee_id
LEFT JOIN HR_EMP_User usr ON e.id = usr.employee_id
LEFT JOIN HR_EMP_Allowances allw ON e.id = allw.employee_id
LEFT JOIN HR_EMP_Expatriate exp ON e.id = exp.employee_id
LEFT JOIN HR_EMP_Accessiability acc ON e.id = acc.employee_id

WHERE e.id = ?
`

type GetEmployeeByIDRow struct {
	EmployeeID              int64          `json:"employee_id"`
	FirstName               string         `json:"first_name"`
	LastName                string         `json:"last_name"`
	Gender                  string         `json:"gender"`
	Dob                     time.Time      `json:"dob"`
	Religion                string         `json:"religion"`
	PrimaryNumber           string         `json:"primary_number"`
	SecondaryNumber         string         `json:"secondary_number"`
	PassportID              string         `json:"passport_id"`
	Nationality             string         `json:"nationality"`
	PassportValidTill       time.Time      `json:"passport_valid_till"`
	Nic                     string         `json:"nic"`
	Country                 string         `json:"country"`
	NicValidTill            time.Time      `json:"nic_valid_till"`
	Address                 string         `json:"address"`
	CurrentCountry          string         `json:"current_country"`
	Email                   string         `json:"email"`
	UpdatedBy               sql.NullInt64  `json:"updated_by"`
	CreatedAt               sql.NullTime   `json:"created_at"`
	UpdatedAt               sql.NullTime   `json:"updated_at"`
	EmergencyFirstName      sql.NullString `json:"emergency_first_name"`
	EmergencyLastName       sql.NullString `json:"emergency_last_name"`
	Relationship            sql.NullString `json:"relationship"`
	EmergencyContact        sql.NullString `json:"emergency_contact"`
	BankName                sql.NullString `json:"bank_name"`
	BranchName              sql.NullString `json:"branch_name"`
	AccountNumber           sql.NullString `json:"account_number"`
	AccountHolder           sql.NullString `json:"account_holder"`
	SalaryType              sql.NullString `json:"salary_type"`
	Amount                  sql.NullString `json:"amount"`
	TotalOfSalaryAllowances sql.NullString `json:"total_of_salary_allowances"`
	PensionEmployer         sql.NullString `json:"pension_employer"`
	PensionEmployee         sql.NullString `json:"pension_employee"`
	TotalNetSalary          sql.NullString `json:"total_net_salary"`
	CertificateDate         sql.NullTime   `json:"certificate_date"`
	CertificateName         sql.NullString `json:"certificate_name"`
	CertificateImage        sql.NullString `json:"certificate_image"`
	Status                  sql.NullString `json:"status"`
	Department              sql.NullString `json:"department"`
	Designation             sql.NullString `json:"designation"`
	StatusValidFrom         sql.NullTime   `json:"status_valid_from"`
	StatusValidTill         sql.NullTime   `json:"status_valid_till"`
	LeaveStatus             sql.NullBool   `json:"leave_status"`
	LeaveType               sql.NullString `json:"leave_type"`
	LeaveCount              sql.NullInt32  `json:"leave_count"`
	HealthInsurance         sql.NullString `json:"health_insurance"`
	InsuranceFrom           sql.NullTime   `json:"insurance_from"`
	InsuranceTill           sql.NullTime   `json:"insurance_till"`
	RetainmentPlan          sql.NullString `json:"retainment_plan"`
	RetainmentPlanFrom      sql.NullTime   `json:"retainment_plan_from"`
	RetainmentPlanTill      sql.NullTime   `json:"retainment_plan_till"`
	Benifits                sql.NullString `json:"benifits"`
	BenifitsFrom            sql.NullTime   `json:"benifits_from"`
	BenifitsTill            sql.NullTime   `json:"benifits_till"`
	UserEmail               sql.NullString `json:"user_email"`
	UserPassword            sql.NullString `json:"user_password"`
	UserBranchID            sql.NullInt64  `json:"user_branch_id"`
	AllowanceName           sql.NullString `json:"allowance_name"`
	AllowanceAmount         sql.NullString `json:"allowance_amount"`
	Expatriate              sql.NullBool   `json:"expatriate"`
	ExpNationality          sql.NullString `json:"exp_nationality"`
	VisaType                sql.NullString `json:"visa_type"`
	VisaFrom                sql.NullTime   `json:"visa_from"`
	VisaTill                sql.NullTime   `json:"visa_till"`
	VisaNumber              sql.NullString `json:"visa_number"`
	VisaFee                 sql.NullString `json:"visa_fee"`
	VisaImagePath           sql.NullString `json:"visa_image_path"`
	Accessibility           sql.NullBool   `json:"accessibility"`
	AccessibilityFrom       sql.NullTime   `json:"accessibility_from"`
	AccessibilityTill       sql.NullTime   `json:"accessibility_till"`
	Enable                  sql.NullBool   `json:"enable"`
}

func (q *Queries) GetEmployeeByID(ctx context.Context, id int64) ([]GetEmployeeByIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getEmployeeByID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetEmployeeByIDRow
	for rows.Next() {
		var i GetEmployeeByIDRow
		if err := rows.Scan(
			&i.EmployeeID,
			&i.FirstName,
			&i.LastName,
			&i.Gender,
			&i.Dob,
			&i.Religion,
			&i.PrimaryNumber,
			&i.SecondaryNumber,
			&i.PassportID,
			&i.Nationality,
			&i.PassportValidTill,
			&i.Nic,
			&i.Country,
			&i.NicValidTill,
			&i.Address,
			&i.CurrentCountry,
			&i.Email,
			&i.UpdatedBy,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.EmergencyFirstName,
			&i.EmergencyLastName,
			&i.Relationship,
			&i.EmergencyContact,
			&i.BankName,
			&i.BranchName,
			&i.AccountNumber,
			&i.AccountHolder,
			&i.SalaryType,
			&i.Amount,
			&i.TotalOfSalaryAllowances,
			&i.PensionEmployer,
			&i.PensionEmployee,
			&i.TotalNetSalary,
			&i.CertificateDate,
			&i.CertificateName,
			&i.CertificateImage,
			&i.Status,
			&i.Department,
			&i.Designation,
			&i.StatusValidFrom,
			&i.StatusValidTill,
			&i.LeaveStatus,
			&i.LeaveType,
			&i.LeaveCount,
			&i.HealthInsurance,
			&i.InsuranceFrom,
			&i.InsuranceTill,
			&i.RetainmentPlan,
			&i.RetainmentPlanFrom,
			&i.RetainmentPlanTill,
			&i.Benifits,
			&i.BenifitsFrom,
			&i.BenifitsTill,
			&i.UserEmail,
			&i.UserPassword,
			&i.UserBranchID,
			&i.AllowanceName,
			&i.AllowanceAmount,
			&i.Expatriate,
			&i.ExpNationality,
			&i.VisaType,
			&i.VisaFrom,
			&i.VisaTill,
			&i.VisaNumber,
			&i.VisaFee,
			&i.VisaImagePath,
			&i.Accessibility,
			&i.AccessibilityFrom,
			&i.AccessibilityTill,
			&i.Enable,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEmployeeDOB = `-- name: GetEmployeeDOB :one
SELECT dob FROM HR_Employee WHERE id = ?
`

func (q *Queries) GetEmployeeDOB(ctx context.Context, id int64) (time.Time, error) {
	row := q.db.QueryRowContext(ctx, getEmployeeDOB, id)
	var dob time.Time
	err := row.Scan(&dob)
	return dob, err
}

const updateEmpAccessiability = `-- name: UpdateEmpAccessiability :exec
UPDATE HR_EMP_Accessiability SET
    accessibility = ?, accessibility_from = ?, accessibility_till = ?, enable = ?, updated_by = ?
WHERE employee_id = ?
`

type UpdateEmpAccessiabilityParams struct {
	Accessibility     bool          `json:"accessibility"`
	AccessibilityFrom time.Time     `json:"accessibility_from"`
	AccessibilityTill time.Time     `json:"accessibility_till"`
	Enable            bool          `json:"enable"`
	UpdatedBy         sql.NullInt64 `json:"updated_by"`
	EmployeeID        int64         `json:"employee_id"`
}

func (q *Queries) UpdateEmpAccessiability(ctx context.Context, arg UpdateEmpAccessiabilityParams) error {
	_, err := q.db.ExecContext(ctx, updateEmpAccessiability,
		arg.Accessibility,
		arg.AccessibilityFrom,
		arg.AccessibilityTill,
		arg.Enable,
		arg.UpdatedBy,
		arg.EmployeeID,
	)
	return err
}

const updateEmpAllowances = `-- name: UpdateEmpAllowances :exec
UPDATE HR_EMP_Allowances SET
    name = ?, amount = ?, updated_by = ?
WHERE employee_id = ?
`

type UpdateEmpAllowancesParams struct {
	Name       string          `json:"name"`
	Amount     decimal.Decimal `json:"amount"`
	UpdatedBy  sql.NullInt64   `json:"updated_by"`
	EmployeeID int64           `json:"employee_id"`
}

func (q *Queries) UpdateEmpAllowances(ctx context.Context, arg UpdateEmpAllowancesParams) error {
	_, err := q.db.ExecContext(ctx, updateEmpAllowances,
		arg.Name,
		arg.Amount,
		arg.UpdatedBy,
		arg.EmployeeID,
	)
	return err
}

const updateEmpBankDetails = `-- name: UpdateEmpBankDetails :exec
UPDATE HR_EMP_Bank_Details SET
    bank_name = ?, branch_name = ?, account_number = ?, account_holder = ?, updated_by = ?
WHERE employee_id = ?
`

type UpdateEmpBankDetailsParams struct {
	BankName      string        `json:"bank_name"`
	BranchName    string        `json:"branch_name"`
	AccountNumber string        `json:"account_number"`
	AccountHolder string        `json:"account_holder"`
	UpdatedBy     sql.NullInt64 `json:"updated_by"`
	EmployeeID    int64         `json:"employee_id"`
}

func (q *Queries) UpdateEmpBankDetails(ctx context.Context, arg UpdateEmpBankDetailsParams) error {
	_, err := q.db.ExecContext(ctx, updateEmpBankDetails,
		arg.BankName,
		arg.BranchName,
		arg.AccountNumber,
		arg.AccountHolder,
		arg.UpdatedBy,
		arg.EmployeeID,
	)
	return err
}

const updateEmpBenifits = `-- name: UpdateEmpBenifits :exec
UPDATE HR_EMP_Benifits SET
    leave_status = ?, leave_type = ?, leave_count = ?, health_insurance = ?, 
    insurance_from = ?, insurance_till = ?, retainment_plan = ?, retainment_plan_from = ?, 
    retainment_plan_till = ?, benifits = ?, benifits_from = ?, benifits_till = ?, updated_by = ?
WHERE employee_id = ?
`

type UpdateEmpBenifitsParams struct {
	LeaveStatus        bool          `json:"leave_status"`
	LeaveType          string        `json:"leave_type"`
	LeaveCount         int32         `json:"leave_count"`
	HealthInsurance    string        `json:"health_insurance"`
	InsuranceFrom      time.Time     `json:"insurance_from"`
	InsuranceTill      time.Time     `json:"insurance_till"`
	RetainmentPlan     string        `json:"retainment_plan"`
	RetainmentPlanFrom time.Time     `json:"retainment_plan_from"`
	RetainmentPlanTill time.Time     `json:"retainment_plan_till"`
	Benifits           string        `json:"benifits"`
	BenifitsFrom       time.Time     `json:"benifits_from"`
	BenifitsTill       time.Time     `json:"benifits_till"`
	UpdatedBy          sql.NullInt64 `json:"updated_by"`
	EmployeeID         int64         `json:"employee_id"`
}

func (q *Queries) UpdateEmpBenifits(ctx context.Context, arg UpdateEmpBenifitsParams) error {
	_, err := q.db.ExecContext(ctx, updateEmpBenifits,
		arg.LeaveStatus,
		arg.LeaveType,
		arg.LeaveCount,
		arg.HealthInsurance,
		arg.InsuranceFrom,
		arg.InsuranceTill,
		arg.RetainmentPlan,
		arg.RetainmentPlanFrom,
		arg.RetainmentPlanTill,
		arg.Benifits,
		arg.BenifitsFrom,
		arg.BenifitsTill,
		arg.UpdatedBy,
		arg.EmployeeID,
	)
	return err
}

const updateEmpCertificates = `-- name: UpdateEmpCertificates :exec
UPDATE HR_EMP_Certificates SET
    date = ?, name = ?, image_path = ?, updated_by = ?
WHERE employee_id = ?
`

type UpdateEmpCertificatesParams struct {
	Date       time.Time     `json:"date"`
	Name       string        `json:"name"`
	ImagePath  string        `json:"image_path"`
	UpdatedBy  sql.NullInt64 `json:"updated_by"`
	EmployeeID int64         `json:"employee_id"`
}

func (q *Queries) UpdateEmpCertificates(ctx context.Context, arg UpdateEmpCertificatesParams) error {
	_, err := q.db.ExecContext(ctx, updateEmpCertificates,
		arg.Date,
		arg.Name,
		arg.ImagePath,
		arg.UpdatedBy,
		arg.EmployeeID,
	)
	return err
}

const updateEmpEmergencyDetails = `-- name: UpdateEmpEmergencyDetails :exec
UPDATE HR_EMP_Emergency_Details SET
    first_name = ?, last_name = ?, relationship = ?, contact = ?, updated_by = ?
WHERE employee_id = ?
`

type UpdateEmpEmergencyDetailsParams struct {
	FirstName    string        `json:"first_name"`
	LastName     string        `json:"last_name"`
	Relationship string        `json:"relationship"`
	Contact      string        `json:"contact"`
	UpdatedBy    sql.NullInt64 `json:"updated_by"`
	EmployeeID   int64         `json:"employee_id"`
}

func (q *Queries) UpdateEmpEmergencyDetails(ctx context.Context, arg UpdateEmpEmergencyDetailsParams) error {
	_, err := q.db.ExecContext(ctx, updateEmpEmergencyDetails,
		arg.FirstName,
		arg.LastName,
		arg.Relationship,
		arg.Contact,
		arg.UpdatedBy,
		arg.EmployeeID,
	)
	return err
}

const updateEmpExpatriate = `-- name: UpdateEmpExpatriate :exec
UPDATE HR_EMP_Expatriate SET
    expatriate = ?, nationality = ?, visa_type = ?, visa_from = ?, visa_till = ?, 
    visa_number = ?, visa_fee = ?, visa_image_path = ?, updated_by = ?
WHERE employee_id = ?
`

type UpdateEmpExpatriateParams struct {
	Expatriate    bool            `json:"expatriate"`
	Nationality   string          `json:"nationality"`
	VisaType      string          `json:"visa_type"`
	VisaFrom      time.Time       `json:"visa_from"`
	VisaTill      time.Time       `json:"visa_till"`
	VisaNumber    string          `json:"visa_number"`
	VisaFee       decimal.Decimal `json:"visa_fee"`
	VisaImagePath string          `json:"visa_image_path"`
	UpdatedBy     sql.NullInt64   `json:"updated_by"`
	EmployeeID    int64           `json:"employee_id"`
}

func (q *Queries) UpdateEmpExpatriate(ctx context.Context, arg UpdateEmpExpatriateParams) error {
	_, err := q.db.ExecContext(ctx, updateEmpExpatriate,
		arg.Expatriate,
		arg.Nationality,
		arg.VisaType,
		arg.VisaFrom,
		arg.VisaTill,
		arg.VisaNumber,
		arg.VisaFee,
		arg.VisaImagePath,
		arg.UpdatedBy,
		arg.EmployeeID,
	)
	return err
}

const updateEmpSalary = `-- name: UpdateEmpSalary :exec
UPDATE HR_EMP_Salary SET
    salary_type = ?, amount = ?, Total_of_salary_allowances = ?, pension_employer = ?, 
    pension_employee = ?, total_net_salary = ?, updated_by = ?
WHERE employee_id = ?
`

type UpdateEmpSalaryParams struct {
	SalaryType              string          `json:"salary_type"`
	Amount                  decimal.Decimal `json:"amount"`
	TotalOfSalaryAllowances decimal.Decimal `json:"total_of_salary_allowances"`
	PensionEmployer         decimal.Decimal `json:"pension_employer"`
	PensionEmployee         decimal.Decimal `json:"pension_employee"`
	TotalNetSalary          decimal.Decimal `json:"total_net_salary"`
	UpdatedBy               sql.NullInt64   `json:"updated_by"`
	EmployeeID              int64           `json:"employee_id"`
}

func (q *Queries) UpdateEmpSalary(ctx context.Context, arg UpdateEmpSalaryParams) error {
	_, err := q.db.ExecContext(ctx, updateEmpSalary,
		arg.SalaryType,
		arg.Amount,
		arg.TotalOfSalaryAllowances,
		arg.PensionEmployer,
		arg.PensionEmployee,
		arg.TotalNetSalary,
		arg.UpdatedBy,
		arg.EmployeeID,
	)
	return err
}

const updateEmpStatus = `-- name: UpdateEmpStatus :exec
UPDATE HR_EMP_Status SET
    status = ?, department = ?, designation = ?, valid_from = ?, valid_till = ?, updated_by = ?
WHERE employee_id = ?
`

type UpdateEmpStatusParams struct {
	Status      string        `json:"status"`
	Department  string        `json:"department"`
	Designation string        `json:"designation"`
	ValidFrom   time.Time     `json:"valid_from"`
	ValidTill   time.Time     `json:"valid_till"`
	UpdatedBy   sql.NullInt64 `json:"updated_by"`
	EmployeeID  int64         `json:"employee_id"`
}

func (q *Queries) UpdateEmpStatus(ctx context.Context, arg UpdateEmpStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateEmpStatus,
		arg.Status,
		arg.Department,
		arg.Designation,
		arg.ValidFrom,
		arg.ValidTill,
		arg.UpdatedBy,
		arg.EmployeeID,
	)
	return err
}

const updateEmpUser = `-- name: UpdateEmpUser :exec
UPDATE HR_EMP_User SET
    email = ?, password = ?, updated_by = ?
WHERE employee_id = ?
`

type UpdateEmpUserParams struct {
	Email      string        `json:"email"`
	Password   string        `json:"password"`
	UpdatedBy  sql.NullInt64 `json:"updated_by"`
	EmployeeID int64         `json:"employee_id"`
}

func (q *Queries) UpdateEmpUser(ctx context.Context, arg UpdateEmpUserParams) error {
	_, err := q.db.ExecContext(ctx, updateEmpUser,
		arg.Email,
		arg.Password,
		arg.UpdatedBy,
		arg.EmployeeID,
	)
	return err
}

const updateEmployee = `-- name: UpdateEmployee :exec
UPDATE HR_Employee SET 
    first_name = ?, last_name = ?, gender = ?, dob = ?, religion = ?, primary_number = ?, secondary_number = ?,
    passport_id = ?, nationality = ?, passport_valid_till = ?, nic = ?, country = ?, nic_valid_till = ?, address = ?, current_country = ?,
    email = ?, updated_by = ?
WHERE id = ?
`

type UpdateEmployeeParams struct {
	FirstName         string        `json:"first_name"`
	LastName          string        `json:"last_name"`
	Gender            string        `json:"gender"`
	Dob               time.Time     `json:"dob"`
	Religion          string        `json:"religion"`
	PrimaryNumber     string        `json:"primary_number"`
	SecondaryNumber   string        `json:"secondary_number"`
	PassportID        string        `json:"passport_id"`
	Nationality       string        `json:"nationality"`
	PassportValidTill time.Time     `json:"passport_valid_till"`
	Nic               string        `json:"nic"`
	Country           string        `json:"country"`
	NicValidTill      time.Time     `json:"nic_valid_till"`
	Address           string        `json:"address"`
	CurrentCountry    string        `json:"current_country"`
	Email             string        `json:"email"`
	UpdatedBy         sql.NullInt64 `json:"updated_by"`
	ID                int64         `json:"id"`
}

func (q *Queries) UpdateEmployee(ctx context.Context, arg UpdateEmployeeParams) error {
	_, err := q.db.ExecContext(ctx, updateEmployee,
		arg.FirstName,
		arg.LastName,
		arg.Gender,
		arg.Dob,
		arg.Religion,
		arg.PrimaryNumber,
		arg.SecondaryNumber,
		arg.PassportID,
		arg.Nationality,
		arg.PassportValidTill,
		arg.Nic,
		arg.Country,
		arg.NicValidTill,
		arg.Address,
		arg.CurrentCountry,
		arg.Email,
		arg.UpdatedBy,
		arg.ID,
	)
	return err
}
