package hr

import (
	"database/sql"
	"log"
	"time"

	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
	rba "github.com/Terracode-Dev/North-Star-Server/internal/pkg/RBA"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/bcrypt"
)

type CreateEmployeeReqModel struct {
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Gender            string `json:"gender"`
	Dob               string `json:"dob"`
	Religion          string `json:"religion"`
	PrimaryNumber     string `json:"primary_number"`
	SecondaryNumber   string `json:"secondary_number"`
	PassportID        string `json:"passport_id"`
	Nationality       string `json:"nationality"`
	PassportValidTill string `json:"passport_valid_till"`
	Nic               string `json:"nic"`
	Country           string `json:"country"`
	NicValidTill      string `json:"nic_valid_till"`
	Address           string `json:"address"`
	CurrentCountry    string `json:"current_country"`
	Email             string `json:"email"`
	UpdatedBy         *int64 `json:"updated_by"`
}

func (M CreateEmployeeReqModel) convertToDbStruct(admin_id int64) (db.CreateEmployeeParams, error) {
	dob, err := time.Parse(time.RFC3339, M.Dob)
	if err != nil {
		log.Printf("Error parsing dob: %v", err)
		return db.CreateEmployeeParams{}, err
	}

	passportValidTill, err := time.Parse(time.RFC3339, M.PassportValidTill)
	if err != nil {
		log.Printf("Error parsing passport valid till: %v", err)
		return db.CreateEmployeeParams{}, err
	}

	nicValidTill, err := time.Parse(time.RFC3339, M.NicValidTill)
	if err != nil {
		log.Printf("Error parsing nic valid till: %v", err)
		return db.CreateEmployeeParams{}, err
	}

	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true

	return db.CreateEmployeeParams{
		FirstName:         M.FirstName,
		LastName:          M.LastName,
		Gender:            M.Gender,
		Dob:               dob,
		Religion:          M.Religion,
		PrimaryNumber:     M.PrimaryNumber,
		SecondaryNumber:   M.SecondaryNumber,
		PassportID:        M.PassportID,
		Nationality:       M.Nationality,
		PassportValidTill: passportValidTill,
		Nic:               M.Nic,
		Country:           M.Country,
		NicValidTill:      nicValidTill,
		Address:           M.Address,
		CurrentCountry:    M.CurrentCountry,
		Email:             M.Email,
		UpdatedBy:         updated_by,
	}, nil
}

func (M CreateEmployeeReqModel) ConvertToUpdateDbStruct(emp_id int64, admin_id int64) (db.UpdateEmployeeParams, error) {
	dob, err := time.Parse(time.RFC3339, M.Dob)
	if err != nil {
		log.Printf("Error parsing dob: %v", err)
		return db.UpdateEmployeeParams{}, err
	}

	passportValidTill, err := time.Parse(time.RFC3339, M.PassportValidTill)
	if err != nil {
		log.Printf("Error parsing passport valid till: %v", err)
		return db.UpdateEmployeeParams{}, err
	}

	nicValidTill, err := time.Parse(time.RFC3339, M.NicValidTill)
	if err != nil {
		log.Printf("Error parsing nic valid till: %v", err)
		return db.UpdateEmployeeParams{}, err
	}

	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true

	return db.UpdateEmployeeParams{
		FirstName:         M.FirstName,
		LastName:          M.LastName,
		Gender:            M.Gender,
		Dob:               dob,
		Religion:          M.Religion,
		PrimaryNumber:     M.PrimaryNumber,
		SecondaryNumber:   M.SecondaryNumber,
		PassportID:        M.PassportID,
		Nationality:       M.Nationality,
		PassportValidTill: passportValidTill,
		Nic:               M.Nic,
		Country:           M.Country,
		NicValidTill:      nicValidTill,
		Address:           M.Address,
		CurrentCountry:    M.CurrentCountry,
		Email:             M.Email,
		UpdatedBy:         updated_by,
		ID:                emp_id,
	}, nil
}

type CreateEmpEmergencyDetailsReqModel struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Relationship string `json:"relationship"`
	Contact      string `json:"contact"`
	EmployeeID   int64  `json:"employee_id"`
	UpdatedBy    *int64 `json:"updated_by"`
}

func (M CreateEmpEmergencyDetailsReqModel) convertToDbStruct(admin_id int64) (db.CreateEmpEmergencyDetailsParams, error) {
	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = admin_id
		updated_by.Valid = true
	}

	return db.CreateEmpEmergencyDetailsParams{
		FirstName:    M.FirstName,
		LastName:     M.LastName,
		Relationship: M.Relationship,
		Contact:      M.Contact,
		EmployeeID:   M.EmployeeID,
		UpdatedBy:    updated_by,
	}, nil
}

func (M CreateEmpEmergencyDetailsReqModel) convertToUpdateDbStruct(admin_id int64) (db.UpdateEmpEmergencyDetailsParams, error) {
	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = admin_id
		updated_by.Valid = true
	}

	return db.UpdateEmpEmergencyDetailsParams{
		FirstName:    M.FirstName,
		LastName:     M.LastName,
		Relationship: M.Relationship,
		Contact:      M.Contact,
		UpdatedBy:    updated_by,
		EmployeeID:   M.EmployeeID,
	}, nil
}

type CreateEmpBankDetailsReqModel struct {
	BankName      string `json:"bank_name"`
	BranchName    string `json:"branch_name"`
	AccountNumber string `json:"account_number"`
	AccountHolder string `json:"account_holder"`
	EmployeeID    int64  `json:"employee_id"`
	UpdatedBy     *int64 `json:"updated_by"`
}

func (M CreateEmpBankDetailsReqModel) convertToDbStruct(admin_id int64) (db.CreateEmpBankDetailsParams, error) {
	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = admin_id
		updated_by.Valid = true
	}

	return db.CreateEmpBankDetailsParams{
		BankName:      M.BankName,
		BranchName:    M.BranchName,
		AccountNumber: M.AccountNumber,
		AccountHolder: M.AccountHolder,
		EmployeeID:    M.EmployeeID,
		UpdatedBy:     updated_by,
	}, nil
}

func (M CreateEmpBankDetailsReqModel) convertToUpdateDbStruct(admin_id int64) (db.UpdateEmpBankDetailsParams, error) {
	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = admin_id
		updated_by.Valid = true
	}

	return db.UpdateEmpBankDetailsParams{
		BankName:      M.BankName,
		BranchName:    M.BranchName,
		AccountNumber: M.AccountNumber,
		AccountHolder: M.AccountHolder,
		UpdatedBy:     updated_by,
		EmployeeID:    M.EmployeeID,
	}, nil
}

type CreateEmpSalaryReqModel struct {
	SalaryType                string          `json:"salary_type"`
	Amount                    string          `json:"amount"`
	SalaryAmountType          string          `json:"salary_amount_type"`
	TotalOfSalaryAllowances   string          `json:"total_of_salary_allowances"`
	TotalSalaryAllowancesType string          `json:"total_salary_allowances_type"`
	PensionEmployer           string          `json:"pension_employer"`
	PensionEmployerType       string          `json:"pension_employer_type"`
	PensionEmployee           string          `json:"pension_employee"`
	PensionEmployeeType       string          `json:"pension_employee_type"`
	TotalNetSalary            string          `json:"total_net_salary"`
	TotalNetSalaryType        string          `json:"total_net_salary_type"`
	EmployeeID                int64           `json:"employee_id"`
	UpdatedBy                 int64           `json:"updated_by"`
	ErID                      int64           `json:"er_id"`
}

func (M CreateEmpSalaryReqModel) convertToDbStruct(admin_id int64) (db.CreateEmpSalaryParams, error) {
	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true

	amount, err := decimal.NewFromString(M.Amount)
	if err != nil {
		return db.CreateEmpSalaryParams{}, err
	}

	total_of_salary_allowances, err := decimal.NewFromString(M.TotalOfSalaryAllowances)
	if err != nil {
		return db.CreateEmpSalaryParams{}, err
	}

	pension_employer, err := decimal.NewFromString(M.PensionEmployer)
	if err != nil {
		return db.CreateEmpSalaryParams{}, err
	}

	pension_employee, err := decimal.NewFromString(M.PensionEmployee)
	if err != nil {
		return db.CreateEmpSalaryParams{}, err
	}

	total_net_salary, err := decimal.NewFromString(M.TotalNetSalary)
	if err != nil {
		return db.CreateEmpSalaryParams{}, err
	}
	var er_id sql.NullInt64
	er_id.Int64 = M.ErID
	er_id.Valid = true

	return db.CreateEmpSalaryParams{
		SalaryType:              M.SalaryType,
		Amount:                  amount,
		SalaryAmountType:        M.SalaryAmountType,
		TotalOfSalaryAllowances: total_of_salary_allowances,
		TotalSalaryAllowancesType: M.TotalSalaryAllowancesType,
		PensionEmployer:         pension_employer,
		PensionEmployerType: M.PensionEmployerType,
		PensionEmployee:         pension_employee,
		PensionEmployeeType: M.PensionEmployeeType,
		TotalNetSalary:          total_net_salary,
		TotalNetSalaryType: M.TotalNetSalaryType,
		EmployeeID:              M.EmployeeID,
		UpdatedBy:               updated_by,
		ErID: 				     er_id,
	}, nil
}

func (M CreateEmpSalaryReqModel) convertToUpdateDbStruct(admin_id int64) (db.UpdateEmpSalaryParams, error) {
	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true

	amount, err := decimal.NewFromString(M.Amount)
	if err != nil {
		return db.UpdateEmpSalaryParams{}, err
	}

	total_of_salary_allowances, err := decimal.NewFromString(M.TotalOfSalaryAllowances)
	if err != nil {
		return db.UpdateEmpSalaryParams{}, err
	}

	pension_employer, err := decimal.NewFromString(M.PensionEmployer)
	if err != nil {
		return db.UpdateEmpSalaryParams{}, err
	}

	pension_employee, err := decimal.NewFromString(M.PensionEmployee)
	if err != nil {
		return db.UpdateEmpSalaryParams{}, err
	}

	total_net_salary, err := decimal.NewFromString(M.TotalNetSalary)
	if err != nil {
		return db.UpdateEmpSalaryParams{}, err
	}

	return db.UpdateEmpSalaryParams{
		SalaryType:              M.SalaryType,
		Amount:                  amount,
		TotalOfSalaryAllowances: total_of_salary_allowances,
		PensionEmployer:         pension_employer,
		PensionEmployee:         pension_employee,
		TotalNetSalary:          total_net_salary,
		UpdatedBy:               updated_by,
		EmployeeID:              M.EmployeeID,
	}, nil
}

type CreateEmpCertificatesReqModel struct {
	Date       string `json:"date" form:"date"`
	Name       string `json:"name" form:"name"`
	UpdatedBy  *int64 `json:"updated_by" form:"updated_by"`
	EmployeeID int64  `json:"employee_id" form:"employee_id"`
}

func (M CreateEmpCertificatesReqModel) convertToDbStruct(admin_id int64) (db.CreateEmpCertificatesParams, error) {
	date, err := time.Parse(time.RFC3339, M.Date)
	if err != nil {
		return db.CreateEmpCertificatesParams{}, err
	}

	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true

	return db.CreateEmpCertificatesParams{
		Date:       date,
		Name:       M.Name,
		UpdatedBy:  updated_by,
		EmployeeID: M.EmployeeID,
	}, nil
}

func (M CreateEmpCertificatesReqModel) convertToUpdateDbStruct(admin_id int64) (db.UpdateEmpCertificatesParams, error) {
	date, err := time.Parse(time.RFC3339, M.Date)
	if err != nil {
		return db.UpdateEmpCertificatesParams{}, err
	}

	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true

	return db.UpdateEmpCertificatesParams{
		Date:       date,
		Name:       M.Name,
		UpdatedBy:  updated_by,
		EmployeeID: M.EmployeeID,
	}, nil
}

type CreateEmpStatusReqModel struct {
	Status      string `json:"status"`
	Department  string `json:"department"`
	Designation string `json:"designation"`
	ValidFrom   string `json:"valid_from"`
	ValidTill   string `json:"valid_till"`
	UpdatedBy   *int64 `json:"updated_by"`
	EmployeeID  int64  `json:"employee_id"`
}

func (M CreateEmpStatusReqModel) convertToDbStruct(admin_id int64) (db.CreateEmpStatusParams, error) {
	validFrom, err := time.Parse(time.RFC3339, M.ValidFrom)
	if err != nil {
		return db.CreateEmpStatusParams{}, err
	}

	validTill, err := time.Parse(time.RFC3339, M.ValidTill)
	if err != nil {
		return db.CreateEmpStatusParams{}, err
	}

	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true

	return db.CreateEmpStatusParams{
		Status:      M.Status,
		Department:  M.Department,
		Designation: M.Designation,
		ValidFrom:   validFrom,
		ValidTill:   validTill,
		UpdatedBy:   updated_by,
		EmployeeID:  M.EmployeeID,
	}, nil
}

func (M CreateEmpStatusReqModel) convertToUpdateDbStruct(admin_id int64) (db.UpdateEmpStatusParams, error) {
	validFrom, err := time.Parse(time.RFC3339, M.ValidFrom)
	if err != nil {
		return db.UpdateEmpStatusParams{}, err
	}

	validTill, err := time.Parse(time.RFC3339, M.ValidTill)
	if err != nil {
		return db.UpdateEmpStatusParams{}, err
	}

	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true

	return db.UpdateEmpStatusParams{
		Status:      M.Status,
		Department:  M.Department,
		Designation: M.Designation,
		ValidFrom:   validFrom,
		ValidTill:   validTill,
		UpdatedBy:   updated_by,
		EmployeeID:  M.EmployeeID,
	}, nil
}

type CreateEmpBenifitsReqModel struct {
	LeaveStatus        bool   `json:"leave_status"`
	LeaveType          string `json:"leave_type"`
	LeaveCount         int32  `json:"leave_count"`
	HealthInsurance    string `json:"health_insurance"`
	InsuranceFrom      string `json:"insurance_from"`
	InsuranceTill      string `json:"insurance_till"`
	RetainmentPlan     string `json:"retainment_plan"`
	RetainmentPlanFrom string `json:"retainment_plan_from"`
	RetainmentPlanTill string `json:"retainment_plan_till"`
	Benifits           string `json:"benifits"`
	BenifitsFrom       string `json:"benifits_from"`
	BenifitsTill       string `json:"benifits_till"`
	UpdatedBy          *int64 `json:"updated_by"`
	EmployeeID         int64  `json:"employee_id"`
}

func (M CreateEmpBenifitsReqModel) convertToDbStruct(admin_id int64) (db.CreateEmpBenifitsParams, error) {
	insuranceFrom, err := time.Parse(time.RFC3339, M.InsuranceFrom)
	if err != nil {
		return db.CreateEmpBenifitsParams{}, err
	}

	insuranceTill, err := time.Parse(time.RFC3339, M.InsuranceTill)
	if err != nil {
		return db.CreateEmpBenifitsParams{}, err
	}

	retainmentPlanFrom, err := time.Parse(time.RFC3339, M.RetainmentPlanFrom)
	if err != nil {
		return db.CreateEmpBenifitsParams{}, err
	}

	retainmentPlanTill, err := time.Parse(time.RFC3339, M.RetainmentPlanTill)
	if err != nil {
		return db.CreateEmpBenifitsParams{}, err
	}

	benifitsFrom, err := time.Parse(time.RFC3339, M.BenifitsFrom)
	if err != nil {
		return db.CreateEmpBenifitsParams{}, err
	}

	benifitsTill, err := time.Parse(time.RFC3339, M.BenifitsTill)
	if err != nil {
		return db.CreateEmpBenifitsParams{}, err
	}

	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true

	return db.CreateEmpBenifitsParams{
		LeaveStatus:        M.LeaveStatus,
		LeaveType:          M.LeaveType,
		LeaveCount:         M.LeaveCount,
		HealthInsurance:    M.HealthInsurance,
		InsuranceFrom:      insuranceFrom,
		InsuranceTill:      insuranceTill,
		RetainmentPlan:     M.RetainmentPlan,
		RetainmentPlanFrom: retainmentPlanFrom,
		RetainmentPlanTill: retainmentPlanTill,
		Benifits:           M.Benifits,
		BenifitsFrom:       benifitsFrom,
		BenifitsTill:       benifitsTill,
		UpdatedBy:          updated_by,
		EmployeeID:         M.EmployeeID,
	}, nil
}

func (M CreateEmpBenifitsReqModel) convertToUpdateDbStruct(admin_id int64) (db.UpdateEmpBenifitsParams, error) {
	insuranceFrom, err := time.Parse(time.RFC3339, M.InsuranceFrom)
	if err != nil {
		return db.UpdateEmpBenifitsParams{}, err
	}

	insuranceTill, err := time.Parse(time.RFC3339, M.InsuranceTill)
	if err != nil {
		return db.UpdateEmpBenifitsParams{}, err
	}

	retainmentPlanFrom, err := time.Parse(time.RFC3339, M.RetainmentPlanFrom)
	if err != nil {
		return db.UpdateEmpBenifitsParams{}, err
	}

	retainmentPlanTill, err := time.Parse(time.RFC3339, M.RetainmentPlanTill)
	if err != nil {
		return db.UpdateEmpBenifitsParams{}, err
	}

	benifitsFrom, err := time.Parse(time.RFC3339, M.BenifitsFrom)
	if err != nil {
		return db.UpdateEmpBenifitsParams{}, err
	}

	benifitsTill, err := time.Parse(time.RFC3339, M.BenifitsTill)
	if err != nil {
		return db.UpdateEmpBenifitsParams{}, err
	}

	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true

	return db.UpdateEmpBenifitsParams{
		LeaveStatus:        M.LeaveStatus,
		LeaveType:          M.LeaveType,
		LeaveCount:         M.LeaveCount,
		HealthInsurance:    M.HealthInsurance,
		InsuranceFrom:      insuranceFrom,
		InsuranceTill:      insuranceTill,
		RetainmentPlan:     M.RetainmentPlan,
		RetainmentPlanFrom: retainmentPlanFrom,
		RetainmentPlanTill: retainmentPlanTill,
		Benifits:           M.Benifits,
		BenifitsFrom:       benifitsFrom,
		BenifitsTill:       benifitsTill,
		UpdatedBy:          updated_by,
		EmployeeID:         M.EmployeeID,
	}, nil
}

type CreateEmpUserReqModel struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	UpdatedBy  *int64 `json:"updated_by"`
	EmployeeID int64  `json:"employee_id"`
}

func (M CreateEmpUserReqModel) convertToDbStruct(id int64, admin_id int64) (db.CreateEmpUserParams, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(M.Password), bcrypt.DefaultCost)
	if err != nil {
		return db.CreateEmpUserParams{}, err
	}

	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true

	return db.CreateEmpUserParams{
		Email:      M.Email,
		Password:   string(hashedPassword),
		UpdatedBy:  updated_by,
		EmployeeID: M.EmployeeID,
		BranchID:   id,
	}, nil
}

func (M CreateEmpUserReqModel) convertToUpdateDbStruct(admin_id int64) (db.UpdateEmpUserParams, error) {
	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(M.Password), bcrypt.DefaultCost)
	if err != nil {
		return db.UpdateEmpUserParams{}, err
	}

	return db.UpdateEmpUserParams{
		Email:      M.Email,
		Password:   string(hashedPassword),
		UpdatedBy:  updated_by,
		EmployeeID: M.EmployeeID,
	}, nil
}

type CreateEmpAllowancesReqModel struct {
	Name       string        `json:"name"`
	Amount     string        `json:"amount"`
	UpdatedBy  sql.NullInt64 `json:"updated_by"`
	EmployeeID int64         `json:"employee_id"`
}

func (M CreateEmpAllowancesReqModel) convertToDbStruct(admin_id int64) (db.CreateEmpAllowancesParams, error) {
	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true

	amount, err := decimal.NewFromString(M.Amount)
	if err != nil {
		return db.CreateEmpAllowancesParams{}, err
	}

	return db.CreateEmpAllowancesParams{
		Name:       M.Name,
		Amount:     amount,
		UpdatedBy:  updated_by,
		EmployeeID: M.EmployeeID,
	}, nil
}

func (M CreateEmpAllowancesReqModel) convertToUpdateDbStruct(admin_id int64) (db.UpdateEmpAllowancesParams, error) {
	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true

	amount, err := decimal.NewFromString(M.Amount)
	if err != nil {
		return db.UpdateEmpAllowancesParams{}, err
	}

	return db.UpdateEmpAllowancesParams{
		Name:       M.Name,
		Amount:     amount,
		UpdatedBy:  updated_by,
		EmployeeID: M.EmployeeID,
	}, nil
}

type CreateEmpExpatriateReqModel struct {
	Expatriate    bool   `json:"expatriate"`
	Nationality   string `json:"nationality"`
	VisaType      string `json:"visa_type"`
	VisaFrom      string `json:"visa_from"`
	VisaTill      string `json:"visa_till"`
	VisaNumber    string `json:"visa_number"`
	VisaFee       string `json:"visa_fee"`
	UpdatedBy     *int64 `json:"updated_by"`
	EmployeeID    int64  `json:"employee_id"`
}

func (M CreateEmpExpatriateReqModel) convertToDbStruct(admin_id int64) (db.CreateEmpExpatriateParams, error) {
	visaFrom, err := time.Parse(time.RFC3339, M.VisaFrom)
	if err != nil {
		return db.CreateEmpExpatriateParams{}, err
	}

	visaTill, err := time.Parse(time.RFC3339, M.VisaTill)
	if err != nil {
		return db.CreateEmpExpatriateParams{}, err
	}

	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true

	visa_amount, err := decimal.NewFromString(M.VisaFee)
	if err != nil {
		return db.CreateEmpExpatriateParams{}, err
	}

	return db.CreateEmpExpatriateParams{
		Expatriate:    M.Expatriate,
		Nationality:   M.Nationality,
		VisaType:      M.VisaType,
		VisaFrom:      visaFrom,
		VisaTill:      visaTill,
		VisaNumber:    M.VisaNumber,
		VisaFee:       visa_amount,
		UpdatedBy:     updated_by,
		EmployeeID:    M.EmployeeID,
	}, nil
}

// func (M CreateEmpExpatriateReqModel) convertToUpdateDbStruct(admin_id int64) (db.UpdateEmpExpatriateParams, error) {
// 	visaFrom, err := time.Parse(time.RFC3339, M.VisaFrom)
// 	if err != nil {
// 		return db.UpdateEmpExpatriateParams{}, err
// 	}

// 	visaTill, err := time.Parse(time.RFC3339, M.VisaTill)
// 	if err != nil {
// 		return db.UpdateEmpExpatriateParams{}, err
// 	}

// 	var updated_by sql.NullInt64
// 	updated_by.Int64 = admin_id
// 	updated_by.Valid = true

// 	visa_amount, err := decimal.NewFromString(M.VisaFee)
// 	if err != nil {
// 		return db.UpdateEmpExpatriateParams{}, err
// 	}

// 	return db.UpdateEmpExpatriateParams{
// 		Expatriate:    M.Expatriate,
// 		Nationality:   M.Nationality,
// 		VisaType:      M.VisaType,
// 		VisaFrom:      visaFrom,
// 		VisaTill:      visaTill,
// 		VisaNumber:    M.VisaNumber,
// 		VisaFee:       visa_amount,
// 		VisaImagePath: M.VisaImagePath,
// 		UpdatedBy:     updated_by,
// 		EmployeeID:    M.EmployeeID,
// 	}, nil
// }

type CreateEmpAccessiabilityReqModel struct {
	Accessibility     bool   `json:"accessibility"`
	AccessibilityFrom string `json:"accessibility_from"`
	AccessibilityTill string `json:"accessibility_till"`
	Enable            bool   `json:"enable"`
	UpdatedBy         *int64 `json:"updated_by"`
	EmployeeID        int64  `json:"employee_id"`
}

func (M CreateEmpAccessiabilityReqModel) convertToDbStruct(admin_id int64) (db.CreateEmpAccessiabilityParams, error) {
	accessibilityFrom, err := time.Parse(time.RFC3339, M.AccessibilityFrom)
	if err != nil {
		return db.CreateEmpAccessiabilityParams{}, err
	}

	accessibilityTill, err := time.Parse(time.RFC3339, M.AccessibilityTill)
	if err != nil {
		return db.CreateEmpAccessiabilityParams{}, err
	}

	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true

	return db.CreateEmpAccessiabilityParams{
		Accessibility:     M.Accessibility,
		AccessibilityFrom: accessibilityFrom,
		AccessibilityTill: accessibilityTill,
		Enable:            M.Enable,
		UpdatedBy:         updated_by,
		EmployeeID:        M.EmployeeID,
	}, nil
}

func (M CreateEmpAccessiabilityReqModel) convertToUpdateDbStruct(admin_id int64) (db.UpdateEmpAccessiabilityParams, error) {
	accessibilityFrom, err := time.Parse(time.RFC3339, M.AccessibilityFrom)
	if err != nil {
		return db.UpdateEmpAccessiabilityParams{}, err
	}

	accessibilityTill, err := time.Parse(time.RFC3339, M.AccessibilityTill)
	if err != nil {
		return db.UpdateEmpAccessiabilityParams{}, err
	}

	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true

	return db.UpdateEmpAccessiabilityParams{
		Accessibility:     M.Accessibility,
		AccessibilityFrom: accessibilityFrom,
		AccessibilityTill: accessibilityTill,
		Enable:            M.Enable,
		UpdatedBy:         updated_by,
		EmployeeID:        M.EmployeeID,
	}, nil
}

type IsTrainerReqModel struct {
	IsTrainer bool `json:"is_trainer"`
	AttendeeId int64 `json:"attendee_id"`
	TrainerID int64 `json:"trainer_id"`
	EmployeeID int64 `json:"employee_id"`
	Commission string `json:"commission"`
}

type CreateFileSubmitReqModel struct {
	EmployeeID int64  `json:"employee_id"`
	FileName   string `json:"file_name"`
	FileType   string `json:"file_type"`
}

type EmpReqModel struct {
	Employee       CreateEmployeeReqModel            `json:"employee"`
	Emergency      CreateEmpEmergencyDetailsReqModel `json:"emergency"`
	Bank           CreateEmpBankDetailsReqModel      `json:"bank"`
	Salary         CreateEmpSalaryReqModel           `json:"salary"`
	Certificates   CreateEmpCertificatesReqModel     `json:"certificates"`
	Status         CreateEmpStatusReqModel           `json:"status"`
	Benifits       CreateEmpBenifitsReqModel         `json:"benifits"`
	User           CreateEmpUserReqModel             `json:"user"`
	Allowances     []CreateEmpAllowancesReqModel     `json:"allowances"`
	Expatriate     CreateEmpExpatriateReqModel       `json:"expatriate"`
	Accessiability CreateEmpAccessiabilityReqModel   `json:"accessiability"`
	FileSubmit     []CreateFileSubmitReqModel		 `json:"file_submit"`
	IsTrainer      IsTrainerReqModel			 	 `json:"is_trainer"`
}

type EmpResponse struct {
	Employee   db.GetEmployeeByIDRow `json:"employee"`
	EmpAllowances []db.GetEmployeeAllowancesRow `json:"allowances"`
	EmpFiles []db.GetEmpFilesRow `json:"files"`
	TrainerCom  TrainerCom `json:"trainer_data"`
}

type TrainerCom struct {
	IsTrainer bool            `json:"is_trainer"`
 	Commission string `json:"commission"`
}

type EmpLoginReqModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetEmployeeReqModel struct {
	Search     string `json:"search"`
	Limit      int32  `json:"limit"`
	PageNumber int32  `json:"page"`
}

func (M GetEmployeeReqModel) convertToDbStruct() (db.GetEmployeeParams, error) {
	offset := (M.PageNumber - 1) * M.Limit

	return db.GetEmployeeParams{
		Limit:  M.Limit,
		Offset: offset,
	}, nil
}

type LoginEmpResponse struct {
	Token string      `json:"token"`
	Data  rba.RBAauth `json:"data"`
}

type CheckTrainerParams struct {
	Email   string `json:"email"`
}

type UpdateCommissionReqModel struct {
	Commission string `json:"commission"`
	EmployeeID int64  `json:"employee_id"`
}

func (M UpdateCommissionReqModel) convertToDbStruct(admin_id int64) (db.UpdateTrainerCommissionParams, error) {
	commission, err := decimal.NewFromString(M.Commission)
	if err != nil {
		return db.UpdateTrainerCommissionParams{}, err
	}

	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true

	return db.UpdateTrainerCommissionParams{
		Commission: commission,
		UpdatedBy:  updated_by,
		EmployeeID: M.EmployeeID,
	}, nil
}

type UpdateEmpCertificatesReqModel struct {
	Date       string `json:"date" form:"date"`
	Name       string `json:"name" form:"name"`
	UpdatedBy  *int64 `json:"updated_by" form:"updated_by"`
	EmployeeID int64  `json:"employee_id" form:"employee_id"`
	FileName   string `json:"file_name"`
	FileType   string `json:"file_type"`
}

func (M UpdateEmpCertificatesReqModel) convertToCertDbStruct(admin_id int64) (db.UpdateEmpCertificatesParams, error) {
	date, err := time.Parse(time.RFC3339, M.Date)
	if err != nil {
		return db.UpdateEmpCertificatesParams{}, err
	}

	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true

	return db.UpdateEmpCertificatesParams{
		Date:       date,
		Name:       M.Name,
		UpdatedBy:  updated_by,
		EmployeeID: M.EmployeeID,
	}, nil
}

func (M UpdateEmpCertificatesReqModel) convertToFileDbStruct() (db.CreateFileSubmitParams, error) {

	return db.CreateFileSubmitParams{
		EmployeeID: M.EmployeeID,
		FileName:   M.FileName,
		FileType:   M.FileType,
	}, nil
}

type UpdateEmpExpatriateAndFilesReqModel struct {
	Expatriate    bool   `json:"expatriate"`
	Nationality   string `json:"nationality"`
	VisaType      string `json:"visa_type"`
	VisaFrom      string `json:"visa_from"`
	VisaTill      string `json:"visa_till"`
	VisaNumber    string `json:"visa_number"`
	VisaFee       string `json:"visa_fee"`
	UpdatedBy     *int64 `json:"updated_by"`
	EmployeeID    int64  `json:"employee_id"`
	FileName   string           `json:"file_name"`
	FileType   string           `json:"file_type"`
}


func (M UpdateEmpExpatriateAndFilesReqModel) convertToExpDbStruct(admin_id int64) (db.UpdateEmpExpatriateParams, error) {
	visaFrom, err := time.Parse(time.RFC3339, M.VisaFrom)
	if err != nil {
		return db.UpdateEmpExpatriateParams{}, err
	}

	visaTill, err := time.Parse(time.RFC3339, M.VisaTill)
	if err != nil {
		return db.UpdateEmpExpatriateParams{}, err
	}

	var updated_by sql.NullInt64
	updated_by.Int64 = admin_id
	updated_by.Valid = true

	visa_amount, err := decimal.NewFromString(M.VisaFee)
	if err != nil {
		return db.UpdateEmpExpatriateParams{}, err
	}
	return db.UpdateEmpExpatriateParams{
		Expatriate:    M.Expatriate,
		Nationality:  M.Nationality,
		VisaType:      M.VisaType,
		VisaFrom:      visaFrom,
		VisaTill:      visaTill,
		VisaNumber:    M.VisaNumber,
		VisaFee:       visa_amount,
		UpdatedBy:     updated_by,
		EmployeeID:    M.EmployeeID,
	}, nil
}

func (M UpdateEmpExpatriateAndFilesReqModel) convertToExpFileDbStruct() (db.CreateFileSubmitParams, error) {

	return db.CreateFileSubmitParams{
		EmployeeID: M.EmployeeID,
		FileName:   M.FileName,
		FileType:   M.FileType,
	}, nil
}
	