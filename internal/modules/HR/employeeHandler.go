package hr

import (
	"strconv"

	"fmt"
	"net/http"

	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
	"github.com/labstack/echo/v4"
)

// createemployee handler
func (S *HRService) createEmployee(c echo.Context) error {
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
	emp.Allowances.EmployeeID = employeeID
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

	userParams, err := emp.User.convertToDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee user to db struct")
	}
	user := qtx.CreateEmpUser(c.Request().Context(), userParams)
	if user != nil {
		return c.JSON(500, "Error creating employee user")
	}

	allowancesParams, err := emp.Allowances.convertToDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee allowances to db struct")
	}
	allowances := qtx.CreateEmpAllowances(c.Request().Context(), allowancesParams)
	if allowances != nil {
		return c.JSON(500, "Error creating employee allowances")
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

// get all employee handler
func (S *HRService) getEmployee(c echo.Context) error {
	var empParams db.GetEmployeeParams
	if err := c.Bind(&empParams); err != nil {
		return err
	}
	emp, err := S.q.GetEmployee(c.Request().Context(), empParams)
	if err != nil {
		return c.JSON(500, "Error getting employee")
	}
	return c.JSON(200, emp)
}

// get one employee handler
func (S *HRService) getEmployeeOne(c echo.Context) error {
	empID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(500, "Error parsing employee id")
	}
	emp, err := S.q.GetEmployeeByID(c.Request().Context(), empID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Database error: %v", err),
		})
	}
	return c.JSON(200, emp)
}

// update employee handler
func (S *HRService) updateEmployee(c echo.Context) error {
	var emp EmpReqModel
	if err := c.Bind(&emp); err != nil {
		return err
	}
	tx, err := S.db.Begin()
	if err != nil {
		return c.JSON(500, "Error starting transaction")
	}
	defer tx.Rollback()
	qtx := S.q.WithTx(tx)

	employeeID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(500, "Error parsing employee id")
	}

	empParams, err := emp.Employee.ConvertToUpdateDbStruct(employeeID)
	if err != nil {
		return c.JSON(500, "Error converting employee to db struct")
	}
	employee := qtx.UpdateEmployee(c.Request().Context(), empParams)
	if employee != nil {
		return c.JSON(500, "Error creating employee: "+err.Error())
	}

	emp.Emergency.EmployeeID = employeeID
	emp.Bank.EmployeeID = employeeID
	emp.Salary.EmployeeID = employeeID
	emp.Certificates.EmployeeID = employeeID
	emp.Status.EmployeeID = employeeID
	emp.Benifits.EmployeeID = employeeID
	emp.User.EmployeeID = employeeID
	emp.Allowances.EmployeeID = employeeID
	emp.Expatriate.EmployeeID = employeeID
	emp.Accessiability.EmployeeID = employeeID

	emergencyParams, err := emp.Emergency.convertToUpdateDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee emergency details to db struct")
	}
	emergency := qtx.UpdateEmpEmergencyDetails(c.Request().Context(), emergencyParams)
	if emergency != nil {
		return c.JSON(500, "Error creating employee emergency details")
	}

	bankParams, err := emp.Bank.convertToUpdateDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee bank details to db struct")
	}
	bank := qtx.UpdateEmpBankDetails(c.Request().Context(), bankParams)
	if bank != nil {
		return c.JSON(500, "Error creating employee bank details")
	}

	salaryParams, err := emp.Salary.convertToUpdateDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee salary to db struct")
	}
	salary := qtx.UpdateEmpSalary(c.Request().Context(), salaryParams)
	if salary != nil {
		return c.JSON(500, "Error creating employee salary"+salary.Error())
	}

	certificatesParams, err := emp.Certificates.convertToUpdateDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee certificates to db struct")
	}
	certificates := qtx.UpdateEmpCertificates(c.Request().Context(), certificatesParams)
	if certificates != nil {
		return c.JSON(500, "Error creating employee certificates")
	}

	statusParams, err := emp.Status.convertToUpdateDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee status to db struct")
	}
	status := qtx.UpdateEmpStatus(c.Request().Context(), statusParams)
	if status != nil {
		return c.JSON(500, "Error creating employee status")
	}

	benifitsParams, err := emp.Benifits.convertToUpdateDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee benifits to db struct")
	}
	benifits := qtx.UpdateEmpBenifits(c.Request().Context(), benifitsParams)
	if benifits != nil {
		return c.JSON(500, "Error creating employee benifits")
	}

	userParams, err := emp.User.convertToUpdateDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee user to db struct")
	}
	user := qtx.UpdateEmpUser(c.Request().Context(), userParams)
	if user != nil {
		return c.JSON(500, "Error creating employee user")
	}

	allowancesParams, err := emp.Allowances.convertToUpdateDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee allowances to db struct")
	}
	allowances := qtx.UpdateEmpAllowances(c.Request().Context(), allowancesParams)
	if allowances != nil {
		return c.JSON(500, "Error creating employee allowances")
	}

	expatriateParams, err := emp.Expatriate.convertToUpdateDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee expatriate to db struct")
	}
	expatriate := qtx.UpdateEmpExpatriate(c.Request().Context(), expatriateParams)
	if expatriate != nil {
		return c.JSON(500, "Error creating employee expatriate")
	}

	accessiabilityParams, err := emp.Accessiability.convertToUpdateDbStruct()
	if err != nil {
		return c.JSON(500, "Error converting employee accessiability to db struct")
	}
	accessiability := qtx.UpdateEmpAccessiability(c.Request().Context(), accessiabilityParams)
	if accessiability != nil {
		return c.JSON(500, "Error creating employee accessiability")
	}

	if err := tx.Commit(); err != nil {
		return c.JSON(500, map[string]string{"error": "Error committing transaction"})
	}

	return c.JSON(200, "Employee created successfully")
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
