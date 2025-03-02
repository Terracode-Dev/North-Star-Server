-- name: CreateServices :exec
INSERT INTO HR_Create_Services (
    category, value, updated_by
) VALUES (
    ?, ?, ?
);

-- name: GetServices :many
SELECT * FROM HR_Create_Services;

-- name: GetService :many
SELECT 
    s.id, 
    s.category, 
    s.value, 
    a.user_name AS updated_by,
    s.created_at, 
    s.updated_at
FROM HR_Create_Services s
LEFT JOIN HR_Admin a ON s.updated_by = a.id
WHERE s.category = ?;


-- name: UpdateService :exec
UPDATE HR_Create_Services
SET category = ?, value = ?, updated_by = ?
WHERE id = ?;

-- name: DeleteService :exec
DELETE FROM HR_Create_Services
WHERE id = ?;