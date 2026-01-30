-- name: CreateLeave :execresult
INSERT INTO HR_EMP_LEAVES (emp_id, leave_type, leave_date, reason, added_by)
VALUES (?, ?, ?, ?, ?);

-- name: DeleteLeave :exec
DELETE FROM HR_EMP_LEAVES 
WHERE id = ?;

-- name: DeleteLeaveByEmpAndDate :exec
DELETE FROM HR_EMP_LEAVES 
WHERE emp_id = ? AND leave_date = ?;

-- name: UpdateLeave :exec
UPDATE HR_EMP_LEAVES 
SET 
    leave_type = ?,
    leave_date = ?,
    reason = ?,
    added_by = ?
WHERE id = ?;

-- name: GetLeaveById :one
SELECT 
    el.id,
    el.emp_id,
    CONCAT(e.first_name, ' ', e.last_name) AS employee_name,
    e.email AS employee_email,
    el.leave_type,
    el.leave_date,
    el.reason,
    el.create_date,
    a.user_name AS added_by_name
FROM HR_EMP_LEAVES el
INNER JOIN HR_Employee e ON el.emp_id = e.id
LEFT JOIN HR_Admin a ON el.added_by = a.id
WHERE el.id = ? AND (el.is_ban = false OR el.is_ban IS NULL);

-- name: GetEmployeeLeaveBenefits :one
SELECT 
    eb.leave_type,
    eb.leave_count,
    eb.leave_status
FROM HR_EMP_Benifits eb
WHERE eb.employee_id = ? AND eb.leave_status = 1
ORDER BY eb.leave_type;

-- name: GetEmployeeApprovedLeaveCount :one
SELECT 
    COUNT(*) AS approved_leave_count
FROM HR_EMP_LEAVES
WHERE emp_id = ? AND leave_type = ?;

-- name: GetEmployeeLeaves :many
SELECT 
    el.id as leave_id,
    el.emp_id,
    el.leave_type,
    el.leave_date,
    el.reason,
    el.create_date
FROM HR_EMP_LEAVES el
WHERE el.emp_id = ?
    AND (el.is_ban = false OR el.is_ban IS NULL)
    AND (? = '' OR el.leave_type LIKE CONCAT('%', ?, '%'))
    AND (? IS NULL OR YEAR(el.leave_date) = ?)
ORDER BY 
    CASE WHEN ? = 'date_asc' THEN el.leave_date END ASC,
    CASE WHEN ? = 'date_desc' THEN el.leave_date END DESC,
    CASE WHEN ? = 'type_asc' THEN el.leave_type END ASC,
    CASE WHEN ? = 'type_desc' THEN el.leave_type END DESC,
    el.leave_date DESC
LIMIT ? OFFSET ?;

-- name: GetAllLeaves :many
SELECT 
    el.id as leave_id,
    el.emp_id,
    e.first_name,
    e.last_name,
    e.email,
    el.leave_type,
    el.leave_date,
    el.create_date,
    el.reason,
    ha.user_name as added_by_name
FROM HR_EMP_LEAVES el
INNER JOIN HR_Employee e ON el.emp_id = e.id
LEFT JOIN HR_Admin ha ON el.added_by = ha.id
WHERE 
    (? = '' OR e.first_name LIKE CONCAT('%', ?, '%') OR e.last_name LIKE CONCAT('%', ?, '%'))
    AND (? = '' OR e.email LIKE CONCAT('%', ?, '%'))
    AND (? = '' OR el.leave_type = ?)
    AND (? IS NULL OR el.leave_date >= ?)
    AND (? IS NULL OR el.leave_date <= ?)
    AND (e.is_ban = false OR e.is_ban IS NULL)
ORDER BY 
    CASE WHEN ? = 'name_asc' THEN e.first_name END ASC,
    CASE WHEN ? = 'name_desc' THEN e.first_name END DESC,
    CASE WHEN ? = 'email_asc' THEN e.email END ASC,
    CASE WHEN ? = 'email_desc' THEN e.email END DESC,
    CASE WHEN ? = 'leave_type_asc' THEN el.leave_type END ASC,
    CASE WHEN ? = 'leave_type_desc' THEN el.leave_type END DESC,
    CASE WHEN ? = 'date_asc' THEN el.leave_date END ASC,
    CASE WHEN ? = 'date_desc' THEN el.leave_date END DESC,
    el.leave_date DESC
LIMIT ? OFFSET ?;

-- name: GetEmployeeLeavesCount :one
SELECT COUNT(*) as total_count
FROM HR_EMP_LEAVES el
WHERE el.emp_id = ?
    AND (? = '' OR el.leave_type LIKE CONCAT('%', ?, '%'))
    AND (? IS NULL OR el.leave_date >= ?)
    AND (? IS NULL OR el.leave_date <= ?);


-- name: UpdateLeaveEMP :exec
UPDATE HR_EMP_LEAVES 
SET 
    leave_date = ?,
    reason = ?
WHERE id = ?;

-- name: GetEmployeeLeavesEMP :many
SELECT 
    el.id as leave_id,
    el.emp_id,
    el.leave_type,
    el.leave_date,
    el.reason,
    el.create_date
FROM HR_EMP_LEAVES el
WHERE el.emp_id = ?
    AND (? = '' OR el.leave_type LIKE CONCAT('%', ?, '%'))
    AND (? IS NULL OR YEAR(el.leave_date) = ?)
ORDER BY 
    CASE WHEN ? = 'date_asc' THEN el.leave_date END ASC,
    CASE WHEN ? = 'date_desc' THEN el.leave_date END DESC,
    CASE WHEN ? = 'type_asc' THEN el.leave_type END ASC,
    CASE WHEN ? = 'type_desc' THEN el.leave_type END DESC,
    el.leave_date DESC
LIMIT ? OFFSET ?
