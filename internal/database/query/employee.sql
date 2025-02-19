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
    salary_type, amount, Total_of_salary_allowances, pension_employer, pension_employee, total_net_salary, employee_id, updated_by
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: CreateEmpCertificates :exec
INSERT INTO HR_EMP_Certificates (
    date, name, image_path, updated_by, employee_id
) VALUES (
    ?, ?, ?, ?, ?
);

-- name: CreateEmpStatus :exec
INSERT INTO HR_EMP_Status (
    status, department, designation, valid_from, valid_till, updated_by, employee_id
) VALUES (
    ?, ?, ?, ?, ?, ?, ?
);

-- name: CreateEmpBenifits :exec
INSERT INTO HR_EMP_Benifits (
    leave_status, leave_type, leave_count, health_insurance, insurance_from, insurance_till, retainment_plan, retainment_plan_from, retainment_plan_till, benifits, benifits_from, benifits_till, updated_by, employee_id
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: CreateEmpUser :exec
INSERT INTO HR_EMP_User (
    email, password, updated_by, employee_id 
) VALUES (
    ?, ?, ?, ?
);

-- name: CreateEmpAllowances :exec
INSERT INTO HR_EMP_Allowances (
   name, amount, updated_by, employee_id 
) VALUES (
    ?, ?, ?, ?
);

-- name: CreateEmpExpatriate :exec
INSERT INTO HR_EMP_Expatriate (
    expatriate, nationality, visa_type, visa_from, visa_till, visa_number, visa_fee, visa_image_path, updated_by, employee_id
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: CreateEmpAccessiability :exec
INSERT INTO HR_EMP_Accessiability (
    accessibility, accessibility_from, accessibility_till, enable, updated_by, employee_id
) VALUES (
    ?, ?, ?, ?, ?, ?
);


-- name: GetEmployee :many
SELECT * FROM HR_Employee
ORDER BY id DESC
LIMIT ? OFFSET ?;

-- name: GetEmployeeDOB :one
SELECT dob FROM HR_Employee WHERE id = ?;

-- name: GetEmployeeByID :one
SELECT 
    e.id AS employee_id, 
    e.first_name, e.last_name, e.gender, e.dob, e.religion, 
    e.primary_number, e.secondary_number, e.passport_id, e.nationality, 
    e.passport_valid_till, e.nic, e.country, e.nic_valid_till, 
    e.address,e.current_country, e.email, e.updated_by, e.created_at, e.updated_at,

    ed.first_name AS emergency_first_name, ed.last_name AS emergency_last_name, 
    ed.relationship, ed.contact AS emergency_contact, 

    bd.bank_name, bd.branch_name, bd.account_number, bd.account_holder,

    s.salary_type, s.amount, s.Total_of_salary_allowances, 
    s.pension_employer, s.pension_employee, s.total_net_salary, 

    cert.date AS certificate_date, cert.name AS certificate_name, 
    cert.image_path AS certificate_image, 

    stat.status, stat.department, stat.designation, 
    stat.valid_from AS status_valid_from, stat.valid_till AS status_valid_till, 

    ben.leave_status, ben.leave_type, ben.leave_count, 
    ben.health_insurance, ben.insurance_from, ben.insurance_till, 
    ben.retainment_plan, ben.retainment_plan_from, ben.retainment_plan_till, 
    ben.benifits, ben.benifits_from, ben.benifits_till, 

    usr.email AS user_email, usr.password AS user_password, 

    allw.name AS allowance_name, allw.amount AS allowance_amount, 

    exp.expatriate, exp.nationality AS exp_nationality, 
    exp.visa_type, exp.visa_from, exp.visa_till, exp.visa_number, 
    exp.visa_fee, exp.visa_image_path, 

    acc.accessibility, acc.accessibility_from, acc.accessibility_till, acc.enable 

FROM HR_Employee e
LEFT JOIN HR_EMP_Emergency_Details ed ON e.id = ed.employee_id
LEFT JOIN HR_EMP_Bank_Details bd ON e.id = bd.employee_id
LEFT JOIN HR_EMP_Salary s ON e.id = s.employee_id
LEFT JOIN HR_EMP_Certificates cert ON e.id = cert.employee_id
LEFT JOIN HR_EMP_Status stat ON e.id = stat.employee_id
LEFT JOIN HR_EMP_Benifits ben ON e.id = ben.employee_id
LEFT JOIN HR_EMP_User usr ON e.id = usr.employee_id
LEFT JOIN HR_EMP_Allowances allw ON e.id = allw.employee_id
LEFT JOIN HR_EMP_Expatriate exp ON e.id = exp.employee_id
LEFT JOIN HR_EMP_Accessiability acc ON e.id = acc.employee_id

WHERE e.id = ?;

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
    date = ?, name = ?, image_path = ?, updated_by = ?
WHERE employee_id = ?;

-- name: UpdateEmpStatus :exec
UPDATE HR_EMP_Status SET
    status = ?, department = ?, designation = ?, valid_from = ?, valid_till = ?, updated_by = ?
WHERE employee_id = ?;

-- name: UpdateEmpBenifits :exec
UPDATE HR_EMP_Benifits SET
    leave_status = ?, leave_type = ?, leave_count = ?, health_insurance = ?, 
    insurance_from = ?, insurance_till = ?, retainment_plan = ?, retainment_plan_from = ?, 
    retainment_plan_till = ?, benifits = ?, benifits_from = ?, benifits_till = ?, updated_by = ?
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
    visa_number = ?, visa_fee = ?, visa_image_path = ?, updated_by = ?
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

-- name: DeleteEmpAccessiability :exec
DELETE FROM HR_EMP_Accessiability WHERE employee_id = ?;





