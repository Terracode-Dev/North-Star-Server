// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: tax.sql

package database

import (
	"context"
	"database/sql"

	"github.com/shopspring/decimal"
)

const createTax = `-- name: CreateTax :exec
INSERT INTO HR_Tax (
    tax_from, tax_to, tax_percentage, updated_by
) VALUES (
    ?, ?, ?, ?
)
`

type CreateTaxParams struct {
	TaxFrom       decimal.Decimal `json:"tax_from"`
	TaxTo         decimal.Decimal `json:"tax_to"`
	TaxPercentage decimal.Decimal `json:"tax_percentage"`
	UpdatedBy     sql.NullInt64   `json:"updated_by"`
}

func (q *Queries) CreateTax(ctx context.Context, arg CreateTaxParams) error {
	_, err := q.db.ExecContext(ctx, createTax,
		arg.TaxFrom,
		arg.TaxTo,
		arg.TaxPercentage,
		arg.UpdatedBy,
	)
	return err
}

const deleteTax = `-- name: DeleteTax :exec
DELETE FROM HR_Tax
WHERE id = ?
`

func (q *Queries) DeleteTax(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTax, id)
	return err
}

const getTax = `-- name: GetTax :many
SELECT t.id, t.tax_from, t.tax_to, t.tax_percentage, t.created_at, a.user_name AS updated_by
FROM HR_Tax t
LEFT JOIN HR_Admin a ON t.updated_by = a.id
`

type GetTaxRow struct {
	ID            int64           `json:"id"`
	TaxFrom       decimal.Decimal `json:"tax_from"`
	TaxTo         decimal.Decimal `json:"tax_to"`
	TaxPercentage decimal.Decimal `json:"tax_percentage"`
	CreatedAt     sql.NullTime    `json:"created_at"`
	UpdatedBy     sql.NullString  `json:"updated_by"`
}

func (q *Queries) GetTax(ctx context.Context) ([]GetTaxRow, error) {
	rows, err := q.db.QueryContext(ctx, getTax)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTaxRow
	for rows.Next() {
		var i GetTaxRow
		if err := rows.Scan(
			&i.ID,
			&i.TaxFrom,
			&i.TaxTo,
			&i.TaxPercentage,
			&i.CreatedAt,
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

const updateTax = `-- name: UpdateTax :exec
UPDATE HR_Tax
SET tax_from = ?, tax_to = ?, tax_percentage = ?
WHERE id = ?
`

type UpdateTaxParams struct {
	TaxFrom       decimal.Decimal `json:"tax_from"`
	TaxTo         decimal.Decimal `json:"tax_to"`
	TaxPercentage decimal.Decimal `json:"tax_percentage"`
	ID            int64           `json:"id"`
}

func (q *Queries) UpdateTax(ctx context.Context, arg UpdateTaxParams) error {
	_, err := q.db.ExecContext(ctx, updateTax,
		arg.TaxFrom,
		arg.TaxTo,
		arg.TaxPercentage,
		arg.ID,
	)
	return err
}
