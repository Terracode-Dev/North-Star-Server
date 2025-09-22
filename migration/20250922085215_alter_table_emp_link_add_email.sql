-- +goose Up
-- +goose StatementBegin
ALTER TABLE emp_link ADD COLUMN email VARCHAR(255) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE emp_link DROP COLUMN email;
-- +goose StatementEnd
