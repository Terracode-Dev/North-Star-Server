// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: tax.sql

package database

import (
	"context"

	"github.com/shopspring/decimal"
)

const createTax = `-- name: CreateTax :exec
INSERT INTO HR_Tax (
    tax_from, tax_to, tax_percentage
) VALUES (
    ?, ?, ?
)
`

type CreateTaxParams struct {
	TaxFrom       decimal.Decimal `json:"tax_from"`
	TaxTo         decimal.Decimal `json:"tax_to"`
	TaxPercentage decimal.Decimal `json:"tax_percentage"`
}

func (q *Queries) CreateTax(ctx context.Context, arg CreateTaxParams) error {
	_, err := q.db.ExecContext(ctx, createTax, arg.TaxFrom, arg.TaxTo, arg.TaxPercentage)
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
SELECT id, tax_from, tax_to, tax_percentage, created_at FROM HR_Tax
`

func (q *Queries) GetTax(ctx context.Context) ([]HrTax, error) {
	rows, err := q.db.QueryContext(ctx, getTax)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []HrTax
	for rows.Next() {
		var i HrTax
		if err := rows.Scan(
			&i.ID,
			&i.TaxFrom,
			&i.TaxTo,
			&i.TaxPercentage,
			&i.CreatedAt,
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
