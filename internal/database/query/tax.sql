-- name: CreateTax :exec
INSERT INTO HR_Tax (
    tax_from, tax_to, tax_percentage, updated_by
) VALUES (
    ?, ?, ?, ?
);

-- name: GetTax :many
SELECT t.id, t.tax_from, t.tax_to, t.tax_percentage, t.created_at, a.user_name AS updated_by
FROM HR_Tax t
LEFT JOIN HR_Admin a ON t.updated_by = a.id;

-- name: UpdateTax :exec
UPDATE HR_Tax
SET tax_from = ?, tax_to = ?, tax_percentage = ?
WHERE id = ?;

-- name: DeleteTax :exec
DELETE FROM HR_Tax
WHERE id = ?;
