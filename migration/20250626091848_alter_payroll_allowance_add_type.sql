-- +goose Up
-- +goose StatementBegin
ALTER TABLE HR_Payroll_Allowances
ADD COLUMN amount_type VARCHAR(100) NOT NULL AFTER amount;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE HR_Payroll_Allowances
DROP COLUMN amount_type;
-- +goose StatementEnd
