package hr

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
	"github.com/shopspring/decimal"
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

func (M CreateEmployeeReqModel) convertToDbStruct() (db.CreateEmployeeParams, error) {
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
	if M.UpdatedBy != nil {
		updated_by.Int64 = *M.UpdatedBy
		updated_by.Valid = true
	}

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

func (M CreateEmployeeReqModel) ConvertToUpdateDbStruct(id int64) (db.UpdateEmployeeParams, error) {
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
	if M.UpdatedBy != nil {
		updated_by.Int64 = *M.UpdatedBy
		updated_by.Valid = true
	}

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
		ID:                id,
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

func (M CreateEmpEmergencyDetailsReqModel) convertToDbStruct() (db.CreateEmpEmergencyDetailsParams, error) {

	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = *M.UpdatedBy
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

func (M CreateEmpEmergencyDetailsReqModel) convertToUpdateDbStruct() (db.UpdateEmpEmergencyDetailsParams, error) {

	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = *M.UpdatedBy
		updated_by.Valid = true
	}

	return db.UpdateEmpEmergencyDetailsParams{
		FirstName:    M.FirstName,
		LastName:     M.LastName,
		Relationship: M.Relationship,
		Contact:      M.Contact,
		UpdatedBy:   updated_by,
		EmployeeID:  M.EmployeeID,
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

func (M CreateEmpBankDetailsReqModel) convertToDbStruct() (db.CreateEmpBankDetailsParams, error) {

	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = *M.UpdatedBy
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

func (M CreateEmpBankDetailsReqModel) convertToUpdateDbStruct() (db.UpdateEmpBankDetailsParams, error) {

	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = *M.UpdatedBy
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
	SalaryType              string `json:"salary_type"`
	Amount                  string `json:"amount"`
	TotalOfSalaryAllowances int32  `json:"total_of_salary_allowances"`
	PensionEmployer         int32  `json:"pension_employer"`
	PensionEmployee         int32  `json:"pension_employee"`
	TotalNetSalary          int32  `json:"total_net_salary"`
	EmployeeID              int64  `json:"employee_id"`
	UpdatedBy               *int64 `json:"updated_by"`
}

func (M CreateEmpSalaryReqModel) convertToDbStruct() (db.CreateEmpSalaryParams, error) {

	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = *M.UpdatedBy
		updated_by.Valid = true
	}
	amount, err := decimal.NewFromString(M.Amount)
	if err != nil {
		return db.CreateEmpSalaryParams{}, fmt.Errorf("invalid amount format: %v", err)
	}

	return db.CreateEmpSalaryParams{
		SalaryType:              M.SalaryType,
		Amount:                  amount,
		TotalOfSalaryAllowances: M.TotalOfSalaryAllowances,
		PensionEmployer:         M.PensionEmployer,
		PensionEmployee:         M.PensionEmployee,
		TotalNetSalary:          M.TotalNetSalary,
		EmployeeID:              M.EmployeeID,
		UpdatedBy:               updated_by,
	}, nil
}

func (M CreateEmpSalaryReqModel) convertToUpdateDbStruct() (db.UpdateEmpSalaryParams, error) {

	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = *M.UpdatedBy
		updated_by.Valid = true
	}
	amount, err := decimal.NewFromString(M.Amount)
	if err != nil {
		return db.UpdateEmpSalaryParams{}, fmt.Errorf("invalid amount format: %v", err)
	}

	return db.UpdateEmpSalaryParams{
		SalaryType:              M.SalaryType,
		Amount:                  amount,
		TotalOfSalaryAllowances: M.TotalOfSalaryAllowances,
		PensionEmployer:         M.PensionEmployer,
		PensionEmployee:         M.PensionEmployee,
		TotalNetSalary:          M.TotalNetSalary,
		UpdatedBy: 			     updated_by,
		EmployeeID:              M.EmployeeID,
	}, nil
}

type CreateEmpCertificatesReqModel struct {
	Date       string `json:"date"`
	Name       string `json:"name"`
	ImagePath  string `json:"image_path"`
	UpdatedBy  *int64 `json:"updated_by"`
	EmployeeID int64  `json:"employee_id"`
}

func (M CreateEmpCertificatesReqModel) convertToDbStruct() (db.CreateEmpCertificatesParams, error) {
	date, err := time.Parse(time.RFC3339, M.Date)
	if err != nil {
		return db.CreateEmpCertificatesParams{}, err
	}

	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = *M.UpdatedBy
		updated_by.Valid = true
	}

	return db.CreateEmpCertificatesParams{
		Date:       date,
		Name:       M.Name,
		ImagePath:  M.ImagePath,
		UpdatedBy:  updated_by,
		EmployeeID: M.EmployeeID,
	}, nil
}

func (M CreateEmpCertificatesReqModel) convertToUpdateDbStruct() (db.UpdateEmpCertificatesParams, error) {
	date, err := time.Parse(time.RFC3339, M.Date)
	if err != nil {
		return db.UpdateEmpCertificatesParams{}, err
	}

	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = *M.UpdatedBy
		updated_by.Valid = true
	}

	return db.UpdateEmpCertificatesParams{
		Date:       date,
		Name:       M.Name,
		ImagePath:  M.ImagePath,
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

func (M CreateEmpStatusReqModel) convertToDbStruct() (db.CreateEmpStatusParams, error) {
	validFrom, err := time.Parse(time.RFC3339, M.ValidFrom)
	if err != nil {
		return db.CreateEmpStatusParams{}, err
	}

	validTill, err := time.Parse(time.RFC3339, M.ValidTill)
	if err != nil {
		return db.CreateEmpStatusParams{}, err
	}

	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = *M.UpdatedBy
		updated_by.Valid = true
	}

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

func (M CreateEmpStatusReqModel) convertToUpdateDbStruct() (db.UpdateEmpStatusParams, error) {
	validFrom, err := time.Parse(time.RFC3339, M.ValidFrom)
	if err != nil {
		return db.UpdateEmpStatusParams{}, err
	}

	validTill, err := time.Parse(time.RFC3339, M.ValidTill)
	if err != nil {
		return db.UpdateEmpStatusParams{}, err
	}

	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = *M.UpdatedBy
		updated_by.Valid = true
	}

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

func (M CreateEmpBenifitsReqModel) convertToDbStruct() (db.CreateEmpBenifitsParams, error) {
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
	if M.UpdatedBy != nil {
		updated_by.Int64 = *M.UpdatedBy
		updated_by.Valid = true
	}

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

func (M CreateEmpBenifitsReqModel) convertToUpdateDbStruct() (db.UpdateEmpBenifitsParams, error) {
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
	if M.UpdatedBy != nil {
		updated_by.Int64 = *M.UpdatedBy
		updated_by.Valid = true
	}

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

func (M CreateEmpUserReqModel) convertToDbStruct() (db.CreateEmpUserParams, error) {

	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = *M.UpdatedBy
		updated_by.Valid = true
	}

	return db.CreateEmpUserParams{
		Email:      M.Email,
		Password:   M.Password,
		UpdatedBy:  updated_by,
		EmployeeID: M.EmployeeID,
	}, nil
}

func (M CreateEmpUserReqModel) convertToUpdateDbStruct() (db.UpdateEmpUserParams, error) {

	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = *M.UpdatedBy
		updated_by.Valid = true
	}

	return db.UpdateEmpUserParams{
		Email:      M.Email,
		Password:   M.Password,
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

func (M CreateEmpAllowancesReqModel) convertToDbStruct() (db.CreateEmpAllowancesParams, error) {

	var updated_by sql.NullInt64
	if M.UpdatedBy.Valid {
		updated_by.Int64 = M.UpdatedBy.Int64
		updated_by.Valid = true
	}
	amount, err := decimal.NewFromString(M.Amount)
	if err != nil {
		return db.CreateEmpAllowancesParams{}, fmt.Errorf("invalid amount format: %v", err)
	}

	return db.CreateEmpAllowancesParams{
		Name:       M.Name,
		Amount:     amount,
		UpdatedBy:  updated_by,
		EmployeeID: M.EmployeeID,
	}, nil
}

func (M CreateEmpAllowancesReqModel) convertToUpdateDbStruct() (db.UpdateEmpAllowancesParams, error) {

	var updated_by sql.NullInt64
	if M.UpdatedBy.Valid {
		updated_by.Int64 = M.UpdatedBy.Int64
		updated_by.Valid = true
	}
	amount, err := decimal.NewFromString(M.Amount)
	if err != nil {
		return db.UpdateEmpAllowancesParams{}, fmt.Errorf("invalid amount format: %v", err)
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
	VisaImagePath string `json:"visa_image_path"`
	UpdatedBy     *int64 `json:"updated_by"`
	EmployeeID    int64  `json:"employee_id"`
}

func (M CreateEmpExpatriateReqModel) convertToDbStruct() (db.CreateEmpExpatriateParams, error) {
	visaFrom, err := time.Parse(time.RFC3339, M.VisaFrom)
	if err != nil {
		return db.CreateEmpExpatriateParams{}, err
	}

	visaTill, err := time.Parse(time.RFC3339, M.VisaTill)
	if err != nil {
		return db.CreateEmpExpatriateParams{}, err
	}

	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = *M.UpdatedBy
		updated_by.Valid = true
	}

	visa_amount, err := decimal.NewFromString(M.VisaFee)
	if err != nil {
		return db.CreateEmpExpatriateParams{}, fmt.Errorf("invalid amount format: %v", err)
	}

	return db.CreateEmpExpatriateParams{
		Expatriate:    M.Expatriate,
		Nationality:   M.Nationality,
		VisaType:      M.VisaType,
		VisaFrom:      visaFrom,
		VisaTill:      visaTill,
		VisaNumber:    M.VisaNumber,
		VisaFee:       visa_amount,
		VisaImagePath: M.VisaImagePath,
		UpdatedBy:     updated_by,
		EmployeeID:    M.EmployeeID,
	}, nil
}

func (M CreateEmpExpatriateReqModel) convertToUpdateDbStruct() (db.UpdateEmpExpatriateParams, error) {
	visaFrom, err := time.Parse(time.RFC3339, M.VisaFrom)
	if err != nil {
		return db.UpdateEmpExpatriateParams{}, err
	}

	visaTill, err := time.Parse(time.RFC3339, M.VisaTill)
	if err != nil {
		return db.UpdateEmpExpatriateParams{}, err
	}

	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = *M.UpdatedBy
		updated_by.Valid = true
	}

	visa_amount, err := decimal.NewFromString(M.VisaFee)
	if err != nil {
		return db.UpdateEmpExpatriateParams{}, fmt.Errorf("invalid amount format: %v", err)
	}

	return db.UpdateEmpExpatriateParams{
		Expatriate:    M.Expatriate,
		Nationality:   M.Nationality,
		VisaType:      M.VisaType,
		VisaFrom:      visaFrom,
		VisaTill:      visaTill,
		VisaNumber:    M.VisaNumber,
		VisaFee:       visa_amount,
		VisaImagePath: M.VisaImagePath,
		UpdatedBy:     updated_by,
		EmployeeID:    M.EmployeeID,
	}, nil
}

type CreateEmpAccessiabilityReqModel struct {
	Accessibility     bool   `json:"accessibility"`
	AccessibilityFrom string `json:"accessibility_from"`
	AccessibilityTill string `json:"accessibility_till"`
	Enable            bool   `json:"enable"`
	UpdatedBy         *int64 `json:"updated_by"`
	EmployeeID        int64  `json:"employee_id"`
}

func (M CreateEmpAccessiabilityReqModel) convertToDbStruct() (db.CreateEmpAccessiabilityParams, error) {
	accessibilityFrom, err := time.Parse(time.RFC3339, M.AccessibilityFrom)
	if err != nil {
		return db.CreateEmpAccessiabilityParams{}, err
	}

	accessibilityTill, err := time.Parse(time.RFC3339, M.AccessibilityTill)
	if err != nil {
		return db.CreateEmpAccessiabilityParams{}, err
	}

	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = *M.UpdatedBy
		updated_by.Valid = true
	}

	return db.CreateEmpAccessiabilityParams{
		Accessibility:     M.Accessibility,
		AccessibilityFrom: accessibilityFrom,
		AccessibilityTill: accessibilityTill,
		Enable:            M.Enable,
		UpdatedBy:         updated_by,
		EmployeeID:        M.EmployeeID,
	}, nil
}

func (M CreateEmpAccessiabilityReqModel) convertToUpdateDbStruct() (db.UpdateEmpAccessiabilityParams, error) {
	accessibilityFrom, err := time.Parse(time.RFC3339, M.AccessibilityFrom)
	if err != nil {
		return db.UpdateEmpAccessiabilityParams{}, err
	}

	accessibilityTill, err := time.Parse(time.RFC3339, M.AccessibilityTill)
	if err != nil {
		return db.UpdateEmpAccessiabilityParams{}, err
	}

	var updated_by sql.NullInt64
	if M.UpdatedBy != nil {
		updated_by.Int64 = *M.UpdatedBy
		updated_by.Valid = true
	}

	return db.UpdateEmpAccessiabilityParams{
		Accessibility:     M.Accessibility,
		AccessibilityFrom: accessibilityFrom,
		AccessibilityTill: accessibilityTill,
		Enable:            M.Enable,
		UpdatedBy:         updated_by,
		EmployeeID:        M.EmployeeID,
	}, nil
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
	Allowances     CreateEmpAllowancesReqModel       `json:"allowances"`
	Expatriate     CreateEmpExpatriateReqModel       `json:"expatriate"`
	Accessiability CreateEmpAccessiabilityReqModel   `json:"accessiability"`
}

// type GetEmployee struct {
// 	Limit  int32 `json:"limit"`
// 	Offset int32 `json:"offset"`
// }

// // update employee
// type UpdateEmployeeReqModel struct {
// 	FirstName         string `json:"first_name"`
// 	LastName          string `json:"last_name"`
// 	Gender            string `json:"gender"`
// 	Dob               string `json:"dob"`
// 	Religion          string `json:"religion"`
// 	PrimaryNumber     string `json:"primary_number"`
// 	SecondaryNumber   string `json:"secondary_number"`
// 	PassportID        string `json:"passport_id"`
// 	Nationality       string `json:"nationality"`
// 	PassportValidTill string `json:"passport_valid_till"`
// 	Nic               string `json:"nic"`
// 	Country           string `json:"country"`
// 	NicValidTill      string `json:"nic_valid_till"`
// 	Address           string `json:"address"`
// 	CurrentCountry    string `json:"current_country"`
// 	Email             string `json:"email"`
// 	UpdatedBy         *int64 `json:"updated_by"`
// 	ID                int64  `json:"id"`
// }

// func (M UpdateEmployeeReqModel) convertToDbStruct() (db.UpdateEmployeeParams, error) {
// 	dob, err := time.Parse(time.RFC3339, M.Dob)
// 	if err != nil {
// 		log.Printf("Error parsing dob: %v", err)
// 		return db.UpdateEmployeeParams{}, err
// 	}

// 	passportValidTill, err := time.Parse(time.RFC3339, M.PassportValidTill)
// 	if err != nil {
// 		log.Printf("Error parsing passport valid till: %v", err)
// 		return db.UpdateEmployeeParams{}, err
// 	}

// 	nicValidTill, err := time.Parse(time.RFC3339, M.NicValidTill)
// 	if err != nil {
// 		log.Printf("Error parsing nic valid till: %v", err)
// 		return db.UpdateEmployeeParams{}, err
// 	}

// 	var updated_by sql.NullInt64
// 	if M.UpdatedBy != nil {
// 		updated_by.Int64 = *M.UpdatedBy
// 		updated_by.Valid = true
// 	}

// 	return db.UpdateEmployeeParams{
// 		FirstName:         M.FirstName,
// 		LastName:          M.LastName,
// 		Gender:            M.Gender,
// 		Dob:               dob,
// 		Religion:          M.Religion,
// 		PrimaryNumber:     M.PrimaryNumber,
// 		SecondaryNumber:   M.SecondaryNumber,
// 		PassportID:        M.PassportID,
// 		Nationality:       M.Nationality,
// 		PassportValidTill: passportValidTill,
// 		Nic:               M.Nic,
// 		Country:           M.Country,
// 		NicValidTill:      nicValidTill,
// 		Address:           M.Address,
// 		CurrentCountry:    M.CurrentCountry,
// 		Email:             M.Email,
// 		UpdatedBy:         updated_by,
// 		ID:                M.ID,
// 	}, nil
// }

// type UpdateEmpEmergencyDetailsReqModel struct {
// 	FirstName    string `json:"first_name"`
// 	LastName     string `json:"last_name"`
// 	Relationship string `json:"relationship"`
// 	Contact      string `json:"contact"`
// 	UpdatedBy    *int64 `json:"updated_by"`
// 	EmployeeID   int64  `json:"employee_id"`
// }

// func (M UpdateEmpEmergencyDetailsReqModel) convertToDbStruct() (db.UpdateEmpEmergencyDetailsParams, error) {

// 	var updated_by sql.NullInt64
// 	if M.UpdatedBy != nil {
// 		updated_by.Int64 = *M.UpdatedBy
// 		updated_by.Valid = true
// 	}

// 	return db.UpdateEmpEmergencyDetailsParams{
// 		FirstName:    M.FirstName,
// 		LastName:     M.LastName,
// 		Relationship: M.Relationship,
// 		Contact:      M.Contact,
// 		UpdatedBy:    updated_by,
// 		EmployeeID:   M.EmployeeID,
// 	}, nil
// }

// type UpdateEmpBankDetailsReqModel struct {
// 	BankName      string `json:"bank_name"`
// 	BranchName    string `json:"branch_name"`
// 	AccountNumber string `json:"account_number"`
// 	AccountHolder string `json:"account_holder"`
// 	UpdatedBy     *int64 `json:"updated_by"`
// 	EmployeeID    int64  `json:"employee_id"`
// }

// func (M UpdateEmpBankDetailsReqModel) convertToDbStruct() (db.UpdateEmpBankDetailsParams, error) {

// 	var updated_by sql.NullInt64
// 	if M.UpdatedBy != nil {
// 		updated_by.Int64 = *M.UpdatedBy
// 		updated_by.Valid = true
// 	}

// 	return db.UpdateEmpBankDetailsParams{
// 		BankName:      M.BankName,
// 		BranchName:    M.BranchName,
// 		AccountNumber: M.AccountNumber,
// 		AccountHolder: M.AccountHolder,
// 		UpdatedBy:     updated_by,
// 		EmployeeID:    M.EmployeeID,
// 	}, nil
// }

// type UpdateEmpSalaryReqModel struct {
// 	SalaryType              string `json:"salary_type"`
// 	Amount                  string `json:"amount"`
// 	TotalOfSalaryAllowances int32  `json:"total_of_salary_allowances"`
// 	PensionEmployer         int32  `json:"pension_employer"`
// 	PensionEmployee         int32  `json:"pension_employee"`
// 	TotalNetSalary          int32  `json:"total_net_salary"`
// 	UpdatedBy               *int64 `json:"updated_by"`
// 	EmployeeID              int64  `json:"employee_id"`
// }

// func (M UpdateEmpSalaryReqModel) convertToDbStruct() (db.UpdateEmpSalaryParams, error) {

// 	var updated_by sql.NullInt64
// 	if M.UpdatedBy != nil {
// 		updated_by.Int64 = *M.UpdatedBy
// 		updated_by.Valid = true
// 	}
// 	amount, err := decimal.NewFromString(M.Amount)
// 	if err != nil {
// 		return db.UpdateEmpSalaryParams{}, fmt.Errorf("invalid amount format: %v", err)
// 	}

// 	return db.UpdateEmpSalaryParams{
// 		SalaryType:              M.SalaryType,
// 		Amount:                  amount,
// 		TotalOfSalaryAllowances: M.TotalOfSalaryAllowances,
// 		PensionEmployer:         M.PensionEmployer,
// 		PensionEmployee:         M.PensionEmployee,
// 		TotalNetSalary:          M.TotalNetSalary,
// 		UpdatedBy:               updated_by,
// 		EmployeeID:              M.EmployeeID,
// 	}, nil
// }

// type UpdateEmpCertificatesReqModel struct {
// 	Date       string `json:"date"`
// 	Name       string `json:"name"`
// 	ImagePath  string `json:"image_path"`
// 	UpdatedBy  *int64 `json:"updated_by"`
// 	EmployeeID int64  `json:"employee_id"`
// }

// func (M UpdateEmpCertificatesReqModel) convertToDbStruct() (db.UpdateEmpCertificatesParams, error) {
// 	date, err := time.Parse(time.RFC3339, M.Date)
// 	if err != nil {
// 		return db.UpdateEmpCertificatesParams{}, err
// 	}

// 	var updated_by sql.NullInt64
// 	if M.UpdatedBy != nil {
// 		updated_by.Int64 = *M.UpdatedBy
// 		updated_by.Valid = true
// 	}

// 	return db.UpdateEmpCertificatesParams{
// 		Date:       date,
// 		Name:       M.Name,
// 		ImagePath:  M.ImagePath,
// 		UpdatedBy:  updated_by,
// 		EmployeeID: M.EmployeeID,
// 	}, nil
// }

// type UpdateEmpStatusReqModel struct {
// 	Status      string `json:"status"`
// 	Department  string `json:"department"`
// 	Designation string `json:"designation"`
// 	ValidFrom   string `json:"valid_from"`
// 	ValidTill   string `json:"valid_till"`
// 	UpdatedBy   *int64 `json:"updated_by"`
// 	EmployeeID  int64  `json:"employee_id"`
// }

// func (M UpdateEmpStatusReqModel) convertToDbStruct() (db.UpdateEmpStatusParams, error) {
// 	validFrom, err := time.Parse(time.RFC3339, M.ValidFrom)
// 	if err != nil {
// 		return db.UpdateEmpStatusParams{}, err
// 	}

// 	validTill, err := time.Parse(time.RFC3339, M.ValidTill)
// 	if err != nil {
// 		return db.UpdateEmpStatusParams{}, err
// 	}

// 	var updated_by sql.NullInt64
// 	if M.UpdatedBy != nil {
// 		updated_by.Int64 = *M.UpdatedBy
// 		updated_by.Valid = true
// 	}

// 	return db.UpdateEmpStatusParams{
// 		Status:      M.Status,
// 		Department:  M.Department,
// 		Designation: M.Designation,
// 		ValidFrom:   validFrom,
// 		ValidTill:   validTill,
// 		UpdatedBy:   updated_by,
// 		EmployeeID:  M.EmployeeID,
// 	}, nil
// }

// type UpdateEmpBenifitsReqModel struct {
// 	LeaveStatus        bool   `json:"leave_status"`
// 	LeaveType          string `json:"leave_type"`
// 	LeaveCount         int32  `json:"leave_count"`
// 	HealthInsurance    string `json:"health_insurance"`
// 	InsuranceFrom      string `json:"insurance_from"`
// 	InsuranceTill      string `json:"insurance_till"`
// 	RetainmentPlan     string `json:"retainment_plan"`
// 	RetainmentPlanFrom string `json:"retainment_plan_from"`
// 	RetainmentPlanTill string `json:"retainment_plan_till"`
// 	Benifits           string `json:"benifits"`
// 	BenifitsFrom       string `json:"benifits_from"`
// 	BenifitsTill       string `json:"benifits_till"`
// 	UpdatedBy          *int64 `json:"updated_by"`
// 	EmployeeID         int64  `json:"employee_id"`
// }

// func (M UpdateEmpBenifitsReqModel) convertToDbStruct() (db.UpdateEmpBenifitsParams, error) {
// 	insuranceFrom, err := time.Parse(time.RFC3339, M.InsuranceFrom)
// 	if err != nil {
// 		return db.UpdateEmpBenifitsParams{}, err
// 	}

// 	insuranceTill, err := time.Parse(time.RFC3339, M.InsuranceTill)
// 	if err != nil {
// 		return db.UpdateEmpBenifitsParams{}, err
// 	}

// 	retainmentPlanFrom, err := time.Parse(time.RFC3339, M.RetainmentPlanFrom)
// 	if err != nil {
// 		return db.UpdateEmpBenifitsParams{}, err
// 	}

// 	retainmentPlanTill, err := time.Parse(time.RFC3339, M.RetainmentPlanTill)
// 	if err != nil {
// 		return db.UpdateEmpBenifitsParams{}, err
// 	}

// 	benifitsFrom, err := time.Parse(time.RFC3339, M.BenifitsFrom)
// 	if err != nil {
// 		return db.UpdateEmpBenifitsParams{}, err
// 	}

// 	benifitsTill, err := time.Parse(time.RFC3339, M.BenifitsTill)
// 	if err != nil {
// 		return db.UpdateEmpBenifitsParams{}, err
// 	}

// 	var updated_by sql.NullInt64
// 	if M.UpdatedBy != nil {
// 		updated_by.Int64 = *M.UpdatedBy
// 		updated_by.Valid = true
// 	}

// 	return db.UpdateEmpBenifitsParams{
// 		LeaveStatus:        M.LeaveStatus,
// 		LeaveType:          M.LeaveType,
// 		LeaveCount:         M.LeaveCount,
// 		HealthInsurance:    M.HealthInsurance,
// 		InsuranceFrom:      insuranceFrom,
// 		InsuranceTill:      insuranceTill,
// 		RetainmentPlan:     M.RetainmentPlan,
// 		RetainmentPlanFrom: retainmentPlanFrom,
// 		RetainmentPlanTill: retainmentPlanTill,
// 		Benifits:           M.Benifits,
// 		BenifitsFrom:       benifitsFrom,
// 		BenifitsTill:       benifitsTill,
// 		UpdatedBy:          updated_by,
// 		EmployeeID:         M.EmployeeID,
// 	}, nil
// }

// type UpdateEmpUserReqModel struct {
// 	Email      string `json:"email"`
// 	Password   string `json:"password"`
// 	UpdatedBy  *int64 `json:"updated_by"`
// 	EmployeeID int64  `json:"employee_id"`
// }

// func (M UpdateEmpUserReqModel) convertToDbStruct() (db.UpdateEmpUserParams, error) {

// 	var updated_by sql.NullInt64
// 	if M.UpdatedBy != nil {
// 		updated_by.Int64 = *M.UpdatedBy
// 		updated_by.Valid = true
// 	}

// 	return db.UpdateEmpUserParams{
// 		Email:      M.Email,
// 		Password:   M.Password,
// 		UpdatedBy:  updated_by,
// 		EmployeeID: M.EmployeeID,
// 	}, nil
// }

// type UpdateEmpAllowancesReqModel struct {
// 	Name       string `json:"name"`
// 	Amount     string `json:"amount"`
// 	UpdatedBy  *int64 `json:"updated_by"`
// 	EmployeeID int64  `json:"employee_id"`
// }

// func (M UpdateEmpAllowancesReqModel) convertToDbStruct() (db.UpdateEmpAllowancesParams, error) {

// 	var updated_by sql.NullInt64
// 	if M.UpdatedBy != nil {
// 		updated_by.Int64 = *M.UpdatedBy
// 		updated_by.Valid = true
// 	}
// 	amount, err := decimal.NewFromString(M.Amount)
// 	if err != nil {
// 		return db.UpdateEmpAllowancesParams{}, fmt.Errorf("invalid amount format: %v", err)
// 	}

// 	return db.UpdateEmpAllowancesParams{
// 		Name:       M.Name,
// 		Amount:     amount,
// 		UpdatedBy:  updated_by,
// 		EmployeeID: M.EmployeeID,
// 	}, nil
// }

// type UpdateEmpExpatriateReqModel struct {
// 	Expatriate    bool   `json:"expatriate"`
// 	Nationality   string `json:"nationality"`
// 	VisaType      string `json:"visa_type"`
// 	VisaFrom      string `json:"visa_from"`
// 	VisaTill      string `json:"visa_till"`
// 	VisaNumber    string `json:"visa_number"`
// 	VisaFee       string `json:"visa_fee"`
// 	VisaImagePath string `json:"visa_image_path"`
// 	UpdatedBy     *int64 `json:"updated_by"`
// 	EmployeeID    int64  `json:"employee_id"`
// }

// func (M UpdateEmpExpatriateReqModel) convertToDbStruct() (db.UpdateEmpExpatriateParams, error) {
// 	visaFrom, err := time.Parse(time.RFC3339, M.VisaFrom)
// 	if err != nil {
// 		return db.UpdateEmpExpatriateParams{}, err
// 	}

// 	visaTill, err := time.Parse(time.RFC3339, M.VisaTill)
// 	if err != nil {
// 		return db.UpdateEmpExpatriateParams{}, err
// 	}

// 	var updated_by sql.NullInt64
// 	if M.UpdatedBy != nil {
// 		updated_by.Int64 = *M.UpdatedBy
// 		updated_by.Valid = true
// 	}

// 	visa_amount, err := decimal.NewFromString(M.VisaFee)
// 	if err != nil {
// 		return db.UpdateEmpExpatriateParams{}, fmt.Errorf("invalid amount format: %v", err)
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

// type UpdateEmpAccessiabilityReqModel struct {
// 	Accessibility     bool   `json:"accessibility"`
// 	AccessibilityFrom string `json:"accessibility_from"`
// 	AccessibilityTill string `json:"accessibility_till"`
// 	Enable            bool   `json:"enable"`
// 	UpdatedBy         *int64 `json:"updated_by"`
// 	EmployeeID        int64  `json:"employee_id"`
// }

// func (M UpdateEmpAccessiabilityReqModel) convertToDbStruct() (db.UpdateEmpAccessiabilityParams, error) {
// 	accessibilityFrom, err := time.Parse(time.RFC3339, M.AccessibilityFrom)
// 	if err != nil {
// 		return db.UpdateEmpAccessiabilityParams{}, err
// 	}

// 	accessibilityTill, err := time.Parse(time.RFC3339, M.AccessibilityTill)
// 	if err != nil {
// 		return db.UpdateEmpAccessiabilityParams{}, err
// 	}

// 	var updated_by sql.NullInt64
// 	if M.UpdatedBy != nil {
// 		updated_by.Int64 = *M.UpdatedBy
// 		updated_by.Valid = true
// 	}

// 	return db.UpdateEmpAccessiabilityParams{
// 		Accessibility:     M.Accessibility,
// 		AccessibilityFrom: accessibilityFrom,
// 		AccessibilityTill: accessibilityTill,
// 		Enable:            M.Enable,
// 		UpdatedBy:         updated_by,
// 		EmployeeID:        M.EmployeeID,
// 	}, nil
// }

// type EmpUpdateReqModel struct {
// 	Employee       UpdateEmployeeReqModel            `json:"employee"`
// 	Emergency      UpdateEmpEmergencyDetailsReqModel `json:"emergency"`
// 	Bank           UpdateEmpBankDetailsReqModel      `json:"bank"`
// 	Salary         UpdateEmpSalaryReqModel           `json:"salary"`
// 	Certificates   UpdateEmpCertificatesReqModel     `json:"certificates"`
// 	Status         UpdateEmpStatusReqModel           `json:"status"`
// 	Benifits       UpdateEmpBenifitsReqModel         `json:"benifits"`
// 	User           UpdateEmpUserReqModel             `json:"user"`
// 	Allowances     UpdateEmpAllowancesReqModel       `json:"allowances"`
// 	Expatriate     UpdateEmpExpatriateReqModel       `json:"expatriate"`
// 	Accessiability UpdateEmpAccessiabilityReqModel   `json:"accessiability"`
// }
