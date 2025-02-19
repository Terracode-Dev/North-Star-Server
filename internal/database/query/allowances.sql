-- name: CreateAllowances :exec
INSERT INTO HR_Create_Allowances (
    allowance_type, amount, updated_by
) VALUES (
    ?, ?, ?
);

-- name: GetAllowances :many
SELECT * FROM HR_Create_Allowances;

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

