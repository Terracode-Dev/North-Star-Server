-- +goose Up
-- +goose StatementBegin
ALTER TABLE emp_confirmation_letter_table
ADD CONSTRAINT fk_emp_confirmation_employee
FOREIGN KEY (emp_id) REFERENCES HR_Employee(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE emp_confirmation_letter_table
DROP CONSTRAINT fk_emp_confirmation_employee;
-- +goose StatementEnd
