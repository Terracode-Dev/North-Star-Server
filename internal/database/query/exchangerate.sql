-- name: CreateExchangeRate :exec
INSERT INTO Exchange_Rate (exchange_rate, currency_type)
VALUES (?, ?);

-- name: GetLatestExchangeRate :many
SELECT exchange_rate, id
FROM Exchange_Rate
WHERE currency_type = ?
ORDER BY created_at DESC
LIMIT 1;

-- name: DeleteExchangeRate :exec
DELETE FROM Exchange_Rate
WHERE id = ?;

-- name: GetExchangeRateAll :many
SELECT exchange_rate, id, currency_type
FROM Exchange_Rate
ORDER BY created_at DESC;

-- name: GetExhangeRateById :one
SELECT exchange_rate, currency_type
FROM Exchange_Rate
WHERE id = ?;