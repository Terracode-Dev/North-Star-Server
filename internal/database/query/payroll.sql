-- name: CreatePayroll :execresult
INSERT INTO HR_Payroll (
    employee, date, salary_type, amount, salary_amount_type, total_of_salary_allowances,total_allowances_type, pension, pension_employer, pension_employer_type, pension_employee, pension_employee_type, total_net_salary, total_net_salary_type, tax, tax_percentage, total_net_salary_after_tax, total_net_salary_after_tax_type, er_id, updated_by
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);
SELECT LAST_INSERT_ID() AS id;

-- name: GetPayrolls :many
SELECT * FROM HR_Payroll
ORDER BY id DESC
LIMIT ? OFFSET ?;

-- name: GetOnePayroll :many
SELECT 
    p.id AS payroll_id,
    p.employee,
    p.date,
    p.salary_type,
    p.amount,
    p.salary_amount_type,
    p.total_of_salary_allowances,
    p.total_allowances_type,
    p.pension,
    p.pension_employer,
    p.pension_employer_type,
    p.pension_employee,
    p.pension_employee_type,
    p.total_net_salary,
    p.total_net_salary_type,
    p.tax,
    p.tax_percentage,
    p.total_net_salary_after_tax,
    p.total_net_salary_after_tax_type,
    p.updated_by AS payroll_updated_by,
    p.created_at AS payroll_created_at,
    p.updated_at AS payroll_updated_at,
    a.id AS allowance_id,
    a.name AS allowance_name,
    a.amount AS allowance_amount,
    a.payroll_id AS allowance_payroll_id,
    a.updated_by AS allowance_updated_by,
    a.created_at AS allowance_created_at
FROM HR_Payroll p
LEFT JOIN HR_Payroll_Allowances a 
    ON p.id = a.payroll_id
WHERE p.id = ?;

-- name: UpdatePayroll :exec
UPDATE HR_Payroll
SET
    employee = ?,
    date = ?,
    salary_type = ?,
    amount = ?,
    salary_amount_type = ?,
    total_of_salary_allowances = ?,
    total_allowances_type = ?,
    pension = ?,
    pension_employer = ?,
    pension_employer_type = ?,
    pension_employee = ?,
    pension_employee_type = ?,
    total_net_salary = ?,
    total_net_salary_type = ?,
    tax = ?,
    tax_percentage = ?,
    total_net_salary_after_tax = ?,
    total_net_salary_after_tax_type = ?,
    updated_by = ?
WHERE id = ?;

-- name: CreatePayrollAllowances :exec
INSERT INTO HR_Payroll_Allowances (
    name, amount, amount_type, payroll_id, updated_by
) VALUES (
    ?, ?, ?, ?, ?
);

-- name: UpdatePayrollAllowance :exec
UPDATE HR_Payroll_Allowances
SET name = ?, amount = ?, amount_type=?, updated_by = ?
WHERE payroll_id = ?;

-- name: GetTrainerEmpDataFromID :one
SELECT trainer_id, employee_id, attendee_id, commission
FROM HR_Trainer_Emp
WHERE employee_id = ?;

-- name: GetTrainerAssingedCount :one
SELECT COUNT(*) AS count
FROM FLM_trainer_assign
WHERE trainer_id = ?
  AND MONTH(CONVERT_TZ(`from`, '+00:00', '+05:00')) = MONTH(CONVERT_TZ(CURRENT_TIMESTAMP(), '+00:00', '+05:00'))
  AND YEAR(CONVERT_TZ(`from`, '+00:00', '+05:00')) = YEAR(CONVERT_TZ(CURRENT_TIMESTAMP(), '+00:00', '+05:00'));

-- name: CreateHRTrainerCom :exec
INSERT INTO HR_Trainer_Com (
    payroll_id,
    trainer_id,
    employee_id,
    commission,
    assigned_count,
    total
) VALUES (
    ?, ?, ?, ?, ?, ?
);