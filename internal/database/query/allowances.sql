-- name: CreateAllowances :exec
INSERT INTO HR_Create_Allowances (
    allowance_type, amount, updated_by
) VALUES (
    ?, ?, ?
);

-- name: GetAllowances :many
SELECT h.id, h.allowance_type, h.amount, h.created_at, h.updated_at, a.user_name as updated_by
FROM HR_Create_Allowances h
LEFT JOIN HR_Admin a ON a.id = h.updated_by;

-- name: GetAllowance :one
SELECT * FROM HR_Create_Allowances
WHERE id = ?;

-- name: UpdateAllowance :exec
UPDATE HR_Create_Allowances
SET allowance_type = ?, amount = ?, updated_by = ?
WHERE id = ?;

-- name: DeleteAllowance :exec
DELETE FROM HR_Create_Allowances
WHERE id = ?;

