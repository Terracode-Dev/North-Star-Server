-- name: CreateSession :exec
INSERT INTO V2Session (
  client_id,
  trainer_id
) VALUES (?, ?);

-- name: DeleteSession :exec
DELETE FROM V2Session
WHERE id = ?;

-- name: SelectSession :one
SELECT id, client_id, trainer_id
FROM V2Session
WHERE id = ?;

-- name: SelectAllSessions :many
SELECT id, client_id, trainer_id
FROM V2Session;
