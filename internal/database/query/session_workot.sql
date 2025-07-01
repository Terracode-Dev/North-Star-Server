-- name: CreateSessionWorkout :exec
INSERT INTO Session_Workout (
  preset_session_id,
  active_day,
  workout_id,
  status,
  session_id
) VALUES (?, ?, ?, ?, ?);

-- name: SelectSessionWorkout :one
SELECT id, preset_session_id, active_day, workout_id, status, session_id
FROM Session_Workout
WHERE id = ?;

-- name: SelectAllSessionWorkouts :many
SELECT id, preset_session_id, active_day, workout_id, status, session_id
FROM Session_Workout;

-- name: UpdateSessionWorkout :exec
UPDATE Session_Workout
SET
  preset_session_id = ?,
  active_day = ?,
  workout_id = ?,
  status = ?,
  session_id = ?
WHERE id = ?;

-- name: DeleteSessionWorkout :exec
DELETE FROM Session_Workout
WHERE id = ?;