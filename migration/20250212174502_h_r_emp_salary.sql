-- +goose Up
-- +goose StatementBegin
CREATE TABLE HR_EMP_Salary(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    salary_type VARCHAR(100) NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    Total_of_salary_allowances INT NOT NULL,
    pension_employer INT NOT NULL,
    pension_employee INT NOT NULL,
    total_net_salary INT NOT NULL,
    updated_by BIGINT,
    employee_id BIGINT NOT NULL,
    FOREIGN KEY (employee_id) REFERENCES HR_Employee(id),
    FOREIGN KEY (updated_by) REFERENCES HR_Admin(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE HR_EMP_Salary;
-- +goose StatementEnd
