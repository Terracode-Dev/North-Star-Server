-- name: CreateServices :exec
INSERT INTO HR_Create_Services (
    category, value, updated_by
) VALUES (
    ?, ?, ?
);

-- name: GetServices :many
SELECT * FROM HR_Create_Services;

-- name: GetService :one
SELECT * FROM HR_Create_Services
WHERE category = ?;

-- name: UpdateService :exec
UPDATE HR_Create_Services
SET category = ?, value = ?, updated_by = ?
WHERE id = ?;

-- name: DeleteService :exec
DELETE FROM HR_Create_Services
WHERE id = ?;