-- name: CreatePresetSession :exec
INSERT INTO V2Preset_Session (
  preset_id,
  client_id,
  assign_session,
  active_day,
  max_day,
  state
) VALUES (?, ?, ?, ?, ?, ?);

-- name: SelectpresetSessionAll :many
SELECT preset_id, client_id, assign_session, active_day, max_day, state
FROM V2Preset_Session;

-- name: UpdatePresetSession :exec
UPDATE V2Preset_Session
SET
  preset_id = ?,
  client_id = ?,
  assign_session = ?,
  active_day = ?,
  max_day = ?,
  state = ?
WHERE id = ?;

-- name: DeletePresetSession :exec
DELETE FROM V2Preset_Session
WHERE id = ?;


