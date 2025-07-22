-- +goose Up
-- +goose StatementBegin
ALTER TABLE HR_Employee
ADD COLUMN is_ban BOOLEAN NOT NULL DEFAULT FALSE AFTER updated_by;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE HR_Employee
DROP COLUMN is_ban;
-- +goose StatementEnd
