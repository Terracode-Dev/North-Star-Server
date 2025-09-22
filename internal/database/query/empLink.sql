-- name: CreateEmpLink :exec
INSERT INTO emp_link (emp_data, preset_id, is_approved, email, updated_by) VALUES (?, ?, ?, ?, ?);

-- name: GetEmpLinkByID :one
SELECT id, emp_data, preset_id, is_approved, create_date, email, updated_by FROM emp_link WHERE id = ?;

-- name: ListEmpLinks :many
SELECT id, emp_data, preset_id, is_approved, create_date, email, updated_by FROM emp_link;

-- name: UpdateEmpLinkApproval :exec
UPDATE emp_link SET is_approved = ?, updated_by = ? WHERE id = ?;