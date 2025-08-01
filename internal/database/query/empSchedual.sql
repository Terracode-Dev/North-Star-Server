-- name: GetEmployeeIdByEmail :one
SELECT employee_id FROM HR_EMP_User WHERE email = ?;

-- name: GetEmployeeByEmail :one
SELECT 
    e.id AS employee_id,
    e.first_name,
    e.last_name,
    s.department,
    s.designation
FROM HR_Employee e
INNER JOIN HR_EMP_Status s ON e.id = s.employee_id
WHERE e.email = ?
ORDER BY s.created_at DESC
LIMIT 1;

-- name: CreateEmployeeSchedule :exec
INSERT INTO HR_EMP_SCHEDUAL (
    emp_id, monday, monday_from, monday_to,
    tuesday, tuesday_from, tuesday_to,
    wednesday, wednesday_from, wednesday_to,
    thursday, thursday_from, thursday_to,
    friday, friday_from, friday_to,
    saturday, saturday_from, saturday_to,
    sunday, sunday_from, sunday_to
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: CreateAdditionalSchedule :exec
INSERT INTO HR_EMP_SCHEDUAL_additional (emp_id, date, from_time, to_time)
VALUES (?, ?, ?, ?);

-- name: UpdateEmployeeSchedule :exec
UPDATE HR_EMP_SCHEDUAL SET
    monday = ?, monday_from = ?, monday_to = ?,
    tuesday = ?, tuesday_from = ?, tuesday_to = ?,
    wednesday = ?, wednesday_from = ?, wednesday_to = ?,
    thursday = ?, thursday_from = ?, thursday_to = ?,
    friday = ?, friday_from = ?, friday_to = ?,
    saturday = ?, saturday_from = ?, saturday_to = ?,
    sunday = ?, sunday_from = ?, sunday_to = ?
WHERE emp_id = ?;

-- name: UpdateAdditionalSchedule :exec
UPDATE HR_EMP_SCHEDUAL_additional 
SET from_time = ?, to_time = ?
WHERE emp_id = ? AND date = ?;

-- name: DeleteEmployeeSchedule :exec
DELETE FROM HR_EMP_SCHEDUAL WHERE emp_id = ?;

-- name: DeleteAdditionalSchedule :exec
DELETE FROM HR_EMP_SCHEDUAL_additional WHERE emp_id = ? AND date = ?;

-- name: DeleteAllAdditionalSchedules :exec
DELETE FROM HR_EMP_SCHEDUAL_additional WHERE emp_id = ?;

-- name: GetEmployeeListWithWorkDays :many
SELECT 
    e.id,
    e.first_name,
    e.last_name,
    u.email,
    -- Calculate base working days from weekly schedule (52 weeks)
    COALESCE(
        (CASE WHEN s.monday THEN 1 ELSE 0 END +
         CASE WHEN s.tuesday THEN 1 ELSE 0 END +
         CASE WHEN s.wednesday THEN 1 ELSE 0 END +
         CASE WHEN s.thursday THEN 1 ELSE 0 END +
         CASE WHEN s.friday THEN 1 ELSE 0 END +
         CASE WHEN s.saturday THEN 1 ELSE 0 END +
         CASE WHEN s.sunday THEN 1 ELSE 0 END) * 52, 0
    ) +
    -- Add additional working days for the specified year
    COALESCE(
        (SELECT COUNT(*) 
         FROM HR_EMP_SCHEDUAL_additional a 
         WHERE a.emp_id = e.id 
         AND YEAR(a.date) = ?
         AND a.from_time IS NOT NULL 
         AND a.to_time IS NOT NULL), 0
    ) as work_days_for_year
FROM HR_Employee e
LEFT JOIN HR_EMP_User u ON e.id = u.employee_id
LEFT JOIN HR_EMP_SCHEDUAL s ON e.id = s.emp_id
WHERE 
    (? = '' OR e.first_name LIKE CONCAT('%', ?, '%'))
    AND (? = '' OR e.last_name LIKE CONCAT('%', ?, '%'))
    AND (? = '' OR u.email LIKE CONCAT('%', ?, '%'))
ORDER BY 
    CASE WHEN ? = 'first_name' THEN e.first_name END ASC,
    CASE WHEN ? = 'last_name' THEN e.last_name END ASC,
    CASE WHEN ? = 'email' THEN u.email END ASC,
    CASE WHEN ? = 'work_days' THEN work_days_for_year END ASC,
    e.id ASC
LIMIT ? OFFSET ?;

-- Count total employees for pagination (with same filters)
-- name: CountEmployeesWithFilters :one
SELECT COUNT(*) as total
FROM HR_Employee e
LEFT JOIN HR_EMP_User u ON e.id = u.employee_id
WHERE 
    (? = '' OR e.first_name LIKE CONCAT('%', ?, '%'))
    AND (? = '' OR e.last_name LIKE CONCAT('%', ?, '%'))
    AND (? = '' OR u.email LIKE CONCAT('%', ?, '%'));

-- name: GetEmployeeWorkDaysBreakdown :one
SELECT 
    e.id,
    e.first_name,
    e.last_name,
    u.email,
    -- Weekly working days count
    COALESCE(
        (CASE WHEN s.monday THEN 1 ELSE 0 END +
         CASE WHEN s.tuesday THEN 1 ELSE 0 END +
         CASE WHEN s.wednesday THEN 1 ELSE 0 END +
         CASE WHEN s.thursday THEN 1 ELSE 0 END +
         CASE WHEN s.friday THEN 1 ELSE 0 END +
         CASE WHEN s.saturday THEN 1 ELSE 0 END +
         CASE WHEN s.sunday THEN 1 ELSE 0 END), 0
    ) as weekly_work_days,
    -- Base yearly working days (weekly * 52)
    COALESCE(
        (CASE WHEN s.monday THEN 1 ELSE 0 END +
         CASE WHEN s.tuesday THEN 1 ELSE 0 END +
         CASE WHEN s.wednesday THEN 1 ELSE 0 END +
         CASE WHEN s.thursday THEN 1 ELSE 0 END +
         CASE WHEN s.friday THEN 1 ELSE 0 END +
         CASE WHEN s.saturday THEN 1 ELSE 0 END +
         CASE WHEN s.sunday THEN 1 ELSE 0 END) * 52, 0
    ) as base_yearly_days,
    -- Additional working days for the year
    COALESCE(
        (SELECT COUNT(*) 
         FROM HR_EMP_SCHEDUAL_additional a 
         WHERE a.emp_id = e.id 
         AND YEAR(a.date) = ?
         AND a.from_time IS NOT NULL 
         AND a.to_time IS NOT NULL), 0
    ) as additional_days,
    -- Total work days for year
    COALESCE(
        (CASE WHEN s.monday THEN 1 ELSE 0 END +
         CASE WHEN s.tuesday THEN 1 ELSE 0 END +
         CASE WHEN s.wednesday THEN 1 ELSE 0 END +
         CASE WHEN s.thursday THEN 1 ELSE 0 END +
         CASE WHEN s.friday THEN 1 ELSE 0 END +
         CASE WHEN s.saturday THEN 1 ELSE 0 END +
         CASE WHEN s.sunday THEN 1 ELSE 0 END) * 52, 0
    ) +
    COALESCE(
        (SELECT COUNT(*) 
         FROM HR_EMP_SCHEDUAL_additional a 
         WHERE a.emp_id = e.id 
         AND YEAR(a.date) = ?
         AND a.from_time IS NOT NULL 
         AND a.to_time IS NOT NULL), 0
    ) as total_work_days_for_year
FROM HR_Employee e
LEFT JOIN HR_EMP_User u ON e.id = u.employee_id
LEFT JOIN HR_EMP_SCHEDUAL s ON e.id = s.emp_id
WHERE e.id = ?;

-- name: GetEmpShedulleByID :one
SELECT 
    monday, monday_from, monday_to,
    tuesday, tuesday_from, tuesday_to,
    wednesday, wednesday_from, wednesday_to,
    thursday, thursday_from, thursday_to,
    friday, friday_from, friday_to,
    saturday, saturday_from, saturday_to,
    sunday, sunday_from, sunday_to
FROM HR_EMP_SCHEDUAL
WHERE emp_id = ?;

-- name: GetEmpAdditionalSheduleByID :many
SELECT
    date, from_time, to_time, created_at, updated_at
FROM HR_EMP_SCHEDUAL_additional 
WHERE emp_id = ?;