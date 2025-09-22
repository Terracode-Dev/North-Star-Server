package hr

import (
	"fmt"
	"encoding/json"

	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
)

type Preset struct {
	Salary         CreateEmpSalaryReqModel           `json:"salary"`
	Status         CreateEmpStatusReqModel           `json:"status"`
	Benifits       CreateEmpBenifitsReqModel         `json:"benifits"`
	Allowances     []CreateEmpAllowancesReqModel     `json:"allowances"`
	Accessiability CreateEmpAccessiabilityReqModel   `json:"accessiability"`
	IsTrainer      IsTrainerReqModel			 	 `json:"is_trainer"`
}

type CreateAdminPresetReqParams struct {
	PresetName  string          `json:"preset_name"`
	PresetValue Preset          `json:"preset_value"`
	Slug        string          `json:"slug"`
}

func (c *CreateAdminPresetReqParams) ToCreateAdminPresetParams() (db.CreateAdminPresetParams, error) {
	presetValue, err := json.Marshal(c.PresetValue)
	if err != nil {
		return db.CreateAdminPresetParams{}, fmt.Errorf("failed to marshal preset value: %w", err)
	}
	return db.CreateAdminPresetParams{
		PresetName:  c.PresetName,
		PresetValue: presetValue,
		Slug:        c.Slug,
	}, nil
}
