-- name: CreateFileSubmit :exec
INSERT INTO HR_FileSubmit (
    employee_id, file_name, file_type
) VALUES (
    ?, ?, ?
);

-- name: UpdateFileSubmit :exec
UPDATE HR_FileSubmit
SET file_name = ?, file_type = ?
WHERE employee_id = ?;