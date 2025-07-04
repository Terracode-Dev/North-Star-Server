// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: filesubmit.sql

package database

import (
	"context"
)

const createFileSubmit = `-- name: CreateFileSubmit :exec
INSERT INTO HR_FileSubmit (
    employee_id, file_name, file_type
) VALUES (
    ?, ?, ?
)
`

type CreateFileSubmitParams struct {
	EmployeeID int64  `json:"employee_id"`
	FileName   string `json:"file_name"`
	FileType   string `json:"file_type"`
}

func (q *Queries) CreateFileSubmit(ctx context.Context, arg CreateFileSubmitParams) error {
	_, err := q.db.ExecContext(ctx, createFileSubmit, arg.EmployeeID, arg.FileName, arg.FileType)
	return err
}

const deleteFileSubmit = `-- name: DeleteFileSubmit :exec
DELETE FROM HR_FileSubmit WHERE employee_id = ?
`

func (q *Queries) DeleteFileSubmit(ctx context.Context, employeeID int64) error {
	_, err := q.db.ExecContext(ctx, deleteFileSubmit, employeeID)
	return err
}

const updateFileSubmit = `-- name: UpdateFileSubmit :exec
UPDATE HR_FileSubmit
SET file_name = ?, file_type = ?
WHERE employee_id = ?
`

type UpdateFileSubmitParams struct {
	FileName   string `json:"file_name"`
	FileType   string `json:"file_type"`
	EmployeeID int64  `json:"employee_id"`
}

func (q *Queries) UpdateFileSubmit(ctx context.Context, arg UpdateFileSubmitParams) error {
	_, err := q.db.ExecContext(ctx, updateFileSubmit, arg.FileName, arg.FileType, arg.EmployeeID)
	return err
}
