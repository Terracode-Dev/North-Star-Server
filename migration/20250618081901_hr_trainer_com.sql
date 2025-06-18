-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS HR_Trainer_Com (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    payroll_id BIGINT NOT NULL,
    trainer_id BIGINT NOT NULL,
    employee_id BIGINT NOT NULL,
    commission DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    assigned_count BIGINT NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (payroll_id) REFERENCES HR_Payroll(id),
    FOREIGN KEY (employee_id) REFERENCES HR_Employee(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS HR_Trainer_Com;
-- +goose StatementEnd
