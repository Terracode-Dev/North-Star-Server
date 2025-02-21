-- name: CreatePayroll :execresult
INSERT INTO HR_Payroll (
    employee, date, salary_type, amount, total_of_salary_allowances, pension, pension_employer, pension_employee, total_net_salary, tax, tax_percentage, total_net_salary_after_tax, updated_by
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
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
    p.total_of_salary_allowances,
    p.pension,
    p.pension_employer,
    p.pension_employee,
    p.total_net_salary,
    p.tax,
    p.tax_percentage,
    p.total_net_salary_after_tax,
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
SET employee = ?, date = ?, salary_type = ?, amount = ?, total_of_salary_allowances = ?, pension = ?, pension_employer = ?, pension_employee = ?, total_net_salary = ?, tax = ?, tax_percentage = ?, total_net_salary_after_tax = ?, updated_by = ?
WHERE id = ?;

-- name: CreatePayrollAllowances :exec
INSERT INTO HR_Payroll_Allowances (
    name, amount, payroll_id, updated_by
) VALUES (
    ?, ?, ?, ?
);

-- name: UpdatePayrollAllowance :exec
UPDATE HR_Payroll_Allowances
SET name = ?, amount = ?, updated_by = ?
WHERE payroll_id = ?;
