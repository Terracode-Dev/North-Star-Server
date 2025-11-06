package hr

import (
	"fmt"
	"strconv"
	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
	"github.com/labstack/echo/v4"
)

func (S *HRService) CreateLoanRequest(c echo.Context) error {
	var req CreateRequestReqParams
	if err := c.Bind(&req); err != nil {
		fmt.Printf("Error binding request %v\n", err.Error())
		return c.JSON(500,"error binding request")
	}
	params, err := req.ToDbParams()
	if err != nil {
		fmt.Printf("Error creating to db struct: %v", err.Error())
		return c.JSON(500, "Error converting to d struct")
	}
	err = S.q.CreateRequest(c.Request().Context(), params)
	if err != nil {
		fmt.Printf("error creating loan request: %v", err.Error())
		return c.JSON(500, "error creating loan request")
	}
	return c.JSON(200, "loan request created successfully")
}

func (S *HRService) CancelRequestByEmployee( c echo.Context) error {
	loanId, err := strconv.Atoi(c.Param("id")) 
	if err != nil {
		fmt.Printf("no id in parameters %v", err.Error())
		return c.JSON(404, "no id found in parameters")
	}
	err = S.q.DeleteRequest(c.Request().Context(), int64(loanId))
	if err != nil {
		fmt.Printf("error canceling request %v", err.Error())
		return c.JSON(500, "error canceling request")
	}
	return c.JSON(200, "request cancelled successfully")
}

func (S *HRService) UpdateRequest(c echo.Context) error {
	var req UpdateRequestReqParams
	if err := c.Bind(&req); err != nil {
		fmt.Printf("Error binding request %v\n", err.Error())
		return c.JSON(500,"error binding request")
	}
	params , err := req.ToDbParams()
	if err != nil {
		fmt.Printf("Error creating to db struct: %v", err.Error())
		return c.JSON(500, "Error converting to d struct")
	}
	err = S.q.UpdateRequest(c.Request().Context(), params)
	if err != nil {
		fmt.Printf("error updating request %v", err.Error())
		return c.JSON(500, "error updating request")
	}
	return c.JSON(200, "request updated successfully")
}

func (S *HRService) GetRequestForEmployee(c echo.Context) error {
	var req db.GetRequestsParams
	if err := c.Bind(&req); err != nil {
		fmt.Printf("Error binding request %v\n", err.Error())
		return c.JSON(500,"error binding request")
	}
	loans, err := S.q.GetRequests(c.Request().Context(), req)
	if err != nil {
		fmt.Printf("Error Getting loan data %v", err.Error())
		return c.JSON(500, "error getting loan data")
	}	
	return c.JSON(200, loans)
}

func (S *HRService) GetRequestForAdmin(c echo.Context) error {
	var req GetRequestsAdminReqParams
	if err := c.Bind(&req); err != nil {
		fmt.Printf("Error binding request %v\n", err.Error())
		return c.JSON(500,"error binding request")
	}
	branchId, ok := c.Get("branch").(int)
	if !ok {
		return c.JSON(500, "no branch id found in cookie")
	}
	params , err := req.ToDbParams(int64(branchId))
	if err != nil {
		fmt.Printf("Error converting to db struct %v", err.Error())
		return c.JSON(500, "error converting to db struct")
	}
	loans, err := S.q.GetRequestsAdmin(c.Request().Context(), params)
	if err != nil {
		fmt.Printf("Error Getting loan data %v", err.Error())
		return c.JSON(500, "error getting loan data")
	}	
	return c.JSON(200, loans)
}

func (S *HRService) UpdateStatus(c echo.Context) error {
	var req UpdateRequestStatusReqParams
	if err := c.Bind(&req); err != nil {
		fmt.Printf("Error binding request %v\n", err.Error())
		return c.JSON(500,"error binding request")
	}
	params, err := req.ToDbParams()
	if err != nil {
		fmt.Printf("Error creating to db struct: %v", err.Error())
		return c.JSON(500, "Error converting to d struct")
	}
	err = S.q.UpdateRequestStatus(c.Request().Context(), params)
	if err != nil {
		fmt.Printf("Error updating loan data %v", err.Error())
		return c.JSON(500, "error updating loan data")		
	}
	return c.JSON(200, "status updated successfully")
}