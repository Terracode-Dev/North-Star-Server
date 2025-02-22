package hr

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Terracode-Dev/North-Star-Server/internal/database"
	rba "github.com/Terracode-Dev/North-Star-Server/internal/pkg/RBA"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// CreateEmployee godoc
// @Summary Create Employee
// @Description Creates a new employee along with all related records (emergency contacts, bank details, salary, etc.) in a transaction
// @Tags employee
// @Accept json
// @Produce json
// @Param employee body EmpReqModel true "Employee data with all related details"
// @Success 200 {string} string "Employee created successfully"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /employee [post]
func (S *HRService) createEmployee(c echo.Context) error {
	branch_id, ok := c.Get("branch").(int64)
	if !ok {
		return c.JSON(301, "Authentication issue")
	}

	var emp EmpReqModel
	if err := c.Bind(&emp); err != nil {
		return err
	}
	fmt.Println(emp)
	tx, err := S.db.Begin()
	if err != nil {
		return c.JSON(500, "Error starting transaction")
	}
	defer tx.Rollback()
	qtx := S.q.WithTx(tx)
	empParams, err := emp.Employee.convertToDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee to db struct")
	}

	employee, err := qtx.CreateEmployee(c.Request().Context(), empParams)
	if err != nil {
		return c.JSON(500, "Error creating employee: "+err.Error())
	}
	employeeID, err := employee.LastInsertId()
	if err != nil {
		return c.JSON(500, "Error getting employee ID: "+err.Error())
	}

	emp.Emergency.EmployeeID = employeeID
	emp.Bank.EmployeeID = employeeID
	emp.Salary.EmployeeID = employeeID
	emp.Certificates.EmployeeID = employeeID
	emp.Status.EmployeeID = employeeID
	emp.Benifits.EmployeeID = employeeID
	emp.User.EmployeeID = employeeID
	// emp.Allowances.EmployeeID = employeeID
	emp.Expatriate.EmployeeID = employeeID
	emp.Accessiability.EmployeeID = employeeID

	emergencyParams, err := emp.Emergency.convertToDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee emergency details to db struct")
	}
	emergency := qtx.CreateEmpEmergencyDetails(c.Request().Context(), emergencyParams)
	if emergency != nil {
		return c.JSON(500, "Error creating employee emergency details")
	}

	bankParams, err := emp.Bank.convertToDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee bank details to db struct")
	}
	bank := qtx.CreateEmpBankDetails(c.Request().Context(), bankParams)
	if bank != nil {
		return c.JSON(500, "Error creating employee bank details")
	}

	salaryParams, err := emp.Salary.convertToDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee salary to db struct")
	}
	salary := qtx.CreateEmpSalary(c.Request().Context(), salaryParams)
	if salary != nil {
		return c.JSON(500, "Error creating employee salary"+salary.Error())
	}

	certificatesParams, err := emp.Certificates.convertToDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee certificates to db struct")
	}
	certificates := qtx.CreateEmpCertificates(c.Request().Context(), certificatesParams)
	if certificates != nil {
		return c.JSON(500, "Error creating employee certificates")
	}

	statusParams, err := emp.Status.convertToDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee status to db struct")
	}
	status := qtx.CreateEmpStatus(c.Request().Context(), statusParams)
	if status != nil {
		return c.JSON(500, "Error creating employee status")
	}

	benifitsParams, err := emp.Benifits.convertToDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee benifits to db struct")
	}
	benifits := qtx.CreateEmpBenifits(c.Request().Context(), benifitsParams)
	if benifits != nil {
		return c.JSON(500, "Error creating employee benifits")
	}

	userParams, err := emp.User.convertToDbStruct(branch_id)
	if err != nil {
		return c.JSON(500, "Error converting employee user to db struct")
	}
	user := qtx.CreateEmpUser(c.Request().Context(), userParams)
	if user != nil {
		return c.JSON(500, user.Error())
	}

	for _, allowance := range emp.Allowances {
		allowance.EmployeeID = employeeID
		allowancesParams, err := allowance.convertToDbStruct()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error converting employee allowance to db struct: %v", err))
		}

		err = qtx.CreateEmpAllowances(c.Request().Context(), allowancesParams)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error creating employee allowance: %v", err))
		}
	}

	expatriateParams, err := emp.Expatriate.convertToDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee expatriate to db struct")
	}
	expatriate := qtx.CreateEmpExpatriate(c.Request().Context(), expatriateParams)
	if expatriate != nil {
		return c.JSON(500, "Error creating employee expatriate")
	}

	accessiabilityParams, err := emp.Accessiability.convertToDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee accessiability to db struct")
	}
	accessiability := qtx.CreateEmpAccessiability(c.Request().Context(), accessiabilityParams)
	if accessiability != nil {
		return c.JSON(500, "Error creating employee accessiability")
	}

	if err := tx.Commit(); err != nil {
		return c.JSON(500, map[string]string{"error": "Error committing transaction"})
	}

	return c.JSON(200, "Employee created successfully")
}

// GetEmployee godoc
// @Summary      Get Employees
// @Description  Fetch employee details with search and pagination. Main branch users can view all employees; others see only employees in their branch.
// @Tags         employee
// @Accept       json
// @Produce      json
// @Param        employee  body      GetEmployeeReqModel  true  "Search criteria and pagination for employees"
// @Success      200       {string}  string "Employees fetched successfully"
// @Failure      400       {string}  string               "Bad request"
// @Failure      500       {string}  string               "Internal server error"
// @Router       /employee/all [post]
func (S *HRService) getEmployee(c echo.Context) error {
	var empReqModel GetEmployeeReqModel
	if err := c.Bind(&empReqModel); err != nil {
		return c.JSON(500, "bind error")
	}
	branch_id, ok := c.Get("branch").(int)
	if !ok {
		return c.JSON(400, "Invalid branch ID")
	}
	var params database.GetEmployeeParams
	if branch_id == int(S.cfg.MainBranchId) {
		log.Println(empReqModel)
		params = database.GetEmployeeParams{
			CONCAT:   empReqModel.Search,
			CONCAT_2: empReqModel.Search,
			CONCAT_3: empReqModel.Search,
			CONCAT_4: empReqModel.Search,
			CONCAT_5: empReqModel.Search,
			Column6:  "",
			ID:       0,
			Limit:    empReqModel.Limit,
			Offset:   (empReqModel.PageNumber - 1) * empReqModel.Limit,
		}
	} else {
		params = database.GetEmployeeParams{
			CONCAT:   empReqModel.Search,
			CONCAT_2: empReqModel.Search,
			CONCAT_3: empReqModel.Search,
			CONCAT_4: empReqModel.Search,
			CONCAT_5: empReqModel.Search,
			Column6:  "1",
			ID:       int64(branch_id),
			Limit:    empReqModel.Limit,
			Offset:   (empReqModel.PageNumber - 1) * empReqModel.Limit,
		}
	}

	emp, err := S.q.GetEmployee(c.Request().Context(), params)
	if err != nil {
		return c.JSON(500, "Error getting employee")
	}
	return c.JSON(200, emp)
}

// get one employee handler
func (S *HRService) getEmployeeOne(c echo.Context) error {
	empID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid employee ID"})
	}

	// Fetch employee data
	emp, err := S.q.GetEmployeeByID(c.Request().Context(), empID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Database error: %v", err)})
	}

	// Check if employee exists
	if len(emp) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Employee not found"})
	}

	// Authentication
	userId, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	role, ok := c.Get("role").(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid role"})
	}

	branchId, ok := c.Get("branch").(int)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid branch ID"})
	}

	// Authorization Logic
	switch role {
	case "admin", "mod":
		if branchId != S.cfg.MainBranchId && emp[0].UserBranchID.Valid && branchId != int(emp[0].UserBranchID.Int64) {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "Unauthorized access"})
		}
	case "emp":
		if branchId != S.cfg.MainBranchId && userId != int(emp[0].EmployeeID) {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "Unauthorized access"})
		}
	default:
		return c.JSON(http.StatusForbidden, map[string]string{"error": "Unknown role, access denied"})
	}

	return c.JSON(http.StatusOK, emp)
}

// update employee handler
// func (S *HRService) updateEmployee(c echo.Context) error {
// 	var emp EmpReqModel
// 	if err := c.Bind(&emp); err != nil {
// 		return err
// 	}
// 	tx, err := S.db.Begin()
// 	if err != nil {
// 		return c.JSON(500, "Error starting transaction")
// 	}
// 	defer tx.Rollback()
// 	qtx := S.q.WithTx(tx)

// 	employeeID, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	if err != nil {
// 		return c.JSON(500, "Error parsing employee id")
// 	}

// 	empParams, err := emp.Employee.ConvertToUpdateDbStruct(employeeID)
// 	if err != nil {
// 		return c.JSON(500, "Error converting employee to db struct")
// 	}
// 	employee := qtx.UpdateEmployee(c.Request().Context(), empParams)
// 	if employee != nil {
// 		return c.JSON(500, "Error creating employee: "+err.Error())
// 	}

// 	emp.Emergency.EmployeeID = employeeID
// 	emp.Bank.EmployeeID = employeeID
// 	emp.Salary.EmployeeID = employeeID
// 	emp.Certificates.EmployeeID = employeeID
// 	emp.Status.EmployeeID = employeeID
// 	emp.Benifits.EmployeeID = employeeID
// 	emp.User.EmployeeID = employeeID
// 	emp.Allowances.EmployeeID = employeeID
// 	emp.Expatriate.EmployeeID = employeeID
// 	emp.Accessiability.EmployeeID = employeeID

// 	emergencyParams, err := emp.Emergency.convertToUpdateDbStruct()
// 	if err != nil {
// 		return c.JSON(500, "Error converting employee emergency details to db struct")
// 	}
// 	emergency := qtx.UpdateEmpEmergencyDetails(c.Request().Context(), emergencyParams)
// 	if emergency != nil {
// 		return c.JSON(500, "Error creating employee emergency details")
// 	}

// 	bankParams, err := emp.Bank.convertToUpdateDbStruct()
// 	if err != nil {
// 		return c.JSON(500, "Error converting employee bank details to db struct")
// 	}
// 	bank := qtx.UpdateEmpBankDetails(c.Request().Context(), bankParams)
// 	if bank != nil {
// 		return c.JSON(500, "Error creating employee bank details")
// 	}

// 	salaryParams, err := emp.Salary.convertToUpdateDbStruct()
// 	if err != nil {
// 		return c.JSON(500, "Error converting employee salary to db struct")
// 	}
// 	salary := qtx.UpdateEmpSalary(c.Request().Context(), salaryParams)
// 	if salary != nil {
// 		return c.JSON(500, "Error creating employee salary"+salary.Error())
// 	}

// 	certificatesParams, err := emp.Certificates.convertToUpdateDbStruct()
// 	if err != nil {
// 		return c.JSON(500, "Error converting employee certificates to db struct")
// 	}
// 	certificates := qtx.UpdateEmpCertificates(c.Request().Context(), certificatesParams)
// 	if certificates != nil {
// 		return c.JSON(500, "Error creating employee certificates")
// 	}

// 	statusParams, err := emp.Status.convertToUpdateDbStruct()
// 	if err != nil {
// 		return c.JSON(500, "Error converting employee status to db struct")
// 	}
// 	status := qtx.UpdateEmpStatus(c.Request().Context(), statusParams)
// 	if status != nil {
// 		return c.JSON(500, "Error creating employee status")
// 	}

// 	benifitsParams, err := emp.Benifits.convertToUpdateDbStruct()
// 	if err != nil {
// 		return c.JSON(500, "Error converting employee benifits to db struct")
// 	}
// 	benifits := qtx.UpdateEmpBenifits(c.Request().Context(), benifitsParams)
// 	if benifits != nil {
// 		return c.JSON(500, "Error creating employee benifits")
// 	}

// 	userParams, err := emp.User.convertToUpdateDbStruct()
// 	if err != nil {
// 		return c.JSON(500, "Error converting employee user to db struct")
// 	}
// 	user := qtx.UpdateEmpUser(c.Request().Context(), userParams)
// 	if user != nil {
// 		return c.JSON(500, "Error creating employee user")
// 	}

// 	allowancesParams, err := emp.Allowances.convertToUpdateDbStruct()
// 	if err != nil {
// 		return c.JSON(500, "Error converting employee allowances to db struct")
// 	}
// 	allowances := qtx.UpdateEmpAllowances(c.Request().Context(), allowancesParams)
// 	if allowances != nil {
// 		return c.JSON(500, "Error creating employee allowances")
// 	}

// 	expatriateParams, err := emp.Expatriate.convertToUpdateDbStruct()
// 	if err != nil {
// 		return c.JSON(500, "Error converting employee expatriate to db struct")
// 	}
// 	expatriate := qtx.UpdateEmpExpatriate(c.Request().Context(), expatriateParams)
// 	if expatriate != nil {
// 		return c.JSON(500, "Error creating employee expatriate")
// 	}

// 	accessiabilityParams, err := emp.Accessiability.convertToUpdateDbStruct()
// 	if err != nil {
// 		return c.JSON(500, "Error converting employee accessiability to db struct")
// 	}
// 	accessiability := qtx.UpdateEmpAccessiability(c.Request().Context(), accessiabilityParams)
// 	if accessiability != nil {
// 		return c.JSON(500, "Error creating employee accessiability")
// 	}

// 	if err := tx.Commit(); err != nil {
// 		return c.JSON(500, map[string]string{"error": "Error committing transaction"})
// 	}

// 	return c.JSON(200, "Employee created successfully")
// }

// update employee handler
func (S *HRService) updateEmployee(c echo.Context) error {
	var emp CreateEmployeeReqModel
	if err := c.Bind(&emp); err != nil {
		return c.JSON(500, err)
	}
	empID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(500, "Error parsing employee id")
	}
	empParams, err := emp.ConvertToUpdateDbStruct(empID)
	if err != nil {
		return c.JSON(500, "Error converting employee to db struct")
	}
	error := S.q.UpdateEmployee(c.Request().Context(), empParams)
	if error != nil {
		return c.JSON(500, "Error updating employee")
	}
	return c.JSON(200, "Employee updated successfully")
}

// update employee emergency details handler
func (S *HRService) updateEmpEmergencyDetails(c echo.Context) error {
	var emg CreateEmpEmergencyDetailsReqModel
	if err := c.Bind(&emg); err != nil {
		return c.JSON(500, err)
	}
	emgParams, err := emg.convertToUpdateDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee emergency details to db struct")
	}
	error := S.q.UpdateEmpEmergencyDetails(c.Request().Context(), emgParams)
	if error != nil {
		return c.JSON(500, "Error updating employee emergency details")
	}
	return c.JSON(200, "Employee emergency details updated successfully")
}

// update employee bank details handler
func (S *HRService) updateEmpBankDetails(c echo.Context) error {
	var bank CreateEmpBankDetailsReqModel
	if err := c.Bind(&bank); err != nil {
		return c.JSON(500, err)
	}
	bankParams, err := bank.convertToUpdateDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee bank details to db struct")
	}
	error := S.q.UpdateEmpBankDetails(c.Request().Context(), bankParams)
	if error != nil {
		return c.JSON(500, "Error updating employee bank details")
	}
	return c.JSON(200, "Employee bank details updated successfully")
}

// update employee salary handler
func (S *HRService) updateEmpSalary(c echo.Context) error {
	var salary CreateEmpSalaryReqModel
	if err := c.Bind(&salary); err != nil {
		return c.JSON(500, err)
	}
	salaryParams, err := salary.convertToUpdateDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee salary to db struct")
	}
	error := S.q.UpdateEmpSalary(c.Request().Context(), salaryParams)
	if error != nil {
		return c.JSON(500, "Error updating employee salary")
	}
	return c.JSON(200, "Employee salary updated successfully")
}

// update employee certificates handler
func (S *HRService) updateEmpCertificates(c echo.Context) error {
	var certificates CreateEmpCertificatesReqModel
	if err := c.Bind(&certificates); err != nil {
		return c.JSON(500, err)
	}
	certificatesParams, err := certificates.convertToUpdateDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee certificates to db struct")
	}
	error := S.q.UpdateEmpCertificates(c.Request().Context(), certificatesParams)
	if error != nil {
		return c.JSON(500, "Error updating employee certificates")
	}
	return c.JSON(200, "Employee certificates updated successfully")
}

// update employee status handler
func (S *HRService) updateEmpStatus(c echo.Context) error {
	var status CreateEmpStatusReqModel
	if err := c.Bind(&status); err != nil {
		return c.JSON(500, err)
	}
	statusParams, err := status.convertToUpdateDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee status to db struct")
	}
	error := S.q.UpdateEmpStatus(c.Request().Context(), statusParams)
	if error != nil {
		return c.JSON(500, "Error updating employee status")
	}
	return c.JSON(200, "Employee status updated successfully")
}

// update employee benifits handler
func (S *HRService) updateEmpBenifits(c echo.Context) error {
	var benifits CreateEmpBenifitsReqModel
	if err := c.Bind(&benifits); err != nil {
		return c.JSON(500, err)
	}
	benifitsParams, err := benifits.convertToUpdateDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee benifits to db struct")
	}
	error := S.q.UpdateEmpBenifits(c.Request().Context(), benifitsParams)
	if error != nil {
		return c.JSON(500, "Error updating employee benifits")
	}
	return c.JSON(200, "Employee benifits updated successfully")
}

// update employee user handler
func (S *HRService) updateEmpUser(c echo.Context) error {
	var user CreateEmpUserReqModel
	if err := c.Bind(&user); err != nil {
		return c.JSON(500, err)
	}
	userParams, err := user.convertToUpdateDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee user to db struct")
	}
	error := S.q.UpdateEmpUser(c.Request().Context(), userParams)
	if error != nil {
		return c.JSON(500, "Error updating employee user")
	}
	return c.JSON(200, "Employee user updated successfully")
}

// update employee allowances handler
func (S *HRService) updateEmpAllowances(c echo.Context) error {
	var allowances CreateEmpAllowancesReqModel
	if err := c.Bind(&allowances); err != nil {
		return c.JSON(500, err)
	}
	allowancesParams, err := allowances.convertToUpdateDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee allowances to db struct")
	}
	error := S.q.UpdateEmpAllowances(c.Request().Context(), allowancesParams)
	if error != nil {
		return c.JSON(500, "Error updating employee allowances")
	}
	return c.JSON(200, "Employee allowances updated successfully")
}

// update employee expatriate handler
func (S *HRService) updateEmpExpatriate(c echo.Context) error {
	var expatriate CreateEmpExpatriateReqModel
	if err := c.Bind(&expatriate); err != nil {
		return c.JSON(500, err)
	}
	expatriateParams, err := expatriate.convertToUpdateDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee expatriate to db struct")
	}
	error := S.q.UpdateEmpExpatriate(c.Request().Context(), expatriateParams)
	if error != nil {
		return c.JSON(500, "Error updating employee expatriate")
	}
	return c.JSON(200, "Employee expatriate updated successfully")
}

// update employee accessiability handler
func (S *HRService) updateEmpAccessiability(c echo.Context) error {
	var accessiability CreateEmpAccessiabilityReqModel
	if err := c.Bind(&accessiability); err != nil {
		return c.JSON(500, err)
	}
	accessiabilityParams, err := accessiability.convertToUpdateDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee accessiability to db struct")
	}
	error := S.q.UpdateEmpAccessiability(c.Request().Context(), accessiabilityParams)
	if error != nil {
		return c.JSON(500, "Error updating employee accessiability")
	}
	return c.JSON(200, "Employee accessiability updated successfully")
}

// delete employee handler
func (S *HRService) deleteEmployee(c echo.Context) error {
	empID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(500, "Error parsing employee id")
	}
	tx, err := S.db.Begin()
	if err != nil {
		return c.JSON(500, "Error starting transaction")
	}
	defer tx.Rollback()
	qtx := S.q.WithTx(tx)

	emergency := qtx.DeleteEmpEmergencyDetails(c.Request().Context(), empID)
	if emergency != nil {
		return c.JSON(500, "Error deleting employee emergency details")
	}

	bank := qtx.DeleteEmpBankDetails(c.Request().Context(), empID)
	if bank != nil {
		return c.JSON(500, "Error deleting employee bank details")
	}

	salary := qtx.DeleteEmpSalary(c.Request().Context(), empID)
	if salary != nil {
		return c.JSON(500, "Error deleting employee salary")
	}

	certificates := qtx.DeleteEmpCertificates(c.Request().Context(), empID)
	if certificates != nil {
		return c.JSON(500, "Error deleting employee certificates")
	}

	status := qtx.DeleteEmpStatus(c.Request().Context(), empID)
	if status != nil {
		return c.JSON(500, "Error deleting employee status")
	}

	benifits := qtx.DeleteEmpBenifits(c.Request().Context(), empID)
	if benifits != nil {
		return c.JSON(500, "Error deleting employee benifits")
	}

	user := qtx.DeleteEmpUser(c.Request().Context(), empID)
	if user != nil {
		return c.JSON(500, "Error deleting employee user")
	}

	allowances := qtx.DeleteEmpAllowances(c.Request().Context(), empID)
	if allowances != nil {
		return c.JSON(500, "Error deleting employee allowances")
	}

	expatriate := qtx.DeleteEmpExpatriate(c.Request().Context(), empID)
	if expatriate != nil {
		return c.JSON(500, "Error deleting employee expatriate")
	}

	accessiability := qtx.DeleteEmpAccessiability(c.Request().Context(), empID)
	if accessiability != nil {
		return c.JSON(500, "Error deleting employee accessiability")
	}

	employee := qtx.DeleteEmployee(c.Request().Context(), empID)
	if employee != nil {
		return c.JSON(500, "Error deleting employee")
	}

	if err := tx.Commit(); err != nil {
		return c.JSON(500, map[string]string{"error": "Error committing transaction"})
	}

	return c.JSON(200, "Employee deleted successfully")
}

// @Summary employee login
// @Description Authenticates user and returns a JWT token in a cookie
// @Tags auth
// @Accept json
// @Produce json
// @Param user body EmpLoginReqModel true "User credentials"
// @Success 200 {string} int
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /employee/login [post]
func (S *HRService) employeeLogin(c echo.Context) error {
	var login EmpLoginReqModel
	if err := c.Bind(&login); err != nil {
		return c.JSON(500, err)
	}
	emp, err := S.q.EmployeeLogin(c.Request().Context(), login.Email)
	if err != nil {
		return c.JSON(301, "invalid email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(emp.Password), []byte(login.Password))
	if err != nil {
		return c.JSON(301, "invalid Password")
	}

	payload := rba.RBAauth{
		Id:     int(emp.EmployeeID),
		Role:   "emp",
		Email:  emp.Email,
		Branch: int(emp.BranchID),
	}

	t, err := rba.GenarateJWTkey(time.Hour*24, payload, []byte(S.cfg.JWTSecret))
	if err != nil {
		log.Println(err)
		return c.JSON(500, t)
	}

	cookie := new(http.Cookie)
	cookie.Name = "auth_token"
	cookie.Value = t
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(time.Hour * time.Duration(S.cfg.JwtExpHour))
	c.SetCookie(cookie)
	return c.JSON(200, "login successfull")
}

// @Summary user loginout
// @Description Authenticates user and returns a JWT token in a cookie
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {string} string "Logout successful"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /logout [get]
func (S *HRService) Logout(c echo.Context) error {
	// Invalidate the cookie by setting its value to empty and expiry to the past
	cookie := new(http.Cookie)
	cookie.Name = "auth_token"
	cookie.Value = ""
	// cookie.Path = "/" // Ensure the path matches the login cookie
	cookie.Expires = time.Unix(0, 0) // Set expiry to a time in the past
	// cookie.HttpOnly = true          // Prevent client-side scripts from accessing it
	// cookie.Secure = true            // Ensure the cookie is only sent over HTTPS
	cookie.SameSite = http.SameSiteStrictMode // Helps mitigate CSRF attacks

	c.SetCookie(cookie)

	// If using a token revocation strategy, add the token to the blacklist here
	// Example (pseudocode):
	// token := getTokenFromContext(c)
	// S.tokenStore.Revoke(token)

	return c.JSON(200, map[string]string{"message": "Logout successful"})
}

func (S *HRService) empOnlyBankDetailsUpdate(c echo.Context) error {
	userId, ok := c.Get("user_id").(int64)
	if !ok {
		return c.JSON(500, "user id convertion issue")
	}
	var bank CreateEmpBankDetailsReqModel
	if err := c.Bind(&bank); err != nil {
		return c.JSON(500, err)
	}
	if userId != bank.EmployeeID {
		return c.JSON(301, "Authentication issue")
	}
	bankParams, err := bank.convertToUpdateDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee bank details to db struct")
	}
	error := S.q.UpdateEmpBankDetails(c.Request().Context(), bankParams)
	if error != nil {
		return c.JSON(500, "Error updating employee bank details")
	}
	return c.JSON(200, "Employee bank details updated successfully")
}
