-- +goose Up
-- +goose StatementBegin
ALTER TABLE emp_link 
MODIFY COLUMN updated_by BIGINT(20) NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE emp_link 
MODIFY COLUMN updated_by BIGINT(20) NOT NULL;
-- +goose StatementEnd
