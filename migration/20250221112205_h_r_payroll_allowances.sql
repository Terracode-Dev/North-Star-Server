-- +goose Up
-- +goose StatementBegin
CREATE TABLE HR_Payroll_Allowances(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    payroll_id BIGINT NOT NULL,
    FOREIGN KEY (payroll_id) REFERENCES HR_Payroll(id),
    updated_by BIGINT,
    FOREIGN KEY (updated_by) REFERENCES HR_Admin(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE HR_Payroll_Allowances;
-- +goose StatementEnd
