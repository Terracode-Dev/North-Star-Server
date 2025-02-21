-- name: CreateHrAdmin :exec
INSERT INTO HR_Admin (user_name, email, password, role, status, branch_id, created_by, updated_by
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: SelectHrAdmin :many
SELECT * FROM HR_Admin;

-- name: SelectOneHrAdmin :one
SELECT * FROM HR_Admin WHERE id = ?;

-- name: UpdateHrAdmin :exec
UPDATE HR_Admin SET user_name = ?, email = ?, password = ?, role = ?, status = ?, branch_id = ? WHERE id = ?;

-- name: DeleteHrAdmin :exec
DELETE FROM HR_Admin WHERE id = ?;

-- name: SuspendedHrAdmin :exec
UPDATE HR_Admin SET status = ? WHERE id = ?;

-- name: AdminLogin :one
SELECT id FROM HR_Admin WHERE email = ? AND password = ?;