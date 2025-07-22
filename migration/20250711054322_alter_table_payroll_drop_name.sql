-- +goose Up
-- +goose StatementBegin
ALTER TABLE HR_Payroll
DROP COLUMN employee;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE HR_Payroll
ADD COLUMN employee VARCHAR(200) NOT NULL;
-- +goose StatementEnd
