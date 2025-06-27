package hr

import (
	"database/sql"
	db "github.com/Terracode-Dev/North-Star-Server/internal/database"
)

type CreatePresetReqModel struct {
	Name        string         `json:"name"`
	Description string `json:"description"`
	TrainerID   int64          `json:"trainer_id"`
}

func (v *CreatePresetReqModel) ToCreatePresetParams() db.CreatePresetParams {
	var description sql.NullString
	description.Valid = true
	description.String = v.Description

	return  db.CreatePresetParams{
		Name:        v.Name,
		Description: description,
		TrainerID:   v.TrainerID,
	}
}

func (v *CreatePresetReqModel) ToUpdatePresetParams(id int64) db.UpdatePresetParams {
	var description sql.NullString
	description.Valid = true
	description.String = v.Description

	return db.UpdatePresetParams{
		ID:          id,
		Name:        v.Name,
		Description: description,
		TrainerID:   v.TrainerID,
	}
}

type CreatePresetWorkoutreqModel struct {
	PresetID  int64          `json:"preset_id"`
	WorkoutID int64          `json:"workout_id"`
	Reps      int64          `json:"reps"`
	Weight    string         `json:"weight"`
	Sets      int64          `json:"sets"`
	Notes     string         `json:"notes"`
}

func (v *CreatePresetWorkoutreqModel) ToCreatePresetWorkoutParams() db.CreatePresetWorkoutParams {
	var weight sql.NullString
	weight.Valid = true
	weight.String = v.Weight

	var notes sql.NullString
	notes.Valid = true
	notes.String = v.Notes

	return db.CreatePresetWorkoutParams{
		PresetID:  v.PresetID,
		WorkoutID: v.WorkoutID,
		Reps:      v.Reps,
		Weight:    weight,
		Sets:      v.Sets,
		Notes:     notes,
	}
}

func (v *CreatePresetWorkoutreqModel) ToUpdatePresetWorkoutParams(id int64) db.UpdatePresetWorkoutParams {
	var weight sql.NullString
	weight.Valid = true
	weight.String = v.Weight

	var notes sql.NullString
	notes.Valid = true
	notes.String = v.Notes

	return db.UpdatePresetWorkoutParams{
		ID:        id,
		PresetID:  v.PresetID,
		WorkoutID: v.WorkoutID,
		Reps:      v.Reps,
		Weight:    weight,
		Sets:      v.Sets,
		Notes:     notes,
	}
}

type CreateSessionWorkoutReqModel struct {
	PresetSessionID int64          `json:"preset_session_id"`
	ActiveDay       int64          `json:"active_day"`
	WorkoutID       int64          `json:"workout_id"`
	Status          string         `json:"status"`
	SessionID       int64          `json:"session_id"`
}

func (v *CreateSessionWorkoutReqModel) ToCreateSessionWorkoutParams() db.CreateSessionWorkoutParams {
	var status sql.NullString
	status.Valid = true
	status.String = v.Status

	return db.CreateSessionWorkoutParams{
		PresetSessionID: v.PresetSessionID,
		ActiveDay:       v.ActiveDay,
		WorkoutID:       v.WorkoutID,
		Status:          status,
		SessionID:       v.SessionID,
	}
}

func (v *CreateSessionWorkoutReqModel) ToUpdateSessionWorkoutParams(id int64) db.UpdateSessionWorkoutParams {
	var status sql.NullString
	status.Valid = true
	status.String = v.Status

	return db.UpdateSessionWorkoutParams{
		ID:              id,
		PresetSessionID: v.PresetSessionID,
		ActiveDay:       v.ActiveDay,
		WorkoutID:       v.WorkoutID,
		Status:          status,
		SessionID:       v.SessionID,
	}
}