-- +goose Up
-- +goose StatementBegin
ALTER TABLE HR_Payroll
ADD COLUMN emp_id BIGINT AFTER employee,
ADD CONSTRAINT fk_payroll_employee
    FOREIGN KEY (emp_id) REFERENCES HR_Employee(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE HR_Payroll
DROP COLUMN emp_id,
DROP FOREIGN KEY fk_payroll_employee;
-- +goose StatementEnd
