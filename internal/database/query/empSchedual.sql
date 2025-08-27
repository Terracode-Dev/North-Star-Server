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
AND (e.is_ban = false OR e.is_ban IS NULL)
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
    (e.is_ban = false OR e.is_ban IS NULL)
    AND(? = '' OR e.first_name LIKE CONCAT('%', ?, '%'))
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
    id,
    emp_id,
    monday,
    COALESCE(TIME_FORMAT(monday_from, '%H:%i:%s'), '') as monday_from,
    COALESCE(TIME_FORMAT(monday_to, '%H:%i:%s'), '') as monday_to,
    tuesday,
    COALESCE(TIME_FORMAT(tuesday_from, '%H:%i:%s'), '') as tuesday_from,
    COALESCE(TIME_FORMAT(tuesday_to, '%H:%i:%s'), '') as tuesday_to,
    wednesday,
    COALESCE(TIME_FORMAT(wednesday_from, '%H:%i:%s'), '') as wednesday_from,
    COALESCE(TIME_FORMAT(wednesday_to, '%H:%i:%s'), '') as wednesday_to,
    thursday,
    COALESCE(TIME_FORMAT(thursday_from, '%H:%i:%s'), '') as thursday_from,
    COALESCE(TIME_FORMAT(thursday_to, '%H:%i:%s'), '') as thursday_to,
    friday,
    COALESCE(TIME_FORMAT(friday_from, '%H:%i:%s'), '') as friday_from,
    COALESCE(TIME_FORMAT(friday_to, '%H:%i:%s'), '') as friday_to,
    saturday,
    COALESCE(TIME_FORMAT(saturday_from, '%H:%i:%s'), '') as saturday_from,
    COALESCE(TIME_FORMAT(saturday_to, '%H:%i:%s'), '') as saturday_to,
    sunday,
    COALESCE(TIME_FORMAT(sunday_from, '%H:%i:%s'), '') as sunday_from,
    COALESCE(TIME_FORMAT(sunday_to, '%H:%i:%s'), '') as sunday_to,
    created_at,
    updated_at
FROM HR_EMP_SCHEDUAL
WHERE emp_id = ?;

-- name: GetEmpAdditionalSheduleByID :many
SELECT
    id,
    emp_id,
    date, 
    COALESCE(TIME_FORMAT(from_time, '%H:%i:%s'), '') as from_time, 
    COALESCE(TIME_FORMAT(to_time, '%H:%i:%s'), '') as to_time, 
    created_at, 
    updated_at
FROM HR_EMP_SCHEDUAL_additional 
WHERE emp_id = ?
ORDER BY date DESC;

-- name: GetInsufficientAttendance :many
SELECT *
FROM (
  SELECT
    DATE(a.create_date) as date,
    a.emp_id,
    TIME_FORMAT(MIN(CASE WHEN a.attendance_type = 'in' THEN TIME(a.create_date) END), '%H:%i:%s') as in_time,
    TIME_FORMAT(MAX(CASE WHEN a.attendance_type = 'out' THEN TIME(a.create_date) END), '%H:%i:%s') as out_time,
    CASE 
      WHEN MIN(CASE WHEN a.attendance_type = 'in' THEN TIME(a.create_date) END) IS NOT NULL 
           AND MAX(CASE WHEN a.attendance_type = 'out' THEN TIME(a.create_date) END) IS NOT NULL 
      THEN TIME_FORMAT(TIMEDIFF(
        MAX(CASE WHEN a.attendance_type = 'out' THEN TIME(a.create_date) END),
        MIN(CASE WHEN a.attendance_type = 'in' THEN TIME(a.create_date) END)
      ), '%H:%i:%s')
      ELSE NULL 
    END as total_time,
    TIME_FORMAT(TIMEDIFF(
      COALESCE(sa.to_time,
        CASE DAYOFWEEK(DATE(a.create_date))
          WHEN 1 THEN s.sunday_to
          WHEN 2 THEN s.monday_to
          WHEN 3 THEN s.tuesday_to
          WHEN 4 THEN s.wednesday_to
          WHEN 5 THEN s.thursday_to
          WHEN 6 THEN s.friday_to
          WHEN 7 THEN s.saturday_to
        END), 
      COALESCE(sa.from_time, 
        CASE DAYOFWEEK(DATE(a.create_date))
          WHEN 1 THEN s.sunday_from
          WHEN 2 THEN s.monday_from  
          WHEN 3 THEN s.tuesday_from
          WHEN 4 THEN s.wednesday_from
          WHEN 5 THEN s.thursday_from
          WHEN 6 THEN s.friday_from
          WHEN 7 THEN s.saturday_from
        END)
    ), '%H:%i:%s') as scheduled_time
  FROM HR_EMP_ATTENDANCE a
  LEFT JOIN HR_EMP_SCHEDUAL s ON a.emp_id = s.emp_id
  LEFT JOIN HR_EMP_SCHEDUAL_additional sa 
    ON a.emp_id = sa.emp_id 
    AND DATE(a.create_date) = sa.date
  WHERE 
    a.emp_id = ?
    AND (? IS NULL OR DATE(a.create_date) = ?)
  GROUP BY DATE(a.create_date), a.emp_id
) t
WHERE t.total_time < t.scheduled_time
ORDER BY t.date DESC
LIMIT ? OFFSET ?;

-- name: GetLateAttendance :many
SELECT *
FROM (
  SELECT
    DATE(a.create_date) as date,
    a.emp_id,
    TIME_FORMAT(MIN(CASE WHEN a.attendance_type = 'in' THEN TIME(a.create_date) END), '%H:%i:%s') as in_time,
    TIME_FORMAT(MAX(CASE WHEN a.attendance_type = 'out' THEN TIME(a.create_date) END), '%H:%i:%s') as out_time,
    TIMEDIFF(
      MAX(CASE WHEN a.attendance_type = 'out' THEN TIME(a.create_date) END),
      MIN(CASE WHEN a.attendance_type = 'in' THEN TIME(a.create_date) END)
    ) as total_time,
    MIN(CASE WHEN a.attendance_type = 'in' THEN TIME(a.create_date) END) as first_in_time,
    COALESCE(sa.from_time, 
      CASE DAYOFWEEK(DATE(a.create_date))
        WHEN 1 THEN s.sunday_from
        WHEN 2 THEN s.monday_from  
        WHEN 3 THEN s.tuesday_from
        WHEN 4 THEN s.wednesday_from
        WHEN 5 THEN s.thursday_from
        WHEN 6 THEN s.friday_from
        WHEN 7 THEN s.saturday_from
      END
    ) as scheduled_in_time
  FROM HR_EMP_ATTENDANCE a
  LEFT JOIN HR_EMP_SCHEDUAL s ON a.emp_id = s.emp_id
  LEFT JOIN HR_EMP_SCHEDUAL_additional sa 
    ON a.emp_id = sa.emp_id 
    AND DATE(a.create_date) = sa.date
  WHERE a.emp_id = ?  -- specific employee
    AND (? IS NULL OR DATE(a.create_date) = ?) -- optional date filter
  GROUP BY DATE(a.create_date), a.emp_id
) t
WHERE t.first_in_time > t.scheduled_in_time
ORDER BY t.date DESC
LIMIT ? OFFSET ?;

-- name: GetNormalAttendance :many
SELECT *
FROM (
  SELECT
    DATE(a.create_date) as date,
    a.emp_id,
    TIME_FORMAT(MIN(CASE WHEN a.attendance_type = 'in' THEN TIME(a.create_date) END), '%H:%i:%s') as in_time,
    TIME_FORMAT(MAX(CASE WHEN a.attendance_type = 'out' THEN TIME(a.create_date) END), '%H:%i:%s') as out_time,
    CASE 
      WHEN MIN(CASE WHEN a.attendance_type = 'in' THEN TIME(a.create_date) END) IS NOT NULL 
           AND MAX(CASE WHEN a.attendance_type = 'out' THEN TIME(a.create_date) END) IS NOT NULL 
      THEN TIME_FORMAT(TIMEDIFF(
        MAX(CASE WHEN a.attendance_type = 'out' THEN TIME(a.create_date) END),
        MIN(CASE WHEN a.attendance_type = 'in' THEN TIME(a.create_date) END)
      ), '%H:%i:%s')
      ELSE NULL 
    END as total_time,
    TIME_FORMAT(TIMEDIFF(
      COALESCE(sa.to_time,
        CASE DAYOFWEEK(DATE(a.create_date))
          WHEN 1 THEN s.sunday_to
          WHEN 2 THEN s.monday_to
          WHEN 3 THEN s.tuesday_to
          WHEN 4 THEN s.wednesday_to
          WHEN 5 THEN s.thursday_to
          WHEN 6 THEN s.friday_to
          WHEN 7 THEN s.saturday_to
        END), 
      COALESCE(sa.from_time, 
        CASE DAYOFWEEK(DATE(a.create_date))
          WHEN 1 THEN s.sunday_from
          WHEN 2 THEN s.monday_from  
          WHEN 3 THEN s.tuesday_from
          WHEN 4 THEN s.wednesday_from
          WHEN 5 THEN s.thursday_from
          WHEN 6 THEN s.friday_from
          WHEN 7 THEN s.saturday_from
        END)
    ), '%H:%i:%s') as scheduled_time
  FROM HR_EMP_ATTENDANCE a
  LEFT JOIN HR_EMP_SCHEDUAL s ON a.emp_id = s.emp_id
  LEFT JOIN HR_EMP_SCHEDUAL_additional sa 
    ON a.emp_id = sa.emp_id 
    AND DATE(a.create_date) = sa.date
  WHERE 
    a.emp_id = ?
    AND (? IS NULL OR DATE(a.create_date) = ?)
  GROUP BY DATE(a.create_date), a.emp_id
) t
WHERE t.total_time = t.scheduled_time
ORDER BY t.date DESC
LIMIT ? OFFSET ?;


-- name: GetAllAttendance :many
SELECT
  DATE(create_date) as date,
  emp_id,
  TIME_FORMAT(MIN(CASE WHEN attendance_type = 'in' THEN TIME(create_date) END), '%H:%i:%s') as in_time,
  TIME_FORMAT(MAX(CASE WHEN attendance_type = 'out' THEN TIME(create_date) END), '%H:%i:%s') as out_time,
  CASE 
    WHEN MIN(CASE WHEN attendance_type = 'in' THEN TIME(create_date) END) IS NOT NULL 
         AND MAX(CASE WHEN attendance_type = 'out' THEN TIME(create_date) END) IS NOT NULL 
    THEN TIME_FORMAT(TIMEDIFF(
      MAX(CASE WHEN attendance_type = 'out' THEN TIME(create_date) END),
      MIN(CASE WHEN attendance_type = 'in' THEN TIME(create_date) END)
    ), '%H:%i:%s')
    ELSE NULL 
  END as total_time
FROM HR_EMP_ATTENDANCE
WHERE emp_id = ?
  AND (? IS NULL OR DATE(create_date) = ?)
GROUP BY DATE(create_date), emp_id
ORDER BY DATE(create_date) DESC
LIMIT ? OFFSET ?;

-- name: GetAllAttendanceForAll :many
SELECT
  DATE(create_date) as date,
  emp_id,
  TIME_FORMAT(MIN(CASE WHEN attendance_type = 'in' THEN TIME(create_date) END), '%H:%i:%s') as in_time,
  TIME_FORMAT(MAX(CASE WHEN attendance_type = 'out' THEN TIME(create_date) END), '%H:%i:%s') as out_time,
  CASE 
    WHEN MIN(CASE WHEN attendance_type = 'in' THEN TIME(create_date) END) IS NOT NULL 
         AND MAX(CASE WHEN attendance_type = 'out' THEN TIME(create_date) END) IS NOT NULL 
    THEN TIME_FORMAT(TIMEDIFF(
      MAX(CASE WHEN attendance_type = 'out' THEN TIME(create_date) END),
      MIN(CASE WHEN attendance_type = 'in' THEN TIME(create_date) END)
    ), '%H:%i:%s')
    ELSE NULL 
  END as total_time
FROM HR_EMP_ATTENDANCE
WHERE (? IS NULL OR DATE(create_date) = ?)
GROUP BY DATE(create_date), emp_id
ORDER BY DATE(create_date) DESC
LIMIT ? OFFSET ?;

-- name: GetNormalAttendanceForAll :many
SELECT *
FROM (
  SELECT
    DATE(a.create_date) as date,
    a.emp_id,
    TIME_FORMAT(MIN(CASE WHEN a.attendance_type = 'in' THEN TIME(a.create_date) END), '%H:%i:%s') as in_time,
    TIME_FORMAT(MAX(CASE WHEN a.attendance_type = 'out' THEN TIME(a.create_date) END), '%H:%i:%s') as out_time,
    CASE 
      WHEN MIN(CASE WHEN a.attendance_type = 'in' THEN TIME(a.create_date) END) IS NOT NULL 
           AND MAX(CASE WHEN a.attendance_type = 'out' THEN TIME(a.create_date) END) IS NOT NULL 
      THEN TIME_FORMAT(TIMEDIFF(
        MAX(CASE WHEN a.attendance_type = 'out' THEN TIME(a.create_date) END),
        MIN(CASE WHEN a.attendance_type = 'in' THEN TIME(a.create_date) END)
      ), '%H:%i:%s')
      ELSE NULL 
    END as total_time,
    TIME_FORMAT(TIMEDIFF(
      COALESCE(sa.to_time,
        CASE DAYOFWEEK(DATE(a.create_date))
          WHEN 1 THEN s.sunday_to
          WHEN 2 THEN s.monday_to
          WHEN 3 THEN s.tuesday_to
          WHEN 4 THEN s.wednesday_to
          WHEN 5 THEN s.thursday_to
          WHEN 6 THEN s.friday_to
          WHEN 7 THEN s.saturday_to
        END), 
      COALESCE(sa.from_time, 
        CASE DAYOFWEEK(DATE(a.create_date))
          WHEN 1 THEN s.sunday_from
          WHEN 2 THEN s.monday_from  
          WHEN 3 THEN s.tuesday_from
          WHEN 4 THEN s.wednesday_from
          WHEN 5 THEN s.thursday_from
          WHEN 6 THEN s.friday_from
          WHEN 7 THEN s.saturday_from
        END)
    ), '%H:%i:%s') as scheduled_time
  FROM HR_EMP_ATTENDANCE a
  LEFT JOIN HR_EMP_SCHEDUAL s ON a.emp_id = s.emp_id
  LEFT JOIN HR_EMP_SCHEDUAL_additional sa 
    ON a.emp_id = sa.emp_id 
    AND DATE(a.create_date) = sa.date
  WHERE  (? IS NULL OR DATE(a.create_date) = ?)
  GROUP BY DATE(a.create_date), a.emp_id
) t
WHERE t.total_time = t.scheduled_time
ORDER BY t.date DESC
LIMIT ? OFFSET ?;


-- name: GetLateAttendanceForAll :many
SELECT *
FROM (
  SELECT
    DATE(a.create_date) as date,
    a.emp_id,
    TIME_FORMAT(MIN(CASE WHEN a.attendance_type = 'in' THEN TIME(a.create_date) END), '%H:%i:%s') as in_time,
    TIME_FORMAT(MAX(CASE WHEN a.attendance_type = 'out' THEN TIME(a.create_date) END), '%H:%i:%s') as out_time,
    TIMEDIFF(
      MAX(CASE WHEN a.attendance_type = 'out' THEN TIME(a.create_date) END),
      MIN(CASE WHEN a.attendance_type = 'in' THEN TIME(a.create_date) END)
    ) as total_time,
    MIN(CASE WHEN a.attendance_type = 'in' THEN TIME(a.create_date) END) as first_in_time,
    COALESCE(sa.from_time, 
      CASE DAYOFWEEK(DATE(a.create_date))
        WHEN 1 THEN s.sunday_from
        WHEN 2 THEN s.monday_from  
        WHEN 3 THEN s.tuesday_from
        WHEN 4 THEN s.wednesday_from
        WHEN 5 THEN s.thursday_from
        WHEN 6 THEN s.friday_from
        WHEN 7 THEN s.saturday_from
      END
    ) as scheduled_in_time
  FROM HR_EMP_ATTENDANCE a
  LEFT JOIN HR_EMP_SCHEDUAL s ON a.emp_id = s.emp_id
  LEFT JOIN HR_EMP_SCHEDUAL_additional sa 
    ON a.emp_id = sa.emp_id 
    AND DATE(a.create_date) = sa.date
  WHERE (? IS NULL OR DATE(a.create_date) = ?) -- optional date filter
  GROUP BY DATE(a.create_date), a.emp_id
) t
WHERE t.first_in_time > t.scheduled_in_time
ORDER BY t.date DESC
LIMIT ? OFFSET ?;


-- name: GetInsufficientAttendanceForAll :many
SELECT *
FROM (
  SELECT
    DATE(a.create_date) as date,
    a.emp_id,
    TIME_FORMAT(MIN(CASE WHEN a.attendance_type = 'in' THEN TIME(a.create_date) END), '%H:%i:%s') as in_time,
    TIME_FORMAT(MAX(CASE WHEN a.attendance_type = 'out' THEN TIME(a.create_date) END), '%H:%i:%s') as out_time,
    CASE 
      WHEN MIN(CASE WHEN a.attendance_type = 'in' THEN TIME(a.create_date) END) IS NOT NULL 
           AND MAX(CASE WHEN a.attendance_type = 'out' THEN TIME(a.create_date) END) IS NOT NULL 
      THEN TIME_FORMAT(TIMEDIFF(
        MAX(CASE WHEN a.attendance_type = 'out' THEN TIME(a.create_date) END),
        MIN(CASE WHEN a.attendance_type = 'in' THEN TIME(a.create_date) END)
      ), '%H:%i:%s')
      ELSE NULL 
    END as total_time,
    TIME_FORMAT(TIMEDIFF(
      COALESCE(sa.to_time,
        CASE DAYOFWEEK(DATE(a.create_date))
          WHEN 1 THEN s.sunday_to
          WHEN 2 THEN s.monday_to
          WHEN 3 THEN s.tuesday_to
          WHEN 4 THEN s.wednesday_to
          WHEN 5 THEN s.thursday_to
          WHEN 6 THEN s.friday_to
          WHEN 7 THEN s.saturday_to
        END), 
      COALESCE(sa.from_time, 
        CASE DAYOFWEEK(DATE(a.create_date))
          WHEN 1 THEN s.sunday_from
          WHEN 2 THEN s.monday_from  
          WHEN 3 THEN s.tuesday_from
          WHEN 4 THEN s.wednesday_from
          WHEN 5 THEN s.thursday_from
          WHEN 6 THEN s.friday_from
          WHEN 7 THEN s.saturday_from
        END)
    ), '%H:%i:%s') as scheduled_time
  FROM HR_EMP_ATTENDANCE a
  LEFT JOIN HR_EMP_SCHEDUAL s ON a.emp_id = s.emp_id
  LEFT JOIN HR_EMP_SCHEDUAL_additional sa 
    ON a.emp_id = sa.emp_id 
    AND DATE(a.create_date) = sa.date
  WHERE (? IS NULL OR DATE(a.create_date) = ?)
  GROUP BY DATE(a.create_date), a.emp_id
) t
WHERE t.total_time < t.scheduled_time
ORDER BY t.date DESC
LIMIT ? OFFSET ?;

