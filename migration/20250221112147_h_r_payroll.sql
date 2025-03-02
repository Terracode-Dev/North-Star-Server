-- +goose Up
-- +goose StatementBegin
CREATE TABLE HR_Payroll(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    employee VARCHAR(200) NOT NULL,
    date DATE NOT NULL,
    salary_type VARCHAR(200) NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    total_of_salary_allowances DECIMAL(10,2) NOT NULL,
    pension BOOLEAN NOT NULL,
    pension_employer DECIMAL(10,2),
    pension_employee DECIMAL(10,2),
    total_net_salary DECIMAL(10,2) NOT NULL,
    tax BOOLEAN NOT NULL,
    tax_percentage DECIMAL(10,2),
    total_net_salary_after_tax DECIMAL(10,2) NOT NULL,
    updated_by BIGINT,
    FOREIGN KEY (updated_by) REFERENCES HR_Admin(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE HR_Payroll;
-- +goose StatementEnd
