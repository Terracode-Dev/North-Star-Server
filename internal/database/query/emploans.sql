-- name: CreateRequest :exec
INSERT INTO emp_loan_req (emp_id, reason, amount, status, declined_by, decline_reason) VALUES (
    ?, ?, ?, ?, ?, ?
);

-- name: DeleteRequest :exec
DELETE FROM emp_loan_req WHERE id = ?;

-- name: GetTotalLoanRows :one
SELECT COUNT(*) AS Total_Rows FROM emp_loan_req;

-- name: UpdateRequest :exec
UPDATE emp_loan_req
SET reason = ?, amount = ?
WHERE id = ?;

-- name: GetRequests :many
SELECT 
    e.id,
    e.emp_id,
    CONCAT(em.first_name, '' , em.last_name) AS name,
    e.reason,
    e.amount,
    e.status,
    e.declined_by,
    e.decline_reason,
    e.requested_date,
    e.status_changed_date
FROM emp_loan_req e
JOIN HR_Employee em ON e.emp_id = em.id
WHERE 
    (? = '' OR em.first_name LIKE CONCAT('%',?,'%'))
    AND(? = '' OR em.last_name LIKE CONCAT('%', ? ,'%'))    
ORDER BY requested_date ASC
LIMIT ? OFFSET ?;

-- name: GetRequestsAdmin :many
SELECT 
    e.id,
    e.emp_id,
    CONCAT(em.first_name, " " , em.last_name) AS name,
    usr.branch_id,
    br.name,
    e.reason,
    e.amount,
    e.status,
    e.declined_by,
    e.decline_reason,
    e.requested_date,
    e.status_changed_date
FROM emp_loan_req e
JOIN HR_EMP_User usr ON e.emp_id = usr.employee_id
JOIN HR_Branch br ON usr.branch_id = br.id
JOIN HR_Employee em ON e.emp_id = em.id
WHERE 
    (1=br.id OR br.id = ?)
    AND(? = '' OR em.first_name LIKE CONCAT('%',?,'%'))
    AND(? = '' OR em.last_name LIKE CONCAT('%', ? ,'%'))
ORDER BY requested_date ASC
LIMIT ? OFFSET ?;

-- name: UpdateRequestStatus :exec
UPDATE emp_loan_req 
SET status = ?, declined_by = ?, decline_reason = ? 
WHERE id = ?;