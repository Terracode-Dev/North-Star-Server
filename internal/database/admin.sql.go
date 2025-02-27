// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: admin.sql

package database

import (
	"context"
	"database/sql"
)

const adminLogin = `-- name: AdminLogin :one
SELECT a.id, a.role, a.status, a.branch_id, a.password, b.name AS branchName
FROM HR_Admin a
LEFT JOIN HR_Branch b ON a.branch_id = b.id
WHERE email = ?
`

type AdminLoginRow struct {
	ID         int64          `json:"id"`
	Role       string         `json:"role"`
	Status     string         `json:"status"`
	BranchID   int64          `json:"branch_id"`
	Password   string         `json:"password"`
	Branchname sql.NullString `json:"branchname"`
}

func (q *Queries) AdminLogin(ctx context.Context, email string) (AdminLoginRow, error) {
	row := q.db.QueryRowContext(ctx, adminLogin, email)
	var i AdminLoginRow
	err := row.Scan(
		&i.ID,
		&i.Role,
		&i.Status,
		&i.BranchID,
		&i.Password,
		&i.Branchname,
	)
	return i, err
}

const createHrAdmin = `-- name: CreateHrAdmin :exec
INSERT INTO HR_Admin (user_name, email, password, role, status, branch_id, created_by, updated_by
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?
)
`

type CreateHrAdminParams struct {
	UserName  string        `json:"user_name"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	Role      string        `json:"role"`
	Status    string        `json:"status"`
	BranchID  int64         `json:"branch_id"`
	CreatedBy sql.NullInt64 `json:"created_by"`
	UpdatedBy sql.NullInt64 `json:"updated_by"`
}

func (q *Queries) CreateHrAdmin(ctx context.Context, arg CreateHrAdminParams) error {
	_, err := q.db.ExecContext(ctx, createHrAdmin,
		arg.UserName,
		arg.Email,
		arg.Password,
		arg.Role,
		arg.Status,
		arg.BranchID,
		arg.CreatedBy,
		arg.UpdatedBy,
	)
	return err
}

const deleteHrAdmin = `-- name: DeleteHrAdmin :exec
DELETE FROM HR_Admin WHERE id = ?
`

func (q *Queries) DeleteHrAdmin(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteHrAdmin, id)
	return err
}

const selectHrAdmin = `-- name: SelectHrAdmin :many
SELECT
  a.id,
  a.user_name,
  a.email,
  a.role,
  a.status,
  b.name AS branch_name,
  a.created_at,
  a.updated_at
FROM HR_Admin a
LEFT JOIN HR_Branch b ON a.branch_id = b.id
WHERE 
  (
    a.user_name LIKE CONCAT('%', ?, '%')
    OR a.email    LIKE CONCAT('%', ?, '%')
    OR a.role     LIKE CONCAT('%', ?, '%')
    OR a.status   LIKE CONCAT('%', ?, '%')
  )
  AND ( ? = '' OR a.branch_id = ? )
ORDER BY a.id DESC
LIMIT ? OFFSET ?
`

type SelectHrAdminParams struct {
	CONCAT   interface{} `json:"CONCAT"`
	CONCAT_2 interface{} `json:"CONCAT_2"`
	CONCAT_3 interface{} `json:"CONCAT_3"`
	CONCAT_4 interface{} `json:"CONCAT_4"`
	Column5  interface{} `json:"column_5"`
	BranchID int64       `json:"branch_id"`
	Limit    int32       `json:"limit"`
	Offset   int32       `json:"offset"`
}

type SelectHrAdminRow struct {
	ID         int64          `json:"id"`
	UserName   string         `json:"user_name"`
	Email      string         `json:"email"`
	Role       string         `json:"role"`
	Status     string         `json:"status"`
	BranchName sql.NullString `json:"branch_name"`
	CreatedAt  sql.NullTime   `json:"created_at"`
	UpdatedAt  sql.NullTime   `json:"updated_at"`
}

func (q *Queries) SelectHrAdmin(ctx context.Context, arg SelectHrAdminParams) ([]SelectHrAdminRow, error) {
	rows, err := q.db.QueryContext(ctx, selectHrAdmin,
		arg.CONCAT,
		arg.CONCAT_2,
		arg.CONCAT_3,
		arg.CONCAT_4,
		arg.Column5,
		arg.BranchID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SelectHrAdminRow
	for rows.Next() {
		var i SelectHrAdminRow
		if err := rows.Scan(
			&i.ID,
			&i.UserName,
			&i.Email,
			&i.Role,
			&i.Status,
			&i.BranchName,
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

const selectOneHrAdmin = `-- name: SelectOneHrAdmin :one
SELECT id, user_name, email, password, role, status, branch_id, created_by, updated_by, created_at, updated_at FROM HR_Admin WHERE id = ?
`

func (q *Queries) SelectOneHrAdmin(ctx context.Context, id int64) (HrAdmin, error) {
	row := q.db.QueryRowContext(ctx, selectOneHrAdmin, id)
	var i HrAdmin
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.Email,
		&i.Password,
		&i.Role,
		&i.Status,
		&i.BranchID,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const suspendedHrAdmin = `-- name: SuspendedHrAdmin :exec
UPDATE HR_Admin SET status = ? WHERE id = ?
`

type SuspendedHrAdminParams struct {
	Status string `json:"status"`
	ID     int64  `json:"id"`
}

func (q *Queries) SuspendedHrAdmin(ctx context.Context, arg SuspendedHrAdminParams) error {
	_, err := q.db.ExecContext(ctx, suspendedHrAdmin, arg.Status, arg.ID)
	return err
}

const updateHrAdmin = `-- name: UpdateHrAdmin :exec
UPDATE HR_Admin SET user_name = ?, email = ?, role = ?, status = ?, branch_id = ? , updated_by = ? WHERE id = ?
`

type UpdateHrAdminParams struct {
	UserName  string        `json:"user_name"`
	Email     string        `json:"email"`
	Role      string        `json:"role"`
	Status    string        `json:"status"`
	BranchID  int64         `json:"branch_id"`
	UpdatedBy sql.NullInt64 `json:"updated_by"`
	ID        int64         `json:"id"`
}

func (q *Queries) UpdateHrAdmin(ctx context.Context, arg UpdateHrAdminParams) error {
	_, err := q.db.ExecContext(ctx, updateHrAdmin,
		arg.UserName,
		arg.Email,
		arg.Role,
		arg.Status,
		arg.BranchID,
		arg.UpdatedBy,
		arg.ID,
	)
	return err
}
