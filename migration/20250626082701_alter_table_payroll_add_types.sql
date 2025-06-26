-- +goose Up
-- +goose StatementBegin
ALTER TABLE HR_Payroll
ADD COLUMN salary_amount_type VARCHAR(100) NOT NULL AFTER amount,
ADD COLUMN total_allowances_type VARCHAR(100) NOT NULL AFTER total_of_salary_allowances,
ADD COLUMN pension_employer_type VARCHAR(100)  AFTER pension_employer,
ADD COLUMN pension_employee_type VARCHAR(100)  AFTER pension_employee,
ADD COLUMN total_net_salary_type VARCHAR(100) NOT NULL AFTER total_net_salary,
ADD COLUMN total_net_salary_after_tax_type VARCHAR(100) NOT NULL AFTER total_net_salary_after_tax;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE HR_Payroll
DROP COLUMN salary_amount_type,
DROP COLUMN total_allowances_type,
DROP COLUMN pension_employer_type,
DROP COLUMN pension_employee_type,
DROP COLUMN total_net_salary_type,
DROP COLUMN total_net_salary_after_tax_type;
-- +goose StatementEnd
