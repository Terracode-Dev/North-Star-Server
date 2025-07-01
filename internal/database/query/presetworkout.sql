-- name: CreatePresetWorkout :exec
INSERT INTO V2Preset_Workouts (
  preset_id,
  workout_id,
  reps,
  weight,
  sets,
  notes
)
VALUES (?, ?, ?, ?, ?, ?);

-- name: SelectAllPresetWorkouts :many
SELECT preset_id, workout_id, reps, weight, sets, notes 
FROM V2Preset_Workouts;

-- name: UpdatePresetWorkout :exec
UPDATE V2Preset_Workouts
SET
  preset_id = ?,
  workout_id = ?,
  reps = ?,
  weight = ?,
  sets = ?,
  notes = ?
WHERE id = ?;

-- name: DeletePresetWorkout :exec
DELETE FROM V2Preset_Workouts
WHERE id = ?;