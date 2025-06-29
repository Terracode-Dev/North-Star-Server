-- name: GetEmployeeLeaveBenefits :many
SELECT 
    id,
    leave_status,
    leave_type,
    leave_count,
    employee_id
FROM HR_EMP_Benifits 
WHERE employee_id = ? AND leave_status = 1;

-- name: CheckLeaveCountForYear :one
SELECT 
    COALESCE(COUNT(*), 0) as used_leaves,
    COALESCE(b.leave_count, 0) as total_allowed
FROM HR_EMP_LEAVES l
RIGHT JOIN HR_EMP_Benifits b ON l.emp_id = b.employee_id 
    AND l.leave_type = b.leave_type
WHERE b.employee_id = ? 
    AND b.leave_type = ? 
    AND b.leave_status = 1
    AND (l.leave_date IS NULL OR YEAR(l.leave_date) = YEAR(CURDATE()));

-- name: CreateLeave :exec
INSERT INTO HR_EMP_LEAVES (emp_id, leave_type, leave_date, reason, added_by)
VALUES (?, ?, ?, ?, ?);

-- name: UpdateLeave :exec
UPDATE HR_EMP_LEAVES 
SET 
    leave_type = ?,
    leave_date = ?,
    reason = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ? AND emp_id = ?;

-- name: DeleteLeave :exec
DELETE FROM HR_EMP_LEAVES 
WHERE id = ? AND emp_id = ?;

-- name: GetEmployeeLeaves :many
SELECT 
    l.id,
    l.emp_id,
    l.leave_type,
    l.leave_date,
    l.reason,
    l.create_date,
    COUNT(*) OVER() as total_count
FROM HR_EMP_LEAVES l
WHERE l.emp_id = ?
ORDER BY l.leave_date DESC
LIMIT ? OFFSET ?;

-- name: GetLeaveSummaryByEmployee :many
SELECT 
    l.emp_id,
    l.leave_type,
    COUNT(*) as used_count,
    b.leave_count as allowed_count,
    (b.leave_count - COUNT(*)) as remaining_count,
    e.first_name,
    e.last_name
FROM HR_EMP_LEAVES l
JOIN HR_Employee e ON l.emp_id = e.id
JOIN HR_EMP_Benifits b ON l.emp_id = b.employee_id AND l.leave_type = b.leave_type
WHERE 
    YEAR(l.leave_date) = YEAR(CURDATE()) AND
    b.leave_status = 1 AND
    (? = 0 OR l.emp_id = ?)
GROUP BY l.emp_id, l.leave_type, b.leave_count, e.first_name, e.last_name
ORDER BY e.first_name, e.last_name, l.leave_type;

-- name: GetLeaveTypesForEmployee :many
SELECT 
    leave_type,
    leave_count,
    leave_status
FROM HR_EMP_Benifits 
WHERE employee_id = ? AND leave_status = 1;
