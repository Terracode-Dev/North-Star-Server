// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: services.sql

package database

import (
	"context"
	"database/sql"
)

const createServices = `-- name: CreateServices :exec
INSERT INTO HR_Create_Services (
    category, value, updated_by
) VALUES (
    ?, ?, ?
)
`

type CreateServicesParams struct {
	Category  string        `json:"category"`
	Value     string        `json:"value"`
	UpdatedBy sql.NullInt64 `json:"updated_by"`
}

func (q *Queries) CreateServices(ctx context.Context, arg CreateServicesParams) error {
	_, err := q.db.ExecContext(ctx, createServices, arg.Category, arg.Value, arg.UpdatedBy)
	return err
}

const deleteService = `-- name: DeleteService :exec
DELETE FROM HR_Create_Services
WHERE id = ?
`

func (q *Queries) DeleteService(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteService, id)
	return err
}

const getService = `-- name: GetService :one
SELECT 
    s.id, 
    s.category, 
    s.value, 
    a.user_name AS updated_by,
    s.created_at, 
    s.updated_at
FROM HR_Create_Services s
LEFT JOIN HR_Admin a ON s.updated_by = a.id
WHERE s.category = ?
`

type GetServiceRow struct {
	ID        int64          `json:"id"`
	Category  string         `json:"category"`
	Value     string         `json:"value"`
	UpdatedBy sql.NullString `json:"updated_by"`
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
}

func (q *Queries) GetService(ctx context.Context, category string) (GetServiceRow, error) {
	row := q.db.QueryRowContext(ctx, getService, category)
	var i GetServiceRow
	err := row.Scan(
		&i.ID,
		&i.Category,
		&i.Value,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getServices = `-- name: GetServices :many
SELECT id, category, value, updated_by, created_at, updated_at FROM HR_Create_Services
`

func (q *Queries) GetServices(ctx context.Context) ([]HrCreateService, error) {
	rows, err := q.db.QueryContext(ctx, getServices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []HrCreateService
	for rows.Next() {
		var i HrCreateService
		if err := rows.Scan(
			&i.ID,
			&i.Category,
			&i.Value,
			&i.UpdatedBy,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateService = `-- name: UpdateService :exec
UPDATE HR_Create_Services
SET category = ?, value = ?, updated_by = ?
WHERE id = ?
`

type UpdateServiceParams struct {
	Category  string        `json:"category"`
	Value     string        `json:"value"`
	UpdatedBy sql.NullInt64 `json:"updated_by"`
	ID        int64         `json:"id"`
}

func (q *Queries) UpdateService(ctx context.Context, arg UpdateServiceParams) error {
	_, err := q.db.ExecContext(ctx, updateService,
		arg.Category,
		arg.Value,
		arg.UpdatedBy,
		arg.ID,
	)
	return err
}
