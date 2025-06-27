-- name: CreatePreset :exec
INSERT INTO V2Presets (name, description, trainer_id)
VALUES (?, ?, ?);

-- name: SelectPreset :one
SELECT name, description, trainer_id 
FROM V2Presets 
WHERE id = ?;

-- name: SelectAllPresets :many
SELECT name, description, trainer_id 
FROM V2Presets; 

-- name: DeletePreset :exec
DELETE FROM V2Presets WHERE id = ?;

-- name: UpdatePreset :exec 
UPDATE V2Presets
SET
  name = ?,
  description = ?,
  trainer_id = ?
WHERE id = ?;

-- name: SelectPresetByname :many
SELECT id, name, description, trainer_id
FROM V2Presets
WHERE name = ?;