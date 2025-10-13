package hr

import (
	"fmt"
	"strconv"
	"encoding/json"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/Terracode-Dev/North-Star-Server/internal/database"
	"github.com/shopspring/decimal"
)

func (h *HRService) CreateEmpLink(c echo.Context) error{
	var req CreateEmpLinkReqParams
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, err.Error())
	}
	updated_by := 1
	params, err := req.ToCreateEmpLinkParams(int64(updated_by))

	if err != nil {
		return c.JSON(500, err.Error())
	}
	err = h.q.CreateEmpLink(c.Request().Context(), params)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, "Employee Link Created Successfully")
}

func (h *HRService) GetEmpLinkByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	data, err := h.q.GetEmpLinkByID(c.Request().Context(), int64(id))
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(200, data)
}

func (h *HRService) ListEmpLinks(c echo.Context) error {
	var req database.ListEmpLinksParams
	err := c.Bind(&req)
	if err != nil{
		return c.JSON(400, "error binding request")
	}
	links, err := h.q.ListEmpLinks(c.Request().Context(), req)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	totalCount, err := h.q.TotalEmpLinksCount(c.Request().Context())
	if err != nil {
		fmt.Printf("error getting total count: %v", err.Error())
		return c.JSON(500, "error getting total count")
	}
	return c.JSON(200, map[string]interface{}{
		"total_rows": totalCount,
		"links": links,
	})
}

func (h *HRService) ApproveEmpLink(c echo.Context) error {
	var req UpdateEmpLinkApprovalReqParams
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, err.Error())
	}
	updated_by, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(301, "Authentication issue")
	}

	linkData, err := h.q.GetEmpLinkData(c.Request().Context(), req.ID)

	var empData EmpDataModel

	var presetData Preset

	err = json.Unmarshal(linkData.EmpData, &empData)
	if err != nil {
		fmt.Println("error Unmarshaling Json", err.Error())
		return c.JSON(500, err)
	}

	err = json.Unmarshal(linkData.PresetValue, &presetData)
	if err != nil {
		fmt.Println("error Unmarshaling Json", err.Error())
		return c.JSON(500, err)
	}

	TrainerData := &TrainerDataModel{
		PresetData: presetData.IsTrainerAdmin,
		LinkData: empData.IsTrainerLink,
	}
	IsTrainerReqModel, err := TrainerData.ToIsTrainerParams()
	if err != nil {
		fmt.Println("error converting to Istrainer model", err.Error())
		return c.JSON(500, err)
	}

	empLink := EmpReqModel{
		Employee : empData.Employee,
		Emergency: empData.Emergency,
		Bank: empData.Bank,
		Salary: presetData.Salary,
		Status: presetData.Status,
		Benifits: presetData.Benifits,
		User: empData.User,
		Allowances: presetData.Allowances,
		Expatriate: empData.Expatriate,
		Accessiability: presetData.Accessiability,
		FileSubmit: empData.FileSubmit,
		IsTrainer: IsTrainerReqModel,
	}
	branch_id, ok := c.Get("branch").(int)
	if !ok {
		return c.JSON(301, "Authentication issue")
	}

	tx, err := h.db.Begin()
	if err != nil {
		return c.JSON(500, "Error starting transaction")
	}
	defer tx.Rollback()
	qtx := h.q.WithTx(tx)

	params, err := req.ToUpdateEmpLinkApprovalParams(int64(updated_by))
	if err != nil {
		fmt.Println("error creating to params", err.Error())
		return c.JSON(500, err)
	}
	
	err = qtx.UpdateEmpLinkApproval(c.Request().Context(), params)
	if err != nil {
		fmt.Println("error approving emplink", err.Error())
		return c.JSON(500, err)
	}

	empParams, err := empLink.Employee.convertToDbStruct(int64(updated_by))
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

	empLink.Emergency.EmployeeID = employeeID
	empLink.Bank.EmployeeID = employeeID
	empLink.Salary.EmployeeID = employeeID
	empLink.Status.EmployeeID = employeeID
	empLink.Benifits.EmployeeID = employeeID
	empLink.User.EmployeeID = employeeID
	empLink.Expatriate.EmployeeID = employeeID
	empLink.Accessiability.EmployeeID = employeeID
	empLink.IsTrainer.EmployeeID = employeeID

	emergencyParams, err := empLink.Emergency.convertToDbStruct(int64(updated_by))
	if err != nil {
		return c.JSON(500, "Error converting employee emergency details to db struct")
	}
	emergency := qtx.CreateEmpEmergencyDetails(c.Request().Context(), emergencyParams)
	if emergency != nil {
		return c.JSON(500, "Error creating employee emergency details")
	}

	bankParams, err := empLink.Bank.convertToDbStruct(int64(updated_by))
	if err != nil {
		return c.JSON(500, "Error converting employee bank details to db struct")
	}
	bank := qtx.CreateEmpBankDetails(c.Request().Context(), bankParams)
	if bank != nil {
		return c.JSON(500, "Error creating employee bank details")
	}

	salaryParams, err := empLink.Salary.convertToDbStruct(int64(updated_by))
	if err != nil {
		return c.JSON(500, "Error converting employee salary to db struct")
	}
	salary := qtx.CreateEmpSalary(c.Request().Context(), salaryParams)
	if salary != nil {
		return c.JSON(500, "Error creating employee salary"+salary.Error())
	}

	statusParams, err := empLink.Status.convertToDbStruct(int64(updated_by))
	if err != nil {
		return c.JSON(500, "Error converting employee status to db struct")
	}
	status := qtx.CreateEmpStatus(c.Request().Context(), statusParams)
	if status != nil {
		return c.JSON(500, "Error creating employee status")
	}

	benifitsParams, err := empLink.Benifits.convertToDbStruct(int64(updated_by))
	if err != nil {
		fmt.Printf("error %v", err)
		return c.JSON(500, "Error converting employee benifits to db struct")
	}
	benifits := qtx.CreateEmpBenifits(c.Request().Context(), benifitsParams)
	if benifits != nil {
		fmt.Printf("error %v", benifits)
		return c.JSON(500, "Error creating employee benifits")
	}

	userParams, err := empLink.User.convertToDbStruct(int64(branch_id), int64(updated_by))
	if err != nil {
		return c.JSON(500, "Error converting employee user to db struct")
	}
	user := qtx.CreateEmpUser(c.Request().Context(), userParams)
	if user != nil {
		return c.JSON(500, user.Error())
	}

	for _, allowance := range empLink.Allowances {
		allowance.EmployeeID = employeeID
		allowancesParams, err := allowance.convertToDbStruct(int64(updated_by))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error converting employee allowance to db struct: %v", err))
		}

		err = qtx.CreateEmpAllowances(c.Request().Context(), allowancesParams)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error creating employee allowance: %v", err))
		}
	}

	expatriateParams, err := empLink.Expatriate.convertToDbStruct(int64(updated_by))
	if err != nil {
		return c.JSON(500, "Error converting employee expatriate to db struct")
	}
	expatriate := qtx.CreateEmpExpatriate(c.Request().Context(), expatriateParams)
	if expatriate != nil {
		return c.JSON(500, "Error creating employee expatriate")
	}

	accessiabilityParams, err := empLink.Accessiability.convertToDbStruct(int64(updated_by))
	if err != nil {
		return c.JSON(500, "Error converting employee accessiability to db struct")
	}
	accessiability := qtx.CreateEmpAccessiability(c.Request().Context(), accessiabilityParams)
	if accessiability != nil {
		return c.JSON(500, "Error creating employee accessiability")
	}

	for _, file := range empLink.FileSubmit {
		file.EmployeeID = employeeID
		err = qtx.CreateFileSubmit(c.Request().Context(), database.CreateFileSubmitParams{
			EmployeeID: file.EmployeeID,
			FileName:   file.FileName,
			FileType:   file.FileType,
		})
		if err != nil {
			return c.JSON(500, "Error creating employee file submit: "+err.Error())
		}
	}

	if empLink.IsTrainer.IsTrainer {
		comValue, err := decimal.NewFromString(empLink.IsTrainer.Commission)
		if err != nil {
			return c.JSON(500, "Error converting commission value: "+err.Error())
		}
		istrainerrow := qtx.CreateTrainerEmp(c.Request().Context(), database.CreateTrainerEmpParams{
			TrainerID:  empLink.IsTrainer.TrainerID,
			EmployeeID: employeeID,
			AttendeeID: empLink.IsTrainer.AttendeeId,
			Commission: comValue,
		},
		)
		if istrainerrow != nil {
			return c.JSON(500, "Error creating employee trainer details: "+istrainerrow.Error())
		}
	}

	if err := tx.Commit(); err != nil {
		return c.JSON(500, map[string]string{"error": "Error committing transaction"})
	}

	return c.JSON(200, "user created successfully")
}

func (h *HRService) DeleteEmpLink(c echo.Context) error {
	id , err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("error converting id: %v", err.Error())
	}
	err = h.q.DeleteEmpLink(c.Request().Context(), int64(id))
	if err != nil {
		fmt.Printf("error deleting emp link: %v", err.Error())
		return c.JSON(500, "error deleting emp link")
	}
	return c.JSON(200, "emplink deleted successfully")
}