-- name: GetExpiredVisaOrReports :many
SELECT 
    e.id as `Employee ID`,
    CONCAT(e.first_name, ' ', e.last_name) as name,
    s.department,
    s.designation,
    s.status as employee_status,
    CONCAT(DATE_FORMAT(s.valid_from, '%Y-%m-%d'), ' / ', DATE_FORMAT(s.valid_till, '%Y-%m-%d')) as status_from_to,
    u.branch_id as branch_id,
    b.name as branch_name,
    ex.nationality,
    ex.visa_type,
    e.passport_id as passport_no,
    ex.visa_number,
    CONCAT(DATE_FORMAT(ex.visa_from, '%Y-%m-%d'), ' / ', DATE_FORMAT(ex.visa_till, '%Y-%m-%d')) as visa_from_to
FROM HR_Employee e
INNER JOIN HR_EMP_Status s ON e.id = s.employee_id
INNER JOIN HR_EMP_Expatriate ex ON e.id = ex.employee_id
INNER JOIN HR_EMP_User u ON e.id = u.employee_id
INNER JOIN HR_Branch b ON u.branch_id = b.id
WHERE (
    ex.visa_till < CONVERT_TZ(NOW(), '+00:00', '+05:00') 
    OR e.passport_valid_till < CONVERT_TZ(NOW(), '+00:00', '+05:00')
)
AND (
    u.branch_id = ? OR u.branch_id = ''
)
AND (? = '' OR e.first_name LIKE CONCAT('%', ?, '%'))
AND (? = '' OR e.last_name LIKE CONCAT('%', ?, '%'))
AND (? = '' OR s.department LIKE CONCAT('%', ?, '%'))
AND (? = '' OR e.passport_id LIKE CONCAT('%', ?, '%'))
AND (? = '' OR ex.visa_number LIKE CONCAT('%', ?, '%'))
ORDER BY e.id
LIMIT ? OFFSET ?;

-- name: GetVisaOrPassportExpiringSoon :many
SELECT 
    e.id as `Employee ID`,
    CONCAT(e.first_name, ' ', e.last_name) as name,
    s.department,
    s.designation,
    s.status as employee_status,
    CONCAT(DATE_FORMAT(s.valid_from, '%Y-%m-%d'), ' / ', DATE_FORMAT(s.valid_till, '%Y-%m-%d')) as status_from_to,
    u.branch_id as branch_id,
    b.name as branch_name,
    ex.nationality,
    ex.visa_type,
    e.passport_id as passport_no,
    ex.visa_number,
    CONCAT(DATE_FORMAT(ex.visa_from, '%Y-%m-%d'), ' / ', DATE_FORMAT(ex.visa_till, '%Y-%m-%d')) as visa_from_to
FROM HR_Employee e
INNER JOIN HR_EMP_Status s ON e.id = s.employee_id
INNER JOIN HR_EMP_Expatriate ex ON e.id = ex.employee_id
INNER JOIN HR_EMP_User u ON e.id = u.employee_id
INNER JOIN HR_Branch b ON u.branch_id = b.id
WHERE (
    (
        ex.visa_till > CONVERT_TZ(NOW(), '+00:00', '+05:00') AND
        ex.visa_till <= DATE_ADD(CONVERT_TZ(NOW(), '+00:00', '+05:00'), INTERVAL 7 DAY)
    )
    OR
    (
        e.passport_valid_till > CONVERT_TZ(NOW(), '+00:00', '+05:00') AND
        e.passport_valid_till <= DATE_ADD(CONVERT_TZ(NOW(), '+00:00', '+05:00'), INTERVAL 7 DAY)
    )
)
AND (
     u.branch_id = ? OR u.branch_id = ''
)
AND (? = '' OR e.first_name LIKE CONCAT('%', ?, '%'))
AND (? = '' OR e.last_name LIKE CONCAT('%', ?, '%'))
AND (? = '' OR s.department LIKE CONCAT('%', ?, '%'))
AND (? = '' OR e.passport_id LIKE CONCAT('%', ?, '%'))
AND (? = '' OR ex.visa_number LIKE CONCAT('%', ?, '%'))
ORDER BY e.id
LIMIT ? OFFSET ?;

-- name: GetStaffPayroll :many
SELECT 
    e.id as `Employee ID`,
    CONCAT(e.first_name, ' ', e.last_name) as name,
    s.department,
    s.designation,
    s.status as employee_status,
    u.branch_id as branch_id,
    b.name as branch_name,
    DATE_FORMAT(p.date, '%Y-%m') as month,
    p.amount as gross_salary,
    p.salary_amount_type as gross_salary_type,
    p.total_of_salary_allowances as allowance,
    p.pension_employee as pension_employee,
    p.pension_employer as pension_employer,
    p.total_net_salary as net_salary,
    p.tax_percentage as tax_percentage,
    p.total_net_salary_after_tax as net_salary_after_tax,
    p.created_at as process_date
FROM HR_Employee e
INNER JOIN HR_EMP_Status s ON e.id = s.employee_id
INNER JOIN HR_Payroll p ON e.id = p.emp_id
INNER JOIN HR_EMP_User u ON e.id = u.employee_id
INNER JOIN HR_Branch b ON u.branch_id = b.id
WHERE 
    DATE_FORMAT(CONVERT_TZ(p.date, '+00:00', '+05:00'), '%Y-%m') = DATE_FORMAT(?, '%Y-%m')
    AND (
        (? = '' OR e.first_name LIKE CONCAT('%', ?, '%'))
        OR (? = '' OR e.last_name LIKE CONCAT('%', ?, '%'))
    )
    AND (
        u.branch_id = ? OR ? = 0
    )
ORDER BY p.date DESC
LIMIT ? OFFSET ?;

-- name: GetAccountDetails :many
SELECT 
    b.account_number as account_number,
    b.account_holder as account_name,
    p.total_net_salary_after_tax as amount_paid,
    u.branch_id as branch_id,
    br.name as branch_name
FROM HR_EMP_Bank_Details b
INNER JOIN HR_Payroll p ON b.employee_id = p.emp_id
INNER JOIN HR_EMP_User u ON b.employee_id = u.employee_id
INNER JOIN HR_Branch br ON u.branch_id = br.id
WHERE (
DATE_FORMAT(CONVERT_TZ(p.date, '+00:00', '+05:00'), '%Y-%m') = DATE_FORMAT(?, '%Y-%m')
    AND (
        u.branch_id = ? OR u.branch_id = ''
    )
);

-- name: GetempployeeInsurance :many
SELECT 
    e.id as `Employee ID`,
    CONCAT(e.first_name, ' ', e.last_name) as name,
    s.department,
    s.designation,
    s.status as employee_status,
    CONCAT(DATE_FORMAT(s.valid_from, '%Y-%m-%d'), ' / ', DATE_FORMAT(s.valid_till, '%Y-%m-%d')) as status_from_to,
    u.branch_id as branch_id,
    b.name as branch_name,
    ex.nationality,
    i.health_insurance,
    CONCAT(DATE_FORMAT(i.insurance_from, '%Y-%m-%d'), ' / ', DATE_FORMAT(i.insurance_till, '%Y-%m-%d')) as insurance_from_to
FROM HR_Employee e
INNER JOIN HR_EMP_Status s ON e.id = s.employee_id
INNER JOIN HR_EMP_Expatriate ex ON e.id = ex.employee_id
INNER JOIN HR_EMP_Benifits i ON e.id = i.employee_id
INNER JOIN HR_EMP_User u ON e.id = u.employee_id
INNER JOIN HR_Branch b ON u.branch_id = b.id
WHERE (
    u.branch_id = ? OR u.branch_id = ''
)
AND (? = '' OR e.first_name LIKE CONCAT('%', ?, '%'))
AND (? = '' OR e.last_name LIKE CONCAT('%', ?, '%'))
AND (? = '' OR s.department LIKE CONCAT('%', ?, '%'))
LIMIT ? OFFSET ?;

