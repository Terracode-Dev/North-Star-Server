-- name: CreateHrAdmin :exec
INSERT INTO HR_Admin (user_name, email, password, role, status, branch_id, created_by, updated_by
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: SelectHrAdmin :many
SELECT
  a.id,
  a.user_name,
  a.email,
  a.role,
  a.status,
  b.name AS branch_name,
  a.branch_id,
  a.created_at,
  a.updated_at
FROM HR_Admin a
LEFT JOIN HR_Branch b ON a.branch_id = b.id
WHERE 
  (
    a.user_name LIKE CONCAT('%', ?, '%')
    OR a.email    LIKE CONCAT('%', ?, '%')
    OR a.role     LIKE CONCAT('%', ?, '%')
    OR a.status   LIKE CONCAT('%', ?, '%')
  )
  AND ( ? = '' OR a.branch_id = ? )
ORDER BY a.id DESC
LIMIT ? OFFSET ?;

-- name: SelectOneHrAdmin :one
SELECT * FROM HR_Admin WHERE id = ?;

-- name: UpdateHrAdmin :exec
UPDATE HR_Admin SET user_name = ?, email = ?, role = ?, status = ?, branch_id = ? , updated_by = ? WHERE id = ?;

-- name: DeleteHrAdmin :exec
DELETE FROM HR_Admin WHERE id = ?;

-- name: SuspendedHrAdmin :exec
UPDATE HR_Admin SET status = ? WHERE id = ?;

-- name: AdminLogin :one
SELECT a.id, a.role, a.status, a.branch_id, a.password, b.name AS branchName
FROM HR_Admin a
LEFT JOIN HR_Branch b ON a.branch_id = b.id
WHERE email = ?;

-- name: AddHRBranch :exec
INSERT INTO HR_Branch (name) VALUES (?);

-- name: GetAllHRBranch :many
SELECT * FROM HR_Branch;

-- name: DeleteHrBranch :exec
DELETE FROM HR_Branch WHERE id = ?;

-- name: GetOneHrBranch :many
SELECT * FROM HR_Branch WHERE id = ?;
