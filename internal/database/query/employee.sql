-- name: CreateEmployee :execresult
INSERT INTO HR_Employee (
    first_name, last_name, gender, dob, religion, primary_number, secondary_number, 
    passport_id, nationality, passport_valid_till, nic, country, nic_valid_till, 
    address, current_country, email, updated_by
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);
SELECT LAST_INSERT_ID() AS id;

-- name: CreateEmpEmergencyDetails :exec
INSERT INTO HR_EMP_Emergency_Details (
    first_name, last_name, relationship, contact, employee_id, updated_by
) VALUES (
    ?, ?, ?, ?, ?, ?
);

-- name: CreateEmpBankDetails :exec
INSERT INTO HR_EMP_Bank_Details (
    bank_name, branch_name, account_number, account_holder, employee_id, updated_by
) VALUES (
    ?, ?, ?, ?, ?, ?
);

-- name: CreateEmpSalary :exec
INSERT INTO HR_EMP_Salary (
    salary_type, amount, salary_amount_type, Total_of_salary_allowances, total_salary_allowances_type, pension_employer, pension_employer_type, pension_employee, pension_employee_type, total_net_salary, total_net_salary_type, employee_id, updated_by, er_id
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: CreateEmpCertificates :exec
INSERT INTO HR_EMP_Certificates (
    date, name,updated_by, employee_id
) VALUES (
    ?, ?, ?, ?
);

-- name: CreateEmpStatus :exec
INSERT INTO HR_EMP_Status (
    status, department, designation, valid_from, valid_till, updated_by, employee_id
) VALUES (
    ?, ?, ?, ?, ?, ?, ?
);

-- name: CreateEmpBenifits :exec
INSERT INTO HR_EMP_Benifits (
    leave_status, leave_type, leave_count, health_insurance, insurance_from, insurance_till, retainment_plan, retainment_plan_from, retainment_plan_till, uniform, uniform_quantity, uniform_renew_months, ticket, ticket_quantity, updated_by, employee_id
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: CreateEmpUser :exec
INSERT INTO HR_EMP_User (
    email, password, updated_by, employee_id, branch_id 
) VALUES (
    ?, ?, ?, ?,?
);

-- name: CreateEmpAllowances :exec
INSERT INTO HR_EMP_Allowances (
   name, amount, updated_by, employee_id 
) VALUES (
    ?, ?, ?, ?
);

-- name: CreateEmpExpatriate :exec
INSERT INTO HR_EMP_Expatriate (
    expatriate, nationality, visa_type, visa_from, visa_till, visa_number, visa_fee, updated_by, employee_id
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: CreateEmpAccessiability :exec
INSERT INTO HR_EMP_Accessiability (
    accessibility, accessibility_from, accessibility_till, enable, updated_by, employee_id
) VALUES (
    ?, ?, ?, ?, ?, ?
);


-- name: GetEmployee :many
SELECT
  e.id AS employee_id,
  e.first_name,
  e.last_name,
  usr.email AS user_email,
  br.name AS branch_name
FROM HR_Employee e
LEFT JOIN HR_EMP_User usr ON e.id = usr.employee_id
LEFT JOIN HR_Branch br ON usr.branch_id = br.id
WHERE 
    (e.is_ban = false OR e.is_ban IS NULL)
    AND
  (
    ? = '' OR 
    e.first_name LIKE CONCAT('%', ?, '%') OR 
    e.last_name LIKE CONCAT('%', ?, '%') OR
    CONCAT(e.first_name, ' ', e.last_name) LIKE CONCAT('%', ?, '%')
  )
  AND (? = 0 OR br.id = ?)
ORDER BY e.id DESC
LIMIT ? OFFSET ?;

-- name: GetEmployeeDOB :one
SELECT dob FROM HR_Employee WHERE id = ?;

-- name: GetEmployeeByID :one
SELECT 
    e.id AS employee_id, 
    e.first_name, 
    e.last_name, 
    e.gender, 
    e.dob, 
    e.religion, 
    e.primary_number, 
    e.secondary_number, 
    e.passport_id, 
    e.nationality, 
    e.passport_valid_till, 
    e.nic, 
    e.country, 
    e.nic_valid_till, 
    e.address,
    e.current_country, 
    e.email, 
    e.updated_by, 
    e.created_at, 
    e.updated_at,

    ed.first_name AS emergency_first_name, 
    ed.last_name AS emergency_last_name, 
    ed.relationship, 
    ed.contact AS emergency_contact, 

    bd.bank_name, 
    bd.branch_name, 
    bd.account_number, 
    bd.account_holder,

    s.salary_type, 
    s.amount, 
    s.salary_amount_type,
    s.Total_of_salary_allowances,
    s.total_salary_allowances_type, 
    s.pension_employer,
    s.pension_employer_type, 
    s.pension_employee, 
    s.pension_employee_type,
    s.total_net_salary,
    s.total_net_salary_type, 

    cert.date AS certificate_date, 
    cert.name AS certificate_name, 
    
    stat.status, 
    stat.department, 
    stat.designation, 
    stat.valid_from AS status_valid_from, 
    stat.valid_till AS status_valid_till, 

    ben.leave_status, 
    ben.leave_type, 
    ben.leave_count, 
    ben.health_insurance, 
    ben.insurance_from, 
    ben.insurance_till, 
    ben.retainment_plan, 
    ben.retainment_plan_from, 
    ben.retainment_plan_till, 
    ben.uniform, 
    ben.uniform_quantity, 
    ben.ticket, 
    ben.ticket_quantity, 
    ben.uniform_renew_months, 

    usr.email AS user_email, 
    usr.password AS user_password, 
    usr.branch_id AS user_branch_id, 

    exp.expatriate, 
    exp.nationality AS exp_nationality, 
    exp.visa_type, 
    exp.visa_from, 
    exp.visa_till, 
    exp.visa_number, 
    exp.visa_fee, 
    
    acc.accessibility, 
    acc.accessibility_from, 
    acc.accessibility_till, 
    acc.enable

FROM HR_Employee e
LEFT JOIN HR_EMP_Emergency_Details ed ON e.id = ed.employee_id
LEFT JOIN HR_EMP_Bank_Details bd ON e.id = bd.employee_id
LEFT JOIN HR_EMP_Salary s ON e.id = s.employee_id
LEFT JOIN HR_EMP_Certificates cert ON e.id = cert.employee_id
LEFT JOIN HR_EMP_Status stat ON e.id = stat.employee_id
LEFT JOIN HR_EMP_Benifits ben ON e.id = ben.employee_id
LEFT JOIN HR_EMP_User usr ON e.id = usr.employee_id
LEFT JOIN HR_EMP_Expatriate exp ON e.id = exp.employee_id
LEFT JOIN HR_EMP_Accessiability acc ON e.id = acc.employee_id

WHERE e.id = ? AND (e.is_ban = false OR e.is_ban IS NULL);

-- name: UpdateEmployee :exec
UPDATE HR_Employee SET 
    first_name = ?, last_name = ?, gender = ?, dob = ?, religion = ?, primary_number = ?, secondary_number = ?,
    passport_id = ?, nationality = ?, passport_valid_till = ?, nic = ?, country = ?, nic_valid_till = ?, address = ?, current_country = ?,
    email = ?, updated_by = ?
WHERE id = ?;

-- name: UpdateEmpEmergencyDetails :exec
UPDATE HR_EMP_Emergency_Details SET
    first_name = ?, last_name = ?, relationship = ?, contact = ?, updated_by = ?
WHERE employee_id = ?;

-- name: UpdateEmpBankDetails :exec
UPDATE HR_EMP_Bank_Details SET
    bank_name = ?, branch_name = ?, account_number = ?, account_holder = ?, updated_by = ?
WHERE employee_id = ?;

-- name: UpdateEmpSalary :exec
UPDATE HR_EMP_Salary SET
    salary_type = ?, amount = ?, Total_of_salary_allowances = ?, pension_employer = ?, 
    pension_employee = ?, total_net_salary = ?, updated_by = ?
WHERE employee_id = ?;

-- name: UpdateEmpCertificates :exec
UPDATE HR_EMP_Certificates SET
    date = ?, name = ?, updated_by = ?
WHERE employee_id = ?;

-- name: UpdateTrainerCommission :exec
UPDATE HR_Trainer_Emp SET
    commission = ?,
    updated_by = ?
WHERE employee_id = ?;

-- name: UpdateEmpStatus :exec
UPDATE HR_EMP_Status SET
    status = ?, department = ?, designation = ?, valid_from = ?, valid_till = ?, updated_by = ?
WHERE employee_id = ?;

-- name: UpdateEmpBenifits :exec
UPDATE HR_EMP_Benifits SET
    leave_status = ?, leave_type = ?, leave_count = ?, health_insurance = ?, 
    insurance_from = ?, insurance_till = ?, retainment_plan = ?, retainment_plan_from = ?, 
    retainment_plan_till = ?, uniform = ?, uniform_quantity = ?, uniform_renew_months = ?, ticket = ?, ticket_quantity = ?, updated_by = ?
WHERE employee_id = ?;

-- name: UpdateEmpUser :exec
UPDATE HR_EMP_User SET
    email = ?, password = ?, updated_by = ?
WHERE employee_id = ?;

-- name: UpdateEmpAllowances :exec
UPDATE HR_EMP_Allowances SET
    name = ?, amount = ?, updated_by = ?
WHERE employee_id = ?;

-- name: UpdateEmpExpatriate :exec
UPDATE HR_EMP_Expatriate SET
    expatriate = ?, nationality = ?, visa_type = ?, visa_from = ?, visa_till = ?, 
    visa_number = ?, visa_fee = ?, updated_by = ?
WHERE employee_id = ?;

-- name: UpdateEmpAccessiability :exec
UPDATE HR_EMP_Accessiability SET
    accessibility = ?, accessibility_from = ?, accessibility_till = ?, enable = ?, updated_by = ?
WHERE employee_id = ?;


-- name: DeleteEmployee :exec
DELETE FROM HR_Employee WHERE id = ?;

-- name: DeleteEmpEmergencyDetails :exec
DELETE FROM HR_EMP_Emergency_Details WHERE employee_id = ?;

-- name: DeleteEmpBankDetails :exec
DELETE FROM HR_EMP_Bank_Details WHERE employee_id = ?;

-- name: DeleteEmpSalary :exec
DELETE FROM HR_EMP_Salary WHERE employee_id = ?;

-- name: DeleteEmpCertificates :exec
DELETE FROM HR_EMP_Certificates WHERE employee_id = ?;

-- name: DeleteEmpStatus :exec
DELETE FROM HR_EMP_Status WHERE employee_id = ?;

-- name: DeleteEmpBenifits :exec
DELETE FROM HR_EMP_Benifits WHERE employee_id = ?;

-- name: DeleteEmpUser :exec
DELETE FROM HR_EMP_User WHERE employee_id = ?;

-- name: DeleteEmpAllowances :exec
DELETE FROM HR_EMP_Allowances WHERE employee_id = ?;

-- name: DeleteEmpExpatriate :exec
DELETE FROM HR_EMP_Expatriate WHERE employee_id = ?;

-- name: DeleteEmpFiles :exec
DELETE FROM HR_FileSubmit WHERE file_name = ? AND employee_id = ? AND file_type = ?;

-- name: DeleteEmpAccessiability :exec
DELETE FROM HR_EMP_Accessiability WHERE employee_id = ?;

-- name: EmployeeLogin :one
SELECT employee_id, password, email, branch_id
FROM HR_EMP_User
WHERE email = ?;

-- name: GetCertificateFile :one
SELECT file_name FROM HR_FileSubmit WHERE employee_id = ? AND file_type = 'certificate';

-- name: GetVisaFile :one
SELECT file_name FROM HR_FileSubmit WHERE employee_id = ? AND file_type = 'visa';

-- name: GetEmployeeFromBranch :many
SELECT e.id, CONCAT(e.first_name, ' ', e.last_name) AS full_name
FROM HR_Employee e
JOIN HR_EMP_User u ON e.id = u.employee_id
WHERE u.branch_id = ? AND e.is_ban = false;

-- name: GetEmployeeSalaryDetails :one
SElECT salary_type, amount, Total_of_salary_allowances, pension_employer, pension_employee, total_net_salary
FROM HR_EMP_Salary
WHERE employee_id = ?;

-- name: GetEmployeeAllowances :many
SELECT name, amount
FROM HR_EMP_Allowances
WHERE employee_id = ?;

-- name: CheckTrainerFromEmail :one
SELECT attendee_id, user_id, branch_id, username, email, phone, nic, role
FROM door_lock_users
WHERE email = ? AND role = 'trainer';

-- name: CreateTrainerEmp :exec
INSERT INTO HR_Trainer_Emp (
    trainer_id, employee_id, attendee_id, commission
) VALUES (
    ?, ?, ?, ?
);

-- name: GetTrainerEmp :one
SELECT
    trainer_id, employee_id, attendee_id, commission
FROM HR_Trainer_Emp
WHERE employee_id = ?;

-- name: GetEmpFiles :many
SELECT file_name, file_type
FROM HR_FileSubmit
WHERE employee_id = ?;

-- name: DeleteTrainerEmp :exec
DELETE FROM HR_Trainer_Emp WHERE employee_id = ?;

-- name: CheckTrainerAssignmentAtTime :one
SELECT 
    EXISTS (
        SELECT 1
        FROM FLM_trainer_assign 
        WHERE trainer_id = ? 
        AND client_id = ?
        AND CONVERT_TZ(?, '+00:00', '+05:00') BETWEEN 
            CONVERT_TZ(`from`, '+00:00', '+05:00') AND 
            CONVERT_TZ(`to`, '+00:00', '+05:00')
    ) as is_assigned;

-- name: BanEmployee :exec
UPDATE HR_Employee SET is_ban = ? WHERE id = ?;

-- name: GetBranchwiseEmpCount :one
SELECT 
    b.id as branch_id,
    b.name as branch_name,
    COUNT(e.id) as employee_count
FROM HR_Branch b
LEFT JOIN HR_EMP_User u ON b.id = u.branch_id
LEFT JOIN HR_Employee e ON u.employee_id = e.id
WHERE b.id = ? AND (e.is_ban = false OR e.is_ban IS NULL)
GROUP BY b.id, b.name;





