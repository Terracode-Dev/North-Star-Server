-- +goose Up
-- +goose StatementBegin
ALTER TABLE emp_link ADD COLUMN updated_by BIGINT(20) NULL, ADD FOREIGN KEY (updated_by) REFERENCES HR_Employee (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE emp_link DROP FOREIGN KEY (updated_by);
ALTER TABLE emp_link DROP COLUMN updated_by;
-- +goose StatementEnd
