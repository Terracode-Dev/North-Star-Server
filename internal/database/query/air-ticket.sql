-- name: CreateEmpAirticketReq :execresult
INSERT INTO emp_airticket_req (
    passenger_name,
    passenger_email,
    passport_number,
    departure_date,
    return_date,
    departure_city,
    arrival_city,
    reason,
    emp_id,
    branch_id,
    status
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: DeleteEmpAirticketReq :exec
DELETE FROM emp_airticket_req
WHERE id = ?;

-- name: UpdateEmpAirticketReq :exec
UPDATE emp_airticket_req
SET 
    passenger_name = ?,
    passenger_email = ?,
    passport_number = ?,
    departure_date = ?,
    return_date = ?,
    departure_city = ?,
    arrival_city = ?,
    reason = ?,
    emp_id = ?,
    branch_id = ?
WHERE id = ?;

-- name: GetEmpAirticketReqByEmpAndBranch :many
SELECT * FROM emp_airticket_req
WHERE emp_id = ? AND branch_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetEmpAirticketReqByBranch :many
SELECT * FROM emp_airticket_req
WHERE branch_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: SetEmpAirticketReqStatus :exec
UPDATE emp_airticket_req
SET status = ?
WHERE id = ?;

-- name: GetEmpAirticketReqByID :one
SELECT * FROM emp_airticket_req
WHERE id = ?;

-- name: CountEmpAirticketReqByEmpAndBranch :one
SELECT COUNT(*) FROM emp_airticket_req
WHERE emp_id = ? AND branch_id = ?;

-- name: CountEmpAirticketReqByBranch :one
SELECT COUNT(*) FROM emp_airticket_req
WHERE branch_id = ?;

-- name: GetEmpAirticketReqByBranchAndStatus :many
SELECT * FROM emp_airticket_req
WHERE branch_id = ? AND status = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetEmpAirticketReqByStatus :many
SELECT * FROM emp_airticket_req
WHERE status = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: CountEmpAirticketReqByBranchAndStatus :one
SELECT COUNT(*) FROM emp_airticket_req
WHERE branch_id = ? AND status = ?;

-- name: CountEmpAirticketReqByStatus :one
SELECT COUNT(*) FROM emp_airticket_req
WHERE status = ?;
