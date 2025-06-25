package hr

import (
	"log"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

func (S *HRService) CreateExchangeRate(c echo.Context) error {
	var req ExchangeRateReqModel
    if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	dbParams, err := req.convertToDbStruct()
	if err != nil {
  		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
 	}
	err = S.q.CreateExchangeRate(c.Request().Context(), dbParams)
 	if err != nil {
	  	log.Printf("Error creating exchange rate: %v", err)
  		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create exchange rate"})
 	}
 	return c.JSON(http.StatusCreated, map[string]string{"message": "Exchange rate created successfully"})
}

func (S *HRService) DeleteExchangeRate(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid exchange rate ID"})
	}
	err = S.q.DeleteExchangeRate(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete exchange rate"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Exchange rate deleted successfully"})
}

func (S *HRService) GetExchangeRate(c echo.Context) error {
	currency_type := c.Param("type")
	if currency_type == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Currency type is required"})
	}
	exchange_rate, err := S.q.GetLatestExchangeRate(c.Request().Context(), currency_type)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve exchange rate"})
	}
	return c.JSON(http.StatusOK, exchange_rate)
}
