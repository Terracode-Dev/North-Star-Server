-- name: CreateTax :exec
INSERT INTO HR_Tax (
    tax_from, tax_to, tax_percentage
) VALUES (
    ?, ?, ?
);

-- name: GetTax :many
SELECT * FROM HR_Tax;

-- name: UpdateTax :exec
UPDATE HR_Tax
SET tax_from = ?, tax_to = ?, tax_percentage = ?
WHERE id = ?;

-- name: DeleteTax :exec
DELETE FROM HR_Tax
WHERE id = ?;
