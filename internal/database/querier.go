// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"context"
	"database/sql"
)

type Querier interface {
	Create_HR_Admin(ctx context.Context, arg Create_HR_AdminParams) (sql.Result, error)
}

var _ Querier = (*Queries)(nil)
