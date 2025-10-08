-- name: CreateEmpLink :exec
INSERT INTO emp_link (emp_data, preset_id, is_approved, email, updated_by) VALUES (?, ?, ?, ?, ?);

-- name: GetEmpLinkByID :one
SELECT id, emp_data, preset_id, is_approved, create_date, email, updated_by FROM emp_link WHERE id = ?;

-- name: ListEmpLinks :many
SELECT id, emp_data, preset_id, is_approved, create_date, email, updated_by FROM emp_link LIMIT ? OFFSET ?;

-- name: TotalEmpLinksCount :one
SELECT COUNT(*) FROM emp_link;

-- name: UpdateEmpLinkApproval :exec
UPDATE emp_link SET is_approved = 1, updated_by = ? WHERE id = ?;

-- name: GetEmpLinkData :one
SELECT 
    el.id AS emp_link_id,
    el.emp_data,
    el.is_approved,
    el.create_date,
    el.updated_by,
    el.email,
    ap.id AS preset_id,
    ap.preset_name,
    ap.preset_value,
    ap.slug
FROM emp_link el
JOIN Admin_Presets ap ON el.preset_id = ap.id
WHERE el.id = ?;

-- name: DeleteEmpLink :exec
DELETE FROM emp_link WHERE id = ?;
