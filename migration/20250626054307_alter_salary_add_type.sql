-- +goose Up
-- +goose StatementBegin
ALTER TABLE HR_EMP_Salary
ADD COLUMN salary_amount_type VARCHAR(100) NOT NULL AFTER amount,
ADD COLUMN total_salary_allowances_type VARCHAR(100) NOT NULL AFTER Total_of_salary_allowances,
ADD COLUMN pension_employer_type VARCHAR(100) NOT NULL AFTER pension_employer,
ADD COLUMN pension_employee_type VARCHAR(100) NOT NULL AFTER pension_employee,
ADD COLUMN total_net_salary_type VARCHAR(100) NOT NULL AFTER total_net_salary;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE HR_EMP_Salary
DROP COLUMN salary_amount_type,
DROP COLUMN total_salary_allowances_type,
DROP COLUMN pension_employer_type,
DROP COLUMN pension_employee_type,
DROP COLUMN total_net_salary_type;
-- +goose StatementEnd
