// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: allowances.sql

package database

import (
	"context"
	"database/sql"

	"github.com/shopspring/decimal"
)

const createAllowances = `-- name: CreateAllowances :exec
INSERT INTO HR_Create_Allowances (
    allowance_type, amount, updated_by
) VALUES (
    ?, ?, ?
)
`

type CreateAllowancesParams struct {
	AllowanceType string          `json:"allowance_type"`
	Amount        decimal.Decimal `json:"amount"`
	UpdatedBy     sql.NullInt64   `json:"updated_by"`
}

func (q *Queries) CreateAllowances(ctx context.Context, arg CreateAllowancesParams) error {
	_, err := q.db.ExecContext(ctx, createAllowances, arg.AllowanceType, arg.Amount, arg.UpdatedBy)
	return err
}

const deleteAllowance = `-- name: DeleteAllowance :exec
DELETE FROM HR_Create_Allowances
WHERE id = ?
`

func (q *Queries) DeleteAllowance(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAllowance, id)
	return err
}

const getAllowance = `-- name: GetAllowance :one
SELECT id, allowance_type, amount, updated_by, created_at, updated_at FROM HR_Create_Allowances
WHERE id = ?
`

func (q *Queries) GetAllowance(ctx context.Context, id int64) (HrCreateAllowance, error) {
	row := q.db.QueryRowContext(ctx, getAllowance, id)
	var i HrCreateAllowance
	err := row.Scan(
		&i.ID,
		&i.AllowanceType,
		&i.Amount,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllowances = `-- name: GetAllowances :many
SELECT h.id, h.allowance_type, h.amount, h.created_at, h.updated_at, a.user_name as updated_by
FROM HR_Create_Allowances h
LEFT JOIN HR_Admin a ON a.id = h.updated_by
`

type GetAllowancesRow struct {
	ID            int64           `json:"id"`
	AllowanceType string          `json:"allowance_type"`
	Amount        decimal.Decimal `json:"amount"`
	CreatedAt     sql.NullTime    `json:"created_at"`
	UpdatedAt     sql.NullTime    `json:"updated_at"`
	UpdatedBy     sql.NullString  `json:"updated_by"`
}

func (q *Queries) GetAllowances(ctx context.Context) ([]GetAllowancesRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllowances)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllowancesRow
	for rows.Next() {
		var i GetAllowancesRow
		if err := rows.Scan(
			&i.ID,
			&i.AllowanceType,
			&i.Amount,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UpdatedBy,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAllowance = `-- name: UpdateAllowance :exec
UPDATE HR_Create_Allowances
SET allowance_type = ?, amount = ?, updated_by = ?
WHERE id = ?
`

type UpdateAllowanceParams struct {
	AllowanceType string          `json:"allowance_type"`
	Amount        decimal.Decimal `json:"amount"`
	UpdatedBy     sql.NullInt64   `json:"updated_by"`
	ID            int64           `json:"id"`
}

func (q *Queries) UpdateAllowance(ctx context.Context, arg UpdateAllowanceParams) error {
	_, err := q.db.ExecContext(ctx, updateAllowance,
		arg.AllowanceType,
		arg.Amount,
		arg.UpdatedBy,
		arg.ID,
	)
	return err
}
