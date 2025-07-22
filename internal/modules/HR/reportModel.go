package hr

import (
	"fmt"

	"time"

	 db "github.com/Terracode-Dev/North-Star-Server/internal/database"

)

type SalaryTransferRes struct {
	AccountData []db.GetAccountDetailsRow `json:"accountdata"`
	Total      float64                   `json:"total"`
}

type ReportsReq struct {
	Date   string `json:"date"`
}

func (s *ReportsReq) ConvertToDbParams() (time.Time, error) {
	date, err := time.Parse("2006-01-02", s.Date)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format: %v", err)
	}
	return date, nil
}
