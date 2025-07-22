-- name: GetExpiredVisaOrReports :many
SELECT 
    e.id as `Employee ID`,
    CONCAT(e.first_name, ' ', e.last_name) as name,
    s.department,
    s.designation,
    s.status as employee_status,
    CONCAT(DATE_FORMAT(s.valid_from, '%Y-%m-%d'), ' / ', DATE_FORMAT(s.valid_till, '%Y-%m-%d')) as status_from_to,
    ex.nationality,
    ex.visa_type,
    e.passport_id as passport_no,
    ex.visa_number,
    CONCAT(DATE_FORMAT(ex.visa_from, '%Y-%m-%d'), ' / ', DATE_FORMAT(ex.visa_till, '%Y-%m-%d')) as visa_from_to
FROM HR_Employee e
INNER JOIN HR_EMP_Status s ON e.id = s.employee_id
INNER JOIN HR_EMP_Expatriate ex ON e.id = ex.employee_id
WHERE (
    ex.visa_till < CONVERT_TZ(NOW(), '+00:00', '+05:00') 
    OR e.passport_valid_till < CONVERT_TZ(NOW(), '+00:00', '+05:00')
);

-- name: GetVisaOrPassportExpiringSoon :many
SELECT 
    e.id as `Employee ID`,
    CONCAT(e.first_name, ' ', e.last_name) as name,
    s.department,
    s.designation,
    s.status as employee_status,
    CONCAT(DATE_FORMAT(s.valid_from, '%Y-%m-%d'), ' / ', DATE_FORMAT(s.valid_till, '%Y-%m-%d')) as status_from_to,
    ex.nationality,
    ex.visa_type,
    e.passport_id as passport_no,
    ex.visa_number,
    CONCAT(DATE_FORMAT(ex.visa_from, '%Y-%m-%d'), ' / ', DATE_FORMAT(ex.visa_till, '%Y-%m-%d')) as visa_from_to
FROM HR_Employee e
INNER JOIN HR_EMP_Status s ON e.id = s.employee_id
INNER JOIN HR_EMP_Expatriate ex ON e.id = ex.employee_id
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
);

-- name: GetStaffPayroll :many
SELECT 
    e.id as `Employee ID`,
    CONCAT(e.first_name, ' ', e.last_name) as name,
    s.department,
    s.designation,
    s.status as employee_status,
    DATE_FORMAT(p.date, '%Y-%m') as month,
    p.amount as gross_salary,
    p.total_of_salary_allowances as allowance,
    p.pension_employee as pension_employee,
    p.pension_employer as pension_employer,
    p.total_net_salary as net_salary,
    p.tax_percentage as tax_percentage,
    p.total_net_salary_after_tax as net_salary_after_tax,
    p.created_at as process_date
FROM HR_Employee e
INNER JOIN HR_EMP_Status s ON e.id = s.employee_id
INNER JOIN HR_Payroll p ON e.id = p.employee
WHERE (
    Month(CONVERT_TZ(p.date,'+00:00', '+05:00'))  = Month(?)
);

-- name: GetAccountDetails :many
SELECT 
    b.account_number as account_number,
    b.account_holder as account_name,
    p.total_net_salary_after_tax as amount_paid
FROM HR_EMP_Bank_Details b
INNER JOIN HR_Payroll p ON b.employee_id = p.emp_id
WHERE (
    Month(CONVERT_TZ(p.date,'+00:00', '+05:00'))  = Month(?)
);

-- name: GetempployeeInsurance :many
SELECT 
    e.id as `Employee ID`,
    CONCAT(e.first_name, ' ', e.last_name) as name,
    s.department,
    s.designation,
    s.status as employee_status,
    CONCAT(DATE_FORMAT(s.valid_from, '%Y-%m-%d'), ' / ', DATE_FORMAT(s.valid_till, '%Y-%m-%d')) as status_from_to,
    ex.nationality,
    i.health_insurance,
    CONCAT(DATE_FORMAT(i.insurance_from, '%Y-%m-%d'), ' / ', DATE_FORMAT(i.insurance_till, '%Y-%m-%d')) as insurance_from_to
FROM HR_Employee e
INNER JOIN HR_EMP_Status s ON e.id = s.employee_id
INNER JOIN HR_EMP_Expatriate ex ON e.id = ex.employee_id
INNER JOIN HR_EMP_Benifits i ON e.id = i.employee_id;

