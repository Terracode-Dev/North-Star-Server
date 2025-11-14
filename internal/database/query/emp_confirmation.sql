-- name: CreateConfirmation :exec
INSERT INTO emp_confirmation_letter_table(emp_id) VALUES (
    ?
);

-- name: GetConfirmation :many
SELECT 
    c.id,
    c.emp_id,
    CONCAT(e.file_name, " " ,e.last_name) AS name,
    c.created_at
FROM emp_confirmation_letter_table c
JOIN HR_Employee e ON c.emp_id = e.id
WHERE 
    (? = '' OR e.first_name LIKE CONCAT('%',?,'%'))
    AND(? = '' OR e.last_name LIKE CONCAT('%', ? ,'%')) 
LIMIT ? OFFSET ?;

-- name: DeleteConfirmation :exec
DELETE FROM emp_confirmation_letter_table WHERE id = ?;