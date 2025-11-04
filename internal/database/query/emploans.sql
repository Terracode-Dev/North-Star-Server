-- name: CreateRequest :exec
INSERT INTO emp_loan_req (emp_id, reason, amount, status, declined_by, decline_reason) VALUES (
    ?, ?, ?, ?, ?, ?
);

-- name: DeleteRequest :exec
DELETE FROM emp_loan_req WHERE id = ?;

-- name: UpdateRequest :exec
UPDATE emp_loan_req
SET reason = ?, amount = ?
WHERE id = ?;

-- name: GetRequests :many
SELECT 
    id,
    emp_id,
    reason,
    amount,
    status,
    declined_by,
    decline_reason,
    requested_date,
    status_changed_date
FROM emp_loan_req 
ORDER BY requested_date ASC
LIMIT ? OFFSET ?;

-- name: GetRequestsAdmin :many
SELECT 
    e.id,
    e.emp_id,
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
WHERE (1=br.id OR br.id = ?)
ORDER BY requested_date ASC
LIMIT ? OFFSET ?;

-- name: UpdateRequestStatus :exec
UPDATE emp_loan_req 
SET status = ?, declined_by = ?, decline_reason = ? 
WHERE id = ?;